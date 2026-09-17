package type_check

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/sym_table"
)

func (tc *typeChecker) addGlobalSymbols(sourceFile *ast.SourceFile) {
	for _, inf := range sourceFile.Interfaces {
		infType, ok := tc.parseInterfaceType(inf)
		if !ok {
			continue
		}

		infScope := tc.currentScope.MakeChild(nodeSpan(inf), infType.Roles().Participants())
		tc.currentScope = infScope

		infSym := sym_table.NewInterfaceSymbol(inf, infScope, infType)
		tc.currentScope.SetInterface(infSym.(*sym_table.InterfaceSymbol))

		tc.currentScope = tc.currentScope.Parent()

		tc.insertSymbol(infSym)
	}

	for _, st := range sourceFile.Structs {
		stType, ok := tc.parseStructType(st)
		if !ok {
			continue
		}

		structScope := tc.currentScope.MakeChild(nodeSpan(st), stType.Roles().Participants())
		tc.currentScope = structScope

		structSym := sym_table.NewStructSymbol(st, structScope, stType)
		tc.currentScope.SetStruct(structSym.(*sym_table.StructSymbol))

		tc.currentScope = tc.currentScope.Parent()

		tc.insertSymbol(structSym)
	}

	for _, fn := range sourceFile.Functions {
		tc.addFuncSymbol(fn.FuncSig, fn)
	}
}
