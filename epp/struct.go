package epp

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/projection"
)

func (epp *epp) eppStruct(st parser.IStructContext) *projection.ChoreographyStruct {
	sym := epp.info.Symbols[st.Ident()]
	result := projection.NewChoreographyStruct(sym.SymbolName())

	for _, role := range sym.Type().Roles().Participants() {
		str := result.AddStruct(role, st)

		for _, field := range st.GetBody().AllStructField() {
			fieldSym := epp.info.Symbols[field.Ident()]
			if fieldSym.Type().Roles().Contains(role) {
				fieldType := epp.eppType(role, fieldSym.Type())
				str.AddField(field, fieldType)
			}
		}

		for _, method := range st.GetBody().AllFunc_() {
			methodSym := epp.info.Symbols[method.FuncSig().Ident()]
			if methodSym.Type().Roles().Contains(role) {
				funcSig := epp.eppFuncSig(role, method.FuncSig())
				m := str.AddMethod(funcSig, method)

				for _, stmt := range method.Scope().AllStmt() {
					eppStmts := epp.EppStmt(role, stmt)
					m.AddStmt(eppStmts...)
				}
			}
		}
	}

	return result
}
