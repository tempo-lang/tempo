package type_error

import (
	"fmt"

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
	return "duplicate roles are not allowed"
}

func (e *DuplicateRoles) Annotations() []Annotation {
	dupRoles := formatList("role", "roles", e.DuplicateRoles, "and")

	return []Annotation{{
		Type:    AnnotationTypeNote,
		Message: "a process can only act as one role at a time in a function, struct or interface.",
	}, {
		Type:    AnnotationTypeHint,
		Message: fmt.Sprintf("the %s %s used multiple times.", dupRoles, toBe(e.DuplicateRoles)),
	}}
}

func (e *DuplicateRoles) ParserRule() antlr.ParserRuleContext {
	return e.Ctx
}

func (e *DuplicateRoles) Code() ErrorCode {
	return CodeDuplicateRoles
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
	roles := formatList("role", "roles", e.UnknownRoles, "and")
	return fmt.Sprintf("%s %s not in scope", roles, toBe(e.UnknownRoles))
}

func (e *RolesNotInScope) ParserRule() antlr.ParserRuleContext {
	return e.RoleType
}

func (e *RolesNotInScope) Code() ErrorCode {
	return CodeRolesNotInScope
}

type UnmergableRoles struct {
	baseError
	Expr  parser.IExprContext
	Roles []*types.Roles
}

func (u *UnmergableRoles) Error() string {
	roles := []string{}
	for _, role := range u.Roles {
		roles = append(roles, role.ToString())
	}

	rolesList := formatList("role", "roles", roles, "and")

	return fmt.Sprintf("cannot merge %s", rolesList)
}

func (u *UnmergableRoles) ParserRule() antlr.ParserRuleContext {
	return u.Expr
}

func (e *UnmergableRoles) Code() ErrorCode {
	return CodeUnmergableRoles
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

func (e *SharedRoleSingleParticipant) Code() ErrorCode {
	return CodeSharedRoleSingleParticipant
}

type MissingRoles struct {
	baseError
	ValueType antlr.ParserRuleContext
}

func NewMissingRoles(valueType antlr.ParserRuleContext) Error {
	return &MissingRoles{
		ValueType: valueType,
	}
}

func (e *MissingRoles) Error() string {
	return "unable to determine roles for this type"
}

func (e *MissingRoles) ParserRule() antlr.ParserRuleContext {
	return e.ValueType
}

func (e *MissingRoles) Annotations() []Annotation {
	return []Annotation{
		{
			Type:    AnnotationTypeHint,
			Message: fmt.Sprintf("specify roles explicitly, like so `%s@(A,B,C)`", e.ValueType.GetText()),
		},
	}
}

func (e *MissingRoles) Code() ErrorCode {
	return CodeMissingRoles
}
