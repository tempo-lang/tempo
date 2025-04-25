package type_check

import (
	"chorego/parser"
	"chorego/type_check/sym_table"
	"chorego/type_check/type_error"
)

func addGlobalSymbols(scope *sym_table.Scope, sourceFile *parser.SourceFileContext) (typeErrors []type_error.Error) {
	typeErrors = []type_error.Error{}

	for _, fn := range sourceFile.AllFunc_() {
		err := scope.InsertSymbol(sym_table.NewFuncSymbol(fn))
		if err != nil {
			typeErrors = append(typeErrors, err)
		}
	}

	return
}
