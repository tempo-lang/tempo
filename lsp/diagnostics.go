package lsp

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/type_check"
	"github.com/tempo-lang/tempo/type_check/type_error"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) analyzeDocument(notify glsp.NotifyFunc, docUri protocol.URI, version int, source string) {
	logger.Infof("Calculating diagnostics of document: %s", docUri)

	diagnostics := []protocol.Diagnostic{}

	// parse source input
	inputStream := antlr.NewInputStream(source)
	sourceFile, syntaxErrors := parser.Parse(inputStream)
	for _, err := range syntaxErrors {
		errorSeverity := protocol.DiagnosticSeverityError

		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      uint32(err.Line() - 1),
					Character: uint32(err.Column()),
				},
				End: protocol.Position{
					Line:      uint32(err.Line() - 1),
					Character: uint32(err.Column()),
				},
			},
			Severity: &errorSeverity,
			Message:  err.Message(),
		})
	}

	// type check ast
	info, typeErrors := type_check.TypeCheck(sourceFile)

	tempoDoc := newTempoDoc(docUri, version, source, sourceFile, info, typeErrors)
	s.UpdateDocument(tempoDoc)

	for _, err := range typeErrors {
		diagnostics = append(diagnostics, typeErrorToDiagnostic(docUri, err))
	}

	notify(protocol.ServerTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
		URI:         docUri,
		Version:     misc.ToPtr(protocol.UInteger(version)),
		Diagnostics: diagnostics,
	})
}

func typeErrorToDiagnostic(docUri protocol.URI, err type_error.Error) protocol.Diagnostic {
	relatedInfo := []protocol.DiagnosticRelatedInformation{}
	for _, related := range err.RelatedInfo() {
		relatedInfo = append(relatedInfo, protocol.DiagnosticRelatedInformation{
			Location: protocol.Location{
				URI:   docUri,
				Range: parserRuleToRange(related.ParserRule),
			},
			Message: related.Message,
		})
	}

	return protocol.Diagnostic{
		Range:              parserRuleToRange(err.ParserRule()),
		Severity:           misc.ToPtr(protocol.DiagnosticSeverityError),
		Message:            err.Error(),
		RelatedInformation: relatedInfo,
		Code: &protocol.IntegerOrString{
			Value: fmt.Sprintf("E%d", err.Code()),
		},
		Source: misc.ToPtr("tempo"),
	}
}
