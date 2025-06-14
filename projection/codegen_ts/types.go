package codegen_ts

import (
	"fmt"

	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenType(t projection.Type) string {
	switch t := t.(type) {
	case *projection.AsyncType:
		return gen.GenAsyncType(t)
	case projection.BuiltinType:
		return gen.GenBuiltinType(t)
	case *projection.ClosureType:
		return gen.GenClosureType(t)
	case *projection.FunctionType:
		return gen.GenFunctionType(t)
	case *projection.InterfaceType:
		return gen.GenInterfaceType(t)
	case *projection.ListType:
		return gen.GenListType(t)
	case *projection.StructType:
		return gen.GenStructType(t)
	default:
	}

	panic(fmt.Sprintf("unexpected projection.Type: %#v", t))
}

func (gen *codegen) GenAsyncType(t *projection.AsyncType) string {
	return fmt.Sprintf("Promise<%s>", gen.GenType(t.Inner))
}

func (gen *codegen) GenBuiltinType(t projection.BuiltinType) string {
	switch t {
	case projection.BoolType:
		return "boolean"
	case projection.FloatType:
		return "number"
	case projection.IntType:
		return "number"
	case projection.StringType:
		return "string"
	}
	panic(fmt.Sprintf("unexpected projection.BuiltinType: %#v", t))
}

func (gen *codegen) GenClosureType(t *projection.ClosureType) string {
	return "[ClosureType]"
}

func (gen *codegen) GenFunctionType(t *projection.FunctionType) string {
	return "[FunctionType]"
}

func (gen *codegen) GenInterfaceType(t *projection.InterfaceType) string {
	return fmt.Sprintf("%s_%s", t.Name(), t.Role())
}

func (gen *codegen) GenListType(t *projection.ListType) string {
	return "[ListType]"
}

func (gen *codegen) GenStructType(t *projection.StructType) string {
	return fmt.Sprintf("%s_%s", t.Name(), t.Role())
}
