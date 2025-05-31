package type_check

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"
)

func (tc *typeChecker) visitValueType(ctx parser.IValueTypeContext) *types.Type {
	if ctx == nil {
		return types.Invalid()
	}

	valType, err := tc.parseValueType(ctx)
	if err != nil {
		tc.reportError(err)
		return types.Invalid()
	}

	if err := tc.checkDuplicateRoles(ctx.RoleType(), valType.Roles()); err != nil {
		tc.reportError(err)
	}

	return valType
}

func (tc *typeChecker) VisitFuncTypeSig(ctx *parser.FuncTypeSigContext) any {
	return nil
}

func (tc *typeChecker) VisitFuncTypeParamList(ctx *parser.FuncTypeParamListContext) any {
	return nil
}

func (tc *typeChecker) VisitValueType(ctx *parser.ValueTypeContext) any {
	return nil
}

func (tc *typeChecker) VisitRoleType(ctx *parser.RoleTypeContext) any {
	return nil
}

func (tc *typeChecker) VisitRoleTypeNormal(ctx *parser.RoleTypeNormalContext) any {
	return nil
}

func (tc *typeChecker) VisitRoleTypeShared(ctx *parser.RoleTypeSharedContext) any {
	return nil
}
