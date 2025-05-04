package types

import "fmt"

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

func (t *Type) CanCoerceTo(other *Type) bool {
	if t.value == Invalid().value {
		return true
	}

	if t.value != other.value {
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
