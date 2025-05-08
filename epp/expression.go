package epp

import (
	"chorego/parser"
	"chorego/projection"
	"chorego/type_check"
	"chorego/types"
	"fmt"
	"strconv"
)

func eppExpression(info *type_check.Info, roleName string, expr parser.IExprContext) (projection.Expression, []projection.Expression) {

	exprType := info.Types[expr]

	switch expr := expr.(type) {
	case *parser.ExprAddContext:
		lhs, aux := eppExpression(info, roleName, expr.Expr(0))
		rhs, rhsAux := eppExpression(info, roleName, expr.Expr(1))
		aux = append(aux, rhsAux...)
		return projection.NewExprBinaryOp(projection.OperatorAdd, lhs, rhs, exprType.Value()), aux
	case *parser.ExprBoolContext:
		value := expr.TRUE() != nil
		return projection.NewExprBool(value), []projection.Expression{}
	case *parser.ExprGroupContext:
		return eppExpression(info, roleName, expr.Expr())
	case *parser.ExprIdentContext:
		return projection.NewExprIdent(expr.Ident().GetText(), exprType.Value()), []projection.Expression{}
	case *parser.ExprNumContext:
		num, err := strconv.Atoi(expr.GetText())
		if err != nil {
			panic(fmt.Sprintf("expected NUMBER to be convertible to int: %s", expr.GetText()))
		}
		return projection.NewExprInt(num), []projection.Expression{}
	case *parser.ExprAwaitContext:
		inner, aux := eppExpression(info, roleName, expr.Expr())
		asyncType := inner.Type().(*types.Async)
		return projection.NewExprAwait(inner, asyncType.Inner()), aux
	case *parser.ExprComContext:
		sender := expr.RoleType(0).(*parser.RoleTypeNormalContext)
		senderRole := sender.Ident(0).GetText()

		receiver := parser.RoleTypeAllIdents(expr.RoleType(1))
		if len(receiver) != 1 {
			panic("sending shared value not implemented yet")
		}
		receiverRole := receiver[0].GetText()
		inner, aux := eppExpression(info, roleName, expr.Expr())

		if roleName == senderRole {
			aux = append(aux, projection.NewExprSend(inner, receiverRole))
		}

		if roleName == receiverRole {
			return projection.NewExprRecv(inner.Type(), senderRole), aux
		}

		return projection.Unit(), aux
	case *parser.ExprContext:
		panic("expr should never be base type")
	}

	panic(fmt.Sprintf("unknown expression: %#v", expr))
}
