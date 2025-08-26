package types

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
)

type InterfaceType struct {
	baseType
	ident        parser.IIdentContext
	participants []string
	substMap     *RoleSubst
}

func (t *InterfaceType) SubstituteRoles(substMap *RoleSubst) Type {
	newParticipants := []string{}
	for _, from := range t.participants {
		newParticipants = append(newParticipants, substMap.Subst(from)[0])
	}

	newInf := Interface(t.ident, newParticipants).(*InterfaceType)
	newInf.substMap = t.substMap.ApplySubst(substMap)

	return newInf
}

func (t *InterfaceType) ReplaceSharedRoles(participants []string) Type {
	return t
}

func (t *InterfaceType) Roles() *Roles {
	return NewRole(t.participants, false)
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
	return true
}

func (t *InterfaceType) ToString() string {
	return fmt.Sprintf("interface@%s %s", t.Roles().ToString(), t.ident.GetText())
}

func Interface(ident parser.IIdentContext, participants []string) Type {
	substMap := NewRoleSubst()
	for _, role := range participants {
		substMap.AddRole(role, role)
	}

	return &InterfaceType{
		ident:        ident,
		participants: participants,
		substMap:     substMap,
	}
}

func (t *InterfaceType) Name() string {
	return t.ident.GetText()
}

func (t *InterfaceType) Ident() parser.IIdentContext {
	return t.ident
}
