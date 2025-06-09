package type_check

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

func (tc *typeChecker) mergeTypes(expr parser.IExprContext, a, b types.Type) (types.Type, bool) {
	if a.Roles().IsDistributedRole() && !b.Roles().IsDistributedRole() {
		tc.reportError(type_error.NewUnmergableRoles(expr, []*types.Roles{a.Roles(), b.Roles()}))
		return types.Invalid(), false
	}

	if b.Roles().IsDistributedRole() && !a.Roles().IsDistributedRole() {
		tc.reportError(type_error.NewUnmergableRoles(expr, []*types.Roles{a.Roles(), b.Roles()}))
		return types.Invalid(), false
	}

	intersectRoles, ok := types.RoleIntersect(a.Roles(), b.Roles())
	if !ok {
		tc.reportError(type_error.NewUnmergableRoles(expr, []*types.Roles{a.Roles(), b.Roles()}))
		return types.Invalid(), false
	}

	intersectParticipants := intersectRoles.Participants()

	newA := a.ReplaceSharedRoles(intersectParticipants)
	newB := b.ReplaceSharedRoles(intersectParticipants)

	if result, ok := newA.CoerceTo(newB); ok {
		return result, true
	}

	if result, ok := newB.CoerceTo(newA); ok {
		return result, true
	}

	tc.reportError(type_error.NewValueMismatch(expr, a, b))
	return types.Invalid(), false
}
