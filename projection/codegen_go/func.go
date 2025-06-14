package codegen_go

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/tempo-lang/tempo/projection"
)

func GenFunc(f *projection.Func) *jen.Statement {
	result := GenFuncSig(f.FuncSig, false)

	result = result.BlockFunc(func(block *jen.Group) {
		for _, bodyStmt := range f.Body {
			for _, stmt := range GenStatement(bodyStmt) {
				block.Add(stmt)
			}
		}
	})

	return result
}

func GenFuncSig(f *projection.FuncSig, isInterfaceMethod bool) *jen.Statement {
	params := []jen.Code{
		jen.Id("env").Add(jen.Op("*").Qual("github.com/tempo-lang/tempo/runtime", "Env")),
	}
	for _, param := range f.Params {
		params = append(params, jen.Id(param.Name).Add(GenType(param.TypeValue)))
	}

	var result *jen.Statement
	if isInterfaceMethod {
		result = jen.Id(f.Name).Params(params...)
	} else {
		result = jen.Func().Id(fmt.Sprintf("%s_%s", f.Name, f.Role)).Params(params...)
	}

	if f.ReturnValue != projection.UnitType() {
		result = result.Add(GenType(f.ReturnValue))
	}

	return result
}

func GenFunctionType(f *projection.FunctionType) jen.Code {
	params := []jen.Code{
		jen.Op("*").Qual("github.com/tempo-lang/tempo/runtime", "Env"),
	}

	for _, param := range f.Params {
		params = append(params, GenType(param))
	}

	result := jen.Func().Params(params...)

	if f.ReturnType != projection.UnitType() {
		result = result.Add(GenType(f.ReturnType))
	}

	return result
}
