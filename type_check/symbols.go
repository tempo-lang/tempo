package type_check

import (
	"tempo/parser"
	"tempo/sym_table"
	"tempo/type_check/type_error"
)

func (tc *typeChecker) lookupSymbol(name parser.IIdentContext) (sym_table.Symbol, bool) {
	sym, ok := tc.currentScope.LookupSymbol(name)
	if !ok {
		tc.reportError(type_error.NewUnknownSymbolError(name))
	}
	return sym, ok
}

func (tc *typeChecker) insertSymbol(symbol sym_table.Symbol) bool {
	tc.info.Symbols[symbol.Ident()] = symbol
	if !tc.currentScope.InsertSymbol(symbol) {
		existing := tc.currentScope.Lookup(symbol.SymbolName())
		tc.reportError(type_error.NewSymbolAlreadyExistsError(existing.Ident(), symbol.Ident()))
		return false
	}
	return true
}
