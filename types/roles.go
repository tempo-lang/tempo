package types

import (
	"fmt"
	"slices"
	"tempo/misc"
)

type RoleType int

const (
	ROLE_LOCAL       RoleType = 0
	ROLE_SHARED      RoleType = 1
	ROLE_DISTRIBUTED RoleType = 2
)

type Roles struct {
	roleType     RoleType
	participants []string
}

type RoleSubst struct {
	Roles []string
	Map   map[string]string
}

func NewRoleSubst() *RoleSubst {
	return &RoleSubst{
		Roles: []string{},
		Map:   map[string]string{},
	}
}

func (r *RoleSubst) AddRole(from, to string) {
	if _, found := r.Map[from]; found {
		return
	}

	r.Roles = append(r.Roles, from)
	r.Map[from] = to
}

func (r *RoleSubst) Subst(from string) string {
	if to, ok := r.Map[from]; ok {
		return to
	}
	return from
}

func (r *RoleSubst) Inverse() *RoleSubst {
	inv := NewRoleSubst()
	for _, role := range r.Roles {
		inv.AddRole(r.Subst(role), role)
	}
	return inv
}

func NewRole(participants []string, isShared bool) *Roles {
	var roleType RoleType
	if len(participants) == 0 {
		participants = nil
		roleType = ROLE_SHARED
	} else if len(participants) == 1 {
		roleType = ROLE_LOCAL
	} else if isShared {
		roleType = ROLE_SHARED
	} else {
		roleType = ROLE_DISTRIBUTED
	}

	return &Roles{
		roleType:     roleType,
		participants: participants,
	}
}

func SingleRole(name string) *Roles {
	return NewRole([]string{name}, false)
}

func EveryoneRole() *Roles {
	return NewRole(nil, true)
}

func RoleIntersect(roles ...*Roles) (*Roles, bool) {
	nonEmptyRoles := []*Roles{}
	for _, role := range roles {
		if len(role.participants) > 0 {
			nonEmptyRoles = append(nonEmptyRoles, role)
		}
	}
	roles = nonEmptyRoles

	if len(roles) == 1 {
		return roles[0], true
	}

	if len(roles) == 0 {
		return EveryoneRole(), true
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
		return SingleRole(participants[0]), true
	}

	if len(participants) > 1 {
		return NewRole(participants, true), true
	}

	// invalid role
	return nil, false
}

func (r *Roles) IsLocalRole() bool {
	return r.roleType == ROLE_LOCAL
}

func (r *Roles) IsSharedRole() bool {
	return r.roleType == ROLE_SHARED
}

func (r *Roles) IsDistributedRole() bool {
	return r.roleType == ROLE_DISTRIBUTED
}

func (r *Roles) Participants() []string {
	return slices.Clone(r.participants)
}

func (r *Roles) ToString() string {
	roles := misc.JoinStrings(r.participants, ",")

	if len(r.participants) == 1 {
		return r.participants[0]
	}

	switch r.roleType {
	case ROLE_LOCAL:
		return r.participants[0]
	case ROLE_SHARED:
		return "[" + roles + "]"
	case ROLE_DISTRIBUTED:
		return "(" + roles + ")"
	default:
		panic(fmt.Sprintf("invalid role type: %d", r.roleType))
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
	if (r.IsLocalRole() || r.IsSharedRole()) && other.IsLocalRole() {
		for _, role := range r.participants {
			if other.participants[0] == role {
				return true
			}
		}
		return false
	}

	if !r.IsSharedRole() || !other.IsSharedRole() {
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
	if len(r.participants) == 0 {
		return true
	}

	return slices.Contains(r.participants, role)
}

func (r *Roles) SubstituteMap(other *Roles) (*RoleSubst, bool) {
	if len(r.participants) != len(other.participants) {
		return nil, false
	}

	roleSubst := NewRoleSubst()
	for i, role := range r.Participants() {
		roleSubst.AddRole(role, other.participants[i])
	}
	return roleSubst, true
}

func (r *Roles) SubstituteRoles(subst *RoleSubst) *Roles {
	roleSubst := r.Participants()
	for i := range roleSubst {
		newRole, found := subst.Map[roleSubst[i]]
		if found {
			roleSubst[i] = newRole
		}
	}

	return NewRole(roleSubst, r.IsSharedRole())
}

func (t *Type) SubstituteRoles(subst *RoleSubst) *Type {
	return New(t.Value().SubstituteRoles(subst), t.Roles().SubstituteRoles(subst))
}
