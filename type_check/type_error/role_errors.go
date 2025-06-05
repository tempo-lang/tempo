package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type DuplicateRoles struct {
	baseError
	Ctx            antlr.ParserRuleContext
	DuplicateRoles []string
}

func NewDuplicateRoles(ctx antlr.ParserRuleContext, duplicateRoles []string) Error {
	return &DuplicateRoles{
		Ctx:            ctx,
		DuplicateRoles: duplicateRoles,
	}
}

func (e *DuplicateRoles) Error() string {
	dupRoles := misc.JoinStrings(e.DuplicateRoles, ",")
	return fmt.Sprintf("duplicate roles '%s'", dupRoles)
}

func (e *DuplicateRoles) ParserRule() antlr.ParserRuleContext {
	return e.Ctx
}

type RolesNotInScope struct {
	baseError
	RoleType     parser.IRoleTypeContext
	UnknownRoles []string
}

func NewRolesNotInScope(roleType parser.IRoleTypeContext, unknownRoles []string) Error {
	return &RolesNotInScope{
		RoleType:     roleType,
		UnknownRoles: unknownRoles,
	}
}

func (e *RolesNotInScope) Error() string {
	return fmt.Sprintf("roles '%s' are not in scope", misc.JoinStrings(e.UnknownRoles, ","))
}

func (e *RolesNotInScope) ParserRule() antlr.ParserRuleContext {
	return e.RoleType
}

type UnmergableRoles struct {
	baseError
	Expr  parser.IExprContext
	Roles []*types.Roles
}

func (u *UnmergableRoles) Error() string {
	roles := misc.JoinStringsFunc(u.Roles, ", ", func(role *types.Roles) string { return role.ToString() })
	return fmt.Sprintf("can not merge roles '%s'", roles)
}

func (u *UnmergableRoles) ParserRule() antlr.ParserRuleContext {
	return u.Expr
}

func NewUnmergableRoles(expr parser.IExprContext, roles []*types.Roles) Error {
	return &UnmergableRoles{
		Expr:  expr,
		Roles: roles,
	}
}

type SharedRoleSingleParticipant struct {
	baseError
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

func (e *SharedRoleSingleParticipant) ParserRule() antlr.ParserRuleContext {
	return e.roleType
}
