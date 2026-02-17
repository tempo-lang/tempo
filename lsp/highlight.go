package lsp

import (
	"github.com/tempo-lang/tempo/misc"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) highlight(context *glsp.Context, params *protocol.DocumentHighlightParams) ([]protocol.DocumentHighlight, error) {
	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	sym, found := findSymFromPos(doc, params.Position)
	if !found {
		return nil, nil
	}

	highlights := []protocol.DocumentHighlight{
		{
			Range: parserRuleToRange(sym.Ident()),
			Kind:  misc.ToPtr(protocol.DocumentHighlightKindWrite),
		},
	}

	for _, read := range sym.AccessReads() {
		highlights = append(highlights, protocol.DocumentHighlight{
			Range: parserRuleToRange(read),
			Kind:  misc.ToPtr(protocol.DocumentHighlightKindRead),
		})
	}

	for _, write := range sym.AccessWrites() {
		highlights = append(highlights, protocol.DocumentHighlight{
			Range: parserRuleToRange(write),
			Kind:  misc.ToPtr(protocol.DocumentHighlightKindWrite),
		})
	}

	return highlights, nil
}
