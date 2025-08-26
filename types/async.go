package types

import "fmt"

type AsyncType struct {
	baseType
	inner Type
}

func (a *AsyncType) SubstituteRoles(substMap *RoleSubst) Type {
	return Async(a.inner.SubstituteRoles(substMap))
}

func (a *AsyncType) ReplaceSharedRoles(participants []string) Type {
	return Async(a.inner.ReplaceSharedRoles(participants))
}

func (t *AsyncType) CoerceTo(other Type) (Type, bool) {
	if value, ok := baseCoerceValue(t, other); ok != nil {
		return value, *ok
	}

	if otherAsync, ok := other.(*AsyncType); ok {
		if newValue, canCoerce := t.inner.CoerceTo(otherAsync.inner); canCoerce {
			return Async(newValue), true
		} else {
			return Async(Invalid()), false
		}
	}
	return Invalid(), false
}

func (t *AsyncType) Roles() *Roles {
	return t.inner.Roles()
}

func (t *AsyncType) IsEquatable() bool {
	return false
}

func (a *AsyncType) ToString() string {
	return fmt.Sprintf("async %s", a.inner.ToString())
}

func (a *AsyncType) Inner() Type {
	return a.inner
}

func Async(inner Type) Type {
	if _, ok := inner.(*AsyncType); ok {
		panic("nested async type")
	}

	return &AsyncType{inner: inner}
}
