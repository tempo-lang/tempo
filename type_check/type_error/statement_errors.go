package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/types"

	"github.com/tempo-lang/tempo/parser/ast"
)

type InvalidAssignType struct {
	baseError
	ExprCtx  ast.Expr
	VarType  types.Type
	ExprType types.Type
}

func (i *InvalidAssignType) Error() string {
	return fmt.Sprintf("cannot assign type `%s` to `%s`", i.ExprType.ToString(), i.VarType.ToString())
}

func (i *InvalidAssignType) Annotations() []Annotation {
	return []Annotation{{
		Type:    AnnotationTypeHint,
		Message: "check that the value type matches the expected variable type.",
	}}
}

func (i *InvalidAssignType) ParserRule() ast.Node {
	return i.ExprCtx
}

func (e *InvalidAssignType) Code() ErrorCode {
	return CodeInvalidAssignType
}

func NewInvalidAssignType(exprCtx ast.Expr, varType types.Type, exprType types.Type) Error {
	return &InvalidAssignType{
		ExprCtx:  exprCtx,
		VarType:  varType,
		ExprType: exprType,
	}
}

type ReturnNotAllRoles struct {
	baseError
	Return       *ast.ReturnStmt
	MissignRoles []string
}

func NewReturnNotAllRoles(ret *ast.ReturnStmt, missingRoles []string) Error {
	return &ReturnNotAllRoles{
		Return:       ret,
		MissignRoles: missingRoles,
	}
}

func (e *ReturnNotAllRoles) Error() string {
	roles := formatList("role", "roles", e.MissignRoles, "and")
	return fmt.Sprintf("%s %s missing from the return statement", roles, toBe(e.MissignRoles))
}

func (e *ReturnNotAllRoles) Annotations() []Annotation {
	return []Annotation{{
		Type:    AnnotationTypeNote,
		Message: "a function can only return if all roles participate in the return.",
	}, {
		Type:    AnnotationTypeHint,
		Message: "ensure the enclosing scope includes all roles.",
	}}
}

func (e *ReturnNotAllRoles) ParserRule() ast.Node {
	return e.Return
}

func (e *ReturnNotAllRoles) Code() ErrorCode {
	return CodeReturnNotAllRoles
}

type AssignUnitValue struct {
	baseError
	Expr ast.Expr
}

func (e *AssignUnitValue) Error() string {
	return "cannot assign a unit value to a variable"
}

func (e *AssignUnitValue) ParserRule() ast.Node {
	return e.Expr
}

func (e *AssignUnitValue) Code() ErrorCode {
	return CodeAssignUnitValue
}

func NewAssignUnitValue(expr ast.Expr) Error {
	return &AssignUnitValue{
		Expr: expr,
	}
}
