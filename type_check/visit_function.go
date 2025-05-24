package type_check

import (
	"tempo/parser"
	"tempo/sym_table"
	"tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

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
					types.Function(fnValue.Params(), types.Invalid()),
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
