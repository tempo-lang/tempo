package type_check

import (
	"chorego/parser"
	"chorego/types"
	"slices"
	"strconv"
)

func (tc *typeChecker) registerType(expr parser.IExprContext, exprType *types.Type) *types.Type {
	tc.info.Types[expr] = exprType
	return exprType
}

func (tc *typeChecker) VisitExprAdd(ctx *parser.ExprAddContext) any {
	lhs := ctx.Expr(0).Accept(tc).(*types.Type)
	rhs := ctx.Expr(1).Accept(tc).(*types.Type)

	if lhs.Value() == types.Int() && rhs.Value() == types.Int() {
		newRoles, err := types.RoleIntersect(ctx, lhs.Roles(), rhs.Roles())
		if err != nil {
			tc.reportError(err)
			return tc.registerType(ctx, types.Invalid())
		} else {
			return tc.registerType(ctx, types.New(types.Int(), newRoles))
		}
	}

	tc.reportError(types.NewTypeMismatchError(ctx, lhs, rhs))
	return tc.registerType(ctx, types.Invalid())
}

func (tc *typeChecker) VisitExprGroup(ctx *parser.ExprGroupContext) any {
	innerType := ctx.Expr().Accept(tc).(*types.Type)
	return tc.registerType(ctx, innerType)
}

func (tc *typeChecker) VisitExprBool(ctx *parser.ExprBoolContext) any {
	return tc.registerType(ctx, types.New(types.Bool(), types.NewRole(nil, true)))
}

func (tc *typeChecker) VisitExprIdent(ctx *parser.ExprIdentContext) any {
	sym, err := tc.currentScope.LookupSymbol(ctx.Ident())
	if err != nil {
		tc.reportError(err)
		return tc.registerType(ctx, types.Invalid())
	}
	return tc.registerType(ctx, sym.Type())
}

func (tc *typeChecker) VisitExprNum(ctx *parser.ExprNumContext) any {
	_, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		tc.reportError(types.NewInvalidNumberError(ctx))
	}

	return tc.registerType(ctx, types.New(types.Int(), types.NewRole(nil, true)))
}

func (tc *typeChecker) VisitExprAwait(ctx *parser.ExprAwaitContext) any {

	exprType := ctx.Expr().Accept(tc).(*types.Type)

	if asyncType, isAsync := exprType.Value().(*types.Async); isAsync {
		return tc.registerType(ctx, types.New(asyncType.Inner(), exprType.Roles()))
	}

	tc.reportError(types.NewExpectedAsyncTypeError(ctx, exprType))
	return tc.registerType(ctx, types.Invalid())
}

func (tc *typeChecker) VisitExprCom(ctx *parser.ExprComContext) any {

	innerExprType := ctx.Expr().Accept(tc).(*types.Type)

	if !innerExprType.Value().IsSendable() {
		tc.reportError(types.NewUnsendableTypeError(ctx, innerExprType))
	}

	fromRoles, err := types.ParseRoleType(ctx.RoleType(0))
	if err != nil {
		tc.reportError(err)
	} else {
		if fromRoles.IsSharedRole() {
			tc.reportError(types.NewComSharedTypeError(ctx, innerExprType))
		}

		tc.checkRolesInScope(ctx.RoleType(0))

		if len(fromRoles.Participants()) > 1 {
			tc.reportError(types.NewComDistributedTypeError(ctx, innerExprType))
		}

		exprHasParticipants := len(innerExprType.Roles().Participants()) > 0
		if exprHasParticipants && !innerExprType.Roles().Contains(fromRoles.Participants()[0]) {
			tc.reportError(types.NewComValueNotAtSenderError(ctx, innerExprType))
		}
	}

	toRoles, err := types.ParseRoleType(ctx.RoleType(1))
	if err != nil {
		tc.reportError(err)
	} else {
		tc.checkRolesInScope(ctx.RoleType(1))
	}

	recvType := types.Invalid()
	if innerExprType.Value().IsSendable() {
		newParticipants := []string{}
		newParticipants = append(newParticipants, innerExprType.Roles().Participants()...)
		for _, role := range toRoles.Participants() {
			if !slices.Contains(newParticipants, role) {
				newParticipants = append(newParticipants, role)
			}
		}

		isShared := len(newParticipants) > 1
		recvType = types.New(types.NewAsync(innerExprType.Value()), types.NewRole(newParticipants, isShared))
	}

	return tc.registerType(ctx, recvType)
}
