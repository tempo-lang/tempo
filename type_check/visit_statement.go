package type_check

import (
	"tempo/parser"
	"tempo/sym_table"
	"tempo/types"
)

func (tc *typeChecker) VisitStmtVarDecl(ctx *parser.StmtVarDeclContext) any {
	exprType := ctx.Expr().Accept(tc).(*types.Type)
	declType, err := types.ParseValueType(ctx.ValueType())
	typeFailed := false
	roleFailed := false

	if err != nil {
		tc.reportError(err)
		typeFailed = true
	} else {
		if !tc.checkRolesExist(ctx.ValueType().RoleType()) {
			roleFailed = true
		}
	}

	if !tc.checkRolesInScope(ctx.Expr(), exprType.Roles()) {
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
