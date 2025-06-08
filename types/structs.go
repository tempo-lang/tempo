package types

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
)

type StructType struct {
	baseValue
	structIdent parser.IIdentContext
	roleSubst   *RoleSubst
}

func (s *StructType) SubstituteRoles(substMap *RoleSubst) Value {
	newRoleSubst := NewRoleSubst()
	for _, from := range s.roleSubst.Roles {
		newRoleSubst.AddRole(from, substMap.Map[from])
	}

	return &StructType{
		structIdent: s.structIdent,
		roleSubst:   newRoleSubst,
	}
}

func (s *StructType) ReplaceSharedRoles(participants []string) Value {
	return s
}

func (s *StructType) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(s, other); ok != nil {
		return value, *ok
	}

	if otherStruct, ok := other.(*StructType); ok {
		if s.structIdent == otherStruct.structIdent {
			return s, true
		}
	}
	return Invalid(), false
}

func (s *StructType) Roles() *Roles {
	participants := []string{}
	for _, r := range s.roleSubst.Roles {
		participants = append(participants, s.roleSubst.Subst(r))
	}
	return NewRole(participants, false)
}

func (s *StructType) IsSendable() bool {
	return true
}

func (t *StructType) IsEquatable() bool {
	return true
}

func (s *StructType) ToString() string {
	return fmt.Sprintf("struct@%s %s", s.Roles().ToString(), s.structIdent.GetText())
}

func NewStructType(structIdent parser.IIdentContext, roles []string) Value {
	roleSubst := NewRoleSubst()
	for _, role := range roles {
		roleSubst.AddRole(role, role)
	}

	return &StructType{structIdent: structIdent, roleSubst: roleSubst}
}

func (s *StructType) Name() string {
	return s.structIdent.GetText()
}
