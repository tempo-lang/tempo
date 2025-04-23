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

func NewExprInt(value int) *ExprInt {
	return &ExprInt{
		Value: value,
	}
}
