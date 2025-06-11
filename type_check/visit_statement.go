package type_check

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

func (tc *typeChecker) VisitStmtVarDecl(ctx *parser.StmtVarDeclContext) any {
	hasExplicitType := ctx.ValueType() != nil
	typeFailed := false
	roleFailed := false
	var declType types.Type = nil

	prevTypeHint := tc.currentTypeHint

	if hasExplicitType {
		var err type_error.Error
		declType, err = tc.parseValueType(ctx.ValueType())
		if err != nil {
			tc.reportError(err)
			typeFailed = true
		} else if !declType.IsInvalid() {
			if !tc.checkRolesInScope(findRoleType(ctx.ValueType())) {
				roleFailed = true
			}

			tc.currentTypeHint = declType
		}
	}

	exprType := tc.visitExpr(ctx.Expr())
	tc.currentTypeHint = prevTypeHint

	// convert function type to closure type when stored in a variable
	if funcType, ok := exprType.(*types.FunctionType); ok {
		exprType = types.Closure(funcType.Params(), funcType.ReturnType(), funcType.Roles().Participants())
	}

	stmtType := exprType

	if !tc.checkExprInScope(ctx.Expr(), exprType.Roles()) {
		roleFailed = true
	}

	if hasExplicitType {
		if !typeFailed {
			newType, ok := exprType.CoerceTo(declType)
			if !ok {
				tc.reportError(type_error.NewInvalidAssignType(ctx.Expr(), declType, exprType))
				typeFailed = true
			} else {
				stmtType = newType
			}
		}

		if typeFailed {
			stmtType = types.Invalid()
		} else if roleFailed {
			stmtType = declType.ReplaceSharedRoles(nil)
		}
	} else {
		if _, isUnit := exprType.(*types.UnitType); isUnit {
			tc.reportError(type_error.NewAssignUnitValue(ctx.Expr()))
			stmtType = types.Invalid()
		}
	}

	if len(stmtType.Roles().Participants()) == 0 {
		newParticipants := tc.currentScope.Roles().Participants()
		stmtType = stmtType.ReplaceSharedRoles(newParticipants)
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
	sym.AddWrite(ctx.Ident())

	if !sym.IsAssignable() {
		tc.reportError(type_error.NewUnassignableSymbol(ctx, sym.Type()))
	} else {
		exprType := tc.visitExpr(ctx.Expr())
		if _, ok := exprType.CoerceTo(sym.Type()); !ok {
			tc.reportError(type_error.NewInvalidAssignType(ctx.Expr(), sym.Type(), exprType))
		}
	}

	tc.checkExprInScope(ctx.Ident(), sym.Type().Roles())

	return false
}

func (tc *typeChecker) VisitStmtIf(ctx *parser.StmtIfContext) any {
	guardType := tc.visitExpr(ctx.Expr())

	boolType := types.Bool(guardType.Roles().Participants())
	if _, ok := guardType.CoerceTo(boolType); !ok {
		tc.reportError(type_error.NewInvalidValue(ctx.Expr(), guardType, boolType))
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

func (tc *typeChecker) VisitStmtWhile(ctx *parser.StmtWhileContext) any {
	condType := tc.visitExpr(ctx.Expr())

	if types.BuiltinKind(condType) != types.BuiltinBool {
		tc.reportError(type_error.NewInvalidValue(ctx.Expr(), condType, types.Bool(nil)))
	}

	tc.checkExprInScope(ctx.Expr(), condType.Roles())

	scopeRoles := condType.Roles().Participants()

	tc.currentScope = tc.currentScope.MakeChild(ctx.Scope().GetStart(), ctx.Scope().GetStop(), scopeRoles)
	ctx.Scope().Accept(tc)
	tc.currentScope = tc.currentScope.Parent()

	return false
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
		tc.reportError(type_error.NewReturnNotAllRoles(ctx, missingRoles))
	}

	if ctx.Expr() == nil {
		if expectedReturnType != types.Unit() {
			tc.reportError(type_error.NewReturnValueMissing(funcSym, ctx))
		}
		return true
	}

	returnType := tc.visitExpr(ctx.Expr())

	if _, ok := returnType.CoerceTo(expectedReturnType); !ok {
		tc.reportError(type_error.NewIncompatibleTypes(ctx.Expr(), returnType, expectedReturnType))
	}

	return true
}
