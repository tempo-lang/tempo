package type_check

import (
	"slices"
	"strconv"
	"tempo/parser"
	"tempo/projection"
	"tempo/sym_table"
	"tempo/type_check/type_error"
	"tempo/types"
)

func (tc *typeChecker) visitExpr(ctx parser.IExprContext) *types.Type {
	if ctx == nil {
		return types.Invalid()
	}
	exprType := ctx.Accept(tc)
	if exprType == nil {
		return types.Invalid()
	}
	return exprType.(*types.Type)
}

func (tc *typeChecker) registerType(expr parser.IExprContext, exprType *types.Type) *types.Type {
	tc.info.Types[expr] = exprType
	return exprType
}

func (tc *typeChecker) VisitExprBinOp(ctx *parser.ExprBinOpContext) any {
	lhs := tc.visitExpr(ctx.Expr(0))
	rhs := tc.visitExpr(ctx.Expr(1))

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
			tc.reportError(type_error.NewInvalidValueError(ctx.Expr(0), lhs.Value(), types.Int()))
			typeError = true
		}

		if !types.ValueCoerseTo(rhs.Value(), types.Int()) {
			tc.reportError(type_error.NewInvalidValueError(ctx.Expr(1), rhs.Value(), types.Int()))
			typeError = true
		}

		newRoles, ok := types.RoleIntersect(lhs.Roles(), rhs.Roles())
		if !ok {
			tc.reportError(type_error.NewUnmergableRolesError(ctx, []*types.Roles{lhs.Roles(), rhs.Roles()}))
			typeError = true
		}

		if op == projection.OpDiv || op == projection.OpMod {
			if numExpr, ok := ctx.Expr(1).(*parser.ExprNumContext); ok {
				if num, err := strconv.Atoi(numExpr.NUMBER().GetText()); err == nil {
					if num == 0 {
						tc.reportError(type_error.NewDivisionByZeroError(ctx.Expr(1)))
					}
				}
			}
		}

		if !typeError {
			return tc.registerType(ctx, types.New(types.Int(), newRoles))
		}

	case slices.Contains(equalityOps, op):
		if !lhs.Value().IsEquatable() {
			tc.reportError(type_error.NewUnequatableTypeError(ctx, lhs.Value()))
		} else if !rhs.Value().IsEquatable() {
			tc.reportError(type_error.NewUnequatableTypeError(ctx, rhs.Value()))
		} else {
			if !types.ValuesEqual(lhs.Value(), rhs.Value()) {
				tc.reportError(type_error.NewValueMismatchError(ctx, lhs.Value(), rhs.Value()))
				typeError = true
			}
		}

		newRoles, ok := types.RoleIntersect(lhs.Roles(), rhs.Roles())
		if !ok {
			tc.reportError(type_error.NewUnmergableRolesError(ctx, []*types.Roles{lhs.Roles(), rhs.Roles()}))
			typeError = true
		}

		if !typeError {
			return tc.registerType(ctx, types.New(types.Bool(), newRoles))
		}
	case slices.Contains(inequalityOps, op):
		if !types.ValueCoerseTo(lhs.Value(), types.Int()) {
			tc.reportError(type_error.NewInvalidValueError(ctx.Expr(0), lhs.Value(), types.Int()))
			typeError = true
		}

		if !types.ValueCoerseTo(rhs.Value(), types.Int()) {
			tc.reportError(type_error.NewInvalidValueError(ctx.Expr(1), rhs.Value(), types.Int()))
			typeError = true
		}

		newRoles, ok := types.RoleIntersect(lhs.Roles(), rhs.Roles())
		if !ok {
			tc.reportError(type_error.NewUnmergableRolesError(ctx, []*types.Roles{lhs.Roles(), rhs.Roles()}))
			typeError = true
		}

		if !typeError {
			return tc.registerType(ctx, types.New(types.Bool(), newRoles))
		}

	case slices.Contains(booleanOps, op):
		if !types.ValueCoerseTo(lhs.Value(), types.Bool()) {
			tc.reportError(type_error.NewInvalidValueError(ctx.Expr(0), lhs.Value(), types.Bool()))
			typeError = true
		}

		if !types.ValueCoerseTo(rhs.Value(), types.Bool()) {
			tc.reportError(type_error.NewInvalidValueError(ctx.Expr(1), rhs.Value(), types.Bool()))
			typeError = true
		}

		newRoles, ok := types.RoleIntersect(lhs.Roles(), rhs.Roles())
		if !ok {
			tc.reportError(type_error.NewUnmergableRolesError(ctx, []*types.Roles{lhs.Roles(), rhs.Roles()}))
			typeError = true
		}

		if !typeError {
			return tc.registerType(ctx, types.New(types.Bool(), newRoles))
		}
	}

	return tc.registerType(ctx, types.Invalid())
}

