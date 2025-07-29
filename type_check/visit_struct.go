package type_check

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
)

func (tc *typeChecker) VisitStruct(ctx *parser.StructContext) any {
	// structs are already resolved by addGlobalSymbols
	structSym, found := tc.info.Symbols[ctx.Ident()].(*sym_table.StructSymbol)
	if !found {
		// was not found if struct has parser errors
		return nil
	}

	tc.currentScope = structSym.Scope()

	ctx.GetBody().Accept(tc)

	tc.currentScope = tc.currentScope.Parent()
	return nil
}

func (tc *typeChecker) VisitStructBody(ctx *parser.StructBodyContext) any {
	for _, field := range ctx.AllStructField() {
		field.Accept(tc)
	}

	for _, method := range ctx.AllFunc_() {
		tc.addFuncSymbol(method.FuncSig(), method.Scope())
	}

	for _, method := range ctx.AllFunc_() {
		method.Accept(tc)
	}

	return nil
}

func (tc *typeChecker) VisitStructField(ctx *parser.StructFieldContext) any {
	fieldType := tc.visitValueType(ctx.ValueType())
	fieldSym := sym_table.NewStructFieldSymbol(ctx, tc.currentScope.GetStruct(), fieldType)
	tc.insertSymbol(fieldSym)

	structSym := tc.currentScope.GetStruct()
	structSym.AddField(fieldSym.(*sym_table.StructFieldSymbol))

	if !fieldType.IsInvalid() {
		tc.checkRolesInScope(findRoleType(ctx.ValueType()))
	}

	return nil
}
