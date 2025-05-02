package sym_table

import (
	"chorego/parser"
	"chorego/types"
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

func (sym *SymTable) IsGlobalScope() bool {
	return len(sym.scope) == 1
}

func (sym *SymTable) LookupSymbol(name parser.IIdentContext) (Symbol, types.Error) {
	for i := len(sym.scope) - 1; i >= 0; i-- {
		symbol := sym.scope[i].Get(name.GetText())
		if symbol != nil {
			return symbol, nil
		}
	}
	return nil, types.NewUnknownSymbolError(name)
}

func (sym *SymTable) InsertSymbol(symbol Symbol) types.Error {
	return sym.Scope().InsertSymbol(symbol)
}

func (scope *Scope) Get(name string) Symbol {
	return scope.symbols[name]
}

func (scope *Scope) InsertSymbol(symbol Symbol) types.Error {
	existing, exists := scope.symbols[symbol.SymbolName()]
	if exists {
		return types.NewSymbolAlreadyExistsError(existing.Ident(), symbol.Ident())
	}

	scope.symbols[symbol.SymbolName()] = symbol
	return nil
}
