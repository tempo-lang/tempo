package projection

import (
	"chorego/types"

	"github.com/dave/jennifer/jen"
)

type Expression interface {
	Codegen() jen.Code
	Type() types.Value
	IsExpression()
}

type ExprInt struct {
	Value int
}

func (e *ExprInt) Codegen() jen.Code {
	return jen.Lit(e.Value)
}

func (e *ExprInt) Type() types.Value {
	return types.Int()
}

func (e *ExprInt) IsExpression() {}

func NewExprInt(value int) Expression {
	return &ExprInt{
		Value: value,
	}
}

type ExprIdent struct {
	Name    string
	typeVal types.Value
}

func (e *ExprIdent) Codegen() jen.Code {
	return jen.Id(e.Name)
}

func (e *ExprIdent) Type() types.Value {
	return e.typeVal
}

func (e *ExprIdent) IsExpression() {}

func NewExprIdent(name string, typeVal types.Value) Expression {
	return &ExprIdent{Name: name, typeVal: typeVal}
}

type ExprBool struct {
	Value bool
}

func (e *ExprBool) Codegen() jen.Code {
	if e.Value {
		return jen.True()
	} else {
		return jen.False()
	}
}

func (e *ExprBool) Type() types.Value {
	return types.Bool()
}

func (e *ExprBool) IsExpression() {}

func NewExprBool(value bool) Expression {
	return &ExprBool{Value: value}
}

type Operator string

const (
	OperatorAdd = "+"
)

type ExprBinaryOp struct {
	operator Operator
	lhs      Expression
	rhs      Expression
	typeVal  types.Value
}

func (e *ExprBinaryOp) Codegen() jen.Code {
	return jen.Add(e.lhs.Codegen()).Op(string(e.operator)).Add(e.rhs.Codegen())
}

func (e *ExprBinaryOp) Type() types.Value {
	return e.typeVal
}

func (e *ExprBinaryOp) IsExpression() {}

func NewExprBinaryOp(operator Operator, lhs Expression, rhs Expression, typeVal types.Value) Expression {
	return &ExprBinaryOp{
		operator: OperatorAdd,
		lhs:      lhs,
		rhs:      rhs,
		typeVal:  typeVal,
	}
}

type ExprAsync struct {
	inner Expression
}

func (e *ExprAsync) Codegen() jen.Code {
	return jen.Qual("chorego/runtime", "FixedAsync").Call(e.inner.Codegen())
}

func (e *ExprAsync) Type() types.Value {
	return types.NewAsync(e.Type())
}

func (e *ExprAsync) IsExpression() {}

func (e *ExprAsync) Inner() Expression {
	return e.inner
}

func NewExprAsync(inner Expression) Expression {
	return &ExprAsync{inner: inner}
}

type ExprAwait struct {
	expr    Expression
	typeVal types.Value
}

func (e *ExprAwait) Inner() Expression {
	return e.expr
}

func (e *ExprAwait) Codegen() jen.Code {

	return jen.Add(e.Inner().Codegen()).Dot("Get").Params().Assert(CodegenType(e.typeVal))
}

func (e *ExprAwait) Type() types.Value {
	return e.typeVal
}

func (e *ExprAwait) IsExpression() {}

func NewExprAwait(inner Expression, typeVal types.Value) Expression {
	return &ExprAwait{expr: inner, typeVal: typeVal}
}
