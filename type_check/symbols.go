package type_check

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check/type_error"
)

func (tc *typeChecker) lookupSymbol(name parser.IIdentContext) (sym_table.Symbol, type_error.Error) {
	sym, ok := tc.currentScope.LookupSymbol(name)
	if !ok {
		return nil, type_error.NewUnknownSymbol(name)
	}
	return sym, nil
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
