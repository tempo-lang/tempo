package type_check

import (
	"fmt"
	"slices"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
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

// IsTypeSendable describes whether the type can be communicated using the `->` operator.
func (tc *typeChecker) IsTypeSendable(t types.Type) bool {
	toCheck := []types.Type{t}
	checkedTypes := []types.Type{}

	for len(toCheck) > 0 {
		t := toCheck[len(toCheck)-1]
		toCheck = toCheck[:len(toCheck)-1]

		if slices.Contains(checkedTypes, t) {
			continue
		}

		checkedTypes = append(checkedTypes, t)

		switch t := t.(type) {
		case *types.AsyncType:
			return false
		case *types.BoolType:
			return true
		case *types.ClosureType:
			return false
		case *types.FloatType:
			return true
		case *types.FunctionType:
			return false
		case *types.IntType:
			return true
		case *types.InterfaceType:
			return false
		case *types.InvalidType:
			return true
		case *types.ListType:
			toCheck = append(toCheck, t.Inner())
		case *types.StringType:
			return true
		case *types.StructType:
			sym := tc.info.Symbols[t.Ident()].(*sym_table.StructSymbol)
			for _, field := range sym.Fields() {
				toCheck = append(toCheck, field.Type())
			}
		case *types.UnitType:
			return false
		default:
			panic(fmt.Sprintf("unexpected types.Type: %#v", t))
		}
	}

	return true
}
