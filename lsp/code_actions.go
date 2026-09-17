package lsp

import (
	"fmt"
	"github.com/tempo-lang/tempo/misc"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) codeAction(context *glsp.Context, params *protocol.CodeActionParams) (any, error) {
	logger.Infof("Calculating code actions for document: %s", params.TextDocument.URI)

	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	actions := []protocol.CodeAction{}
	for _, diagnostic := range doc.syntaxErrors {
		for _, fix := range diagnostic.Fixes {
			fixRange := spanToRange(doc.source, fix.Span)
			if !rangesOverlap(params.Range, spanToRange(doc.source, diagnostic.PrimarySpan)) && !rangesOverlap(params.Range, fixRange) {
				continue
			}
			title := "Apply parser fix"
			if fix.NewText != "" {
				title = fmt.Sprintf("Insert %q", fix.NewText)
			} else if !fix.Span.Empty() {
				title = "Remove unexpected token"
			}
			actions = append(actions, protocol.CodeAction{
				Title: title, Kind: misc.ToPtr(protocol.CodeActionKindQuickFix), IsPreferred: misc.ToPtr(true),
				Diagnostics: []protocol.Diagnostic{syntaxDiagnosticToDiagnostic(doc.source, diagnostic)},
				Edit:        &protocol.WorkspaceEdit{Changes: map[protocol.DocumentUri][]protocol.TextEdit{doc.uri: {{Range: fixRange, NewText: fix.NewText}}}},
			})
		}
	}

	for _, err := range doc.typeErrors {
		action := err.CodeAction()
		if action == nil {
			continue
		}

		if !rangesOverlap(params.Range, parserRuleToRange(doc.source, action.Range)) {
			continue
		}

		textEdit := protocol.TextEdit{
			Range:   parserRuleToRange(doc.source, action.Range),
			NewText: action.NewSource,
		}

		actions = append(actions, protocol.CodeAction{
			Title: action.Title,
			Kind:  misc.ToPtr(protocol.CodeActionKindQuickFix),
			Diagnostics: []protocol.Diagnostic{
				typeErrorToDiagnostic(doc.uri, doc.source, err),
			},
			IsPreferred: misc.ToPtr(true),
			Edit: &protocol.WorkspaceEdit{
				Changes: map[protocol.DocumentUri][]protocol.TextEdit{
					doc.uri: {textEdit},
				},
			},
		})
	}

	return actions, nil
}
