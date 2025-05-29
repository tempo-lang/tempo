package type_error

import (
	"fmt"
	"tempo/parser"
	"tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type InvalidNumberError struct {
	Num parser.IExprContext
}

func NewInvalidNumberError(num parser.IExprContext) Error {
	return &InvalidNumberError{
		Num: num,
	}
}

func (i *InvalidNumberError) Error() string {
	return fmt.Sprintf("invalid number '%s'", i.Num.GetText())
}

func (i *InvalidNumberError) IsTypeError() {}

func (i *InvalidNumberError) ParserRule() antlr.ParserRuleContext {
	return i.Num
}

type ComNonLocalSender struct {
	Com *parser.ExprComContext
}

type ComNonLocalSenderError struct {
	Com *parser.ExprComContext
}

func (e *ComNonLocalSenderError) Error() string {
	return "only a single sender is allowed"
}

func (e *ComNonLocalSenderError) IsTypeError() {}

func (e *ComNonLocalSenderError) ParserRule() antlr.ParserRuleContext {
	return e.Com.GetSender()
}

func NewComNonLocalSenderError(com *parser.ExprComContext) Error {
	return &ComNonLocalSenderError{
		Com: com,
	}
}

type ComValueNotAtSenderError struct {
	Com      *parser.ExprComContext
	ExprType *types.Type
}

func (c *ComValueNotAtSenderError) Error() string {
	sender := parser.RoleTypeAllIdents(c.Com.RoleType(0))[0]
	return fmt.Sprintf("value of type '%s' is not present at sender '%s'", c.ExprType.ToString(), sender.GetText())
}

func (c *ComValueNotAtSenderError) IsTypeError() {}

func (c *ComValueNotAtSenderError) ParserRule() antlr.ParserRuleContext {
	return c.Com.Expr()
}

func NewComValueNotAtSenderError(com *parser.ExprComContext, exprType *types.Type) Error {
	return &ComValueNotAtSenderError{
		Com:      com,
		ExprType: exprType,
	}
}

type DivisionByZeroError struct {
	Expr parser.IExprContext
}

func (e *DivisionByZeroError) Error() string {
	return "invalid operation, division by zero"
}

func (e *DivisionByZeroError) IsTypeError() {}

func (e *DivisionByZeroError) ParserRule() antlr.ParserRuleContext {
	return e.Expr
}

func NewDivisionByZeroError(expr parser.IExprContext) Error {
	return &DivisionByZeroError{
		Expr: expr,
	}
}

type UnequatableTypeError struct {
	BinOp *parser.ExprBinOpContext
	Value types.Value
}

func (e *UnequatableTypeError) Error() string {
	return fmt.Sprintf("type value '%s' is not equatable", e.Value.ToString())
}

func (e *UnequatableTypeError) IsTypeError() {}

func (e *UnequatableTypeError) ParserRule() antlr.ParserRuleContext {
	return e.BinOp
}

func NewUnequatableTypeError(binOp *parser.ExprBinOpContext, value types.Value) Error {
	return &UnequatableTypeError{
		BinOp: binOp,
		Value: value,
	}
}

type StructNotInitialized struct {
	Ident *parser.ExprIdentContext
}

func (e *StructNotInitialized) Error() string {
	return fmt.Sprintf("struct '%s' is not initialized", e.Ident.GetText())
}

func (e *StructNotInitialized) IsTypeError() {}

func (e *StructNotInitialized) ParserRule() antlr.ParserRuleContext {
	return e.Ident
}

func NewStructNotInitialized(ident *parser.ExprIdentContext) Error {
	return &StructNotInitialized{
		Ident: ident,
	}
}
