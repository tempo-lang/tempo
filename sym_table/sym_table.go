package sym_table

import (
	"tempo/parser"
	"tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type Scope struct {
	symbols  map[string]Symbol
	parent   *Scope
	children []*Scope
	pos      antlr.Token
	end      antlr.Token
	roles    []string
	funcSym  *FuncSymbol
}

func NewScope(pos antlr.Token, end antlr.Token, parent *Scope, roles []string) *Scope {
	return &Scope{
		symbols:  map[string]Symbol{},
		parent:   parent,
		children: []*Scope{},
		pos:      pos,
		end:      end,
		roles:    roles,
		funcSym:  nil,
	}
}

func (scope *Scope) LookupParent(name string) Symbol {
	if sym, found := scope.symbols[name]; found {
		return sym
	}

	if scope.parent != nil {
		return scope.parent.LookupParent(name)
	}
	return nil
}

func (scope *Scope) Lookup(name string) Symbol {
	return scope.symbols[name]
}

func (scope *Scope) LookupSymbol(name parser.IIdentContext) (Symbol, types.Error) {
	symbol := scope.LookupParent(name.GetText())
	if symbol != nil {
		return symbol, nil
	}
	return nil, types.NewUnknownSymbolError(name)
}

func (scope *Scope) InsertSymbol(symbol Symbol) types.Error {
	existing, exists := scope.symbols[symbol.SymbolName()]
	if exists {
		return types.NewSymbolAlreadyExistsError(existing.Ident(), symbol.Ident())
	}

	scope.symbols[symbol.SymbolName()] = symbol
	return nil
}

func (scope *Scope) HasParent() bool {
	return scope.parent != nil
}

func (scope *Scope) Parent() *Scope {
	return scope.parent
}

func (scope *Scope) Children() []*Scope {
	return scope.children
}

func (scope *Scope) Roles() []string {
	return scope.roles
}

func (scope *Scope) Pos() antlr.Token {
	return scope.pos
}

func (scope *Scope) End() antlr.Token {
	return scope.end
}

func (scope *Scope) Contains(pos antlr.Token) bool {
	return scope.pos.GetTokenIndex() >= pos.GetTokenIndex() && scope.end.GetTokenIndex() <= pos.GetTokenIndex()
}

func (scope *Scope) Innermost(pos antlr.Token) *Scope {
	for _, child := range scope.children {
		if child.Contains(pos) {
			return child.Innermost(pos)
		}
	}
	return nil
}

func (scope *Scope) MakeChild(pos antlr.Token, end antlr.Token, roles []string) *Scope {
	child := NewScope(pos, end, scope, roles)
	scope.children = append(scope.children, child)
	return child
}

func (scope *Scope) SetFunc(funcSym *FuncSymbol) {
	scope.funcSym = funcSym
}

func (scope *Scope) GetFunc() *FuncSymbol {
	s := scope
	for s.funcSym == nil && s.parent != nil {
		s = s.parent
	}
	return s.funcSym
}
