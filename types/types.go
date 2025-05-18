package types

import (
	"fmt"
	"reflect"
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

func ValueCoerseTo(thisValue, otherValue Value) bool {
	if thisValue == Invalid().Value() || otherValue == Invalid().Value() {
		return true
	}

	// plain types can coerce to async types
	if _, isAsync := thisValue.(*Async); !isAsync {
		if otherAsync, otherIsAsync := otherValue.(*Async); otherIsAsync {
			otherValue = otherAsync.Inner()
		}
	}

	if !ValuesEqual(thisValue, otherValue) {
		return false
	}

	return true
}

func (t *Type) CanCoerceTo(other *Type) bool {
	if !ValueCoerseTo(t.value, other.value) {
		return false
	}

	if t.roles.participants == nil {
		return true
	}

	if other.roles.participants == nil {
		return false
	}

	if t.roles.Encompass(other.roles) {
		return true
	}

	return false
}

func (t *Type) Roles() *Roles {
	return t.roles
}

func (t *Type) Value() Value {
	return t.value
}

func (t *Type) ToString() string {
	return fmt.Sprintf("%s@%s", t.value.ToString(), t.roles.ToString())
}

func (t *Type) IsInvalid() bool {
	return t.Value() == Invalid().value
}

type Value interface {
	IsSendable() bool
	ToString() string
	IsValue()
}

type InvalidValue struct{}

var invalid_type InvalidValue = InvalidValue{}

func (t *InvalidValue) IsSendable() bool {
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

func (u *UnitValue) IsValue() {}

func (u *UnitValue) ToString() string {
	return "()"
}

func (u *UnitValue) IsSendable() bool {
	return true
}

var unit_value UnitValue = UnitValue{}

func Unit() Value {
	return &unit_value
}
