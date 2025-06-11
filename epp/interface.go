package epp

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/sym_table"
)

func (epp *epp) eppInterface(ctx parser.IInterfaceContext) *projection.ChoreographyInterface {
	sym := epp.info.Symbols[ctx.Ident()].(*sym_table.InterfaceSymbol)

	c := projection.NewChoreographyInterface(sym.SymbolName())

	for _, role := range sym.Type().Roles().Participants() {
		inf := c.AddInterface(role, ctx)

		for _, method := range ctx.InterfaceMethodsList().AllInterfaceMethod() {
			funcSig := epp.eppFuncSig(role, method.FuncSig())
			if funcSig != nil {
				inf.AddMethod(funcSig, method)
			}
		}
	}

	return c
}
