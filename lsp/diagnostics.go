package lsp

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/type_check"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) analyzeFile(ctx *glsp.Context, file *tempoFile) {
	logger.Infof("Analyzing file: %s", file.GetUri())

	diagnostics := []protocol.Diagnostic{}

	// parse source input
	inputStream := antlr.NewInputStream(file.GetSource())
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

	file.SetInfo(sourceFile, info)

	for _, err := range typeErrors {
		errorSeverity := protocol.DiagnosticSeverityError

		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range:    parserRuleToRange(err.ParserRule()),
			Severity: &errorSeverity,
			Message:  err.Error(),
		})
	}

	ctx.Notify(protocol.ServerTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
		URI:         file.GetUri(),
		Diagnostics: diagnostics,
	})
}
