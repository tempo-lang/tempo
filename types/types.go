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

func (t *Type) CanCoerceTo(other *Type) bool {
	otherValue := other.value
	if t.value == Invalid().value {
		return true
	}

	// plain types can coerce to async types
	if _, isAsync := t.value.(*Async); !isAsync {
		if otherAsync, otherIsAsync := otherValue.(*Async); otherIsAsync {
			otherValue = otherAsync.Inner()
		}
	}

	if !ValuesEqual(t.value, otherValue) {
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

type Value interface {
	ToString() string
	IsValue()
}

type InvalidValue struct{}

var invalid_type InvalidValue = InvalidValue{}

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

var unit_value UnitValue = UnitValue{}

func Unit() Value {
	return &unit_value
}
