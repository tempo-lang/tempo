package lsp

import (
	"fmt"
	"tempo/parser"
	"tempo/types"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) textDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {

	file, ok := s.files[params.TextDocument.URI]
	if !ok {
		return nil, nil
	}
	file.lock.RLock()
	defer file.lock.RUnlock()

	leaf, _ := astNodeAtPosition(file.ast, params.Position)
	if leaf == nil {
		return nil, nil
	}

	var node antlr.Tree = leaf
	for node != nil {
		switch node := node.(type) {
		case *parser.StmtReturnContext:
			if exprType, ok := file.info.Types[node.Expr()]; ok {

				if len(exprType.Roles().Participants()) == 0 {
					scope := file.info.GlobalScope.Innermost(node.GetStart())

					exprType = types.New(
						exprType.Value(),
						types.NewRole(scope.Roles().Participants(), true),
					)
				}

				stmtRange := parserRuleToRange(node)
				return hoverCode(fmt.Sprintf("return %s", exprType.ToString()), &stmtRange), nil
			}
		case *parser.IdentContext:
			if identSym, ok := file.info.Symbols[node]; ok {
				identRange := parserRuleToRange(node)
				return hoverCode(identSym.Type().ToString(), &identRange), nil
			}
		case parser.IExprContext:
			if exprType, ok := file.info.Types[node]; ok {

				if len(exprType.Roles().Participants()) == 0 {
					scope := file.info.GlobalScope.Innermost(node.GetStart())

					exprType = types.New(
						exprType.Value(),
						types.NewRole(scope.Roles().Participants(), true),
					)
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
