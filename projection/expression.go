package projection

import "github.com/dave/jennifer/jen"

type Expression interface {
	Codegen() jen.Code
	IsExpression()
}

type ExprInt struct {
	Value int
}

func (e *ExprInt) Codegen() jen.Code {
	return jen.Lit(e.Value)
}

func (e *ExprInt) IsExpression() {}

func NewExprInt(value int) Expression {
	return &ExprInt{
		Value: value,
	}
}

type ExprIdent struct {
	Name string
}

func (e *ExprIdent) Codegen() jen.Code {
	return jen.Id(e.Name)
}

func (e *ExprIdent) IsExpression() {}

func NewExprIdent(name string) Expression {
	return &ExprIdent{Name: name}
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
}

func (e *ExprBinaryOp) Codegen() jen.Code {
	return jen.Add(e.lhs.Codegen()).Op(string(e.operator)).Add(e.rhs.Codegen())
}

func (e *ExprBinaryOp) IsExpression() {}

func NewExprBinaryOp(operator Operator, lhs Expression, rhs Expression) Expression {
	return &ExprBinaryOp{
		operator: OperatorAdd,
		lhs:      lhs,
		rhs:      rhs,
	}
}

type ExprAsync struct {
	inner Expression
}

func (e *ExprAsync) Codegen() jen.Code {
	return jen.Qual("chorego/runtime", "FixedAsync").Call(e.inner.Codegen())
}

func (e *ExprAsync) IsExpression() {}

func (e *ExprAsync) Inner() Expression {
	return e.inner
}

func NewExprAsync(inner Expression) Expression {
	return &ExprAsync{inner: inner}
}
