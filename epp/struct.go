package epp

import (
	"tempo/parser"
	"tempo/projection"
)

func (epp *epp) eppStruct(st parser.IStructContext) *projection.ChoreographyStruct {
	sym := epp.info.Symbols[st.Ident()]
	result := projection.NewChoreographyStruct(sym.SymbolName())

	for _, role := range sym.Type().Roles().Participants() {
		str := result.AddStruct(role, st)

		for _, field := range st.StructFieldList().AllStructField() {
			fieldSym := epp.info.Symbols[field.Ident()]
			if fieldSym.Type().Roles().Contains(role) {
				fieldType := epp.eppType(role, fieldSym.Type())
				str.AddField(field, fieldType)
			}
		}
	}

	return result
}
