package epp

import (
	"chorego/parser"
	"slices"
)

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
