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
	symType  *types.Type
}

func NewCallNonFunction(callExpr *parser.ExprCallContext, symType *types.Type) Error {
	return &CallNonFunction{
		callExpr: callExpr,
		symType:  symType,
	}
}

func (e *CallNonFunction) Error() string {
	return fmt.Sprintf("call to non-function '%s'", e.symType.ToString())
}

func (e *CallNonFunction) ParserRule() antlr.ParserRuleContext {
	return e.callExpr
}

type CallWrongArgCount struct {
	baseError
	callExpr *parser.ExprCallContext
}

func NewCallWrongArgCount(callExpr *parser.ExprCallContext) Error {
	return &CallWrongArgCount{
		callExpr: callExpr,
	}
}

func (e *CallWrongArgCount) Error() string {
	return "call has wrong number of arguments"
}

func (e *CallWrongArgCount) ParserRule() antlr.ParserRuleContext {
	return e.callExpr
}

type CallWrongRoleCount struct {
	baseError
	callExpr *parser.ExprCallContext
}

func NewCallWrongRoleCount(callExpr *parser.ExprCallContext) Error {
	return &CallWrongRoleCount{
		callExpr: callExpr,
	}
}

func (e *CallWrongRoleCount) Error() string {
	return "call has wrong number of roles"
}

func (e *CallWrongRoleCount) ParserRule() antlr.ParserRuleContext {
	return e.callExpr
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
	return fmt.Sprintf("can not instantiate roles of '%s' with type '%s', since it is not a function", e.sym.SymbolName(), e.sym.Type().ToString())
}

func (e *InstantiateNonFunction) ParserRule() antlr.ParserRuleContext {
	return e.identAccess.RoleType()
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

func (e *FunctionNotInstantiated) ParserRule() antlr.ParserRuleContext {
	return e.identAccess
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
	return fmt.Sprintf("return is missing value of type '%s'", e.callableEnv.ReturnType().ToString())
}

func (e *ReturnValueMissing) ParserRule() antlr.ParserRuleContext {
	return e.returnCtx
}
