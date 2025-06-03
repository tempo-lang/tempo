package type_check

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

func (tc *typeChecker) VisitStmtVarDecl(ctx *parser.StmtVarDeclContext) any {
	hasExplicitType := ctx.ValueType() != nil

	exprType := tc.visitExpr(ctx.Expr())
	stmtType := exprType

	roleFailed := false

	if !tc.checkExprInScope(ctx.Expr(), exprType.Roles()) {
		roleFailed = true
	}

	if hasExplicitType {
		typeFailed := false
		declType, err := tc.parseValueType(ctx.ValueType())
		if err != nil {
			tc.reportError(err)
			typeFailed = true
		} else if !declType.IsInvalid() {
			if !tc.checkRolesInScope(ctx.ValueType().RoleType()) {
				roleFailed = true
			}
		}

		if !typeFailed {
			newType, ok := exprType.CoerceTo(declType)
			if !ok {
				tc.reportError(type_error.NewInvalidDeclTypeError(ctx.ValueType(), declType, ctx.Expr(), exprType))
				typeFailed = true
			} else {
				stmtType = newType
			}
		}

		if typeFailed {
			stmtType = types.Invalid()
		} else if roleFailed {
			stmtType = types.New(declType.Value(), types.EveryoneRole())
		}
	}

	if len(stmtType.Roles().Participants()) == 0 {
		newParticipants := tc.currentScope.Roles().Participants()
		stmtType = types.New(
			stmtType.Value(),
			types.NewRole(newParticipants, true),
		)
	}

	tc.insertSymbol(sym_table.NewVariableSymbol(ctx, tc.currentScope, stmtType))

	return false
}

func (tc *typeChecker) VisitStmtAssign(ctx *parser.StmtAssignContext) any {
	sym, err := tc.lookupSymbol(ctx.Ident())
	if err != nil {
		tc.reportError(err)
		return nil
	}

	tc.info.Symbols[ctx.Ident()] = sym

	if !sym.IsAssignable() {
		tc.reportError(type_error.NewUnassignableSymbolError(ctx, sym.Type()))
	} else {
		exprType := tc.visitExpr(ctx.Expr())
		if _, ok := exprType.CoerceTo(sym.Type()); !ok {
			tc.reportError(type_error.NewInvalidAssignTypeError(ctx, sym.Type(), exprType))
		}
	}

	tc.checkExprInScope(ctx.Ident(), sym.Type().Roles())

	return false
}

func (tc *typeChecker) VisitStmtIf(ctx *parser.StmtIfContext) any {
	guardType := tc.visitExpr(ctx.Expr())

	if _, ok := guardType.Value().CoerceTo(types.Bool()); !ok {
		tc.reportError(type_error.NewInvalidValueError(ctx.Expr(), guardType.Value(), types.Bool()))
	}

	tc.checkExprInScope(ctx.Expr(), guardType.Roles())

	scopeRoles := guardType.Roles().Participants()

	// Then branch
	thenBranch := ctx.GetThenScope()
	tc.currentScope = tc.currentScope.MakeChild(thenBranch.GetStart(), thenBranch.GetStop(), scopeRoles)
	thenReturn := thenBranch.Accept(tc)
	tc.currentScope = tc.currentScope.Parent()

	if elseBranch := ctx.GetElseScope(); elseBranch != nil {
		// Else branch
		tc.currentScope = tc.currentScope.MakeChild(elseBranch.GetStart(), elseBranch.GetStop(), scopeRoles)
		elseReturn := elseBranch.Accept(tc)
		tc.currentScope = tc.currentScope.Parent()

		if elseReturn == false || thenReturn == false {
			return false
		}
	}

	return thenReturn == true
}

func (tc *typeChecker) VisitStmtExpr(ctx *parser.StmtExprContext) any {
	tc.visitExpr(ctx.Expr())
	return false
}

func (tc *typeChecker) VisitStmtReturn(ctx *parser.StmtReturnContext) any {

	funcSym := tc.currentScope.GetCallableEnv()
	expectedReturnType := funcSym.ReturnType()

	missingRoles := funcSym.Scope().Roles().
		SubtractParticipants(tc.currentScope.Roles().Participants())
	if len(missingRoles) > 0 {
		tc.reportError(type_error.NewReturnNotAllRolesError(ctx, missingRoles))
	}

	if ctx.Expr() == nil {
		if expectedReturnType.Value() != types.Unit() {
			tc.reportError(type_error.NewReturnValueMissing(funcSym, ctx))
		}
		return true
	}

	returnType := tc.visitExpr(ctx.Expr())

	if _, ok := returnType.CoerceTo(expectedReturnType); !ok {
		tc.reportError(type_error.NewIncompatibleTypesError(ctx.Expr(), returnType, expectedReturnType))
	}

	return true
}
