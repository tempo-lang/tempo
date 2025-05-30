package type_check

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"
)

type Info struct {
	Symbols     map[parser.IIdentContext]sym_table.Symbol
	Types       map[parser.IExprContext]*types.Type
	GlobalScope *sym_table.Scope
}

func newInfo() *Info {
	return &Info{
		Symbols:     map[parser.IIdentContext]sym_table.Symbol{},
		Types:       map[parser.IExprContext]*types.Type{},
		GlobalScope: nil,
	}
}
