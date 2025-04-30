package types

import (
	"chorego/parser"
)

func BuiltinTypes() map[string]Type {
	return map[string]Type{
		"Int":    Int(),
		"Float":  Float(),
		"String": String(),
		"Bool":   Bool(),
	}
}

func ParseValueType(ctx parser.IValueTypeContext) (Type, Error) {
	typeName := ctx.Ident()
	builtinType, isBuiltinType := BuiltinTypes()[typeName.GetText()]
	if isBuiltinType {
		return builtinType, nil
	}

	return nil, NewUnknownTypeError(typeName)
}

func ParseFuncType(ctx parser.IFuncContext) (Type, Error) {

	params := []Type{}
	paramErrors := map[int]Error{}

	for i, param := range ctx.FuncParamList().AllFuncParam() {
		paramType, err := ParseValueType(param.ValueType())
		if err != nil {
			paramErrors[i] = err
		}
		params = append(params, paramType)
	}

	if len(paramErrors) > 0 {
		return Function(params), NewInvalidFuncError(ctx, paramErrors)
	}

	return Function(params), nil
}
