package type_check

import (
	"slices"
	"tempo/parser"
	"tempo/types"

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
	funcRoles := tc.currentScope.GetFunc().Type().Roles().Participants()

	unknownRoles := []string{}
	for _, i := range idents {
		if !slices.Contains(funcRoles, i.GetText()) {
			unknownRoles = append(unknownRoles, i.GetText())
		}
	}

	if len(unknownRoles) > 0 {
		tc.reportError(types.NewUnknownRoleError(roleType, unknownRoles))
		return false
	}

	return true
}
