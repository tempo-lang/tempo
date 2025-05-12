package type_check

import (
	"chorego/parser"
	"chorego/types"
	"slices"

	"github.com/antlr4-go/antlr/v4"
)

func (tc *typeChecker) checkRolesInScope(value antlr.ParserRuleContext, roleType *types.Roles) bool {
	idents := roleType.Participants()
	unknownRoles := []string{}

	for _, ident := range idents {
		if !slices.Contains(tc.currentScope.Roles(), ident) {
			unknownRoles = append(unknownRoles, ident)
		}
	}

	if len(unknownRoles) > 0 {
		tc.reportError(types.NewValueRoleNotInScopeError(value, roleType, unknownRoles))
		return false
	}

	return true
}

func (tc *typeChecker) checkRolesExist(roleType parser.IRoleTypeContext) bool {
	idents := parser.RoleTypeAllIdents(roleType)
	unknownRoles := []string{}

	if len(unknownRoles) > 0 {
		tc.reportError(types.NewUnknownRoleError(roleType, unknownRoles))
		return false
	}

	return true
}
