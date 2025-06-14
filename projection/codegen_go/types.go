package codegen_go

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/tempo-lang/tempo/projection"
)

func GenType(t projection.Type) jen.Code {
	switch t := t.(type) {
	case *projection.AsyncType:
		return GenAsyncType(t)
	case projection.BuiltinType:
		return GenBuiltinType(t)
	case *projection.ClosureType:
		return GenClosureType(t)
	case *projection.FunctionType:
		return GenFunctionType(t)
	case *projection.InterfaceType:
		return GenInterfaceType(t)
	case *projection.ListType:
		return GenListType(t)
	case *projection.StructType:
		return GenStructType(t)
	}

	panic(fmt.Sprintf("failed to generate unknown type: %#v", t))
}

func GenBuiltinType(t projection.BuiltinType) jen.Code {
	switch t {
	case projection.IntType:
		return jen.Int()
	case projection.FloatType:
		return jen.Float64()
	case projection.StringType:
		return jen.String()
	case projection.BoolType:
		return jen.Bool()
	default:
		panic(fmt.Sprintf("unknown builtin type: %s", t))
	}
}

func GenAsyncType(t *projection.AsyncType) jen.Code {
	innerType := GenType(t.Inner)
	return jen.Op("*").Qual("github.com/tempo-lang/tempo/runtime", "Async").Types(innerType)
}

func GenListType(l *projection.ListType) jen.Code {
	return jen.Index().Add(GenType(l.Inner))
}
