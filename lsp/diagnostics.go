package lsp

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/type_check"

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

	tempoDoc := newTempoDoc(docUri, version, source, sourceFile, info)
	s.UpdateDocument(tempoDoc)

	for _, err := range typeErrors {
		errorSeverity := protocol.DiagnosticSeverityError

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

		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range:              parserRuleToRange(err.ParserRule()),
			Severity:           &errorSeverity,
			Message:            err.Error(),
			RelatedInformation: relatedInfo,
		})
	}

	notify(protocol.ServerTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
		URI:         docUri,
		Version:     toPtr(protocol.UInteger(version)),
		Diagnostics: diagnostics,
	})
}
