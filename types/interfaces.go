package types

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
)

type InterfaceType struct {
	baseType
	ident    parser.IIdentContext
	roles    *Roles
	substMap *RoleSubst
}

func (t *InterfaceType) SubstituteRoles(substMap *RoleSubst) Type {
	newRoles := t.roles.SubstituteRoles(substMap)

	newInf := Interface(t.ident, newRoles).(*InterfaceType)
	newInf.substMap = t.substMap.ApplySubst(substMap)

	return newInf
}

func (t *InterfaceType) ReplaceSharedRoles(participants []string) Type {
	if t.roles.IsDistributedRole() || len(t.substMap.Roles) == 0 {
		return t
	}

	newRoles := NewRole(participants, true)

	newInf := Interface(t.ident, newRoles).(*InterfaceType)

	newSubst := NewRoleSubst()
	from := t.substMap.Roles[0]
	for _, to := range participants {
		newSubst.AddRole(from, to)
	}

	newInf.substMap = newSubst

	return newInf
}

func (t *InterfaceType) Roles() *Roles {
	return t.roles
}

func (t *InterfaceType) SubstMap() *RoleSubst {
	return t.substMap
}

func (t *InterfaceType) CoerceTo(other Type) (Type, bool) {
	if value, ok := baseCoerceValue(t, other); ok != nil {
		return value, *ok
	}

	if otherInf, ok := other.(*InterfaceType); ok {
		return other, t.ident == otherInf.ident
	}
	return Unit(), false
}

func (t *InterfaceType) IsEquatable() bool {
	return false
}

func (t *InterfaceType) ToString() string {
	if t.Roles().IsUnnamedRole() {
		return fmt.Sprintf("interface %s", t.ident.GetText())
	} else {
		return fmt.Sprintf("interface@%s %s", t.Roles().ToString(), t.ident.GetText())
	}
}

func Interface(ident parser.IIdentContext, roles *Roles) Type {
	substMap, ok := roles.SubstituteMap(roles)
	if !ok {
		panic("should always be ok to substitute with itself")
	}

	return &InterfaceType{
		ident:    ident,
		roles:    roles,
		substMap: substMap,
	}
}

func (t *InterfaceType) Name() string {
	return t.ident.GetText()
}

func (t *InterfaceType) Ident() parser.IIdentContext {
	return t.ident
}
