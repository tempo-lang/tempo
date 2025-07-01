package codegen_go

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/tempo-lang/tempo/projection"
)

func GenExpression(expr projection.Expression) jen.Code {
	switch e := expr.(type) {
	case *projection.ExprAsync:
		return GenExprAsync(e)
	case *projection.ExprAwait:
		return GenExprAwait(e)
	case *projection.ExprBinaryOp:
		return GenExprBinaryOp(e)
	case *projection.ExprBool:
		return GenExprBool(e)
	case *projection.ExprCallClosure:
		return GenExprCallClosure(e)
	case *projection.ExprCallFunc:
		return GenExprCallFunc(e)
	case *projection.ExprClosure:
		return GenExprClosure(e)
	case *projection.ExprFieldAccess:
		return GenExprFieldAccess(e)
	case *projection.ExprFloat:
		return GenExprFloat(e)
	case *projection.ExprIdent:
		return GenExprIdent(e)
	case *projection.ExprIndex:
		return GenExprIndex(e)
	case *projection.ExprInt:
		return GenExprInt(e)
	case *projection.ExprList:
		return GenExprList(e)
	case *projection.ExprListLength:
		return GenExprListLength(e)
	case *projection.ExprRecv:
		return GenExprRecv(e)
	case *projection.ExprSend:
		return GenExprSend(e)
	case *projection.ExprString:
		return GenExprString(e)
	case *projection.ExprStruct:
		return GenExprStruct(e)
	case *projection.ExprPassValue:
		return GenExprCopy(e)
	default:
		panic(fmt.Sprintf("unexpected projection.Expression: %#v", e))
	}
}

func GenExprInt(e *projection.ExprInt) jen.Code {
	return jen.Lit(e.Value)
}

func GenExprFloat(e *projection.ExprFloat) jen.Code {
	return jen.Lit(e.Value)
}

func GenExprString(e *projection.ExprString) jen.Code {
	return jen.Lit(e.Value)
}

func GenExprIdent(e *projection.ExprIdent) jen.Code {
	return jen.Id(e.Name)
}

func GenExprBool(e *projection.ExprBool) jen.Code {
	if e.Value {
		return jen.True()
	} else {
		return jen.False()
	}
}

func GenExprBinaryOp(e *projection.ExprBinaryOp) jen.Code {
	if _, isList := e.Type().(*projection.ListType); isList {
		return jen.Append(GenExpression(e.Lhs), jen.Add(GenExpression(e.Rhs)).Op("..."))
	}

	return jen.Add(GenExpression(e.Lhs)).Op(string(e.Operator)).Add(GenExpression(e.Rhs))
}

func GenExprAsync(e *projection.ExprAsync) jen.Code {
	return RuntimeFunc("FixedAsync").Call(GenExpression(e.Inner()))
}

func GenExprAwait(e *projection.ExprAwait) jen.Code {
	return RuntimeFunc("GetAsync").Call(GenExpression(e.Inner()))
}

func GenExprSend(e *projection.ExprSend) jen.Code {
	args := []jen.Code{
		jen.Id("env"),
		GenExpression(e.Expr),
	}

	for _, role := range e.Receivers {
		args = append(args, jen.Lit(role))
	}

	return RuntimeFunc("Send").Call(args...)
}

func GenExprRecv(e *projection.ExprRecv) jen.Code {
	recvType := GenType(e.RecvType)
	return RuntimeFunc("Recv").Types(recvType).Call(jen.Id("env"), jen.Lit(e.Sender))
}

func GenExprCallFunc(e *projection.ExprCallFunc) jen.Code {
	args := []jen.Code{}

	roleSub := []jen.Code{}
	for _, to := range e.RoleSubs.Roles {
		from := e.RoleSubs.Subst(to)
		roleSub = append(roleSub, jen.Lit(from), jen.Lit(to))
	}

	args = append(args, jen.Id("env").Dot("Subst").Call(roleSub...))

	for _, arg := range e.Args {
		args = append(args, GenExpression(arg))
	}

	return jen.Add(GenExpression(e.FuncExpr)).Call(args...)
}

func GenExprCallClosure(e *projection.ExprCallClosure) jen.Code {
	args := []jen.Code{}

	for _, arg := range e.Args {
		args = append(args, GenExpression(arg))
	}

	return jen.Add(GenExpression(e.ClosureExpr)).Call(args...)
}

func GenExprStruct(e *projection.ExprStruct) jen.Code {
	fields := jen.Dict{}

	for _, fieldName := range e.FieldNames {
		field := e.Fields[fieldName]
		expr := GenExpression(field)
		fields[jen.Id(fieldName)] = expr
	}

	name := fmt.Sprintf("%s_%s", e.StructName, e.StructRole)

	return jen.Id(name).Values(fields)
}

func GenExprFieldAccess(e *projection.ExprFieldAccess) jen.Code {
	return jen.Add(GenExpression(e.BaseExpr)).Dot(e.FieldName)
}

func GenExprCopy(e *projection.ExprPassValue) jen.Code {
	return RuntimeFunc("Copy").Call(GenExpression(e.Inner))
}
