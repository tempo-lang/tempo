package codegen_go

import (
	"github.com/dave/jennifer/jen"
	"github.com/tempo-lang/tempo/projection"
)

func GenExprClosure(e *projection.ExprClosure) jen.Code {
	params := []jen.Code{}
	for _, param := range e.Params {
		params = append(params, jen.Id(param.Name).Add(GenType(param.Type)))
	}

	result := jen.Func().Params(params...)

	if e.ReturnType != projection.UnitType() {
		result = result.Add(GenType(e.ReturnType))
	}

	result = result.BlockFunc(func(block *jen.Group) {
		for _, bodyStmt := range e.Body {
			for _, stmt := range GenStatement(bodyStmt) {
				block.Add(stmt)
			}
		}
	})

	return result
}

func GenClosureType(c *projection.ClosureType) jen.Code {
	params := []jen.Code{}

	for _, param := range c.Params {
		params = append(params, GenType(param))
	}

	result := jen.Func().Params(params...)

	if c.ReturnType != projection.UnitType() {
		result = result.Add(GenType(c.ReturnType))
	}

	return result
}
