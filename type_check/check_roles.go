package type_check

import (
	"slices"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

// coerceExprToScope takes an expression node and its type and returns a supertype of its type that only includes the roles in scope.
// If this coercion was not possible, an error is reported and the [types.Invalid] type is returned instead.
func (tc *typeChecker) coerceExprToScope(value parser.IExprContext, exprType types.Type) types.Type {
	scopeRoles := tc.currentScope.Roles()
	typeParticipants := exprType.Roles().Participants()

	switch exprType.(type) {
	case *types.ClosureType, *types.FunctionType:
		inScope := tc.checkExprInScope(value, exprType.Roles())
		if !inScope {
			return types.Invalid()
		}
	}

	if exprType.Roles().IsSharedRole() {
		rolesInScope := []string{}
		for _, role := range typeParticipants {
			if scopeRoles.Contains(role) {
				rolesInScope = append(rolesInScope, role)
			}
		}

		if len(typeParticipants) > 0 && len(rolesInScope) == 0 {
			tc.reportError(type_error.NewValueRoleNotInScope(value, exprType.Roles(), typeParticipants))
			return types.Invalid()
		}

		return exprType.ReplaceSharedRoles(rolesInScope)
	}

	subst := types.NewRoleSubst()
	for _, role := range typeParticipants {
		if !scopeRoles.Contains(role) {
			subst.AddRole(role, "_")
		}
	}

	return exprType.SubstituteRoles(subst)
}

// limitTypeToRoles if successful, converts a type to a supertype that only includes the given roles.
// It returns the supertype and whether the conversion was successful.
func limitTypeToRoles(exprType types.Type, roles []string) (types.Type, bool) {
	typeParticipants := exprType.Roles().Participants()

	rolesInScope := []string{}
	for _, role := range typeParticipants {
		if slices.Contains(roles, role) {
			rolesInScope = append(rolesInScope, role)
		}
	}

	if len(typeParticipants) > 0 && len(rolesInScope) == 0 {
		return types.Invalid(), false
	}

	switch exprType.(type) {
	case *types.ClosureType, *types.FunctionType:
		if len(roles) != len(rolesInScope) {
			return types.Invalid(), false
		}
	}

	if exprType.Roles().IsSharedRole() {
		return exprType.ReplaceSharedRoles(rolesInScope), true
	}

	subst := types.NewRoleSubst()
	for _, r := range typeParticipants {
		if !slices.Contains(roles, r) {
			subst.AddRole(r, "_")
		}
	}

	return exprType.SubstituteRoles(subst), true
}

// checkExprInScope returns true if the roles in the expression is in scope.
// Otherwise it returns false and reports an appropriate error.
func (tc *typeChecker) checkExprInScope(value antlr.ParserRuleContext, roleType *types.Roles) bool {
	unknownRoles := tc.rolesNotInScope(roleType.Participants())

	if len(unknownRoles) > 0 {
		tc.reportError(type_error.NewValueRoleNotInScope(value, roleType, unknownRoles))
		return false
	}

	return true
}

// checkRolesInScope returns true if the roles explicitly specified in the RoleType node are in scope.
// Otherwise it returns false and reports an appropriate error.
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

// rolesNotInScope returns the participants that are not in scope.
// Hidden roles are removed from the result as well.
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
