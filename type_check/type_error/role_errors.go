package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type DuplicateRolesError struct {
	Ctx            antlr.ParserRuleContext
	DuplicateRoles []string
}

func NewDuplicateRolesError(ctx antlr.ParserRuleContext, duplicateRoles []string) Error {
	return &DuplicateRolesError{
		Ctx:            ctx,
		DuplicateRoles: duplicateRoles,
	}
}

func (e *DuplicateRolesError) IsTypeError() {}

func (e *DuplicateRolesError) Error() string {
	dupRoles := misc.JoinStrings(e.DuplicateRoles, ",")
	return fmt.Sprintf("duplicate roles '%s'", dupRoles)
}

func (e *DuplicateRolesError) ParserRule() antlr.ParserRuleContext {
	return e.Ctx
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
