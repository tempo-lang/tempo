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
