package epp

import (
	"chorego/parser"
	"chorego/projection"
	"chorego/type_check"
	"chorego/types"
	"fmt"
	"strconv"
)

func eppExpression(info *type_check.Info, expr parser.IExprContext) projection.Expression {

	exprType := info.Types[expr]

	switch expr := expr.(type) {
	case *parser.ExprAddContext:
		lhs := eppExpression(info, expr.Expr(0))
		rhs := eppExpression(info, expr.Expr(1))
		return projection.NewExprBinaryOp(projection.OperatorAdd, lhs, rhs, exprType.Value())
	case *parser.ExprBoolContext:
		value := expr.TRUE() != nil
		return projection.NewExprBool(value)
	case *parser.ExprGroupContext:
		return eppExpression(info, expr.Expr())
	case *parser.ExprIdentContext:
		return projection.NewExprIdent(expr.Ident().GetText(), exprType.Value())
	case *parser.ExprNumContext:
		num, err := strconv.Atoi(expr.GetText())
		if err != nil {
			panic(fmt.Sprintf("expected NUMBER to be convertible to int: %s", expr.GetText()))
		}
		return projection.NewExprInt(num)
	case *parser.ExprAwaitContext:
		inner := eppExpression(info, expr.Expr())
		asyncType := inner.Type().(*types.Async)
		return projection.NewExprAwait(inner, asyncType.Inner())
	case *parser.ExprContext:
		panic("expr should never be base type")
	}

	panic(fmt.Sprintf("unknown expression: %#v", expr))
}
