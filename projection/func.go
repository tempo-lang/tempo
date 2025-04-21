package projection

import (
	"chorego/parser"
	"fmt"

	"github.com/dave/jennifer/jen"
)

type Func struct {
	Choreography *Choreography
	FuncCtx      parser.IFuncContext
	Name         string
	Role         string
	Params       []FuncParam
}

type FuncParam struct {
	Func     *Func
	ParamCtx parser.IFuncParamContext
	Name     string
	Type     string
}

func (f *Func) AddParam(param parser.IFuncParamContext, paramType string) *Func {
	f.Params = append(f.Params, FuncParam{
		Func:     f,
		ParamCtx: param,
		Name:     param.Ident().GetText(),
		Type:     paramType,
	})
	return f
}

func (f *Func) Codegen(file *jen.File) {
	file.Func().
		Id(fmt.Sprintf("%s_%s", f.Name, f.Role)).
		ParamsFunc(func(params *jen.Group) {
			for _, param := range f.Params {
				params.Id(param.Name).Id(param.Type)
			}
		}).
		Block()
}
