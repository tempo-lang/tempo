package projection

// Construct a new closure and returns it.
// The closure can capture symbols from parent scopes by reference.
type ExprClosure struct {
	Params      []ClosureParam
	ReturnType  Type
	Body        []Statement
	ClosureType Type
}

type ClosureParam struct {
	Name string
	Type Type
}

func NewClosureParam(name string, typ Type) ClosureParam {
	return ClosureParam{
		Name: name,
		Type: typ,
	}
}

func NewExprClosure(params []ClosureParam, returnType Type, body []Statement) Expression {
	return &ExprClosure{
		Params:     params,
		ReturnType: returnType,
		Body:       body,
	}
}

func (e *ExprClosure) HasSideEffects() bool {
	return true
}

func (e *ExprClosure) IsExpression() {}

func (e *ExprClosure) ReturnsValue() bool {
	return e.ReturnType != UnitType()
}

func (e *ExprClosure) Type() Type {
	params := []Type{}
	for _, param := range e.Params {
		params = append(params, param.Type)
	}
	return NewClosureType(params, e.ReturnType)
}

type ClosureType struct {
	Params     []Type
	ReturnType Type
}

func (c *ClosureType) IsType() {}

func NewClosureType(params []Type, returnType Type) *ClosureType {
	return &ClosureType{
		Params:     params,
		ReturnType: returnType,
	}
}
