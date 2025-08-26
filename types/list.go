package types

import (
	"fmt"
)

type ListType struct {
	baseType
	inner Type
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
		inner: inner,
	}
}
