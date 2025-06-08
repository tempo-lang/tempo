package types

import "fmt"

type List struct {
	baseValue
	inner Value
}

func (l *List) CoerceTo(other Value) (Value, bool) {
	panic("unimplemented")
}

func (l *List) IsEquatable() bool {
	return l.inner.IsEquatable()
}

func (l *List) IsSendable() bool {
	return l.inner.IsSendable()
}

func (l *List) SubstituteRoles(substMap *RoleSubst) Value {
	return NewList(l.inner.SubstituteRoles(substMap))
}

func (l *List) ReplaceSharedRoles(participants []string) Value {
	return NewList(l.inner.ReplaceSharedRoles(participants))
}

func (l *List) Roles() *Roles {
	return l.inner.Roles()
}

func (l *List) ToString() string {
	return fmt.Sprintf("[%s]", l.inner.ToString())
}

func NewList(inner Value) Value {
	return &List{
		inner: inner,
	}
}
