package type_check

import "tempo/parser"

func (tc *typeChecker) VisitStruct(ctx *parser.StructContext) any {
	return nil
}

func (tc *typeChecker) VisitStructFieldList(ctx *parser.StructFieldListContext) any {
	return nil
}

func (tc *typeChecker) VisitStructField(ctx *parser.StructFieldContext) any {
	return nil
}
