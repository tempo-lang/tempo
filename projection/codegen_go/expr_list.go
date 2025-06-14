package codegen_go

import (
	"github.com/dave/jennifer/jen"
	"github.com/tempo-lang/tempo/projection"
)

func GenExprList(e *projection.ExprList) jen.Code {
	items := make([]jen.Code, len(e.Items))
	for i, item := range e.Items {
		items[i] = GenExpression(item)
	}
	return jen.Add(GenType(e.Type())).Values(items...)
}

func GenExprIndex(e *projection.ExprIndex) jen.Code {
	return jen.Add(GenExpression(e.Base)).Index(GenExpression(e.Index))
}

func GenExprListLength(e *projection.ExprListLength) jen.Code {
	return jen.Len(GenExpression(e.List))
}
