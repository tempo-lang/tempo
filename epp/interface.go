package epp

import (
	"tempo/parser"
	"tempo/projection"
	"tempo/sym_table"
)

func (epp *epp) eppInterface(ctx parser.IInterfaceContext) *projection.ChoreographyInterface {
	sym := epp.info.Symbols[ctx.Ident()].(*sym_table.InterfaceSymbol)

	c := projection.NewChoreographyInterface(sym.SymbolName())

	for _, role := range sym.Type().Roles().Participants() {
		inf := c.AddInterface(role, ctx)

		for _, method := range ctx.InterfaceMethodsList().AllInterfaceMethod() {
			funcSig := epp.eppFuncSig(role, method.FuncSig())
			inf.AddMethod(funcSig, method)
		}
	}

	return c
}
