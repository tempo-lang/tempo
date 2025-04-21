package analyzer

import (
	"chorego/analyzer/analyzer_error"
	"chorego/parser"
	"slices"
)

func (a *AnalyzerListener) checkFuncParamUnknownRoles(ctx *parser.Func_paramContext) {
	fn := a.funcScope()

	allRoles := []parser.IIdentContext{}
	if roleType := ctx.Value_type().Role_type().Role_type_normal(); roleType != nil {
		allRoles = roleType.AllIdent()
	}
	if roleType := ctx.Value_type().Role_type().Role_type_shared(); roleType != nil {
		allRoles = roleType.AllIdent()
	}

	for _, role := range allRoles {
		containsRole := slices.ContainsFunc(fn.Role_type_normal().AllIdent(), func(funcRole parser.IIdentContext) bool {
			return funcRole.GetText() == role.GetText()
		})

		if !containsRole {
			a.ErrorListener.ReportAnalyzerError(analyzer_error.NewUnknownRoleError(fn, role))
		}
	}
}
