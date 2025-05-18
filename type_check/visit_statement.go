package type_check

import (
	"tempo/parser"
	"tempo/sym_table"
	"tempo/types"
)

func (tc *typeChecker) VisitStmtVarDecl(ctx *parser.StmtVarDeclContext) any {
	exprType := tc.visitExpr(ctx.Expr())
	declType, err := types.ParseValueType(ctx.ValueType())
	typeFailed := false
	roleFailed := false

	if err != nil {
		tc.reportError(err)
		typeFailed = true
	} else if !declType.IsInvalid() {
		if !tc.checkRolesInScope(ctx.ValueType().RoleType()) {
			roleFailed = true
		}
	}

	if !tc.checkExprInScope(ctx.Expr(), exprType.Roles()) {
		roleFailed = true
	}

	if !typeFailed && !exprType.CanCoerceTo(declType) {
		tc.reportError(types.NewInvalidDeclTypeError(ctx.ValueType(), declType, ctx.Expr(), exprType))
		typeFailed = true
	}

	stmtType := declType
	if typeFailed {
		stmtType = types.Invalid()
	} else if roleFailed {
		stmtType = types.New(declType.Value(), types.EveryoneRole())
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

		if !sym.IsAssignable() {
			tc.reportError(types.NewUnassignableSymbolError(ctx, sym.Type()))
		} else {
			exprType := ctx.Expr().Accept(tc).(*types.Type)
			if !exprType.CanCoerceTo(sym.Type()) {
				tc.reportError(types.NewInvalidAssignTypeError(ctx, sym.Type(), exprType))
			}
		}

		tc.checkExprInScope(ctx.Ident(), sym.Type().Roles())
	}

	return nil
}

func (tc *typeChecker) VisitStmtIf(ctx *parser.StmtIfContext) any {
	guardType := ctx.Expr().Accept(tc).(*types.Type)

	if !types.ValueCoerseTo(guardType.Value(), types.Bool()) {
		tc.reportError(types.NewInvalidValueError(ctx.Expr(), guardType.Value(), types.Bool()))
	}

	tc.checkExprInScope(ctx.Expr(), guardType.Roles())

	scopeRoles := guardType.Roles().Participants()

	// Then branch
	thenBranch := ctx.Scope(0)
	tc.currentScope = tc.currentScope.MakeChild(thenBranch.GetStart(), thenBranch.GetStop(), scopeRoles)
	for _, stmt := range thenBranch.AllStmt() {
		stmt.Accept(tc)
	}
	tc.currentScope = tc.currentScope.Parent()

	if elseBranch := ctx.Scope(1); elseBranch != nil {
		// Else branch
		tc.currentScope = tc.currentScope.MakeChild(elseBranch.GetStart(), elseBranch.GetStop(), scopeRoles)
		for _, stmt := range elseBranch.AllStmt() {
			stmt.Accept(tc)
		}
		tc.currentScope = tc.currentScope.Parent()
	}

	return nil
}

func (tc *typeChecker) VisitStmtExpr(ctx *parser.StmtExprContext) any {
	ctx.Expr().Accept(tc)
	return nil
}

func (tc *typeChecker) VisitStmtReturn(ctx *parser.StmtReturnContext) any {
	returnType := ctx.Expr().Accept(tc).(*types.Type)

	expectedReturnType := tc.currentScope.GetFunc().FuncValue().ReturnType()

	if !returnType.CanCoerceTo(expectedReturnType) {
		tc.reportError(types.NewIncompatibleTypesError(ctx.Expr(), returnType, expectedReturnType))
	}

	missingRoles := tc.currentScope.GetFunc().Roles().
		SubtractParticipants(tc.currentScope.Roles().Participants())

	if len(missingRoles) > 0 {
		tc.reportError(types.NewReturnNotAllRolesError(ctx, missingRoles))
	}

	return nil
}
