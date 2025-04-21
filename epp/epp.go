package epp

import (
	"chorego/parser"
	"chorego/projection"
	"slices"
)

func EppFunc(function parser.IFuncContext) *projection.Choreography {
	func_role := function.Role_type_normal()

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
	for _, param := range function.Func_param_list().AllFunc_param() {
		if roleTypeNormal := param.Value_type().Role_type().Role_type_normal(); roleTypeNormal != nil {
			paramRoles := roleTypeNormal.AllIdent()
			containsRole := slices.ContainsFunc(paramRoles, func(role parser.IIdentContext) bool {
				return role.GetText() == roleName
			})

			if containsRole {
				fn.AddParam(param, param.Value_type().Ident().GetText())
			}
		}

		if roleTypeShared := param.Value_type().Role_type().Role_type_shared(); roleTypeShared != nil {
			panic("shared role type in parameter not supported yet")
		}
	}
}
