package type_check

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

// func BuiltinValues() map[string]types.Value {
// 	return map[string]types.Value{
// 		"Int":    types.Int(),
// 		"Float":  types.Float(),
// 		"String": types.String(),
// 		"Bool":   types.Bool(),
// 	}
// }

func ToBuiltinValue(name string, participants []string) (types.Value, bool) {
	switch name {
	case "Int":
		return types.Int(participants), true
	case "Float":
		return types.Float(participants), true
	case "String":
		return types.String(participants), true
	case "Bool":
		return types.Bool(participants), true
	}

	return nil, false
}

func (tc *typeChecker) parseValueType(ctx parser.IValueTypeContext) (types.Value, type_error.Error) {
	// parser error
	if ctx == nil {
		return types.Invalid(), nil
	}

	switch ctx := ctx.(type) {
	case *parser.NamedTypeContext:
		return tc.parseNamedValueType(ctx)
	case *parser.ListTypeContext:
		return tc.parseListValueType(ctx)
	case *parser.AsyncTypeContext:
		return tc.parseAsyncValueType(ctx)
	case *parser.ClosureTypeContext:
		return tc.parseClosureValueType(ctx)
	case *parser.ValueTypeContext:
		// parser error
		return types.Invalid(), nil
	}
	panic(fmt.Sprintf("parseValueType unexpected ctx: %T", ctx))
}

func (tc *typeChecker) parseAsyncValueType(ctx *parser.AsyncTypeContext) (types.Value, type_error.Error) {
	inner, err := tc.parseValueType(ctx.GetInner())
	if err != nil {
		return inner, err
	}

	if _, isInnerAsync := inner.(*types.Async); isInnerAsync {
		tc.reportError(type_error.NewNestedAsync(ctx))

		// recoverable error
		return inner, nil
	}

	return types.NewAsync(inner), nil
}

func (tc *typeChecker) parseClosureValueType(ctx *parser.ClosureTypeContext) (types.Value, type_error.Error) {
	// parser error
	if ctx.RoleType() == nil || ctx.GetParams() == nil {
		return types.Invalid(), nil
	}

	roles, ok := tc.parseRoleType(ctx.RoleType())
	if !ok {
		return types.Invalid(), nil
	}

	params := []types.Value{}
	for _, param := range ctx.GetParams().AllValueType() {
		paramType, err := tc.parseValueType(param)
		if err != nil {
			return types.Invalid(), err
		}
		params = append(params, paramType)
	}

	var returnType types.Value = types.Unit()
	if ctx.GetReturnType() != nil {
		ret, err := tc.parseValueType(ctx.GetReturnType())
		if err != nil {
			return types.Invalid(), err
		}
		returnType = ret
	}

	closureValue := types.Closure(params, returnType, roles.Participants())

	return closureValue, nil
}

func (tc *typeChecker) parseNamedValueType(ctx *parser.NamedTypeContext) (types.Value, type_error.Error) {
	// parser error
	if ctx == nil || ctx.RoleType() == nil || ctx.Ident() == nil {
		return types.Invalid(), nil
	}

	role, ok := tc.parseRoleType(ctx.RoleType())
	if !ok {
		return types.Invalid(), nil
	}

	typeName := ctx.Ident()
	if builtinType, isBuiltinType := ToBuiltinValue(typeName.GetText(), role.Participants()); isBuiltinType {
		if !role.IsSharedRole() && len(role.Participants()) > 1 {
			return types.Invalid(), type_error.NewNotDistributedType(ctx)
		}
		return builtinType, nil
	}

	var typeError type_error.Error = nil
	sym, err := tc.lookupSymbol(typeName)
	if err != nil {
		return types.Invalid(), err
	}

	sym.AddRead(ctx.Ident())

	expectedRoleCount := len(sym.Type().Roles().Participants())
	if len(role.Participants()) != expectedRoleCount {
		typeError = type_error.NewWrongRoleCount(sym, ctx.RoleType(), role)
	}

	substMap, rolesMatch := sym.Type().Roles().SubstituteMap(role)
	if !rolesMatch {
		return types.Invalid(), type_error.NewWrongRoleCount(sym, ctx.RoleType(), role)
	}
	typeValue := sym.Type().SubstituteRoles(substMap)

	return typeValue, typeError
}

func (tc *typeChecker) parseListValueType(ctx *parser.ListTypeContext) (types.Value, type_error.Error) {
	return types.Invalid(), nil
}

