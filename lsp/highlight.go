package lsp

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) highlight(context *glsp.Context, params *protocol.DocumentHighlightParams) ([]protocol.DocumentHighlight, error) {
	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	leaf, _ := astNodeAtPosition(doc.ast, params.Position)
	if leaf == nil {
		return nil, nil
	}

	var sym sym_table.Symbol

	var node antlr.Tree = leaf
	for node != nil {
		if ident, ok := node.(*parser.IdentContext); ok {
			if identSym, ok := doc.info.Symbols[ident]; ok {
				sym = identSym
				break
			}
			return nil, nil
		}
		node = node.GetParent()
	}

	highlights := []protocol.DocumentHighlight{
		{
			Range: parserRuleToRange(sym.Ident()),
			Kind:  toPtr(protocol.DocumentHighlightKindWrite),
		},
	}

	for _, read := range sym.AccessReads() {
		highlights = append(highlights, protocol.DocumentHighlight{
			Range: parserRuleToRange(read),
			Kind:  toPtr(protocol.DocumentHighlightKindRead),
		})
	}

	for _, write := range sym.AccessWrites() {
		highlights = append(highlights, protocol.DocumentHighlight{
			Range: parserRuleToRange(write),
			Kind:  toPtr(protocol.DocumentHighlightKindWrite),
		})
	}

	return highlights, nil
}
