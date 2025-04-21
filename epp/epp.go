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
		if roleTypeNormal := param.ValueType().RoleType().RoleTypeNormal(); roleTypeNormal != nil {
			paramRoles := roleTypeNormal.AllIdent()
			containsRole := slices.ContainsFunc(paramRoles, func(role parser.IIdentContext) bool {
				return role.GetText() == roleName
			})

			if containsRole {
				fn.AddParam(param, param.ValueType().Ident().GetText())
			}
		}

		if roleTypeShared := param.ValueType().RoleType().RoleTypeShared(); roleTypeShared != nil {
			panic("shared role type in parameter not supported yet")
		}
	}
}
