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

// func CodegenType(t types.Value) jen.Code {
// 	if builtinType, isBuiltin := t.(types.Builtin); isBuiltin {
// 		return CodegenBuiltinType(builtinType)
// 	}

// 	if asyncType, isAsync := t.(*types.Async); isAsync {
// 		return CodegenAsyncType(asyncType)
// 	}

// 	if structType, isStruct := t.(*StructType); isStruct {
// 		return CodegenStructType(structType)
// 	}
// 	if _, ok := t.(*types.StructType); ok {
// 		panic(fmt.Sprintf("struct %#v should be of type projection.StructType instead", t))
// 	}

// 	if infType, isInf := t.(*InterfaceType); isInf {
// 		return CodegenInterfaceType(infType)
// 	}
// 	if _, ok := t.(*types.InterfaceType); ok {
// 		panic(fmt.Sprintf("struct %#v should be of type projection.InterfaceType instead", t))
// 	}

// 	if funcType, isFunc := t.(*FunctionType); isFunc {
// 		return CodegenFuncType(funcType)
// 	}
// 	if _, ok := t.(*types.FunctionType); ok {
// 		panic(fmt.Sprintf("function %#v should be of type projection.FunctionType instead", t))
// 	}

// 	panic(fmt.Sprintf("failed to generate type: %#v", t))
// }

type BuiltinType struct {
	types.Value
}

func (c *BuiltinType) IsType() {}

func (t *BuiltinType) Codegen() jen.Code {
	switch t.Value.(type) {
	case *types.IntType:
		return jen.Int()
	case *types.FloatType:
		return jen.Float64()
	case *types.StringType:
		return jen.String()
	case *types.BoolType:
		return jen.Bool()
	default:
		panic(fmt.Sprintf("unknown builtin type: %s", t.Value.ToString()))
	}
}

type AsyncType struct {
	Inner Type
}

func (c *AsyncType) IsType() {}

func NewAsyncType(inner Type) *AsyncType {
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

type unitType struct{}

var unit unitType = unitType{}

func UnitType() Type {
	return &unit
}

func (c *unitType) IsType() {}

func (t *unitType) Codegen() jen.Code {
	panic("attempt to codegen unit type")
}
