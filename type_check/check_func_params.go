package type_check

import (
	"chorego/parser"
	"chorego/types"
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
			a.errorListener.ReportTypeError(types.NewUnknownRoleError(fn, role))
		}
	}
}
