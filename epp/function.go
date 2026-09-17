package epp

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/types"
)

func (epp *epp) eppFunc(function *ast.Func) *projection.Choreography {
	sym := epp.info.Symbols[function.FuncSig.Name]
	func_role := sym.Type().Roles()

	choreography := projection.NewChoreography(function.FuncSig.Name.Value())

	for _, role := range func_role.Participants() {
		epp.eppFuncRole(choreography, function, role)
	}

	return choreography
}

func (epp *epp) eppFuncSig(role string, ctx *ast.FuncSig) *projection.FuncSig {
	funcSym := epp.info.Symbols[ctx.Name]
	funcType := funcSym.Type().(*types.FunctionType)

	if !funcType.Roles().Contains(role) {
		return nil
	}

	returnValue := epp.eppType(role, funcType.ReturnType())

	funcSig := projection.NewFuncSig(role, ctx, returnValue)

	// project parameters
	for i, param := range ctx.Params.Params {
		paramType := funcType.Params()[i]
		if paramType.Roles().Contains(role) {
			paramValue := epp.eppType(role, paramType)
			funcSig.AddParam(param, paramValue)
		}
	}

	return funcSig
}

func (epp *epp) eppFuncRole(choreography *projection.Choreography, function *ast.Func, roleName string) {
	funcSig := epp.eppFuncSig(roleName, function.FuncSig)
	fn := choreography.AddFunc(funcSig, function)

	// project body
	for _, stmt := range function.Scope.Stmts {
		eppStmts := epp.EppStmt(roleName, stmt)
		fn.AddStmt(eppStmts...)
	}
}
