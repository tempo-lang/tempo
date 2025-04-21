package process

import (
	"chorego/parser"
	"fmt"

	"github.com/dave/jennifer/jen"
)

type Func struct {
	Process *Process
	FuncCtx parser.IFuncContext
	Name    string
	Params  []FuncParam
}

type FuncParam struct {
	Func     *Func
	ParamCtx parser.IFunc_paramContext
	Name     string
	Type     string
}

func (f *Func) AddParam(param parser.IFunc_paramContext, paramType string) *Func {
	f.Params = append(f.Params, FuncParam{
		Func:     f,
		ParamCtx: param,
		Name:     param.Ident().GetText(),
		Type:     paramType,
	})
	return f
}

func (f *Func) Codegen(file *jen.File) {
	// file.Commentf("Projection of function %s at role %s", f.Name, f.Process.Name)

	file.Func().
		Id(fmt.Sprintf("%s_%s", f.Name, f.Process.Name)).
		ParamsFunc(func(params *jen.Group) {
			for _, param := range f.Params {
				params.Id(param.Name).Id(param.Type)
			}
		}).
		Block()
}
