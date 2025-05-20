package type_check

import (
	"tempo/parser"
	"tempo/type_check/type_error"
	"tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

func (tc *typeChecker) checkExprInScope(value antlr.ParserRuleContext, roleType *types.Roles) bool {
	idents := roleType.Participants()
	unknownRoles := []string{}

	for _, ident := range idents {
		if !tc.currentScope.Roles().Contains(ident) {
			unknownRoles = append(unknownRoles, ident)
		}
	}

	if len(unknownRoles) > 0 {
		tc.reportError(type_error.NewValueRoleNotInScopeError(value, roleType, unknownRoles))
		return false
	}

	return true
}

func (tc *typeChecker) checkRolesInScope(roleType parser.IRoleTypeContext) bool {
	idents := parser.RoleTypeAllIdents(roleType)
	scopeRoles := tc.currentScope.Roles()

	unknownRoles := []string{}
	for _, i := range idents {
		if !scopeRoles.Contains(i.GetText()) {
			unknownRoles = append(unknownRoles, i.GetText())
		}
	}

	if len(unknownRoles) > 0 {
		tc.reportError(type_error.NewRolesNotInScopeError(roleType, unknownRoles))
		return false
	}

	return true
}
