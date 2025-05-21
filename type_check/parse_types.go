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

func ParseValueType(ctx parser.IValueTypeContext) (*types.Type, type_error.Error) {
	// parser error
	if ctx == nil || ctx.RoleType() == nil {
		return types.Invalid(), nil
	}

	isAsync := ctx.ASYNC() != nil

	role, err := ParseRoleType(ctx.RoleType())
	if err != nil {
		return types.Invalid(), err
	} else if role == nil {
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

	if typeValue != nil {
		if isAsync {
			typeValue = types.NewAsync(typeValue)
		}
		return types.New(typeValue, role), nil
	} else {
		return types.Invalid(), type_error.NewUnknownTypeError(typeName)
	}
}

func ParseFuncType(ctx parser.IFuncContext) (*types.Type, []type_error.Error) {
	errors := []type_error.Error{}

	funcRoles, err := ParseRoleType(ctx.RoleType())
	if err != nil {
		errors = append(errors, err)
		return types.Invalid(), errors
	} else if funcRoles == nil {
		return types.Invalid(), nil
	}

	params := []*types.Type{}

	for _, param := range ctx.FuncParamList().AllFuncParam() {

		paramType, err := ParseValueType(param.ValueType())
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
		returnType, err = ParseValueType(ctx.ValueType())
		if err != nil {
			errors = append(errors, err)
			return types.Invalid(), errors
		}
	}

	fn := types.New(types.Function(params, returnType), funcRoles)

	return fn, errors
}

func ParseRoleType(ctx parser.IRoleTypeContext) (*types.Roles, type_error.Error) {
	switch ctx := ctx.(type) {
	case *parser.RoleTypeNormalContext:
		return ParseRoleTypeNormal(ctx)
	case *parser.RoleTypeSharedContext:
		return ParseRoleTypeShared(ctx)
	}

	//panic(fmt.Sprintf("unexpected role type: %#v", ctx))
	return nil, nil
}

func ParseRoleTypeNormal(ctx *parser.RoleTypeNormalContext) (*types.Roles, type_error.Error) {
	participants := []string{}
	for _, role := range ctx.AllIdent() {
		participants = append(participants, role.GetText())
	}
	return types.NewRole(participants, false), nil
}

func ParseRoleTypeShared(ctx *parser.RoleTypeSharedContext) (*types.Roles, type_error.Error) {
	participants := []string{}
	for _, role := range ctx.AllIdent() {
		participants = append(participants, role.GetText())
	}
	return types.NewRole(participants, true), nil
}

func ParseStructType(ctx parser.IStructContext) (*types.Type, type_error.Error) {
	// parser error
	if ctx == nil || ctx.Ident() == nil || ctx.RoleType() == nil {
		return types.Invalid(), nil
	}

	name := ctx.Ident().GetText()
	roles, err := ParseRoleType(ctx.RoleType())
	if err != nil {
		return nil, err
	}

	return types.New(types.NewStructType(name), roles), nil
}
