package projection

import (
	"chorego/parser"
	"chorego/types"
	"fmt"

	"github.com/dave/jennifer/jen"
)

type Func struct {
	Choreography *Choreography
	FuncCtx      parser.IFuncContext
	Name         string
	Role         string
	Params       []FuncParam
	Body         []Statement
}

type FuncParam struct {
	Func     *Func
	ParamCtx parser.IFuncParamContext
	Name     string
	Type     *types.Type
}

func (f *Func) AddParam(param parser.IFuncParamContext, paramType *types.Type) *Func {
	f.Params = append(f.Params, FuncParam{
		Func:     f,
		ParamCtx: param,
		Name:     param.Ident().GetText(),
		Type:     paramType,
	})
	return f
}

func (f *Func) AddStmt(stmt ...Statement) *Func {
	f.Body = append(f.Body, stmt...)
	return f
}

func (f *Func) Codegen(file *jen.File) {
	file.Func().
		Id(fmt.Sprintf("%s_%s", f.Name, f.Role)).
		ParamsFunc(func(params *jen.Group) {
			params.Id("env").Add(jen.Op("*").Qual("chorego/runtime", "Env"))
			for _, param := range f.Params {
				params.Id(param.Name).Add(CodegenType(param.Type.Value()))
			}
		}).
		BlockFunc(func(block *jen.Group) {

			for _, bodyStmt := range f.Body {
				for _, stmt := range bodyStmt.Codegen() {
					block.Add(stmt)
				}
			}
		})
}
