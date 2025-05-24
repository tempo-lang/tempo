package type_check

import (
	"cmp"
	"slices"
	"tempo/parser"
	"tempo/type_check/type_error"
)

func (a typeChecker) checkFuncDuplicateRoles(ctx parser.IFuncSigContext) {
	roles := parser.RoleTypeAllIdents(ctx.RoleType())
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
					a.errorListener.ReportTypeError(type_error.NewDuplicateRolesError(ctx, duplicateRoles))
				}
				duplicateRoles = roles[i+1 : i+2]
			}
		}

		// report last error if present
		if len(duplicateRoles) > 1 {
			a.errorListener.ReportTypeError(type_error.NewDuplicateRolesError(ctx, duplicateRoles))
		}
	}
}
