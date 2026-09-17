package type_check

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

func (tc *typeChecker) visitFunc(fn *ast.Func) {
	sym, ok := tc.info.Symbols[fn.FuncSig.Name].(*sym_table.FuncSymbol)
	if !ok {
		return
	}
	tc.currentScope = sym.Scope()
	tc.visitFuncParams(fn.FuncSig.Params)
	if !tc.visitScope(fn.Scope) && sym.FuncType().ReturnType() != types.Unit() {
		tc.reportError(type_error.NewFunctionMissingReturn(sym))
	}
	tc.currentScope = tc.currentScope.Parent()
}

func (tc *typeChecker) visitFuncParams(params *ast.FuncParams) {
	env := tc.currentScope.GetCallableEnv()
	for i, param := range params.Params {
		paramSym := sym_table.NewFuncParamSymbol(param, tc.currentScope, env.CallableType().Params()[i])
		tc.insertSymbol(paramSym)
		env.AddParam(paramSym.(*sym_table.FuncParamSymbol))
	}
}

func (tc *typeChecker) addFuncSymbol(fn *ast.FuncSig, scopeNode ast.Node) (sym_table.Symbol, bool) {
	roles, ok := tc.parseRoleType(fn.RoleType)
	if !ok {
		return nil, false
	}
	if !roles.IsComplete() {
		tc.reportError(type_error.NewUnexpectedHiddenRoles(fn.RoleType))
	}
	funcScope := tc.currentScope.MakeChild(nodeSpan(scopeNode), roles.Participants())
	tc.currentScope = funcScope
	fnType, ok := tc.parseFuncType(fn)
	if !ok {
		params := make([]types.Type, len(fn.Params.Params))
		for i := range params {
			params[i] = types.Invalid()
		}
		fnType = types.Function(fn, params, types.Unit(), roles)
	}
	if fn.ReturnType != nil {
		if rt, found := findRoleType(fn.ReturnType); found && !tc.checkRolesInScope(rt) {
			f := fnType.(*types.FunctionType)
			fnType = types.Function(f.FuncSig(), f.Params(), types.Invalid(), f.Roles())
		}
	}
	tc.currentScope = tc.currentScope.Parent()
	sym := sym_table.NewFuncSymbol(fn, funcScope, fnType)
	funcScope.SetCallableEnv(sym.(*sym_table.FuncSymbol))
	tc.insertSymbol(sym)
	return sym, true
}
