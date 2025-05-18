package type_check

import (
	"slices"
	"strconv"
	"tempo/parser"
	"tempo/projection"
	"tempo/sym_table"
	"tempo/types"
)

func (tc *typeChecker) registerType(expr parser.IExprContext, exprType *types.Type) *types.Type {
	tc.info.Types[expr] = exprType
	return exprType
}

func (tc *typeChecker) VisitExprBinOp(ctx *parser.ExprBinOpContext) any {
	lhs := ctx.Expr(0).Accept(tc).(*types.Type)
	rhs := ctx.Expr(1).Accept(tc).(*types.Type)

	if lhs.IsInvalid() || rhs.IsInvalid() {
		return types.Invalid()
	}

	arithmeticOps := []projection.Operator{
		projection.OpAdd,
		projection.OpSub,
		projection.OpMul,
		projection.OpDiv,
		projection.OpMod,
	}

	equalityOps := []projection.Operator{
		projection.OpEq,
		projection.OpNotEq,
	}

	inequalityOps := []projection.Operator{
		projection.OpLess,
		projection.OpLessEq,
		projection.OpGreater,
		projection.OpGreaterEq,
	}

	booleanOps := []projection.Operator{
		projection.OpAnd,
		projection.OpOr,
	}

	typeError := false
	op := projection.ParseOperator(ctx)

	switch {
	case slices.Contains(arithmeticOps, op):
		if !types.ValueCoerseTo(lhs.Value(), types.Int()) {
			tc.reportError(types.NewInvalidValueError(ctx.Expr(0), lhs.Value(), types.Int()))
			typeError = true
		}

		if !types.ValueCoerseTo(rhs.Value(), types.Int()) {
			tc.reportError(types.NewInvalidValueError(ctx.Expr(1), rhs.Value(), types.Int()))
			typeError = true
		}

		newRoles, err := types.RoleIntersect(ctx, lhs.Roles(), rhs.Roles())
		if err != nil {
			tc.reportError(err)
			typeError = true
		}

		if !typeError {
			return tc.registerType(ctx, types.New(types.Int(), newRoles))
		}

	case slices.Contains(equalityOps, op):
		if !types.ValuesEqual(lhs.Value(), rhs.Value()) {
			tc.reportError(types.NewValueMismatchError(ctx, lhs.Value(), rhs.Value()))
			typeError = true
		}

		newRoles, err := types.RoleIntersect(ctx, lhs.Roles(), rhs.Roles())
		if err != nil {
			tc.reportError(err)
			typeError = true
		}

		if !typeError {
			return tc.registerType(ctx, types.New(types.Bool(), newRoles))
		}

	case slices.Contains(inequalityOps, op):
		if !types.ValueCoerseTo(lhs.Value(), types.Int()) {
			tc.reportError(types.NewInvalidValueError(ctx.Expr(0), lhs.Value(), types.Int()))
			typeError = true
		}

		if !types.ValueCoerseTo(rhs.Value(), types.Int()) {
			tc.reportError(types.NewInvalidValueError(ctx.Expr(1), rhs.Value(), types.Int()))
			typeError = true
		}

		newRoles, err := types.RoleIntersect(ctx, lhs.Roles(), rhs.Roles())
		if err != nil {
			tc.reportError(err)
			typeError = true
		}

		if !typeError {
			return tc.registerType(ctx, types.New(types.Bool(), newRoles))
		}

	case slices.Contains(booleanOps, op):
		if !types.ValueCoerseTo(lhs.Value(), types.Bool()) {
			tc.reportError(types.NewInvalidValueError(ctx.Expr(0), lhs.Value(), types.Bool()))
			typeError = true
		}

		if !types.ValueCoerseTo(rhs.Value(), types.Bool()) {
			tc.reportError(types.NewInvalidValueError(ctx.Expr(1), rhs.Value(), types.Bool()))
			typeError = true
		}

		newRoles, err := types.RoleIntersect(ctx, lhs.Roles(), rhs.Roles())
		if err != nil {
			tc.reportError(err)
			typeError = true
		}

		if !typeError {
			return tc.registerType(ctx, types.New(types.Bool(), newRoles))
		}
	}

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

	tc.checkExprInScope(ctx, sym.Type().Roles())

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

	if exprType.IsInvalid() {
		return tc.registerType(ctx, types.Invalid())
	}

	if asyncType, isAsync := exprType.Value().(*types.Async); isAsync {
		return tc.registerType(ctx, types.New(asyncType.Inner(), exprType.Roles()))
	}

	tc.reportError(types.NewExpectedAsyncTypeError(ctx, exprType))
	return tc.registerType(ctx, types.Invalid())
}

