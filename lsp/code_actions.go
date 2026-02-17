package lsp

import (
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

	for _, err := range doc.typeErrors {
		action := err.CodeAction()
		if action == nil {
			continue
		}

		if !rangesOverlap(params.Range, parserRuleToRange(action.Range)) {
			continue
		}

		textEdit := protocol.TextEdit{
			Range:   parserRuleToRange(action.Range),
			NewText: action.NewSource,
		}

		actions = append(actions, protocol.CodeAction{
			Title: action.Title,
			Kind:  misc.ToPtr(protocol.CodeActionKindQuickFix),
			Diagnostics: []protocol.Diagnostic{
				typeErrorToDiagnostic(doc.uri, err),
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
