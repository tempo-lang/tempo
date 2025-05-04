package type_check

import (
	"chorego/parser"
	"chorego/sym_table"
	"chorego/types"
)

func (tc *typeChecker) VisitStmt(ctx *parser.StmtContext) any {
	if varDecl := ctx.StmtVarDecl(); varDecl != nil {
		return varDecl.Accept(tc)
	}

	panic("unexpected statement")
}

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
		tc.reportError(err)
		stmtType = types.Invalid()
	} else if !exprType.CanCoerceTo(declType) {
		tc.reportError(types.NewInvalidDeclTypeError(ctx.ValueType(), declType, ctx.Expr(), exprType))
		stmtType = types.Invalid()
	}

	tc.insertSymbol(sym_table.NewVariableSymbol(ctx, tc.currentScope, stmtType))

	return tc.VisitChildren(ctx)
}
