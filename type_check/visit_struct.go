package type_check

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

func (tc *typeChecker) VisitStruct(ctx *parser.StructContext) any {
	// structs are already resolved by addGlobalSymbols
	structSym, found := tc.info.Symbols[ctx.Ident()].(*sym_table.StructSymbol)
	if !found {
		// was not found if struct has parser errors
		return nil
	}

	if ctx.GetBody() == nil {
		// parser error
		return nil
	}

	tc.currentScope = structSym.Scope()

	ctx.GetBody().Accept(tc)

	tc.currentScope = tc.currentScope.Parent()

	tc.checkStructImplements(structSym)

	return nil
}

func (tc *typeChecker) VisitStructBody(ctx *parser.StructBodyContext) any {
	for _, field := range ctx.AllStructField() {
		field.Accept(tc)
	}

	structSym := tc.currentScope.GetStruct()

	for _, method := range ctx.AllFunc_() {
		funcSym, ok := tc.addFuncSymbol(method.FuncSig(), method)
		if ok {
			structSym.AddMethod(funcSym.(*sym_table.FuncSymbol))
		}
	}

	for _, method := range ctx.AllFunc_() {

		tc.checkRolesInScope(method.FuncSig().RoleType())

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

func (tc *typeChecker) VisitStructImplements(ctx *parser.StructImplementsContext) any {
	return nil
}

func (tc *typeChecker) checkStructImplements(sym *sym_table.StructSymbol) {
	structType := sym.Type().(*types.StructType)

	for _, impl := range structType.Implements() {
		infType, ok := impl.(*types.InterfaceType)
		if !ok {
			panic(fmt.Sprintf("Struct can only implement interfaces, type was %T", impl))
		}

		infSym := tc.info.Symbols[infType.Ident()].(*sym_table.InterfaceSymbol)

		for _, field := range tc.info.Fields(infType) {
			fn, ok := field.(*types.FunctionType)
			if !ok {
				panic(fmt.Sprintf("Interface can only have function fields, was %T", field))
			}

			method, found := tc.info.Field(structType, fn.NameIdent().GetText())
			if !found {
				tc.reportError(type_error.NewMissingImplementationMethod(sym, infSym, fn.NameIdent().GetText()))
				continue
			}

			_, canCoerce := fn.CoerceTo(method)
			if !canCoerce {
				tc.reportError(type_error.NewIncompatibleImplementationMethod(sym, infSym, fn.NameIdent().GetText()))
				continue
			}
		}
	}

}
