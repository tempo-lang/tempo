package epp

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/types"
)

func (epp *epp) eppStruct(st *ast.Struct) *projection.ChoreographyStruct {
	sym := epp.info.Symbols[st.Name]
	result := projection.NewChoreographyStruct(sym.SymbolName())

	stType := sym.Type().(*types.StructType)

	for _, role := range sym.Type().Roles().Participants() {
		str := result.AddStruct(role, st)

		for _, impl := range stType.Implements() {
			eppImpl := epp.eppType(role, impl)
			if eppImpl != projection.UnitType() {
				str.AddImplements(eppImpl)
			}
		}

		for _, field := range st.Body.Fields {
			fieldSym := epp.info.Symbols[field.Name]
			if fieldSym.Type().Roles().Contains(role) {
				fieldType := epp.eppType(role, fieldSym.Type())
				str.AddField(field, fieldType)
			}
		}

		for _, method := range st.Body.Functions {
			methodSym := epp.info.Symbols[method.FuncSig.Name]
			if methodSym.Type().Roles().Contains(role) {
				funcSig := epp.eppFuncSig(role, method.FuncSig)
				m := str.AddMethod(funcSig, method)

				for _, stmt := range method.Scope.Stmts {
					eppStmts := epp.EppStmt(role, stmt)
					m.AddStmt(eppStmts...)
				}
			}
		}
	}

	return result
}
