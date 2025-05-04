package type_check

import (
	"chorego/parser"
	"chorego/sym_table"
	"chorego/types"
)

func (tc *typeChecker) addGlobalSymbols(sourceFile *parser.SourceFileContext) {
	for _, fn := range sourceFile.AllFunc_() {
		fnType, err := types.ParseFuncType(fn)
		if err != nil {
			tc.reportError(err)
		}

		funcScope := tc.currentScope.MakeChild(fn.GetStart(), fn.GetStop(), fnType.Roles().Participants())
		tc.insertSymbol(sym_table.NewFuncSymbol(fn, funcScope, fnType))
	}

	return
}
