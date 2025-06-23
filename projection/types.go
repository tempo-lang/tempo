package projection

import (
	"fmt"
)

type Type interface {
	IsType()
}

// A primitive builtin type
type BuiltinType string

const (
	// A signed integer with at least 32 bit resolution.
	IntType BuiltinType = "Int"
	// A floating point number with at least 32 bit resolution.
	FloatType BuiltinType = "Float"
	// A string that supports unicode symbols.
	StringType BuiltinType = "String"
	// A boolean value that can only either be `true` or `false`.
	BoolType BuiltinType = "Bool"
)

func (c BuiltinType) IsType() {}

type AsyncType struct {
	Inner Type
}

func (c *AsyncType) IsType() {}

func NewAsyncType(inner Type) Type {
	if _, innerAsync := inner.(*AsyncType); innerAsync {
		panic(fmt.Sprintf("nested async type: %#v", inner))
	}

	return &AsyncType{
		Inner: inner,
	}
}

// A variable size list containing values of the underlying type.
type ListType struct {
	Inner Type
}

func (l *ListType) IsType() {}

func NewListType(inner Type) Type {
	return &ListType{
		Inner: inner,
	}
}

type unitType struct{}

var unit unitType = unitType{}

func UnitType() Type {
	return &unit
}

func (c *unitType) IsType() {}
