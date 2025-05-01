package projection

import (
	"chorego/types"
	"fmt"
)

func CodegenType(t *types.Type) string {
	if builtinType, isBuiltin := t.Value().(types.Builtin); isBuiltin {
		return CodegenBuiltinType(builtinType)
	}

	panic(fmt.Sprintf("failed to generate type: %v", t))
}

func CodegenBuiltinType(builtinType types.Builtin) string {
	switch builtinType.(type) {
	case *types.IntType:
		return "int"
	case *types.FloatType:
		return "float64"
	case *types.StringType:
		return "string"
	case *types.BoolType:
		return "bool"
	default:
		panic(fmt.Sprintf("unknown builtin type: %s", builtinType.ToString()))
	}
}
