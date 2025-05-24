package type_check

import (
	"tempo/parser"
	"tempo/sym_table"
)

func (tc *typeChecker) VisitInterface(ctx *parser.InterfaceContext) any {
	// interfaces are already resolved by addGlobalSymbols
	sym, found := tc.info.Symbols[ctx.Ident()].(*sym_table.InterfaceSymbol)
	if !found {
		// was not found if interface has parser errors
		return nil
	}

	tc.currentScope = sym.Scope()

	ctx.InterfaceMethodsList().Accept(tc)

	tc.currentScope = tc.currentScope.Parent()
	return nil
}

func (tc *typeChecker) VisitInterfaceMethodsList(ctx *parser.InterfaceMethodsListContext) any {

	for _, mtd := range ctx.AllInterfaceMethod() {
		mtd.Accept(tc)
	}

	return nil
}

func (tc *typeChecker) VisitInterfaceMethod(ctx *parser.InterfaceMethodContext) any {
	sym, ok := tc.addFuncSymbol(ctx.FuncSig(), ctx)
	if !ok {
		return nil
	}

	funcSym := sym.(*sym_table.FuncSymbol)

	tc.currentScope = funcSym.Scope()
	ctx.FuncSig().Accept(tc)
	tc.currentScope = tc.currentScope.Parent()

	return nil
}
