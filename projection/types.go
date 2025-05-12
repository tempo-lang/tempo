package projection

import (
	"fmt"
	"tempo/types"

	"github.com/dave/jennifer/jen"
)

func CodegenType(t types.Value) jen.Code {
	if builtinType, isBuiltin := t.(types.Builtin); isBuiltin {
		return CodegenBuiltinType(builtinType)
	}

	if asyncType, isAsync := t.(*types.Async); isAsync {
		return CodegenAsyncType(asyncType)
	}

	panic(fmt.Sprintf("failed to generate type: %v", t))
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
	return jen.Op("*").Qual("tempo/runtime", "Async")
}
