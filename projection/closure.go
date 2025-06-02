package projection

import (
	"github.com/dave/jennifer/jen"
)

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

func (e *ExprClosure) Codegen() jen.Code {
	params := []jen.Code{}
	for _, param := range e.Params {
		params = append(params, jen.Id(param.Name).Add(param.Type.Codegen()))
	}

	result := jen.Func().Params(params...)

	if e.ReturnType != UnitType() {
		result = result.Add(e.ReturnType.Codegen())
	}

	result = result.BlockFunc(func(block *jen.Group) {
		for _, bodyStmt := range e.Body {
			for _, stmt := range bodyStmt.Codegen() {
				block.Add(stmt)
			}
		}
	})

	return result
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

func (c *ClosureType) Codegen() jen.Code {
	params := []jen.Code{}

	for _, param := range c.Params {
		params = append(params, param.Codegen())
	}

	result := jen.Func().Params(params...)

	if c.ReturnType != UnitType() {
		result = result.Add(c.ReturnType.Codegen())
	}

	return result
}

func NewClosureType(params []Type, returnType Type) *ClosureType {
	return &ClosureType{
		Params:     params,
		ReturnType: returnType,
	}
}

// func (c *ClosureType) CoerceTo(other types.Value) (types.Value, bool) {
// 	panic("unimplemented")
// }

// func (c *ClosureType) IsEquatable() bool {
// 	return false
// }

// func (c *ClosureType) IsSendable() bool {
// 	return false
// }

// func (c *ClosureType) IsValue() {}

// func (c *ClosureType) SubstituteRoles(substMap *types.RoleSubst) types.Value {
// 	panic("unimplemented")
// }

// func (c *ClosureType) ToString() string {
// 	panic("unimplemented")
// }
