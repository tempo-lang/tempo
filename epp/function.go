package epp

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/types"
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
	funcType := funcSym.Type().(*types.FunctionType)

	returnValue := epp.eppType(role, funcType.ReturnType())

	funcSig := projection.NewFuncSig(role, ctx, returnValue)

	// project parameters
	for i, param := range ctx.FuncParamList().AllFuncParam() {
		paramType := funcType.Params()[i]
		if paramType.Roles().Contains(role) {
			paramValue := epp.eppType(role, paramType)
			funcSig.AddParam(param, paramValue)
		}
	}

	return funcSig
}

func (epp *epp) eppFuncRole(choreography *projection.Choreography, function parser.IFuncContext, roleName string) {
	funcSig := epp.eppFuncSig(roleName, function.FuncSig())
	fn := choreography.AddFunc(funcSig, function)

	// project body
	for _, stmt := range function.Scope().AllStmt() {
		eppStmts := epp.EppStmt(roleName, stmt)
		fn.AddStmt(eppStmts...)
	}
}
