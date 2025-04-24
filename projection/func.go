package projection

import (
	"chorego/parser"
	"chorego/type_check"
	"fmt"
	"slices"

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

func (f *Func) AddStmt(stmt Statement) *Func {
	f.Body = append(f.Body, stmt)
	return f
}

func (f *Func) Codegen(file *jen.File) {
	file.Func().
		Id(fmt.Sprintf("%s_%s", f.Name, f.Role)).
		ParamsFunc(func(params *jen.Group) {
			for _, param := range f.Params {

				var typeName string
				if slices.Contains(type_check.BuiltinTypes(), param.Type) {
					typeName = BuiltinTypeGo(param.Type)
				}

				params.Id(param.Name).Id(typeName)
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

func BuiltinTypeGo(t string) string {
	switch t {
	case "Int":
		return "int"
	case "Float":
		return "float64"
	case "String":
		return "string"
	case "Bool":
		return "bool"
	default:
		panic(fmt.Sprintf("unknown builtin type: %s", t))
	}
}
