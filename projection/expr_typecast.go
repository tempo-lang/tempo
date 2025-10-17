package projection

type ExprTypeCast struct {
	Inner   Expression
	NewType Type
}

func (e *ExprTypeCast) Type() Type {
	return e.NewType
}

func (e *ExprTypeCast) ReturnsValue() bool {
	return true
}

func (e *ExprTypeCast) HasSideEffects() bool {
	return false
}

func (e *ExprTypeCast) IsExpression() {}

func NewExprTypeCast(inner Expression, newType Type) Expression {
	return &ExprTypeCast{
		Inner:   inner,
		NewType: newType,
	}
}
