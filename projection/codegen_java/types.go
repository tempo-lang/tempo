package codegen_java

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
	gen.AddImport(javaPkgFuture)
	return fmt.Sprintf("Future<%s>", gen.GenType(t.Inner))
}

func (gen *codegen) GenBuiltinType(t projection.BuiltinType) string {
	switch t {
	case projection.BoolType:
		return "Boolean"
	case projection.FloatType:
		return "Double"
	case projection.IntType:
		return "Integer"
	case projection.StringType:
		return "String"
	}
	panic(fmt.Sprintf("unexpected projection.BuiltinType: %#v", t))
}

func (gen *codegen) GenCallableType(t projection.CallableType) string {
	returnsVoid := t.ReturnType() == projection.UnitType()
	params := []string{}

	if !returnsVoid {
		params = append(params, gen.GenType(t.ReturnType()))
	}

	for _, param := range t.Params() {
		params = append(params, gen.GenType(param))
	}

	var javaType string
	if t.ReturnType() != projection.UnitType() {
		javaType = fmt.Sprintf("FnRet%d", len(t.Params()))
	} else {
		javaType = fmt.Sprintf("Fn%d", len(t.Params()))
	}

	gen.AddImport(funcTypeToPkg(javaType))
	return fmt.Sprintf("%s<%s>", javaType, misc.JoinStrings(params, ", "))
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
	gen.AddImport(javaPkgArrayList)
	return fmt.Sprintf("ArrayList<%s>", gen.GenType(t.Inner))
}

func (gen *codegen) GenStructType(t *projection.StructType) string {
	return t.GenName()
}
