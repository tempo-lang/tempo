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

type ComSharedTypeError struct {
	Com       *parser.ExprComContext
	InnerType *types.Type
}

func (e *ComSharedTypeError) Error() string {
	return fmt.Sprintf("can not communicate shared type '%s'", e.InnerType.ToString())
}

func (e *ComSharedTypeError) IsTypeError() {}

func (e *ComSharedTypeError) ParserRule() antlr.ParserRuleContext {
	return e.Com
}

func NewComSharedTypeError(com *parser.ExprComContext, innerType *types.Type) Error {
	return &ComSharedTypeError{
		Com:       com,
		InnerType: innerType,
	}
}

type ComDistributedTypeError struct {
	Com       *parser.ExprComContext
	InnerType *types.Type
}

func (e *ComDistributedTypeError) Error() string {
	return fmt.Sprintf("can not communicate distributed type '%s'", e.InnerType.ToString())
}

func (e *ComDistributedTypeError) IsTypeError() {}

func (e *ComDistributedTypeError) ParserRule() antlr.ParserRuleContext {
	return e.Com
}

func NewComDistributedTypeError(com *parser.ExprComContext, innerType *types.Type) Error {
	return &ComDistributedTypeError{
		Com:       com,
		InnerType: innerType,
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
