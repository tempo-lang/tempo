package projection

import (
	"fmt"

	"github.com/tempo-lang/tempo/types"

	"github.com/dave/jennifer/jen"
)

func CodegenType(t types.Value) jen.Code {
	if builtinType, isBuiltin := t.(types.Builtin); isBuiltin {
		return CodegenBuiltinType(builtinType)
	}

	if asyncType, isAsync := t.(*types.Async); isAsync {
		return CodegenAsyncType(asyncType)
	}

	if structType, isStruct := t.(*StructType); isStruct {
		return CodegenStructType(structType)
	}
	if _, ok := t.(*types.StructType); ok {
		panic(fmt.Sprintf("struct %#v should be of type projection.StructType instead", t))
	}

	if infType, isInf := t.(*InterfaceType); isInf {
		return CodegenInterfaceType(infType)
	}
	if _, ok := t.(*types.InterfaceType); ok {
		panic(fmt.Sprintf("struct %#v should be of type projection.InterfaceType instead", t))
	}

	if funcType, isFunc := t.(*FunctionType); isFunc {
		return CodegenFuncType(funcType)
	}
	if _, ok := t.(*types.FunctionType); ok {
		panic(fmt.Sprintf("function %#v should be of type projection.FunctionType instead", t))
	}

	panic(fmt.Sprintf("failed to generate type: %#v", t))
}

func CodegenFuncType(funcType *FunctionType) jen.Code {
	params := []jen.Code{}
	for _, param := range funcType.Params {
		params = append(params, CodegenType(param))
	}

	result := jen.Func().Params(params...)

	if funcType.ReturnType != types.Unit() {
		result = result.Add(CodegenType(funcType.ReturnType))
	}

	return result
}

func CodegenStructType(structType *StructType) jen.Code {
	return jen.Id(fmt.Sprintf("%s_%s", structType.Name(), structType.Role()))
}

func CodegenInterfaceType(infType *InterfaceType) jen.Code {
	return jen.Id(fmt.Sprintf("%s_%s", infType.Name(), infType.Role()))
}

func CodegenBuiltinType(builtinType types.Builtin) jen.Code {
	switch builtinType.(type) {
	case *types.IntType:
		return jen.Int()
	case *types.FloatType:
		return jen.Float64()
	case *types.StringType:
		return jen.String()
	case *types.BoolType:
		return jen.Bool()
	default:
		panic(fmt.Sprintf("unknown builtin type: %s", builtinType.ToString()))
	}
}

func CodegenAsyncType(asyncType *types.Async) jen.Code {
	innerType := CodegenType(asyncType.Inner())
	return jen.Op("*").Qual("github.com/tempo-lang/tempo/runtime", "Async").Types(innerType)
}
