package types

import (
	"fmt"
	"tempo/parser"
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

	isAsync := ctx.ASYNC() != nil

	role, err := ParseRoleType(ctx.RoleType())
	if err != nil {
		return Invalid(), err
	}

	typeName := ctx.Ident()
	var typeValue Value = nil

	if builtinType, isBuiltinType := BuiltinValues()[typeName.GetText()]; isBuiltinType {
		typeValue = builtinType
		if !role.isSharedRole && len(role.participants) > 1 {
			return Invalid(), NewNotDistributedTypeError(ctx)
		}
	}

	if typeValue != nil {
		if isAsync {
			typeValue = NewAsync(typeValue)
		}
		return New(typeValue, role), nil
	} else {
		return Invalid(), NewUnknownTypeError(typeName)
	}
}

func ParseFuncType(ctx parser.IFuncContext) (*Type, Error) {

	funcRoles, err := ParseRoleType(ctx.RoleType())
	if err != nil {
		return nil, err
	}

	params := []*Type{}
	paramErrors := [][]Error{}

	for i, param := range ctx.FuncParamList().AllFuncParam() {
		paramErrors = append(paramErrors, []Error{})

		paramType, err := ParseValueType(param.ValueType())
		if err != nil {
			paramErrors[i] = append(paramErrors[i], err)
		}

		paramRoles, err := ParseRoleType(param.ValueType().RoleType())
		if err != nil {
			paramErrors[i] = append(paramErrors[i], err)
		}

		if unknownRoles := paramRoles.SubtractParticipants(funcRoles.participants); len(unknownRoles) > 0 {
			paramErrors[i] = append(paramErrors[i], NewRolesNotInScopeError(param.ValueType().RoleType(), unknownRoles))
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
	switch ctx := ctx.(type) {
	case *parser.RoleTypeNormalContext:
		return ParseRoleTypeNormal(ctx)
	case *parser.RoleTypeSharedContext:
		return ParseRoleTypeShared(ctx)
	}

	panic(fmt.Sprintf("unexpected role type: %#v", ctx))
}

func ParseRoleTypeNormal(ctx *parser.RoleTypeNormalContext) (*Roles, Error) {
	participants := []string{}
	for _, role := range ctx.AllIdent() {
		participants = append(participants, role.GetText())
	}
	return NewRole(participants, false), nil
}

func ParseRoleTypeShared(ctx *parser.RoleTypeSharedContext) (*Roles, Error) {
	participants := []string{}
	for _, role := range ctx.AllIdent() {
		participants = append(participants, role.GetText())
	}
	return NewRole(participants, true), nil
}
