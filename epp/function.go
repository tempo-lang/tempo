package epp

import (
	"tempo/parser"
	"tempo/projection"
	"tempo/type_check"
	"tempo/types"
)

func EppFunc(info *type_check.Info, function parser.IFuncContext) *projection.Choreography {
	sym := info.Symbols[function.Ident()]
	func_role := sym.Type().Roles()

	choreography := projection.NewChoreography(function.Ident().GetText())

	for _, role := range func_role.Participants() {
		eppFuncRole(info, choreography, function, role)
	}

	return choreography
}

func eppFuncRole(info *type_check.Info, choreography *projection.Choreography, function parser.IFuncContext, roleName string) {
	fn := choreography.AddFunc(roleName, function)

	funcSym := info.Symbols[function.Ident()]
	funcType := funcSym.Type().Value().(*types.FunctionType)

	// project parameters
	for i, param := range function.FuncParamList().AllFuncParam() {
		paramType := funcType.Params()[i]
		if paramType.Roles().Contains(roleName) {
			fn.AddParam(param, paramType)
		}
	}

	// project body
	for _, stmt := range function.Scope().AllStmt() {
		eppStmts := EppStmt(info, roleName, stmt)
		fn.AddStmt(eppStmts...)
	}
}
