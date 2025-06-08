package lsp

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) textDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {

	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	leaf, _ := astNodeAtPosition(doc.ast, params.Position)
	if leaf == nil {
		return nil, nil
	}

	var node antlr.Tree = leaf
	for node != nil {
		switch node := node.(type) {
		case *parser.StmtReturnContext:
			if exprType, ok := doc.info.Types[node.Expr()]; ok {

				if len(exprType.Roles().Participants()) == 0 {
					scope := doc.info.GlobalScope.Innermost(node.GetStart())
					exprType = exprType.ReplaceSharedRoles(scope.Roles().Participants())
				}

				stmtRange := parserRuleToRange(node)
				return hoverCode(fmt.Sprintf("return %s", exprType.ToString()), &stmtRange), nil
			}
		case *parser.IdentContext:
			if identSym, ok := doc.info.Symbols[node]; ok {
				identRange := parserRuleToRange(node)
				identCode := fmt.Sprintf("let %s: %s", identSym.SymbolName(), identSym.Type().ToString())
				if _, ok := identSym.Type().(*types.FunctionType); ok {
					identCode = identSym.Type().ToString()
				}
				return hoverCode(identCode, &identRange), nil
			}
		case parser.IExprContext:
			if exprType, ok := doc.info.Types[node]; ok {

				if len(exprType.Roles().Participants()) == 0 {
					scope := doc.info.GlobalScope.Innermost(node.GetStart())
					exprType = exprType.ReplaceSharedRoles(scope.Roles().Participants())
				}

				exprRange := parserRuleToRange(node)
				return hoverCode(exprType.ToString(), &exprRange), nil
			}
		}

		node = node.GetParent()
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
