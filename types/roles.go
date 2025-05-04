package types

import "slices"

type Roles struct {
	isSharedRole bool
	participants []string
}

func NewRole(participants []string, isShared bool) *Roles {
	return &Roles{
		isSharedRole: isShared,
		participants: participants,
	}
}

func SingleRole(name string) *Roles {
	return NewRole([]string{name}, false)
}

func (r *Roles) IsSharedRole() bool {
	return r.isSharedRole
}

func (r *Roles) Participants() []string {
	return r.participants
}

func (r *Roles) ToString() string {
	roles := ""
	if len(r.participants) > 0 {
		for _, p := range r.participants {
			roles += p + ","
		}
		roles = roles[:len(roles)-1]
	}

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

func (r *Roles) Contains(other *Roles) bool {
	if r.isSharedRole != other.isSharedRole {
		return false
	}

	for _, o := range other.participants {
		if !slices.Contains(r.participants, o) {
			return false
		}
	}

	return true
}
