package types

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser/new_parser/ast"
)

type StructType struct {
	baseType
	structIdent *ast.Identifier
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

	newStruct := Struct(t.structIdent, newRoles, newImplements).(*StructType)
	newStruct.substMap = t.substMap.ApplySubst(substMap)

	return newStruct
}

func (t *StructType) ReplaceSharedRoles(participants []string) Type {
	if t.roles.IsDistributedRole() || len(t.substMap.Roles) == 0 {
		return t
	}

	newRoles := NewRole(participants, true)

	newImplements := []Type{}
	for _, impl := range t.implements {
		newImplements = append(newImplements, impl.ReplaceSharedRoles(participants))
	}

	newStruct := Struct(t.structIdent, newRoles, newImplements).(*StructType)

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
			return otherStruct, true
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

func (t *StructType) IsEquatable() bool {
	return true
}

func (t *StructType) ToString() string {
	if t.Roles().IsUnnamedRole() {
		return fmt.Sprintf("struct %s", t.structIdent.Value())
	} else {
		return fmt.Sprintf("struct@%s %s", t.Roles().ToString(), t.structIdent.Value())
	}
}

func Struct(structIdent *ast.Identifier, roles *Roles, implements []Type) Type {
	substMap, ok := roles.SubstituteMap(roles)
	if !ok {
		panic("should always be ok to substitute with itself")
	}

	return &StructType{structIdent: structIdent, roles: roles, substMap: substMap, implements: implements}
}

func (t *StructType) Name() string {
	return t.structIdent.Value()
}

func (t *StructType) Ident() *ast.Identifier {
	return t.structIdent
}

func (t *StructType) Implements() []Type {
	return t.implements
}
