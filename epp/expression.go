package epp

import (
	"chorego/parser"
	"chorego/projection"
	"fmt"
	"strconv"
)

func eppExpression(role parser.IIdentContext, expr parser.IExpressionContext) projection.Expression {
	if numToken := expr.NUMBER(); numToken != nil {
		num, err := strconv.Atoi(numToken.GetText())
		if err != nil {
			panic(fmt.Sprintf("expected NUMBER to be convertible to int: %s", numToken.GetText()))
		}

		return projection.NewExprInt(num)
	}

	if ident := expr.Ident(); ident != nil {
		return projection.NewExprIdent(ident.GetText())
	}

	if boolean := expr.ExprBool(); boolean != nil {
		value := boolean.TRUE() != nil
		return projection.NewExprBool(value)
	}

	panic(fmt.Sprintf("unknown expression: %s", expr.GetText()))
}
