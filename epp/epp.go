package epp

import (
	"chorego/parser"
	"chorego/projection"
	"slices"
)

func EppFunc(function parser.IFuncContext) *projection.Choreography {
	func_role := function.RoleTypeNormal()

	choreography := projection.NewChoreography(function.Ident().GetText())

	for _, role := range func_role.AllIdent() {
		eppFuncRole(choreography, function, role)
	}

	return choreography
}

func eppFuncRole(choreography *projection.Choreography, function parser.IFuncContext, role parser.IIdentContext) {
	roleName := role.ID().GetText()
	fn := choreography.AddFunc(roleName, function)

	// project parameters
	for _, param := range function.FuncParamList().AllFuncParam() {
		if ValueExistsAtRole(param.ValueType(), roleName) {
			fn.AddParam(param, param.ValueType().Ident().GetText())
		}
	}

	// project body
	for _, stmt := range function.Scope().AllStatement() {
		if varDecl := stmt.StmtVarDecl(); varDecl != nil {
			if ValueExistsAtRole(varDecl.ValueType(), roleName) {
				variableName := varDecl.Ident().GetText()
				varibleType := varDecl.ValueType().Ident().GetText()

				expr := eppExpression(role, varDecl.Expression())

				fn.AddStmt(projection.NewStmtVarDecl(variableName, varibleType, expr))
			}
		}
	}
}

func ValueExistsAtRole(value parser.IValueTypeContext, roleName string) bool {
	var roles []parser.IIdentContext
	if roleTypeNormal := value.RoleType().RoleTypeNormal(); roleTypeNormal != nil {
		roles = roleTypeNormal.AllIdent()
	}
	if roleTypeShared := value.RoleType().RoleTypeShared(); roleTypeShared != nil {
		roles = roleTypeShared.AllIdent()
	}

	containsRole := slices.ContainsFunc(roles, func(role parser.IIdentContext) bool {
		return role.GetText() == roleName
	})
	return containsRole
}
