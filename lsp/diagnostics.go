package lsp

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	parser "github.com/tempo-lang/tempo/parser/new_parser"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
	"github.com/tempo-lang/tempo/type_check"
	"github.com/tempo-lang/tempo/type_check/type_error"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) analyzeDocument(notify glsp.NotifyFunc, docUri protocol.URI, version int, source string) {
	logger.Infof("Calculating diagnostics of document: %s", docUri)

	diagnostics := []protocol.Diagnostic{}

	// parse source input
	sourceFile := token.SourceFromString(source)
	parsed := parser.ParseSource(sourceFile)
	for _, err := range parsed.Diagnostics {
		diagnostics = append(diagnostics, syntaxDiagnosticToDiagnostic(sourceFile, err))
	}

	// type check ast
	info, typeErrors := type_check.TypeCheck(parsed.Root)

	tempoDoc := newTempoDoc(docUri, version, sourceFile, parsed.Root, info, parsed.Diagnostics, typeErrors)
	s.UpdateDocument(tempoDoc)

	for _, err := range typeErrors {
		diagnostics = append(diagnostics, typeErrorToDiagnostic(docUri, sourceFile, err))
	}

	notify(protocol.ServerTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
		URI:         docUri,
		Version:     misc.ToPtr(protocol.UInteger(version)),
		Diagnostics: diagnostics,
	})
}

func syntaxDiagnosticToDiagnostic(source *token.Source, err parser.Diagnostic) protocol.Diagnostic {
	return protocol.Diagnostic{
		Range:    spanToRange(source, err.PrimarySpan),
		Severity: misc.ToPtr(protocol.DiagnosticSeverityError),
		Message:  err.Message,
		Code:     &protocol.IntegerOrString{Value: fmt.Sprintf("P%d", err.Code)},
		Source:   misc.ToPtr("tempo"),
	}
}

func typeErrorToDiagnostic(docUri protocol.URI, source *token.Source, err type_error.Error) protocol.Diagnostic {
	relatedInfo := []protocol.DiagnosticRelatedInformation{}
	for _, related := range err.RelatedInfo() {
		relatedInfo = append(relatedInfo, protocol.DiagnosticRelatedInformation{
			Location: protocol.Location{
				URI:   docUri,
				Range: parserRuleToRange(source, related.ParserRule),
			},
			Message: related.Message,
		})
	}

	return protocol.Diagnostic{
		Range:              parserRuleToRange(source, err.ParserRule()),
		Severity:           misc.ToPtr(protocol.DiagnosticSeverityError),
		Message:            err.Error(),
		RelatedInformation: relatedInfo,
		Code: &protocol.IntegerOrString{
			Value: fmt.Sprintf("E%d", err.Code()),
		},
		Source: misc.ToPtr("tempo"),
	}
}
