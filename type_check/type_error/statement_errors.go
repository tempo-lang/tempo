package type_error

import (
	"fmt"
	"tempo/misc"
	"tempo/parser"
	"tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type InvalidDeclTypeError struct {
	DeclToken parser.IValueTypeContext
	DeclType  *types.Type
	ExprToken parser.IExprContext
	ExprType  *types.Type
}

func NewInvalidDeclTypeError(declToken parser.IValueTypeContext, declType *types.Type, exprToken parser.IExprContext, exprType *types.Type) Error {
	return &InvalidDeclTypeError{
		DeclToken: declToken,
		DeclType:  declType,
		ExprToken: exprToken,
		ExprType:  exprType,
	}
}

func (e *InvalidDeclTypeError) Error() string {
	return fmt.Sprintf("invalid declaration type, expected %s found %s", e.DeclType.ToString(), e.ExprType.ToString())
}

func (e *InvalidDeclTypeError) IsTypeError() {}

func (e *InvalidDeclTypeError) ParserRule() antlr.ParserRuleContext {
	return e.ExprToken
}

type InvalidAssignTypeError struct {
	Assign   *parser.StmtAssignContext
	VarType  *types.Type
	ExprType *types.Type
}

func (i *InvalidAssignTypeError) Error() string {
	return fmt.Sprintf("invalid assignment type, expected %s found %s", i.VarType.ToString(), i.ExprType.ToString())
}

func (i *InvalidAssignTypeError) IsTypeError() {}

func (i *InvalidAssignTypeError) ParserRule() antlr.ParserRuleContext {
	return i.Assign.Expr()
}

func NewInvalidAssignTypeError(assign *parser.StmtAssignContext, varType *types.Type, exprType *types.Type) Error {
	return &InvalidAssignTypeError{
		Assign:   assign,
		VarType:  varType,
		ExprType: exprType,
	}
}

type ReturnNotAllRolesError struct {
	Return       *parser.StmtReturnContext
	MissignRoles []string
}

func NewReturnNotAllRolesError(ret *parser.StmtReturnContext, missingRoles []string) Error {
	return &ReturnNotAllRolesError{
		Return:       ret,
		MissignRoles: missingRoles,
	}
}

func (e *ReturnNotAllRolesError) Error() string {
	return fmt.Sprintf("return statement is missing roles: %s", misc.JoinStrings(e.MissignRoles, ","))
}

func (e *ReturnNotAllRolesError) IsTypeError() {}

func (e *ReturnNotAllRolesError) ParserRule() antlr.ParserRuleContext {
	return e.Return.Expr()
}
