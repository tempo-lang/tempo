package type_check

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"
)

// Info contains extra type information that can be queried from the original AST passed to the [TypeCheck] function.
type Info struct {
	// Symbols maps every identifier in the AST to its underlying symbol.
	// Only if the identifier references an undefined symbol will the map return `nil`.
	Symbols map[parser.IIdentContext]sym_table.Symbol
	// Types maps every expression (and sub expression) in the AST to its underlying type.
	// The map will return a [types.Invalid] symbol if the expression was wrongly typed.
	Types map[parser.IExprContext]types.Type
	// GlobalScope is a reference to the root scope of the symbol table.
	GlobalScope *sym_table.Scope
}

func newInfo() *Info {
	return &Info{
		Symbols:     map[parser.IIdentContext]sym_table.Symbol{},
		Types:       map[parser.IExprContext]types.Type{},
		GlobalScope: nil,
	}
}
