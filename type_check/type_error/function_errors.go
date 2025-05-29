package type_error

import (
	"fmt"
	"tempo/parser"
	"tempo/sym_table"
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

type InstantiateNonFunction struct {
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

func (e *InstantiateNonFunction) IsTypeError() {}

func (e *InstantiateNonFunction) ParserRule() antlr.ParserRuleContext {
	return e.identAccess.RoleType()
}

type FunctionNotInstantiated struct {
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

func (e *FunctionNotInstantiated) IsTypeError() {}

func (e *FunctionNotInstantiated) ParserRule() antlr.ParserRuleContext {
	return e.identAccess
}

type FunctionMissingReturn struct {
	funcSym *sym_table.FuncSymbol
}

func NewFunctionMissingReturn(funcSym *sym_table.FuncSymbol) Error {
	return &FunctionMissingReturn{
		funcSym: funcSym,
	}
}

func (e *FunctionMissingReturn) Error() string {
	return fmt.Sprintf("function '%s' is missing return statement", e.funcSym.SymbolName())
}

func (e *FunctionMissingReturn) IsTypeError() {}

func (e *FunctionMissingReturn) ParserRule() antlr.ParserRuleContext {
	return e.funcSym.FuncSig().GetReturnType()
}

type ReturnValueMissing struct {
	funcSym   *sym_table.FuncSymbol
	returnCtx *parser.StmtReturnContext
}

func NewReturnValueMissing(funcSym *sym_table.FuncSymbol, returnCtx *parser.StmtReturnContext) Error {
	return &ReturnValueMissing{
		funcSym:   funcSym,
		returnCtx: returnCtx,
	}
}

func (e *ReturnValueMissing) Error() string {
	return fmt.Sprintf("return is missing value of type '%s'", e.funcSym.FuncValue().ReturnType().ToString())
}

func (e *ReturnValueMissing) IsTypeError() {}

func (e *ReturnValueMissing) ParserRule() antlr.ParserRuleContext {
	return e.returnCtx
}
