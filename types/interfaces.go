package types

import (
	"fmt"
	"iter"

	"github.com/tempo-lang/tempo/parser"
)

type InterfaceType struct {
	baseType
	ident        parser.IIdentContext
	participants []string
	fields       TypeFieldMap
}

func (t *InterfaceType) SubstituteRoles(substMap *RoleSubst) Type {
	newParticipants := []string{}
	for _, from := range t.participants {
		newParticipants = append(newParticipants, substMap.Subst(from))
	}

	newFields := TypeFieldMap{}
	for name, fieldType := range t.fields {
		newFields[name] = fieldType.SubstituteRoles(substMap)
	}

	return Interface(t.ident, newParticipants, newFields)
}

func (t *InterfaceType) ReplaceSharedRoles(participants []string) Type {
	return t
}

func (t *InterfaceType) Roles() *Roles {
	return NewRole(t.participants, false)
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

func (s *InterfaceType) IsSendable() bool {
	return false
}

func (t *InterfaceType) IsEquatable() bool {
	return true
}

func (t *InterfaceType) ToString() string {
	return fmt.Sprintf("interface@%s %s", t.Roles().ToString(), t.ident.GetText())
}

func Interface(ident parser.IIdentContext, participants []string, fields TypeFieldMap) Type {
	return &InterfaceType{ident: ident, participants: participants, fields: fields}
}

func (t *InterfaceType) Name() string {
	return t.ident.GetText()
}

func (t *InterfaceType) Ident() parser.IIdentContext {
	return t.ident
}

func (t *InterfaceType) Fields() iter.Seq2[string, Type] {
	return func(yield func(string, Type) bool) {
		for k, v := range t.fields {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (t *InterfaceType) Field(name string) (Type, bool) {
	field, found := t.fields[name]
	return field, found
}

// AddField is a special method that populates the fields of the interface type after it has been instantiated.
// Since types should generally not be mutated, using this method requires special care.
func (t *InterfaceType) AddField(name string, fieldType Type) {
	t.fields[name] = fieldType
}
