package epp

import (
	"tempo/parser"
	"tempo/projection"
	"tempo/types"
)

func (epp *epp) eppFunc(function parser.IFuncContext) *projection.Choreography {
	sym := epp.info.Symbols[function.Ident()]
	func_role := sym.Type().Roles()

	choreography := projection.NewChoreography(function.Ident().GetText())

	for _, role := range func_role.Participants() {
		epp.eppFuncRole(choreography, function, role)
	}

	return choreography
}

func (epp *epp) eppFuncRole(choreography *projection.Choreography, function parser.IFuncContext, roleName string) {
	funcSym := epp.info.Symbols[function.Ident()]
	funcType := funcSym.Type().Value().(*types.FunctionType)

	returnValue := epp.eppType(roleName, funcType.ReturnType())
	fn := choreography.AddFunc(roleName, function, returnValue)

	// project parameters
	for i, param := range function.FuncParamList().AllFuncParam() {
		paramType := funcType.Params()[i]
		if paramType.Roles().Contains(roleName) {
			fn.AddParam(param, paramType)
		}
	}

	// project body
	for _, stmt := range function.Scope().AllStmt() {
		eppStmts := epp.EppStmt(roleName, stmt)
		fn.AddStmt(eppStmts...)
	}
}
