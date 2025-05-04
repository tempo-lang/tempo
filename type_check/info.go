package type_check

import (
	"chorego/parser"
	"chorego/sym_table"
)

type Info struct {
	Symbols     map[parser.IIdentContext]sym_table.Symbol
	GlobalScope *sym_table.Scope
}

func newInfo() *Info {
	return &Info{
		Symbols:     map[parser.IIdentContext]sym_table.Symbol{},
		GlobalScope: nil,
	}
}