func (tc *typeChecker) VisitExprCom(ctx *parser.ExprComContext) any {

	innerExprType := ctx.Expr().Accept(tc).(*types.Type)

	invalidType := false
	invalidRole := false

	if !innerExprType.Value().IsSendable() {
		tc.reportError(types.NewUnsendableTypeError(ctx, innerExprType))
		invalidType = true
	}

	fromRoles, err := types.ParseRoleType(ctx.RoleType(0))
	if err != nil {
		tc.reportError(err)
	} else {
		if fromRoles.IsSharedRole() {
			tc.reportError(types.NewComSharedTypeError(ctx, innerExprType))
		}

		if tc.checkRolesInScope(ctx.RoleType(0)) {
			tc.checkExprInScope(ctx, fromRoles)
		}

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
		if !tc.checkRolesInScope(ctx.RoleType(1)) {
			invalidRole = true
		} else {
			if !tc.checkExprInScope(ctx, toRoles) {
				invalidRole = true
			}
		}
	}

	recvType := types.Invalid()
	if !invalidType && !invalidRole {
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

	if !invalidType && invalidRole {
		recvType = types.New(types.NewAsync(innerExprType.Value()), types.EveryoneRole())
	}

	return tc.registerType(ctx, recvType)
}

func (tc *typeChecker) VisitExprCall(ctx *parser.ExprCallContext) any {
	invalidType := false
	sym, err := tc.currentScope.LookupSymbol(ctx.Ident())
	if err != nil {
		tc.reportError(err)
		return tc.registerType(ctx, types.Invalid())
	}

	tc.info.Symbols[ctx.Ident()] = sym
	funcSym, ok := sym.(*sym_table.FuncSymbol)
	if !ok {
		tc.reportError(types.NewCallNonFunctionError(ctx, sym.Type()))
		return tc.registerType(ctx, types.Invalid())
	}

	callRoles, err := types.ParseRoleType(ctx.RoleType())
	if err != nil {
		tc.reportError(err)
		return tc.registerType(ctx, types.Invalid())
	}

	funcRoles := parser.RoleTypeAllIdents(funcSym.Func().RoleType())
	if len(callRoles.Participants()) != len(funcRoles) {
		tc.reportError(types.NewCallWrongRoleCountError(ctx))
		return tc.registerType(ctx, types.Invalid())
	}

	roleSubst := map[string]string{}
	for i, callRole := range callRoles.Participants() {
		funcRole := funcRoles[i].GetText()
		roleSubst[funcRole] = callRole
	}

	argExprs := ctx.FuncArgList().AllExpr()
	if len(funcSym.Params()) != len(argExprs) {
		tc.reportError(types.NewCallWrongArgCountError(ctx))
		invalidType = true
	}

	for i, param := range funcSym.Params() {
		arg := argExprs[i]

		argType := arg.Accept(tc).(*types.Type)

		paramTypeSubst := param.Type().SubstituteRoles(roleSubst)

		if !argType.CanCoerceTo(paramTypeSubst) {
			tc.reportError(types.NewIncompatibleTypesError(arg, argType, paramTypeSubst))
			invalidType = true
		}
	}

	returnType := types.Invalid()
	if !invalidType {
		returnType = funcSym.FuncValue().ReturnType().SubstituteRoles(roleSubst)
	}

	return tc.registerType(ctx, returnType)
}

func (tc *typeChecker) VisitFuncArgList(ctx *parser.FuncArgListContext) any {
	return nil
}
