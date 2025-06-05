package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type InvalidDeclType struct {
	baseError
	DeclToken parser.IValueTypeContext
	DeclType  *types.Type
	ExprToken parser.IExprContext
	ExprType  *types.Type
}

func NewInvalidDeclType(declToken parser.IValueTypeContext, declType *types.Type, exprToken parser.IExprContext, exprType *types.Type) Error {
	return &InvalidDeclType{
		DeclToken: declToken,
		DeclType:  declType,
		ExprToken: exprToken,
		ExprType:  exprType,
	}
}

func (e *InvalidDeclType) Error() string {
	return fmt.Sprintf("invalid declaration type, expected '%s' found '%s'", e.DeclType.ToString(), e.ExprType.ToString())
}

func (e *InvalidDeclType) ParserRule() antlr.ParserRuleContext {
	return e.ExprToken
}

type InvalidAssignType struct {
	baseError
	Assign   *parser.StmtAssignContext
	VarType  *types.Type
	ExprType *types.Type
}

func (i *InvalidAssignType) Error() string {
	return fmt.Sprintf("invalid assignment type, expected '%s' found '%s'", i.VarType.ToString(), i.ExprType.ToString())
}

func (i *InvalidAssignType) ParserRule() antlr.ParserRuleContext {
	return i.Assign.Expr()
}

func NewInvalidAssignType(assign *parser.StmtAssignContext, varType *types.Type, exprType *types.Type) Error {
	return &InvalidAssignType{
		Assign:   assign,
		VarType:  varType,
		ExprType: exprType,
	}
}

type ReturnNotAllRoles struct {
	baseError
	Return       *parser.StmtReturnContext
	MissignRoles []string
}

func NewReturnNotAllRoles(ret *parser.StmtReturnContext, missingRoles []string) Error {
	return &ReturnNotAllRoles{
		Return:       ret,
		MissignRoles: missingRoles,
	}
}

func (e *ReturnNotAllRoles) Error() string {
	return fmt.Sprintf("return statement is missing roles '%s'", misc.JoinStrings(e.MissignRoles, ","))
}

func (e *ReturnNotAllRoles) ParserRule() antlr.ParserRuleContext {
	return e.Return
}
