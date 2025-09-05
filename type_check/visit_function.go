package type_check

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

func (tc *typeChecker) VisitFunc(ctx *parser.FuncContext) any {
	// functions are already resolved by addGlobalSymbols
	sym, found := tc.info.Symbols[ctx.FuncSig().GetName()].(*sym_table.FuncSymbol)
	if !found {
		// was not found if function has parser errors
		return nil
	}

	tc.currentScope = sym.Scope()

	ctx.FuncSig().Accept(tc)

	// nil if parser error
	if ctx.Scope() != nil {
		returnsValue := ctx.Scope().Accept(tc) == true
		if !returnsValue && sym.FuncType().ReturnType() != types.Unit() {
			tc.reportError(type_error.NewFunctionMissingReturn(sym))
		}
	}

	tc.currentScope = tc.currentScope.Parent()
	return nil
}

func (tc *typeChecker) VisitFuncSig(ctx *parser.FuncSigContext) any {

	ctx.FuncParamList().Accept(tc)

	return nil
}

func (tc *typeChecker) VisitFuncParam(ctx *parser.FuncParamContext) any {
	return nil
}

func (tc *typeChecker) VisitFuncParamList(ctx *parser.FuncParamListContext) any {
	sym := tc.currentScope.GetCallableEnv()

	for i, paramCtx := range ctx.AllFuncParam() {
		paramType := sym.CallableType().Params()[i]
		paramSym := sym_table.NewFuncParamSymbol(paramCtx, tc.currentScope, paramType)
		tc.insertSymbol(paramSym)

		funcSym := tc.currentScope.GetCallableEnv()
		funcSym.AddParam(paramSym.(*sym_table.FuncParamSymbol))
	}

	return nil
}

// addFuncSymbol creates a new function symbol and stores it in the current scope.
// It returns the new symbol along with a success boolean.
func (tc *typeChecker) addFuncSymbol(fn parser.IFuncSigContext, scopeRange antlr.ParserRuleContext) (funcSym sym_table.Symbol, success bool) {
	funcRoles, ok := tc.parseRoleType(fn.RoleType())
	if !ok {
		return nil, false
	}

	funcScope := tc.currentScope.MakeChild(scopeRange.GetStart(), scopeRange.GetStop(), funcRoles.Participants())
	tc.currentScope = funcScope

	fnType, ok := tc.parseFuncType(fn)
	if !ok {
		return nil, false
	}

	// return type
	if fn.GetReturnType() != nil {
		fn.GetReturnType().Accept(tc)
		if !tc.checkRolesInScope(findRoleType(fn.GetReturnType())) {
			if fnValue, ok := fnType.(*types.FunctionType); ok {
				// make return type invalid
				fnType = types.Function(fnValue.NameIdent(), fnValue.Params(), types.Invalid(), fnValue.Roles())
			}
		}
	}

	tc.currentScope = tc.currentScope.Parent()

	funcSym = sym_table.NewFuncSymbol(fn, funcScope, fnType)
	funcScope.SetCallableEnv(funcSym.(*sym_table.FuncSymbol))
	tc.insertSymbol(funcSym)

	return funcSym, true
}
