package type_check

import (
	"chorego/parser"
	"chorego/sym_table"
	"chorego/types"
)

func addGlobalSymbols(scope *sym_table.Scope, sourceFile *parser.SourceFileContext) (typeErrors []types.Error) {
	typeErrors = []types.Error{}

	for _, fn := range sourceFile.AllFunc_() {
		fnType, err := types.ParseFuncType(fn)
		if err != nil {
			typeErrors = append(typeErrors, err)
		}

		err = scope.InsertSymbol(sym_table.NewFuncSymbol(fn, fnType))
		if err != nil {
			typeErrors = append(typeErrors, err)
		}
	}

	return
}
