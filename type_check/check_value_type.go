package type_check

import (
	"chorego/parser"
	"chorego/type_check/type_error"
	"slices"
)

func BuiltinTypes() []string {
	return []string{"Int", "Float", "String", "Bool"}
}

func (tc *typeChecker) checkValueType(ctx parser.IValueTypeContext) {

	typeName := ctx.Ident()
	isBuiltin := slices.Contains(BuiltinTypes(), typeName.GetText())

	if !isBuiltin {
		tc.ErrorListener.ReportTypeError(type_error.NewUnknownTypeError(typeName))
	}
}
