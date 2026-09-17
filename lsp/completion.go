package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) textDocumentCompletion(_ *glsp.Context, params *protocol.CompletionParams) (any, error) {
	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}
	node, _ := astNodeAtPosition(doc.source, doc.ast, params.Position)
	offset := 0
	if node != nil {
		offset = node.StartToken().Span.Start
	}
	scope := doc.info.GlobalScope.Innermost(offset)
	items := []protocol.CompletionItem{}
	seen := map[string]bool{}
	for cur := scope; cur != nil; cur = cur.Parent() {
		for sym := range cur.Symbols() {
			if seen[sym.SymbolName()] || sym.SymbolName() == "" {
				continue
			}
			seen[sym.SymbolName()] = true
			items = append(items, protocol.CompletionItem{Label: sym.SymbolName(), Detail: func() *string { v := sym.Type().ToString(); return &v }()})
		}
	}
	return items, nil
}
