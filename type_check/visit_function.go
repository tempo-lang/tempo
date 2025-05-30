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

	if sym.Type().Roles().IsSharedRole() {
		tc.reportError(type_error.NewUnexpectedSharedTypeError(ctx.FuncSig().GetName(), sym.Type()))
	}

	ctx.FuncSig().Accept(tc)

	// nil if parser error
	if ctx.Scope() != nil {
		returnsValue := ctx.Scope().Accept(tc) == true
		if !returnsValue && sym.FuncValue().ReturnType().Value() != types.Unit() {
			tc.reportError(type_error.NewFunctionMissingReturn(sym))
		}
	}

	tc.currentScope = tc.currentScope.Parent()
	return nil
}

func (tc *typeChecker) VisitFuncSig(ctx *parser.FuncSigContext) any {

	if sym, ok := tc.info.Symbols[ctx.GetName()].(*sym_table.FuncSymbol); ok {
		if err := tc.checkDuplicateRoles(ctx.RoleType(), sym.Type().Roles()); err != nil {
			tc.reportError(err)
		}
	}

	ctx.FuncParamList().Accept(tc)

	return nil
}

func (tc *typeChecker) VisitFuncParam(ctx *parser.FuncParamContext) any {
	paramType := tc.visitValueType(ctx.ValueType())
	paramSym := sym_table.NewFuncParamSymbol(ctx, tc.currentScope, paramType)
	tc.insertSymbol(paramSym)

	funcSym := tc.currentScope.GetFunc()
	funcSym.AddParam(paramSym.(*sym_table.FuncParamSymbol))

	return nil
}

func (tc *typeChecker) VisitFuncParamList(ctx *parser.FuncParamListContext) any {
	for _, param := range ctx.AllFuncParam() {
		param.Accept(tc)
	}
	return nil
}

func (tc *typeChecker) addFuncSymbol(fn parser.IFuncSigContext, scopeRange antlr.ParserRuleContext) (sym_table.Symbol, bool) {
	fnType, errors := tc.parseFuncType(fn)
	for _, err := range errors {
		tc.reportError(err)
	}

	if len(errors) > 0 {
		return nil, false
	}

	if fnType.IsInvalid() {
		return nil, false
	}

	funcScope := tc.currentScope.MakeChild(scopeRange.GetStart(), scopeRange.GetStop(), fnType.Roles().Participants())
	tc.currentScope = funcScope

	// return type
	if fn.GetReturnType() != nil {
		fn.GetReturnType().Accept(tc)
		if !tc.checkRolesInScope(fn.GetReturnType().RoleType()) {
			if fnValue, ok := fnType.Value().(*types.FunctionType); ok {
				// make return type invalid
				fnType = types.New(
					types.Function(fnValue.FuncIdent(), fnValue.Params(), types.Invalid(), fnType.Roles().Participants()),
					fnType.Roles(),
				)
			}
		}
	}

	tc.currentScope = tc.currentScope.Parent()

	funcSym := sym_table.NewFuncSymbol(fn, funcScope, fnType)
	funcScope.SetFunc(funcSym.(*sym_table.FuncSymbol))
	tc.insertSymbol(funcSym)

	return funcSym, true
}
