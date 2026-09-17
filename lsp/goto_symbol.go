package lsp

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) gotoDefinition(context *glsp.Context, params *protocol.DefinitionParams) (any, error) {
	return s.gotoDef(&params.TextDocumentPositionParams)
}

func (s *tempoServer) gotoDeclaration(context *glsp.Context, params *protocol.DeclarationParams) (any, error) {
	return s.gotoDef(&params.TextDocumentPositionParams)
}

func (s *tempoServer) gotoDef(params *protocol.TextDocumentPositionParams) (*protocol.Location, error) {
	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	sym, found := findSymFromPos(doc, params.Position)
	if !found {
		return nil, nil
	}

	location := protocol.Location{
		URI:   doc.uri,
		Range: parserRuleToRange(doc.source, sym.Ident()),
	}

	return &location, nil
}

func (s *tempoServer) gotoReferences(context *glsp.Context, params *protocol.ReferenceParams) ([]protocol.Location, error) {
	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	sym, found := findSymFromPos(doc, params.Position)
	if !found {
		return nil, nil
	}

	references := []protocol.Location{}
	for _, read := range sym.AccessReads() {
		references = append(references, protocol.Location{
			URI:   params.TextDocument.URI,
			Range: parserRuleToRange(doc.source, read),
		})
	}
	for _, write := range sym.AccessWrites() {
		references = append(references, protocol.Location{
			URI:   params.TextDocument.URI,
			Range: parserRuleToRange(doc.source, write),
		})
	}

	return references, nil
}

func findSymFromPos(doc *tempoDoc, pos protocol.Position) (sym_table.Symbol, bool) {
	leaf, _ := astNodeAtPosition(doc.source, doc.ast, pos)
	if leaf == nil {
		return nil, false
	}

	if ident, ok := leaf.(*ast.Identifier); ok {
		sym, found := doc.info.Symbols[ident]
		return sym, found
	}

	return nil, false
}
