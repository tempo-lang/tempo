package types

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
)

type StructType struct {
	baseType
	structIdent  parser.IIdentContext
	participants []string
}

func (s *StructType) SubstituteRoles(substMap *RoleSubst) Type {
	newParticipants := []string{}
	for _, from := range s.participants {
		newParticipants = append(newParticipants, substMap.Subst(from))
	}

	return NewStructType(s.structIdent, newParticipants)
}

func (s *StructType) ReplaceSharedRoles(participants []string) Type {
	return s
}

func (s *StructType) CoerceTo(other Type) (Type, bool) {
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
	return NewRole(s.participants, false)
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

func NewStructType(structIdent parser.IIdentContext, participants []string) Type {
	return &StructType{structIdent: structIdent, participants: participants}
}

func (s *StructType) Name() string {
	return s.structIdent.GetText()
}
