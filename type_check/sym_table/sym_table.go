package sym_table

import (
	"chorego/parser"
	"chorego/type_check/type_error"
)

type SymTable struct {
	scope []*SymTableScope
}

type SymTableScope struct {
	symbols map[string]Symbol
}

func New() *SymTable {
	return &SymTable{
		scope: []*SymTableScope{},
	}
}

func newSymTableScope() *SymTableScope {
	return &SymTableScope{
		symbols: map[string]Symbol{},
	}
}

func (sym *SymTable) EnterScope() {
	sym.scope = append(sym.scope, newSymTableScope())
}

func (sym *SymTable) ExitScope() {
	sym.scope = sym.scope[:len(sym.scope)-1]
}

func (sym *SymTable) Scope() *SymTableScope {
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
	existing, exists := sym.Scope().symbols[symbol.SymbolName()]
	if exists {
		return type_error.NewSymbolAlreadyExistsError(existing.Ident(), symbol.Ident())
	}

	sym.Scope().symbols[symbol.SymbolName()] = symbol
	return nil
}

func (scope *SymTableScope) Get(name string) Symbol {
	return scope.symbols[name]
}

type Symbol interface {
	SymbolName() string
	Ident() parser.IIdentContext
}

type FuncParamSymbol struct {
	Param parser.IFuncParamContext
}

func NewFuncParamSymbol(param parser.IFuncParamContext) *FuncParamSymbol {
	return &FuncParamSymbol{
		Param: param,
	}
}

func (param *FuncParamSymbol) SymbolName() string {
	return param.Param.Ident().GetText()
}

func (param *FuncParamSymbol) Ident() parser.IIdentContext {
	return param.Param.Ident()
}
