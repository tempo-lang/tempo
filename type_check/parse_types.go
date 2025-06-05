package type_check

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

func BuiltinValues() map[string]types.Value {
	return map[string]types.Value{
		"Int":    types.Int(),
		"Float":  types.Float(),
		"String": types.String(),
		"Bool":   types.Bool(),
	}
}

func (tc *typeChecker) parseValueType(ctx parser.IValueTypeContext) (*types.Type, type_error.Error) {
	// parser error
	if ctx == nil || ((ctx.RoleType() == nil || ctx.Ident() == nil) && ctx.ClosureType() == nil) {
		return types.Invalid(), nil
	}

	isAsync := ctx.ASYNC() != nil

	if closure := ctx.ClosureType(); closure != nil {
		closureType, err := tc.parseClosureTypeSig(closure)
		if err != nil {
			return types.Invalid(), err
		}

		if isAsync {
			return types.New(types.NewAsync(closureType.Value()), closureType.Roles()), nil
		}

		return closureType, nil
	}

	role, ok := tc.parseRoleType(ctx.RoleType())
	if !ok {
		return types.Invalid(), nil
	}

	typeName := ctx.Ident()
	var typeValue types.Value = nil

	if builtinType, isBuiltinType := BuiltinValues()[typeName.GetText()]; isBuiltinType {
		typeValue = builtinType
		if !role.IsSharedRole() && len(role.Participants()) > 1 {
			return types.Invalid(), type_error.NewNotDistributedType(ctx)
		}
	}

	if typeValue == nil {
		sym, err := tc.lookupSymbol(typeName)
		if err != nil {
			return types.Invalid(), err
		}

		typeValue = sym.Type().Value()
		sym.AddRead(ctx.Ident())
	}

	if typeValue != nil {
		if isAsync {
			typeValue = types.NewAsync(typeValue)
		}
		return types.New(typeValue, role), nil
	} else {
		return types.Invalid(), type_error.NewUnknownType(typeName)
	}
}

func (tc *typeChecker) parseClosureTypeSig(ctx parser.IClosureTypeContext) (*types.Type, type_error.Error) {
	// parser error
	if ctx.RoleType() == nil || ctx.GetParams() == nil {
		return types.Invalid(), nil
	}

	roles, ok := tc.parseRoleType(ctx.RoleType())
	if !ok {
		return types.Invalid(), nil
	}

	params := []*types.Type{}
	for _, param := range ctx.GetParams().AllValueType() {
		paramType, err := tc.parseValueType(param)
		if err != nil {
			return types.Invalid(), err
		}
		params = append(params, paramType)
	}

	var returnType *types.Type = types.New(types.Unit(), types.EveryoneRole())
	if ctx.GetReturnType() != nil {
		ret, err := tc.parseValueType(ctx.GetReturnType())
		if err != nil {
			return types.Invalid(), err
		}
		returnType = ret
	}

	closureValue := types.Closure(params, returnType)
	funcType := types.New(closureValue, roles)

	return funcType, nil
}

func (tc *typeChecker) parseFuncType(ctx parser.IFuncSigContext) (*types.Type, []type_error.Error) {
	errors := []type_error.Error{}

	funcRoles, ok := tc.parseRoleType(ctx.RoleType())
	if !ok {
		return types.Invalid(), errors
	}

	params := []*types.Type{}

	for _, param := range ctx.FuncParamList().AllFuncParam() {

		paramType, err := tc.parseValueType(param.ValueType())
		if err != nil {
			errors = append(errors, err)
		}

		if unknownRoles := paramType.Roles().SubtractParticipants(funcRoles.Participants()); len(unknownRoles) > 0 {
			errors = append(errors, type_error.NewRolesNotInScope(param.ValueType().RoleType(), unknownRoles))
		}

		params = append(params, paramType)
	}

	returnType := types.New(types.Unit(), types.EveryoneRole())
	if ctx.GetReturnType() != nil {
		var err type_error.Error
		returnType, err = tc.parseValueType(ctx.GetReturnType())
		if err != nil {
			errors = append(errors, err)
			return types.Invalid(), errors
		}
	}

	fn := types.New(types.Function(ctx.Ident(), params, returnType, funcRoles.Participants()), funcRoles)

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

func (tc *typeChecker) parseStructType(ctx parser.IStructContext) (*types.Type, type_error.Error) {
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

	return types.New(types.NewStructType(ctx.Ident(), roles.Participants()), roles), nil
}

func (tc *typeChecker) parseInterfaceType(ctx parser.IInterfaceContext) (*types.Type, type_error.Error) {
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

	return types.New(types.NewInterfaceType(ctx.Ident()), roles), nil
}
