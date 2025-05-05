package type_check

import (
	"chorego/parser"
	"chorego/types"
	"strconv"
)

func (tc *typeChecker) VisitExprAdd(ctx *parser.ExprAddContext) any {
	lhs := ctx.Expr(0).Accept(tc).(*types.Type)
	rhs := ctx.Expr(1).Accept(tc).(*types.Type)

	if lhs.Value() == types.Int() && rhs.Value() == types.Int() {
		newRoles, err := types.RoleIntersect(ctx, lhs.Roles(), rhs.Roles())
		if err != nil {
			tc.reportError(err)
			return types.Invalid()
		} else {
			return types.New(types.Int(), newRoles)
		}
	}

	tc.reportError(types.NewTypeMismatchError(ctx, lhs, rhs))
	return types.Invalid()
}

func (tc *typeChecker) VisitExprGroup(ctx *parser.ExprGroupContext) any {
	return ctx.Expr().Accept(tc)
}

func (tc *typeChecker) VisitExprBool(ctx *parser.ExprBoolContext) any {
	return types.New(types.Bool(), types.NewRole(nil, true))
}

func (tc *typeChecker) VisitExprIdent(ctx *parser.ExprIdentContext) any {
	sym, err := tc.currentScope.LookupSymbol(ctx.Ident())
	if err != nil {
		tc.reportError(err)
		return types.Invalid()
	}
	return sym.Type()
}

func (tc *typeChecker) VisitExprNum(ctx *parser.ExprNumContext) any {
	_, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		tc.reportError(types.NewInvalidNumberError(ctx))
	}

	return types.New(types.Int(), types.NewRole(nil, true))
}

func (tc *typeChecker) VisitExprAwait(ctx *parser.ExprAwaitContext) any {
	panic("unimplemented")
}
