package type_check

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser/ast"
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

func findRoleType(t ast.ValueType) (*ast.RoleType, bool) {
	switch t := t.(type) {
	case *ast.NamedType:
		return t.RoleIdent.RoleType, t.RoleIdent.RoleType != nil
	case *ast.ListType:
		return findRoleType(t.Inner)
	case *ast.AsyncType:
		return findRoleType(t.Inner)
	case *ast.ClosureType:
		return t.RoleType, t.RoleType != nil
	}
	return nil, false
}

func (tc *typeChecker) parseValueType(ctx ast.ValueType) (types.Type, type_error.Error) {
	// parser error
	if ctx == nil {
		return types.Invalid(), nil
	}

	switch ctx := ctx.(type) {
	case *ast.NamedType:
		return tc.parseNamedValueType(ctx)
	case *ast.ListType:
		return tc.parseListValueType(ctx)
	case *ast.AsyncType:
		return tc.parseAsyncValueType(ctx)
	case *ast.ClosureType:
		return tc.parseClosureValueType(ctx)
	case *ast.InvalidType:
		return types.Invalid(), nil
	}
	panic(fmt.Sprintf("parseValueType unexpected ctx: %T", ctx))
}

func (tc *typeChecker) parseAsyncValueType(ctx *ast.AsyncType) (types.Type, type_error.Error) {
	inner, err := tc.parseValueType(ctx.Inner)
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

func (tc *typeChecker) parseClosureValueType(ctx *ast.ClosureType) (types.Type, type_error.Error) {
	// parser error
	if ctx.RoleType == nil || ctx.Params == nil {
		return types.Invalid(), nil
	}

	roles, ok := tc.parseRoleType(ctx.RoleType)
	if !ok {
		return types.Invalid(), nil
	}

	params := []types.Type{}
	for _, param := range ctx.Params.Params {
		paramType, err := tc.parseValueType(param)
		if err != nil {
			return types.Invalid(), err
		}
		params = append(params, paramType)

		if unknownRoles := paramType.Roles().SubtractParticipants(roles.Participants()); len(unknownRoles) > 0 {
			if roleType, found := findRoleType(param); found {
				return types.Invalid(), type_error.NewRolesNotInScope(roleType, unknownRoles)
			} else {
				return types.Invalid(), type_error.NewRolesNotInScope(param, unknownRoles)
			}
		}
	}

	var returnType types.Type = types.Unit()
	if ctx.ReturnType != nil {
		ret, err := tc.parseValueType(ctx.ReturnType)
		if err != nil {
			return types.Invalid(), err
		}
		returnType = ret

		if unknownRoles := returnType.Roles().SubtractParticipants(roles.Participants()); len(unknownRoles) > 0 {
			if roleType, found := findRoleType(ctx.ReturnType); found {
				return types.Invalid(), type_error.NewRolesNotInScope(roleType, unknownRoles)
			} else {
				return types.Invalid(), type_error.NewRolesNotInScope(ctx.ReturnType, unknownRoles)
			}
		}
	}

	closureValue := types.Closure(params, returnType, roles)

	return closureValue, nil
}

func (tc *typeChecker) parseNamedValueType(ctx *ast.NamedType) (types.Type, type_error.Error) {
	// parser error
	if ctx == nil {
		return types.Invalid(), nil
	}

	role, ok := tc.parseRoleType(ctx.RoleIdent.RoleType)
	if !ok {
		return types.Invalid(), nil
	}

	if !tc.currentScope.Roles().IsUnnamedRole() && role.IsUnnamedRole() {
		return types.Invalid(), type_error.NewMissingRoles(ctx)
	}

	typeName := ctx.RoleIdent.Ident
	if builtinType, isBuiltinType := ToBuiltinValue(typeName.Value(), role.Participants()); isBuiltinType {
		if !role.IsSharedRole() && len(role.Participants()) > 1 {
			return types.Invalid(), type_error.NewNotDistributedType(ctx)
		}
		return builtinType, nil
	}

	sym, err := tc.lookupSymbol(typeName)
	if err != nil {
		return types.Invalid(), err
	}

	sym.AddRead(ctx.RoleIdent.Ident)

	substMap, rolesMatch := sym.Type().Roles().SubstituteMap(role)
	if !rolesMatch {
		return types.Invalid(), type_error.NewWrongRoleCount(sym, ctx.RoleIdent, role)
	}
	typeValue := sym.Type().SubstituteRoles(substMap)

	return typeValue, nil
}

func (tc *typeChecker) parseListValueType(ctx *ast.ListType) (types.Type, type_error.Error) {
	inner, err := tc.parseValueType(ctx.Inner)
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

func (tc *typeChecker) parseCallableType(ctx *ast.FuncSig) (*callableTypeProps, bool) {
	if ctx == nil || ctx.Params == nil {
		// parser error
		return nil, false
	}

	funcRoles, ok := tc.parseRoleType(ctx.RoleType)
	if !ok {
		return nil, false
	}

	if funcRoles.IsSharedRole() {
		tc.reportError(type_error.NewUnexpectedSharedType(ctx.RoleType))
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

	for _, param := range ctx.Params.Params {

		paramType, err := tc.parseValueType(param.Type)
		if err != nil {
			tc.reportError(err)
			valid = false
		}

		if unknownRoles := paramType.Roles().SubtractParticipants(funcRoles.Participants()); len(unknownRoles) > 0 {
			if roleType, found := findRoleType(param.Type); found {
				tc.reportError(type_error.NewRolesNotInScope(roleType, unknownRoles))
			} else {
				tc.reportError(type_error.NewRolesNotInScope(param.Type, unknownRoles))
			}
			valid = false
		}

		if len(paramType.Roles().Participants()) == 0 {
			paramType = paramType.ReplaceSharedRoles(funcRoles.Participants())
		}

		params = append(params, paramType)
	}

	returnType := types.Unit()
	if ctx.ReturnType != nil {
		var err type_error.Error
		returnType, err = tc.parseValueType(ctx.ReturnType)
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

func (tc *typeChecker) parseFuncType(ctx *ast.FuncSig) (types.Type, bool) {
	props, ok := tc.parseCallableType(ctx)
	if !ok {
		return types.Invalid(), false
	}

	fn := types.Function(ctx, props.params, props.returnType, props.roles)
	return fn, true
}

func (tc *typeChecker) parseClosureType(ctx *ast.ClosureSig) (types.Type, bool) {
	tmp := &ast.FuncSig{FuncToken: ctx.FuncToken, RoleType: ctx.RoleType, Params: ctx.Params, ReturnType: ctx.ReturnType}
	props, ok := tc.parseCallableType(tmp)
	if !ok {
		return types.Invalid(), false
	}

	closure := types.Closure(props.params, props.returnType, props.roles)
	return closure, true
}

// parseRoleType converts the AST representation of a role type to a [*types.Roles],
// the returned boolean indicates whether the AST role is valid.
func (tc *typeChecker) parseRoleType(ctx *ast.RoleType) (*types.Roles, bool) {
	if ctx == nil {
		return types.UnnamedRole(), true
	}

	if ctx.IsShared() {
		return tc.parseRoleTypeShared(ctx)
	}
	return tc.parseRoleTypeNormal(ctx)
}

func (tc *typeChecker) parseRoleTypeNormal(ctx *ast.RoleType) (*types.Roles, bool) {
	participants := []string{}
	for _, role := range ctx.Roles() {
		participants = append(participants, role.Token.Text)
	}

	role := types.NewRole(participants, false)
	valid := true

	if err := tc.checkDuplicateRoles(ctx, role); err != nil {
		tc.reportError(err)
		valid = false
	}

	return role, valid
}

func (tc *typeChecker) parseRoleTypeShared(ctx *ast.RoleType) (*types.Roles, bool) {
	participants := []string{}
	for _, role := range ctx.Roles() {
		participants = append(participants, role.Token.Text)
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

func (tc *typeChecker) parseStructType(ctx *ast.Struct) (types.Type, bool) {
	// parser error
	if ctx == nil || ctx.Name == nil {
		return types.Invalid(), false
	}

	roles, ok := tc.parseRoleType(ctx.RoleType)
	if !ok {
		return types.Invalid(), false
	}

	if roles.IsSharedRole() {
		tc.reportError(type_error.NewUnexpectedSharedType(ctx.RoleType))
		return types.Invalid(), false
	}

	if !roles.IsComplete() {
		tc.reportError(type_error.NewUnexpectedHiddenRoles(ctx.RoleType))
		return types.Invalid(), false
	}

	implements := []types.Type{}
	if ctx.Implements != nil {
		for _, impl := range ctx.Implements {
			if impl == nil || impl.Ident == nil || impl.Ident.Value() == "" {
				continue // parser error
			}

			implRoles, ok := tc.parseRoleType(impl.RoleType)
			if !ok {
				continue
			}

			if !roles.IsUnnamedRole() && implRoles.IsUnnamedRole() {
				tc.reportError(type_error.NewMissingRoles(impl))
				continue
			}

			if implRoles.IsSharedRole() {
				tc.reportError(type_error.NewUnexpectedSharedType(impl.RoleType))
			} else if unknownRoles := implRoles.SubtractParticipants(roles.Participants()); len(unknownRoles) > 0 {
				tc.reportError(type_error.NewRolesNotInScope(impl.RoleType, unknownRoles))
				continue
			}

			sym, err := tc.lookupSymbol(impl.Ident)
			if err != nil {
				tc.reportError(err)
				continue
			}

			infType, ok := sym.Type().(*types.InterfaceType)
			if !ok {
				tc.reportError(type_error.NewExpectedInterfaceType(sym, impl.Ident))
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

	return types.Struct(ctx.Name, roles, implements), true
}

func (tc *typeChecker) parseInterfaceType(ctx *ast.Interface) (types.Type, bool) {
	// parser error
	if ctx == nil || ctx.Name == nil {
		return types.Invalid(), false
	}

	// name := ctx.Name.GetText()
	roles, ok := tc.parseRoleType(ctx.RoleType)
	if !ok {
		return types.Invalid(), false
	}

	if roles.IsSharedRole() {
		tc.reportError(type_error.NewUnexpectedSharedType(ctx.RoleType))
		return types.Invalid(), false
	}

	if !roles.IsComplete() {
		tc.reportError(type_error.NewUnexpectedHiddenRoles(ctx.RoleType))
		return types.Invalid(), false
	}

	return types.Interface(ctx.Name, roles), true
}
