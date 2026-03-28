package type_check

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

func (tc *typeChecker) coerceExprToScope(value antlr.ParserRuleContext, exprType types.Type) types.Type {
	scopeRoles := tc.currentScope.Roles()

	if exprType.Roles().IsSharedRole() {
		rolesInScope := []string{}
		for _, role := range exprType.Roles().Participants() {
			if scopeRoles.Contains(role) {
				rolesInScope = append(rolesInScope, role)
			}
		}
		return exprType.ReplaceSharedRoles(rolesInScope)
	}

	subst := types.NewRoleSubst()
	for _, role := range exprType.Roles().Participants() {
		if !scopeRoles.Contains(role) {
			subst.AddRole(role, "_")
		}
	}

	return exprType.SubstituteRoles(subst)
}

func (tc *typeChecker) checkExprInScope(value antlr.ParserRuleContext, roleType *types.Roles) bool {
	unknownRoles := tc.rolesNotInScope(roleType.Participants())

	if len(unknownRoles) > 0 {
		tc.reportError(type_error.NewValueRoleNotInScope(value, roleType, unknownRoles))
		return false
	}

	return true
}

func (tc *typeChecker) checkRolesInScope(roleType parser.IRoleTypeContext) bool {

	participants := []string{}
	for _, p := range parser.RoleTypeAllRoles(roleType) {
		participants = append(participants, p.GetText())
	}

	unknownRoles := tc.rolesNotInScope(participants)
	if len(unknownRoles) > 0 {
		tc.reportError(type_error.NewRolesNotInScope(roleType, unknownRoles))
		return false
	}

	return true
}

func (tc *typeChecker) rolesNotInScope(participants []string) []string {
	scopeRoles := tc.currentScope.Roles()
	unknownRoles := []string{}
	for _, p := range participants {
		if p == "_" {
			continue // skip hidden roles
		}

		if !scopeRoles.Contains(p) {
			unknownRoles = append(unknownRoles, p)
		}
	}

	return unknownRoles
}