func (tc *typeChecker) VisitExprGroup(ctx *parser.ExprGroupContext) any {
	innerType := tc.visitExpr(ctx.Expr())
	return tc.registerType(ctx, innerType)
}

func (tc *typeChecker) VisitExprBool(ctx *parser.ExprBoolContext) any {
	return tc.registerType(ctx, types.New(types.Bool(), types.NewRole(nil, true)))
}

func (tc *typeChecker) VisitExprIdent(ctx *parser.ExprIdentContext) any {
	sym, err := tc.lookupSymbol(ctx.Ident())
	if err != nil {
		tc.reportError(err)
		return tc.registerType(ctx, types.Invalid())
	}

	tc.info.Symbols[ctx.Ident()] = sym

	tc.checkExprInScope(ctx, sym.Type().Roles())

	return tc.registerType(ctx, sym.Type())
}

func (tc *typeChecker) VisitExprNum(ctx *parser.ExprNumContext) any {
	_, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		tc.reportError(type_error.NewInvalidNumberError(ctx))
	}

	return tc.registerType(ctx, types.New(types.Int(), types.NewRole(nil, true)))
}

func (tc *typeChecker) VisitExprAwait(ctx *parser.ExprAwaitContext) any {

	exprType := tc.visitExpr(ctx.Expr())

	if exprType.IsInvalid() {
		return tc.registerType(ctx, types.Invalid())
	}

	if asyncType, isAsync := exprType.Value().(*types.Async); isAsync {
		return tc.registerType(ctx, types.New(asyncType.Inner(), exprType.Roles()))
	}

	tc.reportError(type_error.NewExpectedAsyncTypeError(ctx, exprType))
	return tc.registerType(ctx, types.Invalid())
}

