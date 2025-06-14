package projection

type ExprList struct {
	Items    []Expression
	listType Type
}

func (e *ExprList) Type() Type {
	return e.listType
}

func (e *ExprList) ReturnsValue() bool {
	return true
}

func (e *ExprList) HasSideEffects() bool {
	for _, item := range e.Items {
		if item.HasSideEffects() {
			return true
		}
	}
	return false
}

func (e *ExprList) IsExpression() {}

func NewExprList(items []Expression, listType Type) Expression {
	return &ExprList{
		Items:    items,
		listType: listType,
	}
}

type ExprIndex struct {
	Base  Expression
	Index Expression
}

func (e *ExprIndex) Type() Type {
	listType := e.Base.Type().(*ListType)
	return listType.Inner
}

func (e *ExprIndex) ReturnsValue() bool {
	return true
}

func (e *ExprIndex) HasSideEffects() bool {
	return e.Base.HasSideEffects() || e.Index.HasSideEffects()
}

func (e *ExprIndex) IsExpression() {}

func NewExprIndex(base Expression, index Expression) Expression {
	return &ExprIndex{
		Base:  base,
		Index: index,
	}
}

type ExprListLength struct {
	List Expression
}

func (e *ExprListLength) HasSideEffects() bool {
	return e.List.HasSideEffects()
}

func (e *ExprListLength) IsExpression() {}

func (e *ExprListLength) ReturnsValue() bool {
	return true
}

func (e *ExprListLength) Type() Type {
	return IntType
}

func NewExprListLength(list Expression) Expression {
	return &ExprListLength{List: list}
}
