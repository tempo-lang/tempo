package types

import "fmt"

type Async struct {
	baseValue
	inner Value
}

func (a *Async) SubstituteRoles(substMap *RoleSubst) Value {
	return NewAsync(a.inner.SubstituteRoles(substMap))
}

func (a *Async) ReplaceSharedRoles(participants []string) Value {
	return NewAsync(a.inner.ReplaceSharedRoles(participants))
}

func (t *Async) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(t, other); ok != nil {
		return value, *ok
	}

	if otherAsync, ok := other.(*Async); ok {
		if newValue, canCoerce := t.inner.CoerceTo(otherAsync.inner); canCoerce {
			return NewAsync(newValue), true
		} else {
			return NewAsync(Invalid()), false
		}
	}
	return Invalid(), false
}

func (t *Async) Roles() *Roles {
	return t.inner.Roles()
}

func (a *Async) IsSendable() bool {
	return false
}

func (t *Async) IsEquatable() bool {
	return false
}

func (a *Async) ToString() string {
	return fmt.Sprintf("async %s", a.inner.ToString())
}

func (a *Async) Inner() Value {
	return a.inner
}

func NewAsync(inner Value) Value {
	if _, ok := inner.(*Async); ok {
		panic("nested async type")
	}

	return &Async{inner: inner}
}
