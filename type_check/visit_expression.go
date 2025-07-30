package type_check

import (
	"slices"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

func (tc *typeChecker) visitExpr(ctx parser.IExprContext) types.Type {
	if ctx == nil {
		return tc.registerType(ctx, types.Invalid())
	}
	exprType := ctx.Accept(tc)
	if exprType == nil {
		return tc.registerType(ctx, types.Invalid())
	}
	return exprType.(types.Type)
}

func (tc *typeChecker) registerType(expr parser.IExprContext, exprType types.Type) types.Type {
	tc.info.Types[expr] = exprType
	return exprType
}

func (tc *typeChecker) VisitExprBinOp(ctx *parser.ExprBinOpContext) any {
	lhs := tc.visitExpr(ctx.GetLhs())
	rhs := tc.visitExpr(ctx.GetRhs())

	if lhs.IsInvalid() || rhs.IsInvalid() {
		return tc.registerType(ctx, types.Invalid())
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

	mergedType, ok := tc.mergeTypes(ctx, lhs, rhs)
	if !ok {
		return tc.registerType(ctx, types.Invalid())
	}

	isBuiltinTypeOf := func(allowedTypes []types.BuiltinType) bool {
		if slices.Contains(allowedTypes, types.BuiltinKind(mergedType)) {
			return true
		} else {
			tc.reportError(type_error.NewBinOpIncompatibleType(ctx, lhs, allowedTypes))
			return false
		}
	}

	numberTypes := []types.BuiltinType{types.BuiltinInt, types.BuiltinFloat}

	// concatenate lists
	if op == projection.OpAdd {
		if _, ok := mergedType.(*types.ListType); ok {
			return tc.registerType(ctx, mergedType)
		}
	}

	switch {
	case slices.Contains(arithmeticOps, op):
		allowedTypes := []types.BuiltinType{types.BuiltinInt}

		if op != projection.OpMod {
			allowedTypes = append(allowedTypes, types.BuiltinFloat)
		}

		if op == projection.OpAdd {
			allowedTypes = append(allowedTypes, types.BuiltinString)
		}

		if ok := isBuiltinTypeOf(allowedTypes); !ok {
			typeError = true
		}

		newRoles, ok := types.RoleIntersect(lhs.Roles(), rhs.Roles())
		if !ok {
			tc.reportError(type_error.NewUnmergableRoles(ctx, []*types.Roles{lhs.Roles(), rhs.Roles()}))
			typeError = true
		}

		if !typeError {
			return tc.registerType(ctx, lhs.ReplaceSharedRoles(newRoles.Participants()))
		}
	case slices.Contains(equalityOps, op):
		if mergedType, ok := tc.mergeTypes(ctx, lhs, rhs); ok {
			if !mergedType.IsEquatable() {
				tc.reportError(type_error.NewUnequatableType(ctx, mergedType))
			}

			return tc.registerType(ctx, types.Bool(mergedType.Roles().Participants()))
		}
	case slices.Contains(inequalityOps, op):
		if ok := isBuiltinTypeOf(numberTypes); !ok {
			typeError = true
		}

		newRoles, ok := types.RoleIntersect(lhs.Roles(), rhs.Roles())
		if !ok {
			tc.reportError(type_error.NewUnmergableRoles(ctx, []*types.Roles{lhs.Roles(), rhs.Roles()}))
			typeError = true
		}

		if !typeError {
			return tc.registerType(ctx, types.Bool(newRoles.Participants()))
		}
	case slices.Contains(booleanOps, op):
		if ok := isBuiltinTypeOf([]types.BuiltinType{types.BuiltinBool}); !ok {
			typeError = true
		}

		newRoles, ok := types.RoleIntersect(lhs.Roles(), rhs.Roles())
		if !ok {
			tc.reportError(type_error.NewUnmergableRoles(ctx, []*types.Roles{lhs.Roles(), rhs.Roles()}))
			typeError = true
		}

		if !typeError {
			return tc.registerType(ctx, types.Bool(newRoles.Participants()))
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
	_, isFunc := sym.Type().(*types.FunctionType)

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

	if asyncType, isAsync := exprType.(*types.AsyncType); isAsync {
		return tc.registerType(ctx, asyncType.Inner())
	}

	tc.reportError(type_error.NewAwaitNonAsyncType(ctx, exprType))
	return tc.registerType(ctx, types.Invalid())
}

func (tc *typeChecker) VisitExprCom(ctx *parser.ExprComContext) any {

	innerExprType := tc.visitExpr(ctx.Expr())

	invalidType := false
	invalidRole := false

	if !innerExprType.IsSendable() {
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

		recvType = types.Async(innerExprType).ReplaceSharedRoles(newParticipants)
	}

	if !invalidType && invalidRole {
		recvType = types.Async(innerExprType).ReplaceSharedRoles(nil)
	}

	return tc.registerType(ctx, recvType)
}

func (tc *typeChecker) VisitExprCall(ctx *parser.ExprCallContext) any {

	callType := tc.visitExpr(ctx.Expr())
	if callType.IsInvalid() {
		return tc.registerType(ctx, types.Invalid())
	}

	switch callFuncValue := callType.(type) {
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

	defStructType, ok := sym.Type().(*types.StructType)
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

	stSym, ok := sym.(*sym_table.StructSymbol)
	if !ok {
		// parser error
		return tc.registerType(ctx, types.Invalid())
	}

	// check that all struct fields are present
	for _, defField := range stSym.Fields() {
		found := false
		for _, exprField := range ctx.ExprStructField().AllIdent() {
			if exprField.GetText() == defField.SymbolName() {
				found = true
				break
			}
		}

		if !found {
			tc.reportError(type_error.NewMissingStructField(ctx, defField.SymbolName(), defStructType))
		}
	}

	// calculate role substitution map
	structRoles := defStructType.Roles()
	defRoleSubst, ok := structRoles.SubstituteMap(roles)
	if !ok {
		tc.reportError(type_error.NewWrongRoleCount(sym, ctx.RoleType(), roles))
		return tc.registerType(ctx, types.Invalid())
	}

	structType := defStructType.SubstituteRoles(defRoleSubst)

	fieldIdents := ctx.ExprStructField().AllIdent()
	for i, fieldExpr := range ctx.ExprStructField().AllExpr() {
		// check that field exists in struct
		fieldType, fieldFound := tc.info.Field(structType, fieldIdents[i].GetText())
		if !fieldFound {
			tc.reportError(type_error.NewUnexpectedStructField(fieldIdents[i], defStructType))
			continue
		}

		var fieldExprType types.Type
		if fieldFound {
			oldHint := tc.currentTypeHint
			tc.currentTypeHint = fieldType
			fieldExprType = tc.visitExpr(fieldExpr)
			tc.currentTypeHint = oldHint
		} else {
			fieldExprType = tc.visitExpr(fieldExpr)
		}

		// check that field expression can coerce to the expected type
		if _, ok := fieldExprType.CoerceTo(fieldType); !ok {
			tc.reportError(type_error.NewIncompatibleTypes(fieldExpr, fieldExprType, fieldType))
			continue
		}
	}

	return tc.registerType(ctx, structType)
}

func (tc *typeChecker) VisitExprStructField(ctx *parser.ExprStructFieldContext) any {
	return nil
}

func (tc *typeChecker) VisitExprFieldAccess(ctx *parser.ExprFieldAccessContext) any {

	baseType := tc.visitExpr(ctx.Expr())
	if baseType.IsInvalid() {
		return tc.registerType(ctx, types.Invalid())
	}

	fieldName := ctx.Ident().GetText()
	fieldType, found := tc.info.Field(baseType, fieldName)
	if !found {
		tc.reportError(type_error.NewFieldAccessUnknownField(ctx.Ident(), baseType))
		return tc.registerType(ctx, types.Invalid())
	}

	return tc.registerType(ctx, fieldType)
}

func (tc *typeChecker) VisitExprIndex(ctx *parser.ExprIndexContext) any {
	baseType := tc.visitExpr(ctx.GetBaseExpr())

	listType, isList := baseType.(*types.ListType)
	if !isList {
		tc.reportError(type_error.NewIndexWrongBaseType(ctx.GetBaseExpr(), baseType))
		return tc.registerType(ctx, types.Invalid())
	}

	innerType := listType.Inner()

	indexTypeRaw := tc.visitExpr(ctx.GetIndexExpr())
	indexType, ok := indexTypeRaw.CoerceTo(types.Int(indexTypeRaw.Roles().Participants()))
	if !ok {
		tc.reportError(type_error.NewInvalidValue(ctx.GetIndexExpr(), indexTypeRaw, types.Int(indexTypeRaw.Roles().Participants())))
		return tc.registerType(ctx, types.Invalid())
	}

	// find roles
	if innerType.Roles().IsDistributedRole() {
		for _, role := range innerType.Roles().Participants() {
			if !indexType.Roles().Contains(role) {
				tc.reportError(type_error.NewIndexRoleNotEncompassBase(ctx, innerType, indexType.Roles()))
				return tc.registerType(ctx, types.Invalid())
			}
		}

		return tc.registerType(ctx, innerType)
	} else {
		intersectingRoles, ok := types.RoleIntersect(innerType.Roles(), indexType.Roles())
		if !ok {
			tc.reportError(type_error.NewIndexRoleNotEncompassBase(ctx, innerType, indexType.Roles()))
			return tc.registerType(ctx, types.Invalid())
		}

		return tc.registerType(ctx, innerType.ReplaceSharedRoles(intersectingRoles.Participants()))
	}
}

func (tc *typeChecker) VisitExprList(ctx *parser.ExprListContext) any {
	if ctx == nil {
		return tc.registerType(ctx, types.Invalid())
	}

	if len(ctx.AllExpr()) == 0 {
		if tc.currentTypeHint != nil {
			if _, isList := tc.currentTypeHint.(*types.ListType); isList {
				return tc.registerType(ctx, tc.currentTypeHint)
			}
		}

		tc.reportError(type_error.NewUnknownType(ctx))
		return tc.registerType(ctx, types.Invalid())
	}

	var exprType types.Type = nil
	for _, expr := range ctx.AllExpr() {
		if exprType == nil {
			exprType = tc.visitExpr(expr)
		} else {
			nextType := tc.visitExpr(expr)
			newType, _ := tc.mergeTypes(expr, exprType, nextType)
			exprType = newType
		}
	}

	return tc.registerType(ctx, types.List(exprType))
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

	var returnType types.Type = types.Unit()
	if sig.GetReturnType() != nil {
		returnType = tc.visitValueType(sig.GetReturnType())
		tc.checkRolesInScope(findRoleType(sig.GetReturnType()))
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
	if !returnsValue && returnType != types.Unit() {
		tc.reportError(type_error.NewFunctionMissingReturn(closureEnv))
	}

	// exit scope
	tc.currentScope = tc.currentScope.Parent()

	paramTypes := []types.Type{}
	for _, param := range closureEnv.Params() {
		paramTypes = append(paramTypes, param.Type())
	}

	closureType := types.Closure(paramTypes, returnType, closureRoles.Participants())
	return tc.registerType(ctx, closureType)
}

func (tc *typeChecker) VisitClosureSig(ctx *parser.ClosureSigContext) any {
	return nil
}
