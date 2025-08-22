package codegen_ts

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenType(t projection.Type) string {
	if t == projection.UnitType() {
		return "void"
	}

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

func (gen *codegen) GenCallableType(t projection.CallableType) string {
	params := []string{"env: Env"}
	for i, param := range t.Params() {
		params = append(params, fmt.Sprintf("arg%d: %s", i, gen.GenType(param)))
	}
	return fmt.Sprintf("(%s) => Promise<%s>", misc.JoinStrings(params, ", "), gen.GenType(t.ReturnType()))
}

func (gen *codegen) GenClosureType(t *projection.ClosureType) string {
	return gen.GenCallableType(t)
}

func (gen *codegen) GenFunctionType(t *projection.FunctionType) string {
	return gen.GenCallableType(t)
}

func (gen *codegen) GenInterfaceType(t *projection.InterfaceType) string {
	return fmt.Sprintf("%s_%s", t.Name(), t.Role())
}

func (gen *codegen) GenListType(t *projection.ListType) string {
	return gen.GenType(t.Inner) + "[]"
}

func (gen *codegen) GenStructType(t *projection.StructType) string {
	return t.GenName()
}
