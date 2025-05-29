package type_error

import (
	"fmt"
	"tempo/misc"
	"tempo/parser"
	"tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type DuplicateRolesError struct {
	Func           parser.IFuncSigContext
	DuplicateRoles []parser.IIdentContext
}

func NewDuplicateRolesError(function parser.IFuncSigContext, duplicateRoles []parser.IIdentContext) Error {
	return &DuplicateRolesError{
		Func:           function,
		DuplicateRoles: duplicateRoles,
	}
}

func (e *DuplicateRolesError) IsTypeError() {}

func (e *DuplicateRolesError) Error() string {
	return fmt.Sprintf("function '%s' has duplicate role '%s'", e.Func.Ident().GetText(), e.DuplicateRoles[0].GetText())
}

func (e *DuplicateRolesError) ParserRule() antlr.ParserRuleContext {
	return e.Func
}

type RolesNotInScopeError struct {
	RoleType     parser.IRoleTypeContext
	UnknownRoles []string
}

func NewRolesNotInScopeError(roleType parser.IRoleTypeContext, unknownRoles []string) Error {
	return &RolesNotInScopeError{
		RoleType:     roleType,
		UnknownRoles: unknownRoles,
	}
}

func (e *RolesNotInScopeError) IsTypeError() {}

func (e *RolesNotInScopeError) Error() string {
	return fmt.Sprintf("roles '%s' are not in scope", misc.JoinStrings(e.UnknownRoles, ","))
}

func (e *RolesNotInScopeError) ParserRule() antlr.ParserRuleContext {
	return e.RoleType
}

type UnmergableRolesError struct {
	Expr  parser.IExprContext
	Roles []*types.Roles
}

func (u *UnmergableRolesError) Error() string {
	roles := misc.JoinStringsFunc(u.Roles, ", ", func(role *types.Roles) string { return role.ToString() })
	return fmt.Sprintf("can not merge roles '%s'", roles)
}

func (u *UnmergableRolesError) IsTypeError() {}

func (u *UnmergableRolesError) ParserRule() antlr.ParserRuleContext {
	return u.Expr
}

func NewUnmergableRolesError(expr parser.IExprContext, roles []*types.Roles) Error {
	return &UnmergableRolesError{
		Expr:  expr,
		Roles: roles,
	}
}

type SharedRoleSingleParticipant struct {
	roleType *parser.RoleTypeSharedContext
}

func NewSharedRoleSingleParticipant(roleType *parser.RoleTypeSharedContext) Error {
	return &SharedRoleSingleParticipant{
		roleType: roleType,
	}
}

func (e *SharedRoleSingleParticipant) Error() string {
	return "shared role must have more than one participant"
}

func (e *SharedRoleSingleParticipant) IsTypeError() {}

func (e *SharedRoleSingleParticipant) ParserRule() antlr.ParserRuleContext {
	return e.roleType
}
