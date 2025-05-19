package type_check

import (
	"tempo/parser"
	"tempo/sym_table"
	"tempo/types"
)

func (tc *typeChecker) addGlobalSymbols(sourceFile *parser.SourceFileContext) {
	for _, fn := range sourceFile.AllFunc_() {
		fnType, err := types.ParseFuncType(fn)
		if err != nil {
			tc.reportError(err)
			continue
		}

		if fnType.IsInvalid() {
			continue
		}

		funcScope := tc.currentScope.MakeChild(fn.GetStart(), fn.GetStop(), fnType.Roles().Participants())
		tc.currentScope = funcScope

		// return type
		if fn.ValueType() != nil {
			fn.ValueType().Accept(tc)
			if !tc.checkRolesInScope(fn.ValueType().RoleType()) {
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
	}
}
