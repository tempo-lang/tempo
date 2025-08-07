package type_check

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
)

func (tc *typeChecker) addGlobalSymbols(sourceFile *parser.SourceFileContext) {
	for _, inf := range sourceFile.AllInterface_() {
		infType, err := tc.parseInterfaceType(inf)
		if err != nil {
			tc.reportError(err)
			continue
		}

		if infType.IsInvalid() {
			continue
		}

		infScope := tc.currentScope.MakeChild(inf.GetStart(), inf.GetStop(), infType.Roles().Participants())
		tc.currentScope = infScope

		infSym := sym_table.NewInterfaceSymbol(inf, infScope, infType)
		tc.currentScope.SetInterface(infSym.(*sym_table.InterfaceSymbol))

		tc.currentScope = tc.currentScope.Parent()

		tc.insertSymbol(infSym)
	}

	for _, st := range sourceFile.AllStruct_() {
		stType, err := tc.parseStructType(st)
		if err != nil {
			tc.reportError(err)
			continue
		}

		if stType.IsInvalid() {
			continue
		}

		structScope := tc.currentScope.MakeChild(st.GetStart(), st.GetStop(), stType.Roles().Participants())
		tc.currentScope = structScope

		structSym := sym_table.NewStructSymbol(st, structScope, stType)
		tc.currentScope.SetStruct(structSym.(*sym_table.StructSymbol))

		tc.currentScope = tc.currentScope.Parent()

		tc.insertSymbol(structSym)
	}

	for _, fn := range sourceFile.AllFunc_() {
		tc.addFuncSymbol(fn.FuncSig(), fn)
	}
}
