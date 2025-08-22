package codegen_java

import (
	"fmt"
	"slices"

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
	case *projection.ExprPassValue:
		return gen.GenExprCopy(e)
	case *projection.ExprSelf:
		return gen.GenExprSelf(e)
	}

	panic(fmt.Sprintf("unexpected projection.Expression: %#v", expr))
}

func (gen *codegen) GenExprAsync(e *projection.ExprAsync) string {
	gen.AddImport(javaPkgComplatableFuture)
	return fmt.Sprintf("CompletableFuture.completedFuture(%s)", gen.GenExpr(e.Inner()))
}

func (gen *codegen) GenExprAwait(e *projection.ExprAwait) string {
	// Future::get()
	return fmt.Sprintf("%s.get()", gen.GenExpr(e.Inner()))
}

func (gen *codegen) GenExprBinaryOp(e *projection.ExprBinaryOp) string {
	op := e.Operator

	switch op {
	case projection.OpEq:
		op = "==="
	case projection.OpNotEq:
		op = "!=="
	}

	if _, isList := e.Type().(*projection.ListType); isList && op == projection.OpAdd {
		gen.AddImport(javaPkgStream, javaPkgCollectors, javaPkgArrayList)
		return fmt.Sprintf("Stream.concat(%s.stream(), %s.stream()).collect(Collectors.toCollection(ArrayList::new))", gen.GenExpr(e.Lhs), gen.GenExpr(e.Rhs))
	}

	result := fmt.Sprintf("%s %s %s", gen.GenExpr(e.Lhs), op, gen.GenExpr(e.Rhs))

	return result
}

func (gen *codegen) GenExprBool(e *projection.ExprBool) string {
	if e.Value {
		return "true"
	} else {
		return "false"
	}
}

func (gen *codegen) GenExprCallClosure(e *projection.ExprCallClosure) string {
	return "CALL_CLOSURE"
}

func (gen *codegen) GenExprCallFunc(e *projection.ExprCallFunc) string {
	roleSub := []string{}
	for _, to := range e.RoleSubst.Roles {
		from := e.RoleSubst.Subst(to)[0]
		if from != to {
			roleSub = append(roleSub, fmt.Sprintf("\"%s\"", from), fmt.Sprintf("\"%s\"", to))
		}
	}

	env := "env"
	if len(roleSub) > 0 {
		env = fmt.Sprintf("env.subst(%s)", misc.JoinStrings(roleSub, ", "))
	}

	args := []string{env}
	for _, arg := range e.Args {
		args = append(args, gen.GenExpr(arg))
	}

	out := fmt.Sprintf("%s(%s)", gen.GenExpr(e.FuncExpr), misc.JoinStrings(args, ", "))

	return out
}

func (gen *codegen) GenExprClosure(e *projection.ExprClosure) string {
	return "MAKE_CLOSURE"
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

	gen.AddImport(javaPkgArrayList, javaPkgList)
	return fmt.Sprintf("new ArrayList<>(List.of(%s))", misc.JoinStrings(items, ", "))
}

func (gen *codegen) GenExprListLength(e *projection.ExprListLength) string {
	return fmt.Sprintf("%s.size()", gen.GenExpr(e.List))
}

func (gen *codegen) GenExprRecv(e *projection.ExprRecv) string {
	return fmt.Sprintf("env.<%s>recv(\"%s\")", gen.GenType(e.RecvType), e.Sender)
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
	return fmt.Sprintf("new %s_%s(%s)", e.StructName, e.StructRole, misc.JoinStrings(e.FieldNames, ", "))
}

func (gen *codegen) GenExprSelf(e *projection.ExprSelf) string {
	return "this"
}

func (gen *codegen) GenExprCopy(e *projection.ExprPassValue) string {
	if CopyNecessary(e.Inner) && !ValueBasedType(e.Inner.Type()) {
		return fmt.Sprintf("%s.clone()", gen.GenExpr(e.Inner))
	} else {
		return gen.GenExpr(e.Inner)
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
		return true
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
