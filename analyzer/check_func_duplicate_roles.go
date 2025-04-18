package analyzer

import (
	"chorego/parser"
	"cmp"
	"slices"
)

func (a AnalyzerListener) checkFuncDuplicateRoles(ctx parser.IFuncContext) {
	roles := ctx.Role_type_normal().AllIdent()
	slices.SortFunc(roles, func(a, b parser.IIdentContext) int {
		return cmp.Compare(a.GetText(), b.GetText())
	})

	if len(roles) > 1 {
		duplicateRoles := roles[0:1]
		for i, role := range roles[1:] {
			if duplicateRoles[0].GetText() == role.GetText() {
				duplicateRoles = append(duplicateRoles, role)
			} else {
				// report collected duplicates error if any
				if len(duplicateRoles) > 1 {
					a.ErrorListener.ReportAnalyzerError(NewDuplicateRolesError(ctx, duplicateRoles))
					duplicateRoles = roles[i+1 : i+2]
				}
			}
		}

		// report last error if present
		if len(duplicateRoles) > 1 {
			a.ErrorListener.ReportAnalyzerError(NewDuplicateRolesError(ctx, duplicateRoles))
		}
	}
}