func (tc *typeChecker) VisitExprCom(ctx *parser.ExprComContext) any {

	innerExprType := tc.visitExpr(ctx.Expr())

	invalidType := false
	invalidRole := false

	if !innerExprType.Value().IsSendable() {
		tc.reportError(type_error.NewUnsendableTypeError(ctx, innerExprType))
		invalidType = true
	}

	if fromRoles, ok := ParseRoleType(ctx.RoleType(0)); ok {
		if fromRoles.IsSharedRole() {
			tc.reportError(type_error.NewComSharedTypeError(ctx, innerExprType))
		}

		if tc.checkRolesInScope(ctx.RoleType(0)) {
			tc.checkExprInScope(ctx, fromRoles)
		}

		if len(fromRoles.Participants()) > 1 {
			tc.reportError(type_error.NewComDistributedTypeError(ctx, innerExprType))
		}

		exprHasParticipants := len(innerExprType.Roles().Participants()) > 0
		if exprHasParticipants && !innerExprType.Roles().Contains(fromRoles.Participants()[0]) {
			tc.reportError(type_error.NewComValueNotAtSenderError(ctx, innerExprType))
		}
	} else {
		invalidRole = true
	}

	toRoles, ok := ParseRoleType(ctx.RoleType(1))
	if ok {
		if !tc.checkRolesInScope(ctx.RoleType(1)) {
			invalidRole = true
		} else {
			if !tc.checkExprInScope(ctx, toRoles) {
				invalidRole = true
			}
		}
	} else {
		invalidRole = true
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
	sym, err := tc.lookupSymbol(ctx.Ident())
	if err != nil {
		tc.reportError(err)
		return tc.registerType(ctx, types.Invalid())
	}

	tc.info.Symbols[ctx.Ident()] = sym
	funcSym, ok := sym.(*sym_table.FuncSymbol)
	if !ok {
		tc.reportError(type_error.NewCallNonFunctionError(ctx, sym.Type()))
		return tc.registerType(ctx, types.Invalid())
	}

	callRoles, ok := ParseRoleType(ctx.RoleType())
	if !ok {
		return tc.registerType(ctx, types.Invalid())
	}

	funcRoles := parser.RoleTypeAllIdents(funcSym.Func().RoleType())
	if len(callRoles.Participants()) != len(funcRoles) {
		tc.reportError(type_error.NewCallWrongRoleCountError(ctx))
		return tc.registerType(ctx, types.Invalid())
	}

	roleSubst := map[string]string{}
	for i, callRole := range callRoles.Participants() {
		funcRole := funcRoles[i].GetText()
		roleSubst[funcRole] = callRole
	}

	argExprs := ctx.FuncArgList().AllExpr()
	if len(funcSym.Params()) != len(argExprs) {
		tc.reportError(type_error.NewCallWrongArgCountError(ctx))
		invalidType = true
	} else {
		for i, param := range funcSym.Params() {
			arg := argExprs[i]

			argType := tc.visitExpr(arg)

			paramTypeSubst := param.Type().SubstituteRoles(roleSubst)

			if !argType.CanCoerceTo(paramTypeSubst) {
				tc.reportError(type_error.NewIncompatibleTypesError(arg, argType, paramTypeSubst))
				invalidType = true
			}
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

func (tc *typeChecker) VisitExprStruct(ctx *parser.ExprStructContext) any {
	if ctx.ExprStructField() == nil {
		return tc.registerType(ctx, types.Invalid())
	}

	sym, err := tc.lookupSymbol(ctx.Ident())
	if err != nil {
		tc.reportError(err)
		return tc.registerType(ctx, types.Invalid())
	}

	structSym, ok := sym.(*sym_table.StructSymbol)
	if !ok {
		tc.reportError(type_error.NewExpectedStructTypeError(sym, ctx))
		return tc.registerType(ctx, types.Invalid())
	}

	tc.info.Symbols[ctx.Ident()] = sym

	roles, ok := ParseRoleType(ctx.RoleType())
	if !ok {
		return tc.registerType(ctx, types.Invalid())
	}

	structRoles := structSym.Type().Roles().Participants()
	if len(roles.Participants()) != len(structRoles) {
		tc.reportError(type_error.NewStructWrongRoleCountError(ctx.RoleType(), structSym))
		return tc.registerType(ctx, types.Invalid())
	}

	for _, defField := range structSym.Fields() {
		found := false
		for _, exprField := range ctx.ExprStructField().AllIdent() {
			if exprField.GetText() == defField.SymbolName() {
				found = true
				break
			}
		}

		if !found {
			tc.reportError(type_error.NewMissingStructFieldError(ctx, defField))
		}
	}

	defRoleSubst := map[string]string{}
	for i, role := range roles.Participants() {
		defRoleSubst[structRoles[i]] = role
	}

	defFields := map[string]*types.Type{}
	for _, field := range structSym.Fields() {
		defFields[field.SymbolName()] = field.Type()
	}

	fieldIdents := ctx.ExprStructField().AllIdent()

	exprFieldsExpr := map[string]parser.IExprContext{}
	exprFieldsType := map[string]*types.Type{}
	exprFieldsIdent := map[string]parser.IIdentContext{}
	for i, field := range ctx.ExprStructField().AllExpr() {
		fieldType := tc.visitExpr(field)
		exprFieldsExpr[fieldIdents[i].GetText()] = field
		exprFieldsType[fieldIdents[i].GetText()] = fieldType
		exprFieldsIdent[fieldIdents[i].GetText()] = fieldIdents[i]
	}

	foundError := false

	for name, exprType := range exprFieldsType {
		defField, found := defFields[name]
		if !found {
			tc.reportError(type_error.NewUnexpectedStructFieldError(exprFieldsIdent[name], structSym))
			foundError = true
			continue
		}

		defSubstType := defField.SubstituteRoles(defRoleSubst)

		if !exprType.CanCoerceTo(defSubstType) {
			tc.reportError(type_error.NewIncompatibleTypesError(exprFieldsExpr[name], exprType, defField))
			foundError = true
			continue
		}
	}

	if foundError {
		return tc.registerType(ctx, types.Invalid())
	}

	structType := structSym.Type().SubstituteRoles(defRoleSubst)
	return tc.registerType(ctx, structType)
}

func (tc *typeChecker) VisitExprStructField(ctx *parser.ExprStructFieldContext) any {
	return nil
}

func (tc *typeChecker) VisitExprFieldAccess(ctx *parser.ExprFieldAccessContext) any {

	objectType := tc.visitExpr(ctx.Expr())
	if objectType.IsInvalid() {
		return tc.registerType(ctx, types.Invalid())
	}

	if structType, ok := objectType.Value().(*types.StructType); ok {
		structSym := tc.currentScope.LookupParent(structType.Name()).(*sym_table.StructSymbol)

		var field *sym_table.StructFieldSymbol = nil
		for _, f := range structSym.Fields() {
			if ctx.Ident().GetText() == f.SymbolName() {
				field = f
				break
			}
		}

		if field == nil {
			tc.reportError(type_error.NewFieldAccessUnknownField(ctx, structSym))
			return tc.registerType(ctx, types.Invalid())
		}

		substMap, rolesMatch := structSym.Type().Roles().SubstituteMap(objectType.Roles())
		if !rolesMatch {
			return tc.registerType(ctx, types.Invalid())
		}

		fieldType := field.Type().SubstituteRoles(substMap)
		return tc.registerType(ctx, fieldType)
	}

	tc.reportError(type_error.NewFieldAccessUnexpectedType(ctx, objectType))
	return tc.registerType(ctx, types.Invalid())
}
