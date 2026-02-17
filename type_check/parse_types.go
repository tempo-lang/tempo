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

func ToBuiltinValue(name string, participants []string) (types.Type, bool) {
	// if len(participants) == 1 && participants[0] == "" {
	// 	participants = []string{}
	// }

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

func (tc *typeChecker) parseValueType(ctx parser.IValueTypeContext) (types.Type, type_error.Error) {
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

func (tc *typeChecker) parseAsyncValueType(ctx *parser.AsyncTypeContext) (types.Type, type_error.Error) {
	inner, err := tc.parseValueType(ctx.GetInner())
	if err != nil {
		return inner, err
	}

	if _, isInnerAsync := inner.(*types.AsyncType); isInnerAsync {
		tc.reportError(type_error.NewNestedAsync(ctx))

		// recoverable error
		return inner, nil
	}

	return types.Async(inner), nil
}

func (tc *typeChecker) parseClosureValueType(ctx *parser.ClosureTypeContext) (types.Type, type_error.Error) {
	// parser error
	if ctx.RoleType() == nil || ctx.GetParams() == nil {
		return types.Invalid(), nil
	}

	roles, ok := tc.parseRoleType(ctx.RoleType())
	if !ok {
		return types.Invalid(), nil
	}

	params := []types.Type{}
	for _, param := range ctx.GetParams().AllValueType() {
		paramType, err := tc.parseValueType(param)
		if err != nil {
			return types.Invalid(), err
		}
		params = append(params, paramType)

		if unknownRoles := paramType.Roles().SubtractParticipants(roles.Participants()); len(unknownRoles) > 0 {
			if roleType, found := parser.FindRoleType(param); found {
				return types.Invalid(), type_error.NewRolesNotInScope(roleType, unknownRoles)
			} else {
				return types.Invalid(), type_error.NewRolesNotInScope(param, unknownRoles)
			}
		}
	}

	var returnType types.Type = types.Unit()
	if ctx.GetReturnType() != nil {
		ret, err := tc.parseValueType(ctx.GetReturnType())
		if err != nil {
			return types.Invalid(), err
		}
		returnType = ret

		if unknownRoles := returnType.Roles().SubtractParticipants(roles.Participants()); len(unknownRoles) > 0 {
			if roleType, found := parser.FindRoleType(ctx.GetReturnType()); found {
				return types.Invalid(), type_error.NewRolesNotInScope(roleType, unknownRoles)
			} else {
				return types.Invalid(), type_error.NewRolesNotInScope(ctx.GetReturnType(), unknownRoles)
			}
		}
	}

	closureValue := types.Closure(params, returnType, roles)

	return closureValue, nil
}

func (tc *typeChecker) parseNamedValueType(ctx *parser.NamedTypeContext) (types.Type, type_error.Error) {
	// parser error
	if ctx == nil {
		return types.Invalid(), nil
	}

	role, ok := tc.parseRoleType(ctx.RoleIdent().RoleType())
	if !ok {
		return types.Invalid(), nil
	}

	if !tc.currentScope.Roles().IsUnnamedRole() && role.IsUnnamedRole() {
		return types.Invalid(), type_error.NewMissingRoles(&ctx.ValueTypeContext)
	}

	typeName := ctx.RoleIdent().Ident()
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

	sym.AddRead(ctx.RoleIdent().Ident())

	substMap, rolesMatch := sym.Type().Roles().SubstituteMap(role)
	if !rolesMatch {
		return types.Invalid(), type_error.NewWrongRoleCount(sym, ctx.RoleIdent(), role)
	}
	typeValue := sym.Type().SubstituteRoles(substMap)

	return typeValue, typeError
}

func (tc *typeChecker) parseListValueType(ctx *parser.ListTypeContext) (types.Type, type_error.Error) {
	inner, err := tc.parseValueType(ctx.GetInner())
	if err != nil {
		return inner, err
	}

	return types.List(inner), nil
}

type callableTypeProps struct {
	roles      *types.Roles
	params     []types.Type
	returnType types.Type
}

func (tc *typeChecker) parseCallableType(ctx parser.CallableSigContext) (*callableTypeProps, bool) {
	if ctx == nil || ctx.FuncParamList() == nil {
		// parser error
		return nil, false
	}

	funcRoles, ok := tc.parseRoleType(ctx.RoleType())
	if !ok {
		return nil, false
	}

	if funcRoles.IsSharedRole() {
		tc.reportError(type_error.NewUnexpectedSharedType(ctx.RoleType()))
		funcRoles = types.NewRole(funcRoles.Participants(), false) // recover gracefully
	}

	// If function is defined within a scope, and roles are not explicitly stated,
	// then inherit the roles from the scope
	if funcRoles.IsUnnamedRole() {
		scopeRoles := tc.currentScope.Roles().Participants()
		if len(scopeRoles) > 0 {
			funcRoles = types.NewRole(scopeRoles, false)
		}
	}

	valid := true
	params := []types.Type{}

	for _, param := range ctx.FuncParamList().AllFuncParam() {

		paramType, err := tc.parseValueType(param.ValueType())
		if err != nil {
			tc.reportError(err)
			valid = false
		}

		if unknownRoles := paramType.Roles().SubtractParticipants(funcRoles.Participants()); len(unknownRoles) > 0 {
			if roleType, found := parser.FindRoleType(param.ValueType()); found {
				tc.reportError(type_error.NewRolesNotInScope(roleType, unknownRoles))
			} else {
				tc.reportError(type_error.NewRolesNotInScope(param.ValueType(), unknownRoles))
			}
			valid = false
		}

		if len(paramType.Roles().Participants()) == 0 {
			paramType = paramType.ReplaceSharedRoles(funcRoles.Participants())
		}

		params = append(params, paramType)
	}

	returnType := types.Unit()
	if ctx.GetReturnType() != nil {
		var err type_error.Error
		returnType, err = tc.parseValueType(ctx.GetReturnType())
		if err != nil {
			tc.reportError(err)
			returnType = types.Invalid()
			valid = false
		}

		if len(returnType.Roles().Participants()) == 0 {
			returnType = returnType.ReplaceSharedRoles(funcRoles.Participants())
		}
	}

	return &callableTypeProps{
		roles:      funcRoles,
		params:     params,
		returnType: returnType,
	}, valid
}

func (tc *typeChecker) parseFuncType(ctx parser.IFuncSigContext) (types.Type, bool) {
	props, ok := tc.parseCallableType(ctx)
	if !ok {
		return types.Invalid(), false
	}

	fn := types.Function(ctx, props.params, props.returnType, props.roles)
	return fn, true
}

func (tc *typeChecker) parseClosureType(ctx parser.IClosureSigContext) (types.Type, bool) {
	props, ok := tc.parseCallableType(ctx)
	if !ok {
		return types.Invalid(), false
	}

	closure := types.Closure(props.params, props.returnType, props.roles)
	return closure, true
}

// parseRoleType converts the AST representation of a role type to a [*types.Roles],
// the returned boolean indicates whether the AST role is valid.
func (tc *typeChecker) parseRoleType(ctx parser.IRoleTypeContext) (*types.Roles, bool) {
	if ctx == nil {
		return types.UnnamedRole(), true
	}

	switch ctx := ctx.(type) {
	case *parser.RoleTypeNormalContext:
		return tc.parseRoleTypeNormal(ctx)
	case *parser.RoleTypeSharedContext:
		return tc.parseRoleTypeShared(ctx)
	}

	//panic(fmt.Sprintf("unexpected role type: %#v", ctx))
	return types.EveryoneRole(), false
}

func (tc *typeChecker) parseRoleTypeNormal(ctx *parser.RoleTypeNormalContext) (*types.Roles, bool) {
	participants := []string{}
	for _, role := range ctx.AllIdent() {
		participants = append(participants, role.GetText())
	}

	role := types.NewRole(participants, false)
	valid := true

	if err := tc.checkDuplicateRoles(ctx, role); err != nil {
		tc.reportError(err)
		valid = false
	}

	return role, valid
}

func (tc *typeChecker) parseRoleTypeShared(ctx *parser.RoleTypeSharedContext) (*types.Roles, bool) {
	participants := []string{}
	for _, role := range ctx.AllIdent() {
		participants = append(participants, role.GetText())
	}

	role := types.NewRole(participants, true)
	valid := true

	if err := tc.checkDuplicateRoles(ctx, role); err != nil {
		tc.reportError(err)
		valid = false
	}

	if len(participants) == 1 {
		tc.reportError(type_error.NewSharedRoleSingleParticipant(ctx))
		valid = false
	}

	return role, valid
}

func (tc *typeChecker) parseStructType(ctx parser.IStructContext) (types.Type, bool) {
	// parser error
	if ctx == nil || ctx.Ident() == nil {
		return types.Invalid(), false
	}

	roles, ok := tc.parseRoleType(ctx.RoleType())
	if !ok {
		return types.Invalid(), false
	}

	if roles.IsSharedRole() {
		tc.reportError(type_error.NewUnexpectedSharedType(ctx.RoleType()))
		return types.Invalid(), false
	}

	implements := []types.Type{}
	if ctx.StructImplements() != nil {
		for _, impl := range ctx.StructImplements().AllRoleIdent() {
			if impl == nil || impl.Ident() == nil || impl.Ident().GetText() == "" {
				continue // parser error
			}

			implRoles, ok := tc.parseRoleType(impl.RoleType())
			if !ok {
				continue
			}

			if !roles.IsUnnamedRole() && implRoles.IsUnnamedRole() {
				tc.reportError(type_error.NewMissingRoles(impl))
				continue
			}

			if implRoles.IsSharedRole() {
				tc.reportError(type_error.NewUnexpectedSharedType(impl.RoleType()))
			} else if unknownRoles := implRoles.SubtractParticipants(roles.Participants()); len(unknownRoles) > 0 {
				tc.reportError(type_error.NewRolesNotInScope(impl.RoleType(), unknownRoles))
				continue
			}

			sym, err := tc.lookupSymbol(impl.Ident())
			if err != nil {
				tc.reportError(err)
				continue
			}

			infType, ok := sym.Type().(*types.InterfaceType)
			if !ok {
				tc.reportError(type_error.NewExpectedInterfaceType(sym, impl.Ident()))
				continue
			}

			infRoleSubst, ok := infType.Roles().SubstituteMap(implRoles)
			if !ok {
				tc.reportError(type_error.NewWrongRoleCount(sym, impl, implRoles))
				continue
			}

			implements = append(implements, infType.SubstituteRoles(infRoleSubst))
		}
	}

	return types.Struct(ctx.Ident(), roles, implements), true
}

func (tc *typeChecker) parseInterfaceType(ctx parser.IInterfaceContext) (types.Type, bool) {
	// parser error
	if ctx == nil || ctx.Ident() == nil {
		return types.Invalid(), false
	}

	// name := ctx.Ident().GetText()
	roles, ok := tc.parseRoleType(ctx.RoleType())
	if !ok {
		return types.Invalid(), false
	}

	if roles.IsSharedRole() {
		tc.reportError(type_error.NewUnexpectedSharedType(ctx.RoleType()))
		return types.Invalid(), false
	}

	return types.Interface(ctx.Ident(), roles), true
}
