package codegen_go

import (
	"fmt"
	"slices"

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
	case *projection.ExprSelf:
		return GenExprSelf(e)
	case *projection.ExprTypeCast:
		return GenExprTypeCast(e)
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
	for _, to := range e.RoleSubst.Roles {
		fromSubst := e.RoleSubst.Subst(to)
		if len(fromSubst) != 1 {
			panic(fmt.Sprintf("expected role substitution to be unique: %s -> %v", to, fromSubst))
		}

		from := e.RoleSubst.Subst(to)[0]
		if from != to {
			roleSub = append(roleSub, jen.Lit(from), jen.Lit(to))
		}
	}

	if len(roleSub) == 0 {
		args = append(args, jen.Id("env"))
	} else {
		args = append(args, jen.Id("env").Dot("Subst").Call(roleSub...))
	}

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

	return jen.Id(e.StructName()).Values(fields)
}

func GenExprFieldAccess(e *projection.ExprFieldAccess) jen.Code {
	return jen.Add(GenExpression(e.BaseExpr)).Dot(e.FieldName)
}

func GenExprSelf(e *projection.ExprSelf) jen.Code {
	return jen.Id("self")
}

func GenExprCopy(e *projection.ExprPassValue) jen.Code {
	if CopyNecessary(e.Inner) && !ValueBasedType(e.Inner.Type()) {
		return RuntimeFunc("Copy").Call(GenExpression(e.Inner))
	} else {
		return GenExpression(e.Inner)
	}
}

func GenExprTypeCast(e *projection.ExprTypeCast) jen.Code {
	switch e.NewType {
	case projection.StringType:
		switch e.Inner.Type() {
		case projection.IntType:
			return RuntimeFunc("IntToString").Call(GenExpression(e.Inner))
		case projection.FloatType:
			return RuntimeFunc("FloatToString").Call(GenExpression(e.Inner))
		case projection.BoolType:
			return RuntimeFunc("BoolToString").Call(GenExpression(e.Inner))
		default:
			panic(fmt.Sprintf("unsupported conversion: %v to %v", e.Inner.Type(), e.NewType))
		}
	default:
		return jen.Add(GenType(e.NewType)).Call(GenExpression(e.Inner))
	}
}

func CopyNecessary(e projection.Expression) bool {
	switch e := e.(type) {
	case *projection.ExprAsync:
		return false
	case *projection.ExprAwait:
		return false
	case *projection.ExprBinaryOp:
		return CopyNecessary(e.Lhs) || CopyNecessary(e.Rhs)
	case *projection.ExprBool:
		return false
	case *projection.ExprCallClosure:
		return false
	case *projection.ExprCallFunc:
		return false
	case *projection.ExprClosure:
		return false
	case *projection.ExprFieldAccess:
		return true
	case *projection.ExprFloat:
		return false
	case *projection.ExprIdent:
		return true
	case *projection.ExprIndex:
		return true
	case *projection.ExprInt:
		return false
	case *projection.ExprList:
		return slices.ContainsFunc(e.Items, CopyNecessary)
	case *projection.ExprListLength:
		return false
	case *projection.ExprPassValue:
		return CopyNecessary(e.Inner)
	case *projection.ExprRecv:
		return false
	case *projection.ExprSend:
		return false
	case *projection.ExprString:
		return false
	case *projection.ExprStruct:
		for _, expr := range e.Fields {
			if CopyNecessary(expr) {
				return true
			}
		}
		return false
	case *projection.ExprTypeCast:
		return CopyNecessary(e.Inner)
	default:
		panic(fmt.Sprintf("unexpected projection.Expression: %#v", e))
	}
}

func ValueBasedType(t projection.Type) bool {
	switch t.(type) {
	case *projection.AsyncType:
		return true
	case projection.BuiltinType:
		return true
	case *projection.ClosureType:
		return false
	case *projection.FunctionType:
		return true
	case *projection.InterfaceType:
		return false
	case *projection.ListType:
		return false
	case *projection.StructType:
		return false
	}
	panic(fmt.Sprintf("unexpected projection.Type: %#v", t))
}
