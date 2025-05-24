package epp

import (
	"tempo/parser"
	"tempo/projection"
	"tempo/types"
)

func (epp *epp) eppFunc(function parser.IFuncContext) *projection.Choreography {
	sym := epp.info.Symbols[function.FuncSig().Ident()]
	func_role := sym.Type().Roles()

	choreography := projection.NewChoreography(function.FuncSig().Ident().GetText())

	for _, role := range func_role.Participants() {
		epp.eppFuncRole(choreography, function, role)
	}

	return choreography
}

func (epp *epp) eppFuncSig(role string, ctx parser.IFuncSigContext) *projection.FuncSig {
	funcSym := epp.info.Symbols[ctx.Ident()]
	funcType := funcSym.Type().Value().(*types.FunctionType)

	returnValue := epp.eppType(role, funcType.ReturnType())

	funcSig := projection.NewFuncSig(role, ctx, returnValue)

	// project parameters
	for i, param := range ctx.FuncParamList().AllFuncParam() {
		paramType := funcType.Params()[i]
		if paramType.Roles().Contains(role) {
			funcSig.AddParam(param, paramType)
		}
	}

	return funcSig
}

func (epp *epp) eppFuncRole(choreography *projection.Choreography, function parser.IFuncContext, roleName string) {
	funcSym := epp.info.Symbols[function.FuncSig().Ident()]
	funcType := funcSym.Type().Value().(*types.FunctionType)

	funcSig := epp.eppFuncSig(roleName, function.FuncSig())
	fn := choreography.AddFunc(funcSig, function)

	// project parameters
	for i, param := range function.FuncSig().FuncParamList().AllFuncParam() {
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
