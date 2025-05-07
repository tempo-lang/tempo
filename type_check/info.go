package type_check

import (
	"chorego/parser"
	"chorego/sym_table"
	"chorego/types"
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
