package type_check

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/sym_table"
)

func (tc *typeChecker) visitInterface(inf *ast.Interface) {
	// interfaces are already resolved by addGlobalSymbols
	sym, ok := tc.info.Symbols[inf.Name].(*sym_table.InterfaceSymbol)
	if !ok {
		return
	}
	tc.currentScope = sym.Scope()
	for _, method := range inf.Methods.Methods {
		s, ok := tc.addFuncSymbol(method.FuncSig, method)
		if !ok {
			continue
		}
		fn := s.(*sym_table.FuncSymbol)
		sym.AddMethod(fn)
		if method.FuncSig.RoleType != nil {
			tc.checkRolesInScope(method.FuncSig.RoleType)
		}
		tc.currentScope = fn.Scope()
		tc.visitFuncParams(method.FuncSig.Params)
		tc.currentScope = tc.currentScope.Parent()
	}
	tc.currentScope = tc.currentScope.Parent()
}
