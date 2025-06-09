package type_check

import (
	"strconv"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

func (tc *typeChecker) checkLiteralRoles(roleTypeCtx parser.IRoleTypeContext) *types.Roles {
	if roleTypeCtx == nil {
		return types.EveryoneRole()
	}

	roleType, ok := tc.parseRoleType(roleTypeCtx)
	if !ok {
		return types.EveryoneRole()
	}

	hasError := false
	if roleType.IsDistributedRole() {
		tc.reportError(type_error.NewNotDistributedType(roleTypeCtx))
		hasError = true
	}

	if !tc.checkRolesInScope(roleTypeCtx) {
		hasError = true
	}

	if hasError {
		return types.EveryoneRole()
	}

	return roleType
}

func (tc *typeChecker) VisitExprPrimitive(ctx *parser.ExprPrimitiveContext) any {
	roleType := tc.checkLiteralRoles(ctx.RoleType())
	value := ctx.Literal().Accept(tc).(types.Type)
	roleSubstMap, _ := value.Roles().SubstituteMap(roleType)
	return tc.registerType(ctx, value.SubstituteRoles(roleSubstMap))
}

func (tc *typeChecker) VisitInt(ctx *parser.IntContext) any {
	_, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		tc.reportError(type_error.NewInvalidNumber(ctx, err))
	}

	return types.Int(nil)
}

func (tc *typeChecker) VisitFloat(ctx *parser.FloatContext) any {
	_, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		tc.reportError(type_error.NewInvalidNumber(ctx, err))
	}

	return types.Float(nil)
}

func (tc *typeChecker) VisitBool(ctx *parser.BoolContext) any {
	return types.Bool(nil)
}

func (tc *typeChecker) VisitString(ctx *parser.StringContext) any {
	return types.String(nil)
}
