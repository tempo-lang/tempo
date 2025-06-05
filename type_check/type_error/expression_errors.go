package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type InvalidNumber struct {
	baseError
	Num parser.ILiteralContext
}

func NewInvalidNumber(num parser.ILiteralContext) Error {
	return &InvalidNumber{
		Num: num,
	}
}

func (i *InvalidNumber) Error() string {
	return fmt.Sprintf("invalid number '%s'", i.Num.GetText())
}

func (i *InvalidNumber) ParserRule() antlr.ParserRuleContext {
	return i.Num
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
	return fmt.Sprintf("value of type '%s' is not present at sender '%s'", c.ExprType.ToString(), sender.GetText())
}

func (c *ComValueNotAtSender) ParserRule() antlr.ParserRuleContext {
	return c.Com.Expr()
}

func (c *ComValueNotAtSender) Annotations() []Annotation {
	roles := c.ExprType.Roles()
	if roles.IsLocalRole() {
		return []Annotation{
			{
				Type:    AnnotationTypeHint,
				Message: fmt.Sprintf("change sender to '%s', so it matches the role of the value.", roles.Participants()[0]),
			},
		}
	} else {
		return []Annotation{
			{
				Type:    AnnotationTypeHint,
				Message: fmt.Sprintf("change sender to one of '%s', so it matches one of the roles having the value.", misc.JoinStrings(roles.Participants(), ", ")),
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
	return fmt.Sprintf("type value '%s' is not equatable", e.Value.ToString())
}

func (e *UnequatableType) ParserRule() antlr.ParserRuleContext {
	return e.BinOp
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
	return fmt.Sprintf("struct '%s' is not initialized", e.Ident.GetText())
}

func (e *StructNotInitialized) ParserRule() antlr.ParserRuleContext {
	return e.Ident
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
