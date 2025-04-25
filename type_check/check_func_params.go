package type_check

import (
	"chorego/parser"
	"chorego/type_check/type_error"
	"slices"
)

func (a *typeChecker) checkFuncParamUnknownRoles(ctx *parser.FuncParamContext) {
	fn := a.funcScope()

	allRoles := []parser.IIdentContext{}
	if roleType := ctx.ValueType().RoleType().RoleTypeNormal(); roleType != nil {
		allRoles = roleType.AllIdent()
	}
	if roleType := ctx.ValueType().RoleType().RoleTypeShared(); roleType != nil {
		allRoles = roleType.AllIdent()
	}

	for _, role := range allRoles {
		containsRole := slices.ContainsFunc(fn.RoleTypeNormal().AllIdent(), func(funcRole parser.IIdentContext) bool {
			return funcRole.GetText() == role.GetText()
		})

		if !containsRole {
			a.ErrorListener.ReportTypeError(type_error.NewUnknownRoleError(fn, role))
		}
	}
}
