package type_check

import (
	"tempo/parser"
	"tempo/type_check/type_error"
	"tempo/types"
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
	if ctx == nil || ctx.RoleType() == nil {
		return types.Invalid(), nil
	}

	isAsync := ctx.ASYNC() != nil

	role, ok := ParseRoleType(ctx.RoleType())
	if !ok {
		return types.Invalid(), nil
	}

	typeName := ctx.Ident()
	var typeValue types.Value = nil

	if builtinType, isBuiltinType := BuiltinValues()[typeName.GetText()]; isBuiltinType {
		typeValue = builtinType
		if !role.IsSharedRole() && len(role.Participants()) > 1 {
			return types.Invalid(), type_error.NewNotDistributedTypeError(ctx)
		}
	}

	if typeValue == nil {
		sym, err := tc.lookupSymbol(typeName)
		if err != nil {
			return types.Invalid(), err
		}

		typeValue = sym.Type().Value()
	}

	if typeValue != nil {
		if isAsync {
			typeValue = types.NewAsync(typeValue)
		}
		return types.New(typeValue, role), nil
	} else {
		return types.Invalid(), type_error.NewUnknownTypeError(typeName)
	}
}

func (tc *typeChecker) parseFuncType(ctx parser.IFuncContext) (*types.Type, []type_error.Error) {
	errors := []type_error.Error{}

	funcRoles, ok := ParseRoleType(ctx.RoleType())
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
			errors = append(errors, type_error.NewRolesNotInScopeError(param.ValueType().RoleType(), unknownRoles))
		}

		params = append(params, paramType)
	}

	returnType := types.New(types.Unit(), types.EveryoneRole())
	if ctx.ValueType() != nil {
		var err type_error.Error
		returnType, err = tc.parseValueType(ctx.ValueType())
		if err != nil {
			errors = append(errors, err)
			return types.Invalid(), errors
		}
	}

	fn := types.New(types.Function(params, returnType), funcRoles)

	return fn, errors
}

func ParseRoleType(ctx parser.IRoleTypeContext) (*types.Roles, bool) {
	switch ctx := ctx.(type) {
	case *parser.RoleTypeNormalContext:
		return ParseRoleTypeNormal(ctx), true
	case *parser.RoleTypeSharedContext:
		return ParseRoleTypeShared(ctx), true
	}

	//panic(fmt.Sprintf("unexpected role type: %#v", ctx))
	return types.EveryoneRole(), false
}

func ParseRoleTypeNormal(ctx *parser.RoleTypeNormalContext) *types.Roles {
	participants := []string{}
	for _, role := range ctx.AllIdent() {
		participants = append(participants, role.GetText())
	}
	return types.NewRole(participants, false)
}

func ParseRoleTypeShared(ctx *parser.RoleTypeSharedContext) *types.Roles {
	participants := []string{}
	for _, role := range ctx.AllIdent() {
		participants = append(participants, role.GetText())
	}
	return types.NewRole(participants, true)
}

func ParseStructType(ctx parser.IStructContext) (*types.Type, type_error.Error) {
	// parser error
	if ctx == nil || ctx.Ident() == nil || ctx.RoleType() == nil {
		return types.Invalid(), nil
	}

	name := ctx.Ident().GetText()
	roles, ok := ParseRoleType(ctx.RoleType())
	if !ok {
		return types.Invalid(), nil
	}

	return types.New(types.NewStructType(name), roles), nil
}
