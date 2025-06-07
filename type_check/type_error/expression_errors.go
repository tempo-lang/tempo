package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type InvalidNumber struct {
	baseError
	Num    parser.ILiteralContext
	Reason error
}

func NewInvalidNumber(num parser.ILiteralContext, reason error) Error {
	return &InvalidNumber{
		Num:    num,
		Reason: reason,
	}
}

func (i *InvalidNumber) Error() string {
	return fmt.Sprintf("value `%s` is an invalid number", i.Num.GetText())
}

func (i *InvalidNumber) ParserRule() antlr.ParserRuleContext {
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
	Com *parser.ExprComContext
}

func (e *ComNonLocalSender) Error() string {
	return "only a single sender is allowed"
}

func (e *ComNonLocalSender) ParserRule() antlr.ParserRuleContext {
	return e.Com.GetSender()
}

func (e *ComNonLocalSender) Code() ErrorCode {
	return CodeComNonLocalSender
}

func NewComNonLocalSender(com *parser.ExprComContext) Error {
	return &ComNonLocalSender{
		Com: com,
	}
}

type ComValueNotAtSender struct {
	baseError
	Com      *parser.ExprComContext
	ExprType *types.Type
}

func (c *ComValueNotAtSender) Error() string {
	sender := parser.RoleTypeAllIdents(c.Com.RoleType(0))[0]
	return fmt.Sprintf("value of type `%s` is not present at sender `%s`", c.ExprType.ToString(), sender.GetText())
}

func (c *ComValueNotAtSender) ParserRule() antlr.ParserRuleContext {
	return c.Com.Expr()
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

func NewComValueNotAtSender(com *parser.ExprComContext, exprType *types.Type) Error {
	return &ComValueNotAtSender{
		Com:      com,
		ExprType: exprType,
	}
}

type UnequatableType struct {
	baseError
	BinOp *parser.ExprBinOpContext
	Value types.Value
}

func (e *UnequatableType) Error() string {
	return fmt.Sprintf("values of type `%s` cannot be compared", e.Value.ToString())
}

func (e *UnequatableType) ParserRule() antlr.ParserRuleContext {
	return e.BinOp
}

func (e *UnequatableType) Code() ErrorCode {
	return CodeUnequatableType
}

func NewUnequatableType(binOp *parser.ExprBinOpContext, value types.Value) Error {
	return &UnequatableType{
		BinOp: binOp,
		Value: value,
	}
}

type StructNotInitialized struct {
	baseError
	Ident *parser.ExprIdentContext
}

func (e *StructNotInitialized) Error() string {
	return fmt.Sprintf("struct `%s` is not initialized", e.Ident.GetText())
}

func (e *StructNotInitialized) ParserRule() antlr.ParserRuleContext {
	return e.Ident
}

func (e *StructNotInitialized) Code() ErrorCode {
	return CodeStructNotInitialized
}

func (e *StructNotInitialized) Annotations() []Annotation {
	return []Annotation{{
		Type:    "hint",
		Message: fmt.Sprintf("add roles after the name of the structure, like %s@(A,B,C)", e.Ident.GetText()),
	}}
}

func NewStructNotInitialized(ident *parser.ExprIdentContext) Error {
	return &StructNotInitialized{
		Ident: ident,
	}
}
