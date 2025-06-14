package types

import (
	"fmt"
	"iter"

	"github.com/tempo-lang/tempo/parser"
)

type StructType struct {
	baseType
	structIdent  parser.IIdentContext
	participants []string
	fields       TypeFieldMap
}

func (t *StructType) SubstituteRoles(substMap *RoleSubst) Type {
	newParticipants := []string{}
	for _, from := range t.participants {
		newParticipants = append(newParticipants, substMap.Subst(from))
	}

	newFields := TypeFieldMap{}
	for name, fieldType := range t.fields {
		newFields[name] = fieldType.SubstituteRoles(substMap)
	}

	return NewStructType(t.structIdent, newParticipants, newFields)
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

func (t *StructType) IsSendable() bool {
	return true
}

func (t *StructType) IsEquatable() bool {
	return true
}

func (t *StructType) ToString() string {
	return fmt.Sprintf("struct@%s %s", t.Roles().ToString(), t.structIdent.GetText())
}

func NewStructType(structIdent parser.IIdentContext, participants []string, fields TypeFieldMap) Type {
	return &StructType{structIdent: structIdent, participants: participants, fields: fields}
}

func (t *StructType) Name() string {
	return t.structIdent.GetText()
}

func (t *StructType) Ident() parser.IIdentContext {
	return t.structIdent
}

func (t *StructType) Fields() iter.Seq2[string, Type] {
	return func(yield func(string, Type) bool) {
		for k, v := range t.fields {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (t *StructType) Field(name string) (Type, bool) {
	field, found := t.fields[name]
	return field, found
}

// AddField is a special method that populates the fields of the struct type after it has been instantiated.
// Since types should generally not be mutated, using this method requires special care.
func (t *StructType) AddField(name string, fieldType Type) {
	t.fields[name] = fieldType
}
