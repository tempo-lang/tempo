package lsp

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check"
	"github.com/tempo-lang/tempo/types"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) signatureHelp(context *glsp.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	leaf, _ := astNodeAtPosition(doc.ast, params.Position)
	if leaf == nil {
		return nil, nil
	}

	var exprCall *parser.ExprCallContext = nil
	var node antlr.Tree = leaf
	for node != nil {
		if callNode, ok := node.(*parser.ExprCallContext); ok {
			exprCall = callNode
			break
		}
		node = node.GetParent()
	}

	if exprCall == nil {
		return nil, nil
	}

	funcType := doc.info.Types[exprCall.Expr()]

	switch funcValue := funcType.Value().(type) {
	case *types.FunctionType:
		return funcSignatureHelp(params, doc.info, funcValue, exprCall)
	case *types.ClosureType:
		return closureSignatureHelp(params, funcType, exprCall)
	default:
		logger.Debugf("can not give type hint for non-function: %T", funcValue)
		return nil, nil
	}
}

func funcSignatureHelp(params *protocol.SignatureHelpParams, info *type_check.Info, funcValue *types.FunctionType, exprCall *parser.ExprCallContext) (*protocol.SignatureHelp, error) {
	funcSym := info.Symbols[funcValue.NameIdent()].(*sym_table.FuncSymbol)

	parameters := []protocol.ParameterInformation{}

	roleSubst, _ := funcSym.Roles().SubstituteMap(funcValue.Roles())

	paramLabels := []string{}
	for _, param := range funcSym.Params() {
		paramLabel := fmt.Sprintf("%s: %s", param.SymbolName(), param.Type().SubstituteRoles(roleSubst).ToString())
		paramLabels = append(paramLabels, paramLabel)
		parameters = append(parameters, protocol.ParameterInformation{
			Label:         paramLabel,
			Documentation: nil,
		})
	}

	var activeParameter *protocol.UInteger = nil
	for i, arg := range exprCall.FuncArgList().AllExpr() {
		argRange := parserRuleToRange(arg)
		if posWithinRange(params.Position, argRange) {
			p := uint32(i)
			activeParameter = &p
		}
	}

	funcLabel := fmt.Sprintf("func@%s %s(%s)", funcSym.Roles().SubstituteRoles(roleSubst).ToString(), funcSym.SymbolName(), misc.JoinStrings(paramLabels, ", "))
	if funcSym.FuncValue().ReturnType().Value() != types.Unit() {
		funcLabel = funcLabel + " " + funcSym.FuncValue().ReturnType().SubstituteRoles(roleSubst).ToString()
	}

	return &protocol.SignatureHelp{
		Signatures: []protocol.SignatureInformation{
			{
				Label:           funcLabel,
				Documentation:   nil,
				Parameters:      parameters,
				ActiveParameter: activeParameter,
			},
		},
	}, nil
}

func closureSignatureHelp(params *protocol.SignatureHelpParams, closureType *types.Type, exprCall *parser.ExprCallContext) (*protocol.SignatureHelp, error) {
	closureValue := closureType.Value().(*types.ClosureType)

	parameters := []protocol.ParameterInformation{}
	for _, param := range closureValue.Params() {
		paramLabel := param.ToString()
		parameters = append(parameters, protocol.ParameterInformation{
			Label:         paramLabel,
			Documentation: nil,
		})
	}

	var activeParameter *protocol.UInteger = nil
	for i, arg := range exprCall.FuncArgList().AllExpr() {
		argRange := parserRuleToRange(arg)
		if posWithinRange(params.Position, argRange) {
			p := uint32(i)
			activeParameter = &p
		}
	}

	funcLabel := closureType.ToString()

	return &protocol.SignatureHelp{
		Signatures: []protocol.SignatureInformation{
			{
				Label:           funcLabel,
				Documentation:   nil,
				Parameters:      parameters,
				ActiveParameter: activeParameter,
			},
		},
	}, nil
}
