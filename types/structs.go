package types

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
)

type StructType struct {
	baseType
	structIdent parser.IIdentContext
	roles       *Roles
	substMap    *RoleSubst
	implements  []Type
}

func (t *StructType) SubstituteRoles(substMap *RoleSubst) Type {

	newRoles := t.roles.SubstituteRoles(substMap)

	newImplements := []Type{}
	for _, impl := range t.implements {
		newImplements = append(newImplements, impl.SubstituteRoles(substMap))
	}

	newStruct := NewStructType(t.structIdent, newRoles, newImplements).(*StructType)
	newStruct.substMap = t.substMap.ApplySubst(substMap)

	return newStruct
}

func (t *StructType) ReplaceSharedRoles(participants []string) Type {
	if t.roles.IsDistributedRole() {
		return t
	}

	newRoles := NewRole(participants, true)

	newImplements := []Type{}
	for _, impl := range t.implements {
		newImplements = append(newImplements, impl.ReplaceSharedRoles(participants))
	}

	newStruct := NewStructType(t.structIdent, newRoles, newImplements).(*StructType)

	newSubst := NewRoleSubst()
	from := t.substMap.Roles[0]
	for _, to := range participants {
		newSubst.AddRole(from, to)
	}

	newStruct.substMap = newSubst

	return newStruct
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

	for _, impl := range t.implements {
		if implType, ok := impl.CoerceTo(other); ok {
			return implType, true
		}
	}

	return Invalid(), false
}

func (t *StructType) Roles() *Roles {
	return t.roles
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

func NewStructType(structIdent parser.IIdentContext, roles *Roles, implements []Type) Type {
	substMap, ok := roles.SubstituteMap(roles)
	if !ok {
		panic("should always be ok to substitute with itself")
	}

	return &StructType{structIdent: structIdent, roles: roles, substMap: substMap, implements: implements}
}

func (t *StructType) Name() string {
	return t.structIdent.GetText()
}

func (t *StructType) Ident() parser.IIdentContext {
	return t.structIdent
}

func (t *StructType) Implements() []Type {
	return t.implements
}
