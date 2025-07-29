package types

import (
	"fmt"
	"slices"

	"github.com/tempo-lang/tempo/misc"
)

type RoleType int

const (
	// ROLE_LOCAL is a type with only a single role participant.
	ROLE_LOCAL RoleType = 0
	// ROLE_SHARED is a type with its value shared across multiple role participants.
	ROLE_SHARED RoleType = 1
	// ROLE_DISTRIBUTED is a type which spans across multiple roles,
	// such that different parts of the type live at different roles.
	ROLE_DISTRIBUTED RoleType = 2
)

// Roles describes the role participants of a [Type].
type Roles struct {
	roleType     RoleType
	participants []string
}

// RoleSubst is a substitution map which is used to substitute the roles in a type with a new set of roles.
type RoleSubst struct {
	Roles []string
	Map   map[string]string
}

// NewRoleSubst constructs a new role substitution.
func NewRoleSubst() *RoleSubst {
	return &RoleSubst{
		Roles: []string{},
		Map:   map[string]string{},
	}
}

// AddRole adds a new substitution to a role substitution map.
func (r *RoleSubst) AddRole(from, to string) {
	if _, found := r.Map[from]; found {
		return
	}

	r.Roles = append(r.Roles, from)
	r.Map[from] = to
}

// Subst returns the substitution of the given `from` role participant.
func (r *RoleSubst) Subst(from string) string {
	if to, ok := r.Map[from]; ok {
		return to
	}
	return from
}

// Inverse returns a new role substitution map where all substitutions are reversed.
func (r *RoleSubst) Inverse() *RoleSubst {
	inv := NewRoleSubst()
	for _, role := range r.Roles {
		inv.AddRole(r.Subst(role), role)
	}
	return inv
}

// ApplySubst returns a new role substitution which has the samme effect as applying the original followed by the other.
func (r *RoleSubst) ApplySubst(other *RoleSubst) *RoleSubst {
	result := NewRoleSubst()
	for _, role := range r.Roles {
		result.AddRole(role, other.Subst(r.Map[role]))
	}
	return result
}

// NewRole constructs a new roles object, given a set of participants and whether the role is shared or not.
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

// SingleRole is a convenience constructor to create a roles object with a single participant.
func SingleRole(name string) *Roles {
	return NewRole([]string{name}, false)
}

// EveryoneRole constructs a special roles object which describes a shared role across all active participants in scope.
func EveryoneRole() *Roles {
	return NewRole(nil, true)
}

// RoleIntersect calculates the intersection of participants between a list of roles.
// It returns a new roles object containing the intersection and a boolean that is false if the intersection is empty.
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

// IsLocalRole returns true if the [RoleType] of the roles object is [ROLE_LOCAL].
func (r *Roles) IsLocalRole() bool {
	return r.roleType == ROLE_LOCAL
}

// IsSharedRole returns true if the [RoleType] of the roles object is [ROLE_SHARED].
func (r *Roles) IsSharedRole() bool {
	return r.roleType == ROLE_SHARED
}

// IsDistributedRole returns true if the [RoleType] of the roles object is [ROLE_DISTRIBUTED].
func (r *Roles) IsDistributedRole() bool {
	return r.roleType == ROLE_DISTRIBUTED
}

// Participants returns a copy of the participants involved in the roles object.
func (r *Roles) Participants() []string {
	return slices.Clone(r.participants)
}

// ToString formats the roles object as how it would be written in source code.
func (r *Roles) ToString() string {
	roles := misc.JoinStrings(r.participants, ",")

	if len(r.participants) == 1 {
		return r.participants[0]
	}

	switch r.roleType {
	case ROLE_LOCAL:
		return r.participants[0]
	case ROLE_SHARED:
		if len(r.participants) == 0 {
			return "[..]"
		}
		return "[" + roles + "]"
	case ROLE_DISTRIBUTED:
		return "(" + roles + ")"
	default:
		panic(fmt.Sprintf("invalid role type: %d", r.roleType))
	}
}

// SubtractParticipants returns a list of the participants in this roles object,
// without any of the roles in the `other` list.
func (r *Roles) SubtractParticipants(other []string) []string {
	result := []string{}

	for _, role := range r.participants {
		if !slices.Contains(other, role) {
			result = append(result, role)
		}
	}

	return result
}

// Encompass returns true if `this` roles object can be coerced into `other`.
//
//   - For [ROLE_SHARED] roles, this is the case if `this` is a superset of `other`.
//   - For [ROLE_DISTRIBUTED] roles, this is the case if `this` and `other` has the same number of participants,
//     but each pairwise participant may be different.
func (r *Roles) Encompass(other *Roles) bool {
	if (r.IsLocalRole() || r.IsSharedRole()) && other.IsLocalRole() {
		if r.participants == nil {
			return true
		}

		for _, role := range r.participants {
			if other.participants[0] == role {
				return true
			}
		}
		return false
	}

	if r.IsDistributedRole() && other.IsDistributedRole() {
		if len(r.participants) != len(other.participants) {
			return false
		}

		for i, role := range r.participants {
			if role != other.participants[i] {
				return false
			}
		}

		return true
	}

	if r.IsSharedRole() && other.IsSharedRole() {
		if r.participants == nil {
			return true
		}

		for _, o := range other.participants {
			if !slices.Contains(r.participants, o) {
				return false
			}
		}

		return true
	}

	return false
}

// Contains returns whether the given role is contained in this roles object.
func (r *Roles) Contains(role string) bool {
	if len(r.participants) == 0 {
		return true
	}

	return slices.Contains(r.participants, role)
}

// SubstituteMap calculates a role substitution map from `this` roles object to `other`.
// The substitution fails if the set of participants for each roles object has different lengths.
// The returned boolean indicates whether the substitution was successfull.
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

// SubstituteRoles returns a new roles object with substitutions for the roles mention in the given role substitution map.
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
