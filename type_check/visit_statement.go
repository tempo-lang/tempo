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
		tc.checkRolesInScope(ctx.Expr(), exprType.Roles())
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

		if !sym.IsAssignable() {
			tc.reportError(types.NewUnassignableSymbolError(ctx, sym.Type()))
		} else {
			exprType := ctx.Expr().Accept(tc).(*types.Type)
			if !exprType.CanCoerceTo(sym.Type()) {
				tc.reportError(types.NewInvalidAssignTypeError(ctx, sym.Type(), exprType))
			}
		}

		tc.checkRolesInScope(ctx.Ident(), sym.Type().Roles())
	}

	return nil
}

func (tc *typeChecker) VisitStmtIf(ctx *parser.StmtIfContext) any {
	guardType := ctx.Expr().Accept(tc).(*types.Type)

	if !types.ValueCoerseTo(guardType.Value(), types.Bool()) {
		tc.reportError(types.NewInvalidValueError(ctx.Expr(), guardType.Value(), types.Bool()))
	}

	tc.checkRolesInScope(ctx.Expr(), guardType.Roles())

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
