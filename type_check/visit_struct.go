package type_check

import (
	"tempo/parser"
	"tempo/sym_table"
)

func (tc *typeChecker) VisitStruct(ctx *parser.StructContext) any {
	// structs are already resolved by addGlobalSymbols
	sym, found := tc.info.Symbols[ctx.Ident()].(*sym_table.StructSymbol)
	if !found {
		// was not found if struct has parser errors
		return nil
	}

	tc.currentScope = sym.Scope()

	ctx.StructFieldList().Accept(tc)

	tc.currentScope = tc.currentScope.Parent()

	return nil
}

func (tc *typeChecker) VisitStructFieldList(ctx *parser.StructFieldListContext) any {
	for _, field := range ctx.AllStructField() {
		field.Accept(tc)
	}

	return nil
}

func (tc *typeChecker) VisitStructField(ctx *parser.StructFieldContext) any {
	fieldType := tc.visitValueType(ctx.ValueType())
	fieldSym := sym_table.NewStructFieldSymbol(ctx, tc.currentScope, fieldType)
	tc.insertSymbol(fieldSym)

	structSym := tc.currentScope.GetStruct()
	structSym.AddField(fieldSym.(*sym_table.StructFieldSymbol))

	tc.checkRolesInScope(ctx.ValueType().RoleType())

	return nil
}
