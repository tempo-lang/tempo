package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/types"

	"github.com/tempo-lang/tempo/parser/new_parser/ast"
)

type InvalidNumber struct {
	baseError
	Num    ast.Literal
	Reason error
}

func NewInvalidNumber(num ast.Literal, reason error) Error {
	return &InvalidNumber{
		Num:    num,
		Reason: reason,
	}
}

func (i *InvalidNumber) Error() string {
	return fmt.Sprintf("value `%s` is an invalid number", ast.Text(i.Num))
}

func (i *InvalidNumber) ParserRule() ast.Node {
	return i.Num
}

func (i *InvalidNumber) Annotations() []Annotation {
	return []Annotation{{
		Type:    AnnotationTypeNote,
		Message: "reason: " + i.Reason.Error(),
	}}
}

func (e *InvalidNumber) Code() ErrorCode {
	return CodeInvalidNumber
}

type ComNonLocalSender struct {
	baseError
	Com *ast.ComExpr
}

func (e *ComNonLocalSender) Error() string {
	return "only a single sender is allowed"
}

func (e *ComNonLocalSender) ParserRule() ast.Node {
	return e.Com.Sender
}

func (e *ComNonLocalSender) Code() ErrorCode {
	return CodeComNonLocalSender
}

func NewComNonLocalSender(com *ast.ComExpr) Error {
	return &ComNonLocalSender{
		Com: com,
	}
}

type ComValueNotAtSender struct {
	baseError
	Com      *ast.ComExpr
	ExprType types.Type
}

func (c *ComValueNotAtSender) Error() string {
	sender := ast.Text(c.Com.Sender)
	return fmt.Sprintf("value of type `%s` is not present at sender `%s`", c.ExprType.ToString(), sender)
}

func (c *ComValueNotAtSender) ParserRule() ast.Node {
	return c.Com.Expr
}

func (e *ComValueNotAtSender) Code() ErrorCode {
	return CodeComValueNotAtSender
}

func (c *ComValueNotAtSender) Annotations() []Annotation {
	roles := c.ExprType.Roles()
	if roles.IsLocalRole() {
		return []Annotation{
			{
				Type:    AnnotationTypeHint,
				Message: fmt.Sprintf("consider changing the sender to `%s`, so it matches the role of the value.", roles.Participants()[0]),
			},
		}
	} else {
		formattedRoles := formatList("", "", roles.Participants(), "or")
		return []Annotation{
			{
				Type:    AnnotationTypeHint,
				Message: fmt.Sprintf("consider changing the sender to %s, so it matches a role with the value.", formattedRoles),
			},
		}
	}
}

func NewComValueNotAtSender(com *ast.ComExpr, exprType types.Type) Error {
	return &ComValueNotAtSender{
		Com:      com,
		ExprType: exprType,
	}
}

type UnequatableType struct {
	baseError
	BinOp *ast.BinaryExpr
	Value types.Type
}

func (e *UnequatableType) Error() string {
	return fmt.Sprintf("values of type `%s` cannot be compared", e.Value.ToString())
}

func (e *UnequatableType) ParserRule() ast.Node {
	return e.BinOp
}

func (e *UnequatableType) Code() ErrorCode {
	return CodeUnequatableType
}

func NewUnequatableType(binOp *ast.BinaryExpr, value types.Type) Error {
	return &UnequatableType{
		BinOp: binOp,
		Value: value,
	}
}

type StructNotInitialized struct {
	baseError
	Ident *ast.IdentAccessExpr
}

func (e *StructNotInitialized) Error() string {
	return fmt.Sprintf("struct `%s` is not initialized", e.Ident.Ident.Value())
}

func (e *StructNotInitialized) ParserRule() ast.Node {
	return e.Ident
}

func (e *StructNotInitialized) Code() ErrorCode {
	return CodeStructNotInitialized
}

func (e *StructNotInitialized) Annotations() []Annotation {
	return []Annotation{{
		Type:    "hint",
		Message: fmt.Sprintf("add roles after the name of the structure, like %s@(A,B,C)", e.Ident.Ident.Value()),
	}}
}

func NewStructNotInitialized(ident *ast.IdentAccessExpr) Error {
	return &StructNotInitialized{
		Ident: ident,
	}
}

type TypeNotAnExpression struct {
	baseError
	Ident *ast.IdentAccessExpr
}

func (e *TypeNotAnExpression) Error() string {
	return fmt.Sprintf("type `%s` is not an expression", e.Ident.Ident.Value())
}

func (e *TypeNotAnExpression) ParserRule() ast.Node {
	return e.Ident
}

func (e *TypeNotAnExpression) Code() ErrorCode {
	return CodeTypeNotAnExpression
}

func NewTypeNotAnExpression(ident *ast.IdentAccessExpr) Error {
	return &TypeNotAnExpression{
		Ident: ident,
	}
}

type IncompatibleTypeCast struct {
	baseError
	CastFrom types.Type
	CastTo   types.Type
	CastExpr *ast.CallExpr
}

func (e *IncompatibleTypeCast) Error() string {
	return fmt.Sprintf("cannot cast value of type `%s` to `%s`", e.CastFrom.ToString(), e.CastTo.ToString())
}

func (e *IncompatibleTypeCast) ParserRule() ast.Node {
	return e.CastExpr
}

func (e *IncompatibleTypeCast) Code() ErrorCode {
	return CodeIncompatibleTypeCast
}

func NewIncompatibleTypeCast(castFrom types.Type, castTo types.Type, castExpr *ast.CallExpr) Error {
	return &IncompatibleTypeCast{
		CastFrom: castFrom,
		CastTo:   castTo,
		CastExpr: castExpr,
	}
}

type HiddenExpression struct {
	baseError
	Expr ast.Expr
	Type types.Type
}

func NewHiddenExpression(expr ast.Expr, exprType types.Type) Error {
	return &HiddenExpression{
		Expr: expr,
		Type: exprType,
	}
}

func (e *HiddenExpression) Error() string {
	return fmt.Sprintf("hidden expression of type `%s`", e.Type.ToString())
}

func (e *HiddenExpression) Annotations() []Annotation {
	return []Annotation{
		{
			Type:    AnnotationTypeNote,
			Message: "all the roles in the expression are hidden",
		},
	}
}

func (e *HiddenExpression) ParserRule() ast.Node {
	return e.Expr
}

func (e *HiddenExpression) Code() ErrorCode {
	return CodeHiddenExpression
}
