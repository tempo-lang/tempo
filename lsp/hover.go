package lsp

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser/ast"
	"github.com/tempo-lang/tempo/types"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) textDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {

	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	leaf, _ := astNodeAtPosition(doc.source, doc.ast, params.Position)
	if leaf == nil {
		return nil, nil
	}

	switch node := leaf.(type) {
	case *ast.Identifier:
		if identSym, ok := doc.info.Symbols[node]; ok {
			identRange := parserRuleToRange(doc.source, node)
			identCode := fmt.Sprintf("let %s: %s", identSym.SymbolName(), identSym.Type().ToString())

			switch identSym.Type().(type) {
			case *types.FunctionType, *types.StructType, *types.InterfaceType:
				identCode = identSym.Type().ToString()
			}

			return hoverCode(identCode, &identRange), nil
		}
	case ast.Expr:
		if exprType, ok := doc.info.Types[node]; ok {

			if len(exprType.Roles().Participants()) == 0 {
				scope := doc.info.GlobalScope.Innermost(node.StartToken().Span.Start)
				exprType = exprType.ReplaceSharedRoles(scope.Roles().Participants())
			}

			exprRange := parserRuleToRange(doc.source, node)
			return hoverCode(exprType.ToString(), &exprRange), nil
		}
	}

	return nil, nil
}

func hoverCode(code string, highlightRange *protocol.Range) *protocol.Hover {
	return &protocol.Hover{
		Contents: protocol.MarkedStringStruct{
			Language: "tempo",
			Value:    code,
		},
		Range: highlightRange,
	}
}
