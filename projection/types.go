package projection

import (
	"fmt"
)

type Type interface {
	IsType()
}

type BuiltinType string

const (
	IntType    BuiltinType = "Int"
	FloatType  BuiltinType = "Float"
	StringType BuiltinType = "String"
	BoolType   BuiltinType = "Bool"
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
