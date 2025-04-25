package sym_table

import (
	"chorego/type_check/type_error"
)

type SymTable struct {
	scope []*Scope
}

type Scope struct {
	symbols map[string]Symbol
}

func New() *SymTable {
	return &SymTable{
		scope: []*Scope{},
	}
}

func newScope() *Scope {
	return &Scope{
		symbols: map[string]Symbol{},
	}
}

func (sym *SymTable) EnterScope() *Scope {
	sym.scope = append(sym.scope, newScope())
	return sym.Scope()
}

func (sym *SymTable) ExitScope() {
	sym.scope = sym.scope[:len(sym.scope)-1]
}

func (sym *SymTable) Scope() *Scope {
	return sym.scope[len(sym.scope)-1]
}

func (sym *SymTable) LookupSymbol(name string) Symbol {
	for i := len(sym.scope) - 1; i >= 0; i-- {
		symbol := sym.scope[i].Get(name)
		if symbol != nil {
			return symbol
		}
	}
	return nil
}

func (sym *SymTable) InsertSymbol(symbol Symbol) *type_error.SymbolAlreadyExists {
	return sym.Scope().InsertSymbol(symbol)
}

func (scope *Scope) Get(name string) Symbol {
	return scope.symbols[name]
}

func (scope *Scope) InsertSymbol(symbol Symbol) *type_error.SymbolAlreadyExists {
	existing, exists := scope.symbols[symbol.SymbolName()]
	if exists {
		return type_error.NewSymbolAlreadyExistsError(existing.Ident(), symbol.Ident())
	}

	scope.symbols[symbol.SymbolName()] = symbol
	return nil
}
