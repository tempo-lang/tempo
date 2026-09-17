package type_check

import (
	"fmt"
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

func (tc *typeChecker) visitStruct(st *ast.Struct) {
	sym, ok := tc.info.Symbols[st.Name].(*sym_table.StructSymbol)
	if !ok {
		return
	}
	tc.currentScope = sym.Scope()
	for _, field := range st.Body.Fields {
		fieldType := tc.visitValueType(field.Type)
		fs := sym_table.NewStructFieldSymbol(field, sym, fieldType)
		tc.insertSymbol(fs)
		sym.AddField(fs.(*sym_table.StructFieldSymbol))
		if !fieldType.IsInvalid() {
			if rt, found := findRoleType(field.Type); found {
				tc.checkRolesInScope(rt)
			}
		}
	}
	for _, method := range st.Body.Functions {
		if fs, ok := tc.addFuncSymbol(method.FuncSig, method); ok {
			sym.AddMethod(fs.(*sym_table.FuncSymbol))
		}
	}
	for _, method := range st.Body.Functions {
		if method.FuncSig.RoleType != nil {
			tc.checkRolesInScope(method.FuncSig.RoleType)
		}
		tc.visitFunc(method)
	}
	tc.currentScope = tc.currentScope.Parent()
	tc.checkStructImplementsConform(sym)
}

func (tc *typeChecker) checkStructImplementsConform(sym *sym_table.StructSymbol) {
	st := sym.Type().(*types.StructType)
	for _, impl := range st.Implements() {
		inf, ok := impl.(*types.InterfaceType)
		if !ok {
			panic(fmt.Sprintf("struct implementation is %T", impl))
		}
		infSym := tc.info.Symbols[inf.Ident()].(*sym_table.InterfaceSymbol)
		for _, field := range tc.info.Fields(inf) {
			fn, ok := field.(*types.FunctionType)
			if !ok {
				continue
			}
			method, found := tc.info.Field(st, fn.NameIdent().Value())
			if !found {
				tc.reportError(type_error.NewMissingImplementationMethod(sym, infSym, fn.NameIdent().Value()))
				continue
			}
			if _, ok := fn.CoerceTo(method); !ok {
				if m, ok := method.(*types.FunctionType); ok {
					tc.reportError(type_error.NewIncompatibleImplementationMethod(sym, infSym, m, fn))
				}
			}
		}
	}
}
