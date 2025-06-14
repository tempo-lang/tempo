package codegen_ts

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenExpr(expr projection.Expression) string {
	switch e := expr.(type) {
	case *projection.ExprAsync:
		return gen.GenExprAsync(e)
	case *projection.ExprAwait:
		return gen.GenExprAwait(e)
	case *projection.ExprBinaryOp:
		return gen.GenExprBinaryOp(e)
	case *projection.ExprBool:
		return gen.GenExprBool(e)
	case *projection.ExprCallClosure:
		return gen.GenExprCallClosure(e)
	case *projection.ExprCallFunc:
		return gen.GenExprCallFunc(e)
	case *projection.ExprClosure:
		return gen.GenExprClosure(e)
	case *projection.ExprFieldAccess:
		return gen.GenExprFieldAccess(e)
	case *projection.ExprFloat:
		return gen.GenExprFloat(e)
	case *projection.ExprIdent:
		return gen.GenExprIdent(e)
	case *projection.ExprIndex:
		return gen.GenExprIndex(e)
	case *projection.ExprInt:
		return gen.GenExprInt(e)
	case *projection.ExprList:
		return gen.GenExprList(e)
	case *projection.ExprListLength:
		return gen.GenExprListLength(e)
	case *projection.ExprRecv:
		return gen.GenExprRecv(e)
	case *projection.ExprSend:
		return gen.GenExprSend(e)
	case *projection.ExprString:
		return gen.GenExprString(e)
	case *projection.ExprStruct:
		return gen.GenExprStruct(e)
	}

	panic(fmt.Sprintf("unexpected projection.Expression: %#v", expr))
}

func (gen *codegen) GenExprAsync(e *projection.ExprAsync) string {
	return fmt.Sprintf("Promise.resolve(%s)", gen.GenExpr(e.Inner()))
}

func (gen *codegen) GenExprAwait(e *projection.ExprAwait) string {
	return fmt.Sprintf("await %s", gen.GenExpr(e.Inner()))
}

func (gen *codegen) GenExprBinaryOp(e *projection.ExprBinaryOp) string {
	return fmt.Sprintf("%s %s %s", gen.GenExpr(e.Lhs), e.Operator, gen.GenExpr(e.Rhs))
}

func (gen *codegen) GenExprBool(e *projection.ExprBool) string {
	if e.Value {
		return "true"
	} else {
		return "false"
	}
}

func (gen *codegen) GenExprCallClosure(e *projection.ExprCallClosure) string {
	args := []string{}
	for _, arg := range e.Args {
		args = append(args, gen.GenExpr(arg))
	}

	return fmt.Sprintf("%s(%s)", gen.GenExpr(e.ClosureExpr), misc.JoinStrings(args, ", "))
}

func (gen *codegen) GenExprCallFunc(e *projection.ExprCallFunc) string {
	args := []string{}
	for _, arg := range e.Args {
		args = append(args, gen.GenExpr(arg))
	}

	return fmt.Sprintf("%s(%s)", gen.GenExpr(e.FuncExpr), misc.JoinStrings(args, ", "))
}

func (gen *codegen) GenExprClosure(e *projection.ExprClosure) string {
	params := []string{}
	for _, param := range e.Params {
		if gen.opts.DisableTypes {
			params = append(params, param.Name)
		} else {
			params = append(params, fmt.Sprintf("%s: %s", param.Name, gen.GenType(param.Type)))
		}
	}

	result := fmt.Sprintf("(%s)", misc.JoinStrings(params, ", "))

	if e.ReturnType != projection.UnitType() && !gen.opts.DisableTypes {
		result += fmt.Sprintf(": %s", gen.GenType(e.ReturnType))
	}

	result += " => {"
	gen.IncIndent()

	// builder := strings.Builder{}
	// for _, stmt := range e.Body {
	// 	gen.GenStmt(stmt)
	// }
	// result += builder.String()

	gen.DecIndent()

	result += "}"

	return result
}

func (gen *codegen) GenExprFieldAccess(e *projection.ExprFieldAccess) string {
	return fmt.Sprintf("%s.%s", gen.GenExpr(e.BaseExpr), e.FieldName)
}

func (gen *codegen) GenExprFloat(e *projection.ExprFloat) string {
	return fmt.Sprintf("%f", e.Value)
}

func (gen *codegen) GenExprIdent(e *projection.ExprIdent) string {
	return e.Name
}

func (gen *codegen) GenExprIndex(e *projection.ExprIndex) string {
	return fmt.Sprintf("%s[%s]", gen.GenExpr(e.Base), gen.GenExpr(e.Index))
}

func (gen *codegen) GenExprInt(e *projection.ExprInt) string {
	return fmt.Sprintf("%d", e.Value)
}

func (gen *codegen) GenExprList(e *projection.ExprList) string {
	items := []string{}
	for _, item := range e.Items {
		items = append(items, gen.GenExpr(item))
	}

	return fmt.Sprintf("[%s]", misc.JoinStrings(items, ", "))
}

func (gen *codegen) GenExprListLength(e *projection.ExprListLength) string {
	return fmt.Sprintf("%s.length", gen.GenExpr(e.List))
}

func (gen *codegen) GenExprRecv(e *projection.ExprRecv) string {
	return fmt.Sprintf("env.recv(\"%s\")", e.Sender)
}

func (gen *codegen) GenExprSend(e *projection.ExprSend) string {
	receivers := misc.JoinStringsFunc(e.Receivers, ", ", func(role string) string {
		return fmt.Sprintf("\"%s\"", role)
	})

	return fmt.Sprintf("env.send(%s, %s)", gen.GenExpr(e.Expr), receivers)
}

func (gen *codegen) GenExprString(e *projection.ExprString) string {
	return fmt.Sprintf("\"%s\"", e.Value)
}

func (gen *codegen) GenExprStruct(e *projection.ExprStruct) string {
	fields := []string{}

	for _, field := range e.FieldNames {
		fields = append(fields, fmt.Sprintf("%s: %s", field, gen.GenExpr(e.Fields[field])))
	}

	return fmt.Sprintf("{ %s }", misc.JoinStrings(fields, ", "))
}
