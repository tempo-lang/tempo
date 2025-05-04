package type_check

import (
	"chorego/parser"
	"chorego/sym_table"
	"chorego/types"
)

func (tc *typeChecker) VisitStmtVarDecl(ctx *parser.StmtVarDeclContext) any {
	exprType := ctx.Expr().Accept(tc).(*types.Type)
	declType, err := types.ParseValueType(ctx.ValueType())
	if err != nil {
		tc.reportError(err)
	} else {
		activeRoles := tc.currentScope.Roles()

		if unknownRoles := declType.Roles().SubtractParticipants(activeRoles); len(unknownRoles) > 0 {
			tc.reportError(types.NewUnknownRoleError(ctx.ValueType(), unknownRoles))
		}
	}

	stmtType := declType

	if err != nil {
		stmtType = types.Invalid()
	} else if !exprType.CanCoerceTo(declType) {
		tc.reportError(types.NewInvalidDeclTypeError(ctx.ValueType(), declType, ctx.Expr(), exprType))
		stmtType = types.Invalid()
	}

	tc.insertSymbol(sym_table.NewVariableSymbol(ctx, tc.currentScope, stmtType))

	return nil
}

func (tc *typeChecker) VisitStmtAssign(ctx *parser.StmtAssignContext) any {
	sym, err := tc.currentScope.LookupSymbol(ctx.Ident())
	if err != nil {
		tc.reportError(err)
	} else {
		tc.info.Symbols[ctx.Ident()] = sym

		exprType := ctx.Expr().Accept(tc).(*types.Type)
		if !exprType.CanCoerceTo(sym.Type()) {
			tc.reportError(types.NewInvalidAssignTypeError(ctx, sym.Type(), exprType))
		}
	}

	return nil
}
