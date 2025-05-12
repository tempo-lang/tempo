package type_check

import (
	"cmp"
	"slices"
	"tempo/parser"
	"tempo/types"
)

func (a typeChecker) checkFuncDuplicateRoles(ctx parser.IFuncContext) {
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
					a.errorListener.ReportTypeError(types.NewDuplicateRolesError(ctx, duplicateRoles))
				}
				duplicateRoles = roles[i+1 : i+2]
			}
		}

		// report last error if present
		if len(duplicateRoles) > 1 {
			a.errorListener.ReportTypeError(types.NewDuplicateRolesError(ctx, duplicateRoles))
		}
	}
}
