package epp

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/sym_table"
)

func (epp *epp) eppInterface(ctx *ast.Interface) *projection.ChoreographyInterface {
	sym := epp.info.Symbols[ctx.Name].(*sym_table.InterfaceSymbol)

	c := projection.NewChoreographyInterface(sym.SymbolName())

	for _, role := range sym.Type().Roles().Participants() {
		inf := c.AddInterface(role, ctx)

		for _, method := range ctx.Methods.Methods {
			funcSig := epp.eppFuncSig(role, method.FuncSig)
			if funcSig != nil {
				inf.AddMethod(funcSig, method)
			}
		}
	}

	return c
}
