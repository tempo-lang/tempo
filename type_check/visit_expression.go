package type_check

import (
	"slices"

	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

func (tc *typeChecker) visitExpr(ctx ast.Expr) types.Type {
	if ctx == nil {
		return types.Invalid()
	}
	var result types.Type
	switch e := ctx.(type) {
	case *ast.InvalidExpr:
		result = types.Invalid()
	case *ast.PrimitiveExpr:
		return tc.visitPrimitive(e)
	case *ast.BinaryExpr:
		result = tc.visitBinary(e)
	case *ast.GroupExpr:
		result = tc.visitExpr(e.Expr)
	case *ast.IdentAccessExpr:
		result = tc.visitIdent(e)
	case *ast.AwaitExpr:
		result = tc.visitAwait(e)
	case *ast.ComExpr:
		result = tc.visitCom(e)
	case *ast.CallExpr:
		result = tc.visitCall(e)
	case *ast.FieldAccessExpr:
		result = tc.visitFieldAccess(e)
	case *ast.IndexExpr:
		result = tc.visitIndex(e)
	case *ast.ListExpr:
		result = tc.visitList(e)
	case *ast.StructExpr:
		result = tc.visitStructExpr(e)
	case *ast.ClosureExpr:
		result = tc.visitClosure(e)
	default:
		result = types.Invalid()
	}
	if result.Roles().IsHidden() {
		tc.reportError(type_error.NewHiddenExpression(ctx, result))
	}
	return tc.registerType(ctx, result)
}

func (tc *typeChecker) registerType(expr ast.Expr, exprType types.Type) types.Type {
	tc.info.Types[expr] = exprType
	return exprType
}

func (tc *typeChecker) visitBinary(ctx *ast.BinaryExpr) types.Type {
	lhs := tc.visitExpr(ctx.Left)
	rhs := tc.visitExpr(ctx.Right)

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
	op := projection.Operator(ctx.Operator.Text)

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

func (tc *typeChecker) visitIdent(ctx *ast.IdentAccessExpr) types.Type {
	sym, err := tc.lookupSymbol(ctx.Ident)
	if err != nil {
		tc.reportError(err)
		return tc.registerType(ctx, types.Invalid())
	}

	if _, ok := sym.(*sym_table.TypeSymbol); ok {
		tc.reportError(type_error.NewTypeNotAnExpression(ctx))
		return tc.registerType(ctx, types.Invalid())
	}

	tc.info.Symbols[ctx.Ident] = sym
	sym.AddRead(ctx.Ident)

	identType := sym.Type()

	if _, isStructDef := sym.(*sym_table.StructSymbol); isStructDef {
		tc.reportError(type_error.NewStructNotInitialized(ctx))
		return tc.registerType(ctx, types.Invalid())
	}

	if _, isFunc := identType.(*types.FunctionType); isFunc {
		if ctx.RoleType != nil {
			// function role substitution is defined explicitly

			roles, ok := tc.parseRoleType(ctx.RoleType)
			if !ok {
				return tc.registerType(ctx, types.Invalid())
			}

			roleSubst, ok := identType.Roles().SubstituteMap(roles)
			if !ok {
				tc.reportError(type_error.NewUnmergableRoles(ctx, []*types.Roles{identType.Roles(), roles}))
				return tc.registerType(ctx, types.Invalid())
			}

			identType = identType.SubstituteRoles(roleSubst)
		} else {
			scopeRoles := tc.currentScope.Roles()

			// if both definition and call scope is local role, do conversion automatically.
			if identType.Roles().IsLocalRole() && scopeRoles.IsLocalRole() {
				roleSubst, ok := identType.Roles().SubstituteMap(scopeRoles)
				if !ok {
					panic("two local roles should always be substitutable")
				}
				identType = identType.SubstituteRoles(roleSubst)
			} else {
				tc.reportError(type_error.NewFunctionNotInstantiated(ctx, sym))
				return tc.registerType(ctx, types.Invalid())
			}
		}
	} else {
		if ctx.RoleType != nil {
			tc.reportError(type_error.NewInstantiateNonFunction(ctx, sym))
			return tc.registerType(ctx, types.Invalid())
		}
	}

	identType = tc.coerceExprToScope(ctx, identType)

	return tc.registerType(ctx, identType)
}

func (tc *typeChecker) visitAwait(ctx *ast.AwaitExpr) types.Type {

	exprType := tc.visitExpr(ctx.Expr)

	if exprType.IsInvalid() {
		return tc.registerType(ctx, types.Invalid())
	}

	if asyncType, isAsync := exprType.(*types.AsyncType); isAsync {
		return tc.registerType(ctx, asyncType.Inner())
	}

	tc.reportError(type_error.NewAwaitNonAsyncType(ctx, exprType))
	return tc.registerType(ctx, types.Invalid())
}

func (tc *typeChecker) visitCom(ctx *ast.ComExpr) types.Type {

	innerExprType := tc.visitExpr(ctx.Expr)

	invalidType := false
	invalidRole := false

	if !tc.IsTypeSendable(innerExprType) {
		tc.reportError(type_error.NewUnsendableType(ctx, innerExprType))
		invalidType = true
	}

	if fromRoles, ok := tc.parseRoleType(ctx.Sender); ok {
		if !fromRoles.IsLocalRole() {
			tc.reportError(type_error.NewComNonLocalSender(ctx))
		}

		if tc.checkRolesInScope(ctx.Sender) {
			tc.checkExprInScope(ctx, fromRoles)
		}

		exprHasParticipants := len(innerExprType.Roles().Participants()) > 0
		if exprHasParticipants && !innerExprType.Roles().Contains(fromRoles.Participants()[0]) {
			tc.reportError(type_error.NewComValueNotAtSender(ctx, innerExprType))
		}
	} else {
		invalidRole = true
	}

	toRoles, ok := tc.parseRoleType(ctx.Receiver)
	if ok {
		if !tc.checkRolesInScope(ctx.Receiver) {
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

func (tc *typeChecker) visitCall(ctx *ast.CallExpr) types.Type {

	// Special case for type casting
	if exprIdent, isIdent := ctx.Function.(*ast.IdentAccessExpr); isIdent {
		if sym, err := tc.lookupSymbol(exprIdent.Ident); err == nil {
			if typeSym, ok := sym.(*sym_table.TypeSymbol); ok {
				args := ctx.Args
				if len(args) != 1 {
					tc.reportError(type_error.NewCallWrongArgCount(ctx, 1, len(args)))
					return tc.registerType(ctx, types.Invalid())
				}

				argType := tc.visitExpr(args[0])

				tc.info.Symbols[exprIdent.Ident] = typeSym
				sym.AddRead(exprIdent.Ident)

				validType := true
				switch typeSym.Type().(type) {
				case *types.FloatType:
					switch argType.(type) {
					case *types.FloatType:
					case *types.IntType:
					default:
						validType = false
					}
				case *types.IntType:
					switch argType.(type) {
					case *types.IntType:
					case *types.FloatType:
					default:
						validType = false
					}
				case *types.StringType:
					switch argType.(type) {
					case *types.StringType:
					case *types.IntType:
					case *types.FloatType:
					case *types.BoolType:
					default:
						validType = false
					}
				}

				if !validType {
					tc.reportError(type_error.NewIncompatibleTypeCast(argType, typeSym.Type(), ctx))
					return tc.registerType(ctx, types.Invalid())
				}

				newType := typeSym.Type().ReplaceSharedRoles(argType.Roles().Participants())

				if exprIdent.RoleType != nil {
					if explicitRoles, ok := tc.parseRoleType(exprIdent.RoleType); ok {
						newRoles, ok := types.RoleIntersect(newType.Roles(), explicitRoles)
						if !ok {
							tc.reportError(type_error.NewUnmergableRoles(ctx, []*types.Roles{newType.Roles(), explicitRoles}))
							return tc.registerType(ctx, types.Invalid())
						}

						newType = newType.ReplaceSharedRoles(newRoles.Participants())
					}
				}

				return tc.registerType(ctx, newType)
			}
		}
	}

	callType := tc.visitExpr(ctx.Function)
	if callType.IsInvalid() {
		return tc.registerType(ctx, types.Invalid())
	}

	switch callFuncValue := callType.(type) {
	case *types.FunctionType:
		funcParamCount := len(callFuncValue.Params())
		callArgCount := len(ctx.Args)
		if funcParamCount != callArgCount {
			tc.reportError(type_error.NewCallWrongArgCount(ctx, funcParamCount, callArgCount))
		} else {
			for i, arg := range ctx.Args {
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
		callArgCount := len(ctx.Args)
		if funcParamCount != callArgCount {
			tc.reportError(type_error.NewCallWrongArgCount(ctx, funcParamCount, callArgCount))
		} else {
			for i, arg := range ctx.Args {
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

func (tc *typeChecker) visitStructExpr(ctx *ast.StructExpr) types.Type {
	sym, err := tc.lookupSymbol(ctx.RoleIdent.Ident)
	if err != nil {
		tc.reportError(err)
		return tc.registerType(ctx, types.Invalid())
	}

	defStructType, ok := sym.Type().(*types.StructType)
	if !ok {
		tc.reportError(type_error.NewExpectedStructType(sym, ctx))
		return tc.registerType(ctx, types.Invalid())
	}

	tc.info.Symbols[ctx.RoleIdent.Ident] = sym
	sym.AddRead(ctx.RoleIdent.Ident)

	roles, ok := tc.parseRoleType(ctx.RoleIdent.RoleType)
	if !ok {
		return tc.registerType(ctx, types.Invalid())
	}

	stSym, ok := sym.(*sym_table.StructSymbol)
	if !ok {
		// parser error
		return tc.registerType(ctx, types.Invalid())
	}

	// check that no fields are duplicated
	fieldsCount := map[string]*ast.Identifier{}
	for _, field := range ctx.StructFields.Fields {
		fieldName := field.Name
		if first, found := fieldsCount[fieldName.Value()]; found {
			fieldsCount[fieldName.Value()] = nil
			if first != nil {
				tc.reportError(type_error.NewDuplicateStructField(ctx, first, fieldName))
			}
		} else {
			fieldsCount[fieldName.Value()] = fieldName
		}
	}

	// calculate role substitution map
	structRoles := defStructType.Roles()
	defRoleSubst, ok := structRoles.SubstituteMap(roles)
	if !ok {
		tc.reportError(type_error.NewWrongRoleCount(sym, ctx.RoleIdent, roles))
		return tc.registerType(ctx, types.Invalid())
	}

	structType := defStructType.SubstituteRoles(defRoleSubst)

	// check that all struct fields are present
	for _, defField := range stSym.Fields() {
		var exprFieldIdent *ast.Identifier
		for _, field := range ctx.StructFields.Fields {
			if field.Name.Value() == defField.SymbolName() {
				exprFieldIdent = field.Name
				break
			}
		}

		fieldType, ok := tc.info.Field(structType, defField.SymbolName())
		if !ok {
			panic("type should have same fields as definition")
		}

		if exprFieldIdent != nil && fieldType.Roles().IsHidden() {
			tc.reportError(type_error.NewHiddenStructField(ctx, exprFieldIdent, fieldType))
		}

		if exprFieldIdent == nil && !fieldType.Roles().IsHidden() {
			tc.reportError(type_error.NewMissingStructField(ctx, defField.SymbolName(), defStructType))
		}
	}

	for _, field := range ctx.StructFields.Fields {
		fieldIdent, fieldExpr := field.Name, field.Expr
		// check that field exists in struct
		fieldType, fieldFound := tc.info.Field(structType, fieldIdent.Value())
		if !fieldFound {
			tc.reportError(type_error.NewUnexpectedStructField(fieldIdent, defStructType))
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

func (tc *typeChecker) visitFieldAccess(ctx *ast.FieldAccessExpr) types.Type {

	baseType := tc.visitExpr(ctx.Object)
	if baseType.IsInvalid() {
		return tc.registerType(ctx, types.Invalid())
	}

	fieldName := ctx.Field.Value()
	fieldType, found := tc.info.Field(baseType, fieldName)
	if !found {
		tc.reportError(type_error.NewFieldAccessUnknownField(ctx.Field, baseType))
		return tc.registerType(ctx, types.Invalid())
	}

	switch fieldType.(type) {
	case *types.FunctionType, *types.ClosureType:
		if !fieldType.Roles().IsComplete() {
			tc.reportError(type_error.NewIncompleteFunction(ctx.Field, fieldType))
		}
	}

	return tc.registerType(ctx, fieldType)
}

func (tc *typeChecker) visitIndex(ctx *ast.IndexExpr) types.Type {
	baseType := tc.visitExpr(ctx.Object)

	listType, isList := baseType.(*types.ListType)
	if !isList {
		tc.reportError(type_error.NewIndexWrongBaseType(ctx.Object, baseType))
		return tc.registerType(ctx, types.Invalid())
	}

	innerType := listType.Inner()

	indexTypeRaw := tc.visitExpr(ctx.Index)
	indexType, ok := indexTypeRaw.CoerceTo(types.Int(indexTypeRaw.Roles().Participants()))
	if !ok {
		tc.reportError(type_error.NewInvalidValue(ctx.Index, indexTypeRaw, types.Int(indexTypeRaw.Roles().Participants())))
		return tc.registerType(ctx, types.Invalid())
	}

	indexRoles := indexType.Roles().Participants()
	if len(indexRoles) == 0 {
		indexRoles = tc.currentScope.Roles().Participants()
	}

	limitedInnerType, ok := limitTypeToRoles(innerType, indexRoles)
	if !ok {
		tc.reportError(type_error.NewIndexRoleNotEncompassBase(ctx, innerType, types.NewRole(indexRoles, true)))
		return tc.registerType(ctx, types.Invalid())
	}

	return tc.registerType(ctx, limitedInnerType)
}

func (tc *typeChecker) visitList(ctx *ast.ListExpr) types.Type {
	if len(ctx.Elements) == 0 {
		if tc.currentTypeHint != nil {
			if _, isList := tc.currentTypeHint.(*types.ListType); isList {
				return tc.registerType(ctx, tc.currentTypeHint)
			}
		}

		tc.reportError(type_error.NewUnknownType(ctx))
		return tc.registerType(ctx, types.Invalid())
	}

	var exprType types.Type = nil
	for _, expr := range ctx.Elements {
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

func (tc *typeChecker) visitClosure(ctx *ast.ClosureExpr) types.Type {
	sig := ctx.ClosureSig
	if sig == nil || sig.Params == nil {
		// parser error
		return tc.registerType(ctx, types.Invalid())
	}

	t, ok := tc.parseClosureType(sig)
	if !ok {
		return tc.registerType(ctx, types.Invalid())
	}
	closureType := t.(*types.ClosureType)

	// enter scope
	tc.currentScope = tc.currentScope.MakeChild(nodeSpan(ctx.Scope), closureType.Roles().Participants())
	closureEnv := sym_table.NewClosureEnv(tc.currentScope, closureType, sig.ReturnType)
	tc.currentScope.SetCallableEnv(closureEnv)

	// add params to scope
	tc.visitFuncParams(sig.Params)

	returnsValue := tc.visitScope(ctx.Scope)
	if !returnsValue && closureType.ReturnType() != types.Unit() {
		tc.reportError(type_error.NewFunctionMissingReturn(closureEnv))
	}

	// exit scope
	tc.currentScope = tc.currentScope.Parent()

	paramTypes := []types.Type{}
	for _, param := range closureEnv.Params() {
		paramTypes = append(paramTypes, param.Type())
	}

	return tc.registerType(ctx, closureType)
}
