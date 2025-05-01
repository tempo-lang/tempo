package types

import (
	"chorego/parser"
	"fmt"
	"slices"
)

func BuiltinValues() map[string]Value {
	return map[string]Value{
		"Int":    Int(),
		"Float":  Float(),
		"String": String(),
		"Bool":   Bool(),
	}
}

func ParseValueType(ctx parser.IValueTypeContext) (*Type, Error) {

	role, err := ParseRoleType(ctx.RoleType())
	if err != nil {
		return nil, err
	}

	typeName := ctx.Ident()
	builtinType, isBuiltinType := BuiltinValues()[typeName.GetText()]
	if isBuiltinType {
		return New(builtinType, role), nil
	}

	return nil, NewUnknownTypeError(typeName)
}

func ParseFuncType(ctx parser.IFuncContext) (*Type, Error) {

	funcRoles, err := ParseRoleTypeNormal(ctx.RoleTypeNormal())
	if err != nil {
		return nil, err
	}

	params := []*Type{}
	paramErrors := map[int][]Error{}

	for i, param := range ctx.FuncParamList().AllFuncParam() {
		paramErrors[i] = []Error{}

		paramType, err := ParseValueType(param.ValueType())
		if err != nil {
			paramErrors[i] = append(paramErrors[i], err)
		}

		var roleIdents []parser.IIdentContext
		if roleNormal := param.ValueType().RoleType().RoleTypeNormal(); roleNormal != nil {
			roleIdents = roleNormal.AllIdent()
		}
		if roleShared := param.ValueType().RoleType().RoleTypeShared(); roleShared != nil {
			roleIdents = roleShared.AllIdent()
		}

		for _, r := range roleIdents {
			foundRole := slices.ContainsFunc(funcRoles.participants, func(role string) bool {
				return role == r.GetText()
			})
			if !foundRole {
				paramErrors[i] = append(paramErrors[i], NewUnknownRoleError(ctx, r))
			}
		}

		params = append(params, paramType)
	}

	fn := New(Function(params), funcRoles)

	for _, errList := range paramErrors {
		if len(errList) > 0 {
			return fn, NewInvalidFuncError(ctx, paramErrors)
		}
	}

	return fn, nil
}

func ParseRoleType(ctx parser.IRoleTypeContext) (*Roles, Error) {
	if roleNormal := ctx.RoleTypeNormal(); roleNormal != nil {
		return ParseRoleTypeNormal(roleNormal)
	}
	if roleShared := ctx.RoleTypeShared(); roleShared != nil {
		return ParseRoleTypeShared(roleShared)
	}
	panic(fmt.Sprintf("unexpected role type: %v", ctx))
}

func ParseRoleTypeNormal(ctx parser.IRoleTypeNormalContext) (*Roles, Error) {
	participants := []string{}
	for _, role := range ctx.AllIdent() {
		participants = append(participants, role.GetText())
	}
	return NewRole(participants, false), nil
}

func ParseRoleTypeShared(ctx parser.IRoleTypeSharedContext) (*Roles, Error) {
	participants := []string{}
	for _, role := range ctx.AllIdent() {
		participants = append(participants, role.GetText())
	}
	return NewRole(participants, false), nil
}
