package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"

	"github.com/tempo-lang/tempo/parser/new_parser/ast"
)

type CallNonFunction struct {
	baseError
	callExpr *ast.CallExpr
	symType  types.Type
}

func NewCallNonFunction(callExpr *ast.CallExpr, symType types.Type) Error {
	return &CallNonFunction{
		callExpr: callExpr,
		symType:  symType,
	}
}

func (e *CallNonFunction) Error() string {
	return fmt.Sprintf("cannot call value of type `%s`, since it is not a function", e.symType.ToString())
}

func (e *CallNonFunction) ParserRule() ast.Node {
	return e.callExpr
}

func (e *CallNonFunction) Code() ErrorCode {
	return CodeCallNonFunction
}

type CallWrongArgCount struct {
	baseError
	callExpr *ast.CallExpr
	expected int
	actual   int
}

func NewCallWrongArgCount(callExpr *ast.CallExpr, expected, actual int) Error {
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

func (e *CallWrongArgCount) ParserRule() ast.Node {
	return &ast.TokenRange{First: e.callExpr.OpenParen, Last: e.callExpr.CloseParen}
}

func (e *CallWrongArgCount) Code() ErrorCode {
	return CodeCallWrongArgCount
}

type InstantiateNonFunction struct {
	baseError
	identAccess *ast.IdentAccessExpr
	sym         sym_table.Symbol
}

func NewInstantiateNonFunction(identAccess *ast.IdentAccessExpr, sym sym_table.Symbol) Error {
	return &InstantiateNonFunction{
		identAccess: identAccess,
		sym:         sym,
	}
}

func (e *InstantiateNonFunction) Error() string {
	return fmt.Sprintf("cannot instantiate roles of `%s` with type `%s`, since it is not a function", e.sym.SymbolName(), e.sym.Type().ToString())
}

func (e *InstantiateNonFunction) ParserRule() ast.Node {
	return e.identAccess.RoleType
}

func (e *InstantiateNonFunction) Code() ErrorCode {
	return CodeInstantiateNonFunction
}

type FunctionNotInstantiated struct {
	baseError
	identAccess *ast.IdentAccessExpr
	sym         sym_table.Symbol
}

func NewFunctionNotInstantiated(identAccess *ast.IdentAccessExpr, sym sym_table.Symbol) Error {
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
		Type:    AnnotationTypeHint,
		Message: fmt.Sprintf("add roles after the name of the function, like %s@(A,B,C)", e.identAccess.Ident.Value()),
	}}
}

func (e *FunctionNotInstantiated) ParserRule() ast.Node {
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

func (e *FunctionMissingReturn) ParserRule() ast.Node {
	return e.callableEnv.ReturnCtx()
}

func (e *FunctionMissingReturn) Code() ErrorCode {
	return CodeFunctionMissingReturn
}

type ReturnValueMissing struct {
	baseError
	callableEnv sym_table.CallableEnv
	returnCtx   *ast.ReturnStmt
}

func NewReturnValueMissing(callableEnv sym_table.CallableEnv, returnCtx *ast.ReturnStmt) Error {
	return &ReturnValueMissing{
		callableEnv: callableEnv,
		returnCtx:   returnCtx,
	}
}

func (e *ReturnValueMissing) Error() string {
	return fmt.Sprintf("return is missing value of type `%s`", e.callableEnv.ReturnType().ToString())
}

func (e *ReturnValueMissing) ParserRule() ast.Node {
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

type IncompleteFunction struct {
	baseError
	FnIdent *ast.Identifier
	FnType  types.Type
}

func NewIncompleteFunction(fn *ast.Identifier, fnType types.Type) Error {
	return &IncompleteFunction{
		FnIdent: fn,
		FnType:  fnType,
	}
}

func (e *IncompleteFunction) Error() string {
	return fmt.Sprintf("function of type `%s` is incomplete", e.FnType.ToString())
}

func (e *IncompleteFunction) ParserRule() ast.Node {
	return e.FnIdent
}

func (e *IncompleteFunction) Code() ErrorCode {
	return CodeIncompleteFunction
}

func (e *IncompleteFunction) Annotations() []Annotation {
	return []Annotation{{
		Type:    AnnotationTypeNote,
		Message: "the function is incomplete because some of its roles are hidden `_`",
	}}
}
