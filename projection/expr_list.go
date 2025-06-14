package projection

import (
	"github.com/dave/jennifer/jen"
)

type ExprList struct {
	items    []Expression
	listType Type
}

func (e *ExprList) Codegen() jen.Code {
	items := make([]jen.Code, len(e.items))
	for i, item := range e.items {
		items[i] = item.Codegen()
	}
	return jen.Add(e.listType.Codegen()).Values(items...)
}

func (e *ExprList) Type() Type {
	return e.listType
}

func (e *ExprList) ReturnsValue() bool {
	return true
}

func (e *ExprList) HasSideEffects() bool {
	for _, item := range e.items {
		if item.HasSideEffects() {
			return true
		}
	}
	return false
}

func (e *ExprList) IsExpression() {}

func NewExprList(items []Expression, listType Type) Expression {
	return &ExprList{
		items:    items,
		listType: listType,
	}
}

type ExprIndex struct {
	base  Expression
	index Expression
}

func (e *ExprIndex) Codegen() jen.Code {
	return jen.Add(e.base.Codegen()).Index(e.index.Codegen())
}

func (e *ExprIndex) Type() Type {
	listType := e.base.Type().(*ListType)
	return listType.Inner
}

func (e *ExprIndex) ReturnsValue() bool {
	return true
}

func (e *ExprIndex) HasSideEffects() bool {
	return e.base.HasSideEffects() || e.index.HasSideEffects()
}

func (e *ExprIndex) IsExpression() {}

func NewExprIndex(base Expression, index Expression) Expression {
	return &ExprIndex{
		base:  base,
		index: index,
	}
}

type ExprListLength struct {
	list Expression
}

func (e *ExprListLength) Codegen() jen.Code {
	return jen.Len(e.list.Codegen())
}

func (e *ExprListLength) HasSideEffects() bool {
	return e.list.HasSideEffects()
}

func (e *ExprListLength) IsExpression() {}

func (e *ExprListLength) ReturnsValue() bool {
	return true
}

func (e *ExprListLength) Type() Type {
	return IntType
}

func NewExprListLength(list Expression) Expression {
	return &ExprListLength{list: list}
}
