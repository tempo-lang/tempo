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

	panic(fmt.Sprintf("unknown expression: %s", expr.GetText()))
}
