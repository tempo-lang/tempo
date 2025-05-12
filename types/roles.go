package types

import (
	"slices"
	"tempo/misc"
	"tempo/parser"
)

type Roles struct {
	isSharedRole bool
	participants []string
}

func NewRole(participants []string, isShared bool) *Roles {
	if len(participants) == 0 {
		participants = nil
	}

	return &Roles{
		isSharedRole: isShared,
		participants: participants,
	}
}

func SingleRole(name string) *Roles {
	return NewRole([]string{name}, false)
}

func EveryoneRole() *Roles {
	return NewRole(nil, true)
}

func RoleIntersect(expr parser.IExprContext, roles ...*Roles) (*Roles, Error) {

	nonEmptyRoles := []*Roles{}
	for _, role := range roles {
		if len(role.participants) > 0 {
			nonEmptyRoles = append(nonEmptyRoles, role)
		}
	}
	roles = nonEmptyRoles

	if len(roles) == 1 {
		return roles[0], nil
	}

	if len(roles) == 0 {
		return EveryoneRole(), nil
	}

	participants := roles[0].participants
	for _, role := range roles[1:] {
		// zero participants, means that the type can coerce to all roles
		if len(role.participants) == 0 {
			continue
		}

		newParticipants := []string{}
		for _, p := range role.participants {
			if slices.Contains(participants, p) {
				newParticipants = append(newParticipants, p)
			}
		}
		participants = newParticipants
	}

	if len(participants) == 1 {
		return SingleRole(participants[0]), nil
	}

	if len(participants) > 1 {
		return NewRole(participants, true), nil
	}

	return nil, NewUnmergableRolesError(expr, roles)
}

func (r *Roles) IsSharedRole() bool {
	return r.isSharedRole
}

func (r *Roles) Participants() []string {
	return r.participants
}

func (r *Roles) ToString() string {
	roles := misc.JoinStrings(r.participants, ",")

	if r.isSharedRole {
		return "[" + roles + "]"
	} else {
		return "(" + roles + ")"
	}
}

func (r *Roles) SubtractParticipants(other []string) []string {
	result := []string{}

	for _, role := range r.participants {
		if !slices.Contains(other, role) {
			result = append(result, role)
		}
	}

	return result
}

func (r *Roles) Encompass(other *Roles) bool {
	if len(other.participants) > 1 && r.isSharedRole != other.isSharedRole {
		return false
	}

	for _, o := range other.participants {
		if !slices.Contains(r.participants, o) {
			return false
		}
	}

	return true
}

func (r *Roles) Contains(role string) bool {
	return slices.Contains(r.participants, role)
}
