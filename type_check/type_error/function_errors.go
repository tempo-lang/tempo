package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type CallNonFunction struct {
	baseError
	callExpr *parser.ExprCallContext
	symType  types.Type
}

func NewCallNonFunction(callExpr *parser.ExprCallContext, symType types.Type) Error {
	return &CallNonFunction{
		callExpr: callExpr,
		symType:  symType,
	}
}

func (e *CallNonFunction) Error() string {
	return fmt.Sprintf("cannot call value of type `%s`, since it is not a function", e.symType.ToString())
}

func (e *CallNonFunction) ParserRule() antlr.ParserRuleContext {
	return e.callExpr
}

func (e *CallNonFunction) Code() ErrorCode {
	return CodeCallNonFunction
}

type CallWrongArgCount struct {
	baseError
	callExpr *parser.ExprCallContext
	expected int
	actual   int
}

func NewCallWrongArgCount(callExpr *parser.ExprCallContext, expected, actual int) Error {
	return &CallWrongArgCount{
		callExpr: callExpr,
		expected: expected,
		actual:   actual,
	}
}

func (e *CallWrongArgCount) Error() string {
	expectedArgs := fmt.Sprintf("%d arguments", e.expected)
	if e.expected == 1 {
		expectedArgs = fmt.Sprintf("%d argument", e.expected)
	}

	actualArgs := fmt.Sprintf("%d", e.actual)
	if e.actual < e.expected {
		actualArgs = fmt.Sprintf("only %d", e.actual)
	}

	return fmt.Sprintf("function expected %s, but %s was given", expectedArgs, actualArgs)
}

func (e *CallWrongArgCount) ParserRule() antlr.ParserRuleContext {
	return e.callExpr.FuncArgList()
}

func (e *CallWrongArgCount) Code() ErrorCode {
	return CodeCallWrongArgCount
}

type InstantiateNonFunction struct {
	baseError
	identAccess parser.IIdentAccessContext
	sym         sym_table.Symbol
}

func NewInstantiateNonFunction(identAccess parser.IIdentAccessContext, sym sym_table.Symbol) Error {
	return &InstantiateNonFunction{
		identAccess: identAccess,
		sym:         sym,
	}
}

func (e *InstantiateNonFunction) Error() string {
	return fmt.Sprintf("cannot instantiate roles of `%s` with type `%s`, since it is not a function", e.sym.SymbolName(), e.sym.Type().ToString())
}

func (e *InstantiateNonFunction) ParserRule() antlr.ParserRuleContext {
	return e.identAccess.RoleType()
}

func (e *InstantiateNonFunction) Code() ErrorCode {
	return CodeInstantiateNonFunction
}

type FunctionNotInstantiated struct {
	baseError
	identAccess parser.IIdentAccessContext
	sym         sym_table.Symbol
}

func NewFunctionNotInstantiated(identAccess parser.IIdentAccessContext, sym sym_table.Symbol) Error {
	return &FunctionNotInstantiated{
		identAccess: identAccess,
		sym:         sym,
	}
}

func (e *FunctionNotInstantiated) Error() string {
	return "roles of function must be instantiated"
}

func (e *FunctionNotInstantiated) Annotations() []Annotation {
	return []Annotation{{
		Type:    "hint",
		Message: fmt.Sprintf("add roles after the name of the function, like %s@(A,B,C)", e.identAccess.GetText()),
	}}
}

func (e *FunctionNotInstantiated) ParserRule() antlr.ParserRuleContext {
	return e.identAccess
}

func (e *FunctionNotInstantiated) Code() ErrorCode {
	return CodeFunctionNotInstantiated
}

type FunctionMissingReturn struct {
	baseError
	callableEnv sym_table.CallableEnv
}

func NewFunctionMissingReturn(callableEnv sym_table.CallableEnv) Error {
	return &FunctionMissingReturn{
		callableEnv: callableEnv,
	}
}

func (e *FunctionMissingReturn) Error() string {
	return "missing return statement"
}

func (e *FunctionMissingReturn) ParserRule() antlr.ParserRuleContext {
	return e.callableEnv.ReturnCtx()
}

func (e *FunctionMissingReturn) Code() ErrorCode {
	return CodeFunctionMissingReturn
}

type ReturnValueMissing struct {
	baseError
	callableEnv sym_table.CallableEnv
	returnCtx   *parser.StmtReturnContext
}

func NewReturnValueMissing(callableEnv sym_table.CallableEnv, returnCtx *parser.StmtReturnContext) Error {
	return &ReturnValueMissing{
		callableEnv: callableEnv,
		returnCtx:   returnCtx,
	}
}

func (e *ReturnValueMissing) Error() string {
	return fmt.Sprintf("return is missing value of type `%s`", e.callableEnv.ReturnType().ToString())
}

func (e *ReturnValueMissing) ParserRule() antlr.ParserRuleContext {
	return e.returnCtx
}

func (e *ReturnValueMissing) Code() ErrorCode {
	return CodeReturnValueMissing
}

func (e *ReturnValueMissing) RelatedInfo() []RelatedInfo {
	return []RelatedInfo{{
		Message:    "return type is specified here",
		ParserRule: e.callableEnv.ReturnCtx(),
	}}
}
