package types

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
)

type StructType struct {
	baseType
	structIdent  parser.IIdentContext
	participants []string
	substMap     *RoleSubst
}

func (t *StructType) SubstituteRoles(substMap *RoleSubst) Type {
	newParticipants := []string{}
	for _, from := range t.participants {
		newParticipants = append(newParticipants, substMap.Subst(from))
	}

	newStruct := NewStructType(t.structIdent, newParticipants).(*StructType)
	newStruct.substMap = t.substMap.ApplySubst(substMap)

	return newStruct
}

func (t *StructType) ReplaceSharedRoles(participants []string) Type {
	return t
}

func (t *StructType) CoerceTo(other Type) (Type, bool) {
	if value, ok := baseCoerceValue(t, other); ok != nil {
		return value, *ok
	}

	if otherStruct, ok := other.(*StructType); ok {
		if t.structIdent == otherStruct.structIdent {
			return t, true
		}
	}
	return Invalid(), false
}

func (t *StructType) Roles() *Roles {
	return NewRole(t.participants, false)
}

func (t *StructType) SubstMap() *RoleSubst {
	return t.substMap
}

func (t *StructType) IsSendable() bool {
	return true
}

func (t *StructType) IsEquatable() bool {
	return true
}

func (t *StructType) ToString() string {
	return fmt.Sprintf("struct@%s %s", t.Roles().ToString(), t.structIdent.GetText())
}

func NewStructType(structIdent parser.IIdentContext, participants []string) Type {
	substMap := NewRoleSubst()
	for _, role := range participants {
		substMap.AddRole(role, role)
	}

	return &StructType{structIdent: structIdent, participants: participants, substMap: substMap}
}

func (t *StructType) Name() string {
	return t.structIdent.GetText()
}

func (t *StructType) Ident() parser.IIdentContext {
	return t.structIdent
}
