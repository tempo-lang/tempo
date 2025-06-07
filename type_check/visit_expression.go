package type_check

import (
	"slices"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
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
	lhs := tc.visitExpr(ctx.GetLhs())
	rhs := tc.visitExpr(ctx.GetRhs())

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

	numberTypes := []types.Value{types.Int(), types.Float()}

	areSameTypes := func(allowedTypes []types.Value) (types.Value, bool) {
		if types.ValuesEqual(lhs.Value(), rhs.Value()) {
			for _, allowed := range allowedTypes {
				if types.ValuesEqual(allowed, lhs.Value()) {
					return allowed, true
				}
			}
			tc.reportError(type_error.NewBinOpIncompatibleType(ctx, lhs.Value(), allowedTypes))
		} else {
			tc.reportError(type_error.NewValueMismatch(ctx, lhs.Value(), rhs.Value()))
		}
		return types.Invalid().Value(), false
	}

	switch {
	case slices.Contains(arithmeticOps, op):
		allowedTypes := []types.Value{types.Int()}

		if op != projection.OpMod {
			allowedTypes = append(allowedTypes, types.Float())
		}

		if op == projection.OpAdd {
			allowedTypes = append(allowedTypes, types.String())
		}

		value, ok := areSameTypes(allowedTypes)
		if !ok {
			typeError = true
		}

		newRoles, ok := types.RoleIntersect(lhs.Roles(), rhs.Roles())
		if !ok {
			tc.reportError(type_error.NewUnmergableRoles(ctx, []*types.Roles{lhs.Roles(), rhs.Roles()}))
			typeError = true
		}

		if !typeError {
			return tc.registerType(ctx, types.New(value, newRoles))
		}
	case slices.Contains(equalityOps, op):
		if !lhs.Value().IsEquatable() {
			tc.reportError(type_error.NewUnequatableType(ctx, lhs.Value()))
		} else if !rhs.Value().IsEquatable() {
			tc.reportError(type_error.NewUnequatableType(ctx, rhs.Value()))
		} else {
			if !types.ValuesEqual(lhs.Value(), rhs.Value()) {
				tc.reportError(type_error.NewValueMismatch(ctx, lhs.Value(), rhs.Value()))
				typeError = true
			}
		}

		newRoles, ok := types.RoleIntersect(lhs.Roles(), rhs.Roles())
		if !ok {
			tc.reportError(type_error.NewUnmergableRoles(ctx, []*types.Roles{lhs.Roles(), rhs.Roles()}))
			typeError = true
		}

		if !typeError {
			return tc.registerType(ctx, types.New(types.Bool(), newRoles))
		}
	case slices.Contains(inequalityOps, op):
		if _, ok := areSameTypes(numberTypes); !ok {
			typeError = true
		}

		newRoles, ok := types.RoleIntersect(lhs.Roles(), rhs.Roles())
		if !ok {
			tc.reportError(type_error.NewUnmergableRoles(ctx, []*types.Roles{lhs.Roles(), rhs.Roles()}))
			typeError = true
		}

		if !typeError {
			return tc.registerType(ctx, types.New(types.Bool(), newRoles))
		}
	case slices.Contains(booleanOps, op):
		if _, ok := areSameTypes([]types.Value{types.Bool()}); !ok {
			typeError = true
		}

		newRoles, ok := types.RoleIntersect(lhs.Roles(), rhs.Roles())
		if !ok {
			tc.reportError(type_error.NewUnmergableRoles(ctx, []*types.Roles{lhs.Roles(), rhs.Roles()}))
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

func (tc *typeChecker) VisitExprIdent(ctx *parser.ExprIdentContext) any {
	sym, err := tc.lookupSymbol(ctx.IdentAccess().Ident())
	if err != nil {
		tc.reportError(err)
		return tc.registerType(ctx, types.Invalid())
	}

	tc.info.Symbols[ctx.IdentAccess().Ident()] = sym
	sym.AddRead(ctx.IdentAccess().Ident())

	identType := sym.Type()
	_, isFunc := sym.Type().Value().(*types.FunctionType)

	if _, isStructDef := sym.(*sym_table.StructSymbol); isStructDef {
		tc.reportError(type_error.NewStructNotInitialized(ctx))
		return tc.registerType(ctx, types.Invalid())
	}

	if ctx.IdentAccess().RoleType() != nil {
		roles, ok := tc.parseRoleType(ctx.IdentAccess().RoleType())
		if !ok {
			return tc.registerType(ctx, types.Invalid())
		}

		if err := tc.checkDuplicateRoles(ctx.IdentAccess().RoleType(), roles); err != nil {
			tc.reportError(err)
		}

		if isFunc {
			roleSubst, ok := identType.Roles().SubstituteMap(roles)
			if !ok {
				tc.reportError(type_error.NewUnmergableRoles(ctx, []*types.Roles{identType.Roles(), roles}))
				return tc.registerType(ctx, types.Invalid())
			}

			identType = identType.SubstituteRoles(roleSubst)
		} else {
			tc.reportError(type_error.NewInstantiateNonFunction(ctx.IdentAccess(), sym))
			return tc.registerType(ctx, types.Invalid())
		}
	} else if isFunc {
		tc.reportError(type_error.NewFunctionNotInstantiated(ctx.IdentAccess(), sym))
		return tc.registerType(ctx, types.Invalid())
	}

	tc.checkExprInScope(ctx, identType.Roles())

	return tc.registerType(ctx, identType)
}

func (tc *typeChecker) VisitExprAwait(ctx *parser.ExprAwaitContext) any {

	exprType := tc.visitExpr(ctx.Expr())

	if exprType.IsInvalid() {
		return tc.registerType(ctx, types.Invalid())
	}

	if asyncType, isAsync := exprType.Value().(*types.Async); isAsync {
		return tc.registerType(ctx, types.New(asyncType.Inner(), exprType.Roles()))
	}

	tc.reportError(type_error.NewAwaitNonAsyncType(ctx, exprType))
	return tc.registerType(ctx, types.Invalid())
}

func (tc *typeChecker) VisitExprCom(ctx *parser.ExprComContext) any {

	innerExprType := tc.visitExpr(ctx.Expr())

	invalidType := false
	invalidRole := false

	if !innerExprType.Value().IsSendable() {
		tc.reportError(type_error.NewUnsendableType(ctx, innerExprType))
		invalidType = true
	}

	if fromRoles, ok := tc.parseRoleType(ctx.RoleType(0)); ok {
		if !fromRoles.IsLocalRole() {
			tc.reportError(type_error.NewComNonLocalSender(ctx))
		}

		if tc.checkRolesInScope(ctx.RoleType(0)) {
			tc.checkExprInScope(ctx, fromRoles)
		}

		exprHasParticipants := len(innerExprType.Roles().Participants()) > 0
		if exprHasParticipants && !innerExprType.Roles().Contains(fromRoles.Participants()[0]) {
			tc.reportError(type_error.NewComValueNotAtSender(ctx, innerExprType))
		}
	} else {
		invalidRole = true
	}

	toRoles, ok := tc.parseRoleType(ctx.RoleType(1))
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

	callType := tc.visitExpr(ctx.Expr())
	if callType.IsInvalid() {
		return tc.registerType(ctx, types.Invalid())
	}

	switch callFuncValue := callType.Value().(type) {
	case *types.FunctionType:
		funcParamCount := len(callFuncValue.Params())
		callArgCount := len(ctx.FuncArgList().AllExpr())
		if funcParamCount != callArgCount {
			tc.reportError(type_error.NewCallWrongArgCount(ctx, funcParamCount, callArgCount))
		} else {
			for i, arg := range ctx.FuncArgList().AllExpr() {
				argType := tc.visitExpr(arg)
				paramType := callFuncValue.Params()[i]

				if _, ok := argType.CoerceTo(paramType); !ok {
					tc.reportError(type_error.NewIncompatibleTypes(arg, argType, paramType))
				}
			}
		}
		return tc.registerType(ctx, callFuncValue.ReturnType())
	case *types.ClosureType:
		funcParamCount := len(callFuncValue.Params())
		callArgCount := len(ctx.FuncArgList().AllExpr())
		if funcParamCount != callArgCount {
			tc.reportError(type_error.NewCallWrongArgCount(ctx, funcParamCount, callArgCount))
		} else {
			for i, arg := range ctx.FuncArgList().AllExpr() {
				argType := tc.visitExpr(arg)
				paramType := callFuncValue.Params()[i]

				if _, ok := argType.CoerceTo(paramType); !ok {
					tc.reportError(type_error.NewIncompatibleTypes(arg, argType, paramType))
				}
			}
		}
		return tc.registerType(ctx, callFuncValue.ReturnType())
	default:
		tc.reportError(type_error.NewCallNonFunction(ctx, callType))
		return tc.registerType(ctx, types.Invalid())
	}
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
		tc.reportError(type_error.NewExpectedStructType(sym, ctx))
		return tc.registerType(ctx, types.Invalid())
	}

	tc.info.Symbols[ctx.Ident()] = sym
	sym.AddRead(ctx.Ident())

	roles, ok := tc.parseRoleType(ctx.RoleType())
	if !ok {
		return tc.registerType(ctx, types.Invalid())
	}

	if err := tc.checkDuplicateRoles(ctx.RoleType(), roles); err != nil {
		tc.reportError(err)
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
			tc.reportError(type_error.NewMissingStructField(ctx, defField))
		}
	}

	structRoles := structSym.Type().Roles()
	defRoleSubst, ok := structRoles.SubstituteMap(roles)
	if !ok {
		tc.reportError(type_error.NewWrongRoleCount(structSym, ctx.RoleType(), roles))
		return tc.registerType(ctx, types.Invalid())
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
			tc.reportError(type_error.NewUnexpectedStructField(exprFieldsIdent[name], structSym))
			foundError = true
			continue
		}

		defSubstType := defField.SubstituteRoles(defRoleSubst)

		if _, ok := exprType.CoerceTo(defSubstType); !ok {
			tc.reportError(type_error.NewIncompatibleTypes(exprFieldsExpr[name], exprType, defField))
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

	switch value := objectType.Value().(type) {
	case *types.StructType:
		structSym, ok := tc.currentScope.LookupParent(value.Name()).(*sym_table.StructSymbol)
		if !ok {
			return tc.registerType(ctx, types.Invalid())
		}

		field := structSym.Scope().Lookup(ctx.Ident().GetText())
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

	case *types.InterfaceType:
		infSym, ok := tc.currentScope.LookupParent(value.Name()).(*sym_table.InterfaceSymbol)
		if !ok {
			return tc.registerType(ctx, types.Invalid())
		}

		field := infSym.Scope().Lookup(ctx.Ident().GetText())
		if field == nil {
			tc.reportError(type_error.NewFieldAccessUnknownField(ctx, infSym))
			return tc.registerType(ctx, types.Invalid())
		}

		substMap, rolesMatch := field.Type().Roles().SubstituteMap(objectType.Roles())
		if !rolesMatch {
			return tc.registerType(ctx, types.Invalid())
		}

		fieldType := field.Type().SubstituteRoles(substMap)
		return tc.registerType(ctx, fieldType)
	}

	tc.reportError(type_error.NewFieldAccessUnexpectedType(ctx, objectType))
	return tc.registerType(ctx, types.Invalid())
}

func (tc *typeChecker) VisitIdentAccess(ctx *parser.IdentAccessContext) any {
	return nil
}

func (tc *typeChecker) VisitExprClosure(ctx *parser.ExprClosureContext) any {
	sig := ctx.ClosureSig()
	if sig == nil || sig.FuncParamList() == nil {
		// parser error
		return tc.registerType(ctx, types.Invalid())
	}

	var returnType *types.Type = types.New(types.Unit(), types.EveryoneRole())
	if sig.GetReturnType() != nil {
		returnType = tc.visitValueType(sig.GetReturnType())
		tc.checkRolesInScope(sig.GetReturnType().RoleType())
	}

	closureRoles, _ := tc.parseRoleType(sig.RoleType())
	if closureRoles.IsSharedRole() {
		tc.reportError(type_error.NewUnexpectedSharedType(sig.RoleType()))
	}

	// enter scope
	tc.currentScope = tc.currentScope.MakeChild(ctx.Scope().GetStart(), ctx.Scope().GetStop(), closureRoles.Participants())
	closureEnv := sym_table.NewClosureEnv(tc.currentScope, returnType, sig.GetReturnType())
	tc.currentScope.SetCallableEnv(closureEnv)

	// add params to scope
	sig.FuncParamList().Accept(tc)

	returnsValue := ctx.Scope().Accept(tc) == true
	if !returnsValue && returnType.Value() != types.Unit() {
		tc.reportError(type_error.NewFunctionMissingReturn(closureEnv))
	}

	// exit scope
	tc.currentScope = tc.currentScope.Parent()

	paramTypes := []*types.Type{}
	for _, param := range closureEnv.Params() {
		paramTypes = append(paramTypes, param.Type())
	}

	closureType := types.New(types.Closure(paramTypes, returnType), closureRoles)
	return tc.registerType(ctx, closureType)
}

func (tc *typeChecker) VisitClosureSig(ctx *parser.ClosureSigContext) any {
	return nil
}
