package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) prepareRename(context *glsp.Context, params *protocol.PrepareRenameParams) (any, error) {
	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	sym, found := findSymFromPos(doc, params.Position)
	if !found {
		return nil, nil
	}

	return parserRuleToRange(sym.Ident()), nil
}

func (s *tempoServer) renameSymbol(context *glsp.Context, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	sym, found := findSymFromPos(doc, params.Position)
	if !found {
		return nil, nil
	}

	edits := []protocol.TextEdit{{
		Range:   parserRuleToRange(sym.Ident()),
		NewText: params.NewName,
	}}

	for _, read := range sym.AccessReads() {
		edits = append(edits, protocol.TextEdit{
			Range:   parserRuleToRange(read),
			NewText: params.NewName,
		})
	}
	for _, write := range sym.AccessWrites() {
		edits = append(edits, protocol.TextEdit{
			Range:   parserRuleToRange(write),
			NewText: params.NewName,
		})
	}

	docChanges := map[protocol.DocumentUri][]protocol.TextEdit{}
	docChanges[doc.uri] = edits

	return &protocol.WorkspaceEdit{
		Changes: docChanges,
	}, nil
}
