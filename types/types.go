package types

import (
	"fmt"
	"reflect"

	"github.com/tempo-lang/tempo/misc"
)

type Type struct {
	value Value
	roles *Roles
}

func New(value Value, roles *Roles) *Type {
	return &Type{
		value: value,
		roles: roles,
	}
}

func ValuesEqual(a, b Value) bool {
	return reflect.DeepEqual(a, b)
}

func baseCoerceValue(thisValue, otherValue Value) (Value, bool) {
	if thisValue == Invalid().Value() {
		return otherValue, true
	}

	if otherValue == Invalid().Value() {
		return Invalid().Value(), true
	}

	// plain types can coerce to async types
	if _, isAsync := thisValue.(*Async); !isAsync {
		if otherAsync, otherIsAsync := otherValue.(*Async); otherIsAsync {
			innerCoerce, ok := thisValue.CoerceTo(otherAsync.Inner())
			return NewAsync(innerCoerce), ok
		}
	}

	return nil, false
}

func (t *Type) CoerceTo(other *Type) (*Type, bool) {
	newValue, ok := t.value.CoerceTo(other.value)
	if !ok {
		return Invalid(), false
	}

	newType := New(newValue, other.Roles())

	if t.roles.participants == nil {
		return newType, true
	}

	if other.roles.participants == nil {
		return newType, false
	}

	if t.roles.IsSharedRole() {
		return newType, t.roles.Encompass(other.roles)
	}

	if len(t.roles.participants) != len(other.roles.participants) {
		return newType, false
	}

	for i, role := range t.roles.participants {
		if role != other.roles.participants[i] {
			return newType, false
		}
	}

	return newType, true
}

func (t *Type) Roles() *Roles {
	return t.roles
}

func (t *Type) Value() Value {
	return t.value
}

func (t *Type) ToString() string {

	switch value := t.value.(type) {
	case *FunctionType:
		params := misc.JoinStringsFunc(value.params, ", ", func(param *Type) string { return param.ToString() })
		returnType := ""
		if value.returnType.Value() != Unit() {
			returnType = value.returnType.ToString()
		}
		return fmt.Sprintf("func@%s(%s)%s", t.roles.ToString(), params, returnType)
	case *StructType:
		return fmt.Sprintf("struct@%s %s", t.roles.ToString(), value.Name())
	case *InterfaceType:
		return fmt.Sprintf("interface@%s %s", t.roles.ToString(), value.Name())
	}

	return fmt.Sprintf("%s@%s", t.value.ToString(), t.roles.ToString())
}

func (t *Type) IsInvalid() bool {
	return t.Value() == Invalid().value
}

type Value interface {
	IsSendable() bool
	IsEquatable() bool
	ToString() string
	IsValue()
	SubstituteRoles(substMap *RoleSubst) Value
	CoerceTo(other Value) (Value, bool)
}

type InvalidValue struct{}

func (t *InvalidValue) SubstituteRoles(substMap *RoleSubst) Value {
	return t
}

func (t *InvalidValue) CoerceTo(other Value) (Value, bool) {
	return other, true
}

var invalid_type InvalidValue = InvalidValue{}

func (t *InvalidValue) IsSendable() bool {
	return true
}

func (t *InvalidValue) IsEquatable() bool {
	return true
}

func (t *InvalidValue) ToString() string {
	return "ERROR"
}

func (t *InvalidValue) IsValue() {}

func Invalid() *Type {
	return New(&invalid_type, NewRole(nil, false))
}

type UnitValue struct{}

func (u *UnitValue) SubstituteRoles(substMap *RoleSubst) Value {
	return u
}

func (t *UnitValue) CoerceTo(other Value) (Value, bool) {
	return Unit(), other == Unit()
}

func (u *UnitValue) IsValue() {}

func (u *UnitValue) ToString() string {
	return "()"
}

func (u *UnitValue) IsSendable() bool {
	return true
}

func (t *UnitValue) IsEquatable() bool {
	return false
}

var unit_value UnitValue = UnitValue{}

func Unit() Value {
	return &unit_value
}
