package sym_table

import (
	"iter"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type Scope struct {
	symbols      map[string]Symbol
	parent       *Scope
	children     []*Scope
	pos          antlr.Token
	end          antlr.Token
	roles        []string
	funcSym      *FuncSymbol
	structSym    *StructSymbol
	interfaceSym *InterfaceSymbol
}

func NewScope(pos antlr.Token, end antlr.Token, parent *Scope, roles []string) *Scope {
	return &Scope{
		symbols:      map[string]Symbol{},
		parent:       parent,
		children:     []*Scope{},
		pos:          pos,
		end:          end,
		roles:        roles,
		funcSym:      nil,
		structSym:    nil,
		interfaceSym: nil,
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

func (scope *Scope) LookupSymbol(name parser.IIdentContext) (Symbol, bool) {
	symbol := scope.LookupParent(name.GetText())
	if symbol != nil {
		return symbol, true
	}
	return nil, false
}

func (scope *Scope) InsertSymbol(symbol Symbol) bool {
	_, exists := scope.symbols[symbol.SymbolName()]
	if exists {
		return false
	}

	scope.symbols[symbol.SymbolName()] = symbol
	return true
}

func (scope *Scope) HasParent() bool {
	return scope.parent != nil
}

func (scope *Scope) Parent() *Scope {
	return scope.parent
}

func (scope *Scope) Global() *Scope {
	if scope.Parent() != nil {
		return scope.Parent().Global()
	}
	return scope
}

func (scope *Scope) Children() []*Scope {
	return scope.children
}

func (scope *Scope) Symbols() iter.Seq[Symbol] {
	return func(yield func(Symbol) bool) {
		for _, sym := range scope.symbols {
			if !yield(sym) {
				return
			}
		}
	}
}

func (scope *Scope) Roles() *types.Roles {
	return types.NewRole(scope.roles, false)
}

func (scope *Scope) Pos() antlr.Token {
	return scope.pos
}

func (scope *Scope) End() antlr.Token {
	return scope.end
}

func (scope *Scope) Contains(pos antlr.Token) bool {
	return scope.pos.GetTokenIndex() <= pos.GetTokenIndex() && pos.GetTokenIndex() <= scope.end.GetTokenIndex()
}

func (scope *Scope) Innermost(pos antlr.Token) *Scope {
	for _, child := range scope.children {
		if child.Contains(pos) {
			inner := child.Innermost(pos)
			if inner != nil {
				return inner
			} else {
				return child
			}
		}
	}
	return scope
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

func (scope *Scope) SetStruct(structSym *StructSymbol) {
	scope.structSym = structSym
}

func (scope *Scope) GetStruct() *StructSymbol {
	s := scope
	for s.structSym == nil && s.parent != nil {
		s = s.parent
	}
	return s.structSym
}

func (scope *Scope) SetInterface(interfaceSym *InterfaceSymbol) {
	scope.interfaceSym = interfaceSym
}

func (scope *Scope) GetInterface() *InterfaceSymbol {
	s := scope
	for s.interfaceSym == nil && s.parent != nil {
		s = s.parent
	}
	return s.interfaceSym
}
