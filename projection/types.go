package projection

import (
	"fmt"

	"github.com/tempo-lang/tempo/types"

	"github.com/dave/jennifer/jen"
)

type Type interface {
	IsType()
	Codegen() jen.Code
}

type BuiltinType struct {
	types.Type
}

func (c *BuiltinType) IsType() {}

func (t *BuiltinType) Codegen() jen.Code {
	switch t.Type.(type) {
	case *types.IntType:
		return jen.Int()
	case *types.FloatType:
		return jen.Float64()
	case *types.StringType:
		return jen.String()
	case *types.BoolType:
		return jen.Bool()
	default:
		panic(fmt.Sprintf("unknown builtin type: %s", t.Type.ToString()))
	}
}

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

func (t *AsyncType) Codegen() jen.Code {
	innerType := t.Inner.Codegen()
	return jen.Op("*").Qual("github.com/tempo-lang/tempo/runtime", "Async").Types(innerType)
}

type ListType struct {
	Inner Type
}

func (l *ListType) Codegen() jen.Code {
	return jen.Index().Add(l.Inner.Codegen())
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

func (t *unitType) Codegen() jen.Code {
	panic("attempt to codegen unit type")
}
