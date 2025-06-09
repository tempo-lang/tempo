package types

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
)

type InterfaceType struct {
	baseValue
	ident        parser.IIdentContext
	participants []string
}

func (s *InterfaceType) SubstituteRoles(substMap *RoleSubst) Value {
	newParticipants := []string{}
	for _, from := range s.participants {
		newParticipants = append(newParticipants, substMap.Subst(from))
	}

	return NewInterfaceType(s.ident, newParticipants)
}

func (t *InterfaceType) ReplaceSharedRoles(participants []string) Value {
	return t
}

func (t *InterfaceType) Roles() *Roles {
	return NewRole(t.participants, false)
}

func (s *InterfaceType) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(s, other); ok != nil {
		return value, *ok
	}

	if otherInf, ok := other.(*InterfaceType); ok {
		return other, s.ident == otherInf.ident
	}
	return Unit(), false
}

func (s *InterfaceType) IsSendable() bool {
	return false
}

func (t *InterfaceType) IsEquatable() bool {
	return true
}

func (s *InterfaceType) ToString() string {
	return fmt.Sprintf("interface@%s %s", s.Roles().ToString(), s.ident.GetText())
}

func NewInterfaceType(ident parser.IIdentContext, participants []string) Value {
	return &InterfaceType{ident: ident, participants: participants}
}

func (s *InterfaceType) Name() string {
	return s.ident.GetText()
}

func (s *InterfaceType) Ident() parser.IIdentContext {
	return s.ident
}
