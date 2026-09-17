package type_check

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
	"strconv"
)

func (tc *typeChecker) checkLiteralRoles(rt *ast.RoleType) *types.Roles {
	if rt == nil {
		return types.EveryoneRole()
	}
	roles, ok := tc.parseRoleType(rt)
	if !ok {
		return types.EveryoneRole()
	}
	bad := false
	if roles.IsDistributedRole() {
		tc.reportError(type_error.NewNotDistributedType(rt))
		bad = true
	}
	if !tc.checkRolesInScope(rt) {
		bad = true
	}
	if bad {
		return types.EveryoneRole()
	}
	return roles
}

func (tc *typeChecker) visitPrimitive(p *ast.PrimitiveExpr) types.Type {
	roles := tc.checkLiteralRoles(p.RoleType)
	var value types.Type
	switch lit := p.Literal.(type) {
	case *ast.IntLit:
		if _, err := strconv.Atoi(lit.IntToken.Text); err != nil {
			tc.reportError(type_error.NewInvalidNumber(lit, err))
		}
		value = types.Int(nil)
	case *ast.FloatLit:
		if _, err := strconv.ParseFloat(lit.FloatToken.Text, 64); err != nil {
			tc.reportError(type_error.NewInvalidNumber(lit, err))
		}
		value = types.Float(nil)
	case *ast.BoolLit:
		value = types.Bool(nil)
	case *ast.StringLit:
		value = types.String(nil)
	default:
		value = types.Invalid()
	}
	return tc.registerType(p, value.ReplaceSharedRoles(roles.Participants()))
}
