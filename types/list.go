package types

import (
	"fmt"
	"iter"
)

type ListType struct {
	baseType
	inner  Type
	fields TypeFieldMap
}

func (l *ListType) CoerceTo(other Type) (Type, bool) {
	if other, isOtherList := other.(*ListType); isOtherList {
		coerceType, ok := l.inner.CoerceTo(other.inner)
		return List(coerceType), ok
	}
	return Invalid(), false
}

func (l *ListType) IsEquatable() bool {
	return l.inner.IsEquatable()
}

func (l *ListType) IsSendable() bool {
	return l.inner.IsSendable()
}

func (l *ListType) SubstituteRoles(substMap *RoleSubst) Type {
	return List(l.inner.SubstituteRoles(substMap))
}

func (l *ListType) ReplaceSharedRoles(participants []string) Type {
	return List(l.inner.ReplaceSharedRoles(participants))
}

func (l *ListType) Roles() *Roles {
	return l.inner.Roles()
}

func (l *ListType) ToString() string {
	return fmt.Sprintf("[%s]", l.inner.ToString())
}

func (l *ListType) Inner() Type {
	return l.inner
}

func List(inner Type) Type {
	return &ListType{
		inner:  inner,
		fields: listFields(inner),
	}
}

func listFields(innerType Type) TypeFieldMap {
	result := TypeFieldMap{}

	participants := innerType.Roles().participants

	result["length"] = Int(participants)

	return result
}

func (l *ListType) Fields() iter.Seq2[string, Type] {
	return func(yield func(string, Type) bool) {
		for k, v := range l.fields {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (l *ListType) Field(name string) (Type, bool) {
	fieldType, found := l.fields[name]
	return fieldType, found
}
