package type_check

import (
	"chorego/parser"
	"chorego/type_check/type_error"
	"slices"
)

func BuiltinTypes() []string {
	return []string{"Int", "Float", "String", "Bool"}
}

func (tc *TypeChecker) checkValueType(ctx parser.IValueTypeContext) {

	typeName := ctx.Ident()
	isBuiltin := slices.Contains(BuiltinTypes(), typeName.GetText())

	if !isBuiltin {
		tc.ErrorListener.ReportAnalyzerError(type_error.NewUnknownTypeError(typeName))
	}
}