func (tc *typeChecker) parseFuncType(ctx parser.IFuncSigContext) (types.Value, []type_error.Error) {
	errors := []type_error.Error{}

	funcRoles, ok := tc.parseRoleType(ctx.RoleType())
	if !ok {
		return types.Invalid(), errors
	}

	if funcRoles.IsSharedRole() {
		tc.reportError(type_error.NewUnexpectedSharedType(ctx.RoleType()))
	}

	params := []types.Value{}

	for _, param := range ctx.FuncParamList().AllFuncParam() {

		paramType, err := tc.parseValueType(param.ValueType())
		if err != nil {
			errors = append(errors, err)
		}

		if unknownRoles := paramType.Roles().SubtractParticipants(funcRoles.Participants()); len(unknownRoles) > 0 {
			errors = append(errors, type_error.NewRolesNotInScope(findRoleType(param.ValueType()), unknownRoles))
		}

		params = append(params, paramType)
	}

	returnType := types.Unit()
	if ctx.GetReturnType() != nil {
		var err type_error.Error
		returnType, err = tc.parseValueType(ctx.GetReturnType())
		if err != nil {
			errors = append(errors, err)
			return types.Invalid(), errors
		}
	}

	fn := types.Function(ctx.Ident(), params, returnType, funcRoles.Participants())

	return fn, errors
}

func (tc *typeChecker) parseRoleType(ctx parser.IRoleTypeContext) (*types.Roles, bool) {
	switch ctx := ctx.(type) {
	case *parser.RoleTypeNormalContext:
		return tc.parseRoleTypeNormal(ctx), true
	case *parser.RoleTypeSharedContext:
		return tc.parseRoleTypeShared(ctx), true
	}

	//panic(fmt.Sprintf("unexpected role type: %#v", ctx))
	return types.EveryoneRole(), false
}

func (tc *typeChecker) parseRoleTypeNormal(ctx *parser.RoleTypeNormalContext) *types.Roles {
	participants := []string{}
	for _, role := range ctx.AllIdent() {
		participants = append(participants, role.GetText())
	}
	return types.NewRole(participants, false)
}

func (tc *typeChecker) parseRoleTypeShared(ctx *parser.RoleTypeSharedContext) *types.Roles {
	participants := []string{}
	for _, role := range ctx.AllIdent() {
		participants = append(participants, role.GetText())
	}

	if len(participants) == 1 {
		tc.reportError(type_error.NewSharedRoleSingleParticipant(ctx))
	}

	return types.NewRole(participants, true)
}

func (tc *typeChecker) parseStructType(ctx parser.IStructContext) (types.Value, type_error.Error) {
	// parser error
	if ctx == nil || ctx.Ident() == nil || ctx.RoleType() == nil {
		return types.Invalid(), nil
	}

	roles, ok := tc.parseRoleType(ctx.RoleType())
	if !ok {
		return types.Invalid(), nil
	}

	if err := tc.checkDuplicateRoles(ctx.RoleType(), roles); err != nil {
		tc.reportError(err)
	}

	return types.NewStructType(ctx.Ident(), roles.Participants()), nil
}

func (tc *typeChecker) parseInterfaceType(ctx parser.IInterfaceContext) (types.Value, type_error.Error) {
	// parser error
	if ctx == nil || ctx.Ident() == nil || ctx.RoleType() == nil {
		return types.Invalid(), nil
	}

	// name := ctx.Ident().GetText()
	roles, ok := tc.parseRoleType(ctx.RoleType())
	if !ok {
		return types.Invalid(), nil
	}

	if err := tc.checkDuplicateRoles(ctx.RoleType(), roles); err != nil {
		tc.reportError(err)
	}

	return types.NewInterfaceType(ctx.Ident(), roles.Participants()), nil
}

func findRoleType(ctx parser.IValueTypeContext) parser.IRoleTypeContext {
	// parser error
	if ctx == nil {
		return nil
	}

	switch ctx := ctx.(type) {
	case *parser.AsyncTypeContext:
		return findRoleType(ctx.GetInner())
	case *parser.ClosureTypeContext:
		return ctx.RoleType()
	case *parser.ListTypeContext:
		return findRoleType(ctx.GetInner())
	case *parser.NamedTypeContext:
		return ctx.RoleType()
	case *parser.ValueTypeContext:
		// parser error
		return nil
	}

	panic(fmt.Sprintf("findRoleType unknown type: %T", ctx))
}
