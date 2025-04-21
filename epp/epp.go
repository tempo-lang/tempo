package epp

import (
	"chorego/parser"
	"chorego/process"
	"slices"
)

func EppFunc(function parser.IFuncContext) *process.Network {
	func_role := function.Role_type_normal()

	network := process.NewNetwork()

	for _, role := range func_role.AllIdent() {
		eppFuncRole(network, function, role)
	}

	return network
}

func eppFuncRole(network *process.Network, function parser.IFuncContext, role parser.IIdentContext) {
	roleName := role.ID().GetText()
	proc := network.Process(roleName)
	fn := proc.AddFunc(function)

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
