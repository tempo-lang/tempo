package type_check

import (
	"chorego/parser"
	"chorego/type_check/type_error"
	"chorego/type_check/types"
)

func BuiltinTypes() map[string]types.Type {
	return map[string]types.Type{
		"Int":    types.Int(),
		"Float":  types.Float(),
		"String": types.String(),
		"Bool":   types.Bool(),
	}
}

func ParseValueType(ctx parser.IValueTypeContext) (types.Type, *type_error.UnknownTypeError) {
	typeName := ctx.Ident()
	builtinType, isBuiltinType := BuiltinTypes()[typeName.GetText()]
	if isBuiltinType {
		return builtinType, nil
	}

	return nil, type_error.NewUnknownTypeError(typeName)
}
