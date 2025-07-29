// The symbol table is used by the [tempo/type_checker] package in order to keep track of symbols.
package sym_table

import (
	"iter"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

// Scope contains all information for a particular scope in the symbol table.
// It contains a pointer to the parent scope, as well as pointers to all child scopes.
type Scope struct {
	symbols      map[string]Symbol
	parent       *Scope
	children     []*Scope
	pos          antlr.Token
	end          antlr.Token
	roles        []string
	callableEnv  CallableEnv
	structSym    *StructSymbol
	interfaceSym *InterfaceSymbol
}

// NewScope constructs a new child scope to `parent`.
//
// The parent scope is not automatically updated to include the new child scope.
// For that, use [Scope.MakeChild] instead.
func NewScope(pos antlr.Token, end antlr.Token, parent *Scope, roles []string) *Scope {
	return &Scope{
		symbols:      map[string]Symbol{},
		parent:       parent,
		children:     []*Scope{},
		pos:          pos,
		end:          end,
		roles:        roles,
		callableEnv:  nil,
		structSym:    nil,
		interfaceSym: nil,
	}
}

// LookupParent will look for the symbol in the current scope,
// and if it is not found then it will recursively look through all parent scopes, until the root is reached.
//
// If no symbol was found, it returns `nil`.
func (scope *Scope) LookupParent(name string) Symbol {
	if sym, found := scope.symbols[name]; found {
		return sym
	}

	if scope.parent != nil {
		return scope.parent.LookupParent(name)
	}
	return nil
}

// Lookup returns the symbol with the given `name` only looking in this scope.
//
// If no symbol was found, it returns `nil`.
func (scope *Scope) Lookup(name string) Symbol {
	return scope.symbols[name]
}

// LookupSymbol is a safe version of [Scope.LookupParent] which returns a boolean indicating whether the symbol was found or not.
func (scope *Scope) LookupSymbol(name parser.IIdentContext) (Symbol, bool) {
	symbol := scope.LookupParent(name.GetText())
	if symbol != nil {
		return symbol, true
	}
	return nil, false
}

// InsertSymbol inserts the given symbol into the current scope.
//
// If a symbol with the same name already exists,
// then the function will return `false`, indicating that the insertion failed.
func (scope *Scope) InsertSymbol(symbol Symbol) bool {
	_, exists := scope.symbols[symbol.SymbolName()]
	if exists {
		return false
	}

	scope.symbols[symbol.SymbolName()] = symbol
	return true
}

// HasParent returns true if [Scope.Parent] is not `nil`.
func (scope *Scope) HasParent() bool {
	return scope.parent != nil
}

// Parent returns the parent scope, or `nil` if the scope has no parent.
func (scope *Scope) Parent() *Scope {
	return scope.parent
}

// Global returns the global scope, the scope without a parent.
func (scope *Scope) Global() *Scope {
	if scope.Parent() != nil {
		return scope.Parent().Global()
	}
	return scope
}

// Children returns a slice of all children of this scope.
func (scope *Scope) Children() []*Scope {
	return scope.children
}

// Symbols returns an iterator of all symbols directly in this scope.
//
// It will not yield the symbols in the child scopes.
func (scope *Scope) Symbols() iter.Seq[Symbol] {
	return func(yield func(Symbol) bool) {
		for _, sym := range scope.symbols {
			if !yield(sym) {
				return
			}
		}
	}
}

// Roles returns the roles that participate in this scope.
func (scope *Scope) Roles() *types.Roles {
	return types.NewRole(scope.roles, false)
}

// Pos returns the AST token of the first token in the source code which is a part of this scope.
func (scope *Scope) Pos() antlr.Token {
	return scope.pos
}

// End returns the AST token of the last token in the source chode which is a part of this scope.
func (scope *Scope) End() antlr.Token {
	return scope.end
}

// Contains returns whether the provided AST token lies within this scope, as well as any child scopes.
func (scope *Scope) Contains(pos antlr.Token) bool {
	return scope.pos.GetTokenIndex() <= pos.GetTokenIndex() && pos.GetTokenIndex() <= scope.end.GetTokenIndex()
}

// Innermost finds the innermost scope containing the provided AST token.
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

// MakeChild creates a new child scope and adds it to the list of children of this scope.
// The newly created child scope is returned.
func (scope *Scope) MakeChild(pos antlr.Token, end antlr.Token, roles []string) *Scope {
	child := NewScope(pos, end, scope, roles)
	scope.children = append(scope.children, child)
	return child
}

// SetCallableEnv sets the [CallableEnv] associated with this scope.
func (scope *Scope) SetCallableEnv(callableEnv CallableEnv) {
	scope.callableEnv = callableEnv
}

// GetCallableEnv gets the [CallableEnv] associated with this scope.
func (scope *Scope) GetCallableEnv() CallableEnv {
	s := scope
	for s.callableEnv == nil && s.parent != nil {
		s = s.parent
	}
	return s.callableEnv
}

// SetStruct sets the [StructSymbol] if this scope was created for a struct.
func (scope *Scope) SetStruct(structSym *StructSymbol) {
	scope.structSym = structSym
}

// GetStruct gets the [StructSymbol] if this scope was created for a struct.
// Returns nil otherwise.
func (scope *Scope) GetStruct() *StructSymbol {
	s := scope
	for s.structSym == nil && s.parent != nil {
		s = s.parent
	}
	return s.structSym
}

// SetInterface sets the [InterfaceSymbol] if this scope was created for an interface.
func (scope *Scope) SetInterface(interfaceSym *InterfaceSymbol) {
	scope.interfaceSym = interfaceSym
}

// GetInterface gets the [InterfaceSymbol] if this scope was created for an interface.
// Returns nil otherwise.
func (scope *Scope) GetInterface() *InterfaceSymbol {
	s := scope
	for s.interfaceSym == nil && s.parent != nil {
		s = s.parent
	}
	return s.interfaceSym
}
