package type_check

import "tempo/parser"

func (tc *typeChecker) VisitInterface(ctx *parser.InterfaceContext) any {
	return nil
}

func (tc *typeChecker) VisitInterfaceMethodsList(ctx *parser.InterfaceMethodsListContext) any {
	return nil
}

func (tc *typeChecker) VisitInterfaceMethod(ctx *parser.InterfaceMethodContext) any {
	return nil
}
