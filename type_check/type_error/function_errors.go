package type_error

import (
	"fmt"
	"tempo/parser"
	"tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type CallNonFunctionError struct {
	callExpr *parser.ExprCallContext
	symType  *types.Type
}

func NewCallNonFunctionError(callExpr *parser.ExprCallContext, symType *types.Type) Error {
	return &CallNonFunctionError{
		callExpr: callExpr,
		symType:  symType,
	}
}

func (e *CallNonFunctionError) Error() string {
	return fmt.Sprintf("call to non-function '%s'", e.symType.ToString())
}

func (e *CallNonFunctionError) IsTypeError() {}

func (e *CallNonFunctionError) ParserRule() antlr.ParserRuleContext {
	return e.callExpr
}

type CallWrongArgCountError struct {
	callExpr *parser.ExprCallContext
}

func NewCallWrongArgCountError(callExpr *parser.ExprCallContext) Error {
	return &CallWrongArgCountError{
		callExpr: callExpr,
	}
}

func (e *CallWrongArgCountError) Error() string {
	return "call has wrong number of arguments"
}

func (e *CallWrongArgCountError) IsTypeError() {}

func (e *CallWrongArgCountError) ParserRule() antlr.ParserRuleContext {
	return e.callExpr
}

type CallWrongRoleCountError struct {
	callExpr *parser.ExprCallContext
}

func NewCallWrongRoleCountError(callExpr *parser.ExprCallContext) Error {
	return &CallWrongRoleCountError{
		callExpr: callExpr,
	}
}

func (e *CallWrongRoleCountError) Error() string {
	return "call has wrong number of roles"
}

func (e *CallWrongRoleCountError) IsTypeError() {}

func (e *CallWrongRoleCountError) ParserRule() antlr.ParserRuleContext {
	return e.callExpr
}
