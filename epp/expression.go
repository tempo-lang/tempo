package epp

import (
	"chorego/parser"
	"chorego/projection"
	"fmt"
	"strconv"
)

func eppExpression(expr parser.IExprContext) projection.Expression {

	switch expr := expr.(type) {
	case *parser.ExprAddContext:
		lhs := eppExpression(expr.Expr(0))
		rhs := eppExpression(expr.Expr(1))
		return projection.NewExprBinaryOp(projection.OperatorAdd, lhs, rhs)
	case *parser.ExprBoolContext:
		value := expr.TRUE() != nil
		return projection.NewExprBool(value)
	case *parser.ExprGroupContext:
		return eppExpression(expr.Expr())
	case *parser.ExprIdentContext:
		return projection.NewExprIdent(expr.Ident().GetText())
	case *parser.ExprNumContext:
		num, err := strconv.Atoi(expr.GetText())
		if err != nil {
			panic(fmt.Sprintf("expected NUMBER to be convertible to int: %s", expr.GetText()))
		}
		return projection.NewExprInt(num)
	case *parser.ExprContext:
		panic("expr should never be base type")
	}

	panic(fmt.Sprintf("unknown expression: %#v", expr))
}
