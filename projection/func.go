package projection

import (
	"fmt"
	"tempo/parser"
	"tempo/types"

	"github.com/dave/jennifer/jen"
)

type FuncSig struct {
	FuncSigCtx  parser.IFuncSigContext
	Name        string
	Role        string
	Params      []FuncParam
	ReturnValue types.Value
}

type Func struct {
	*FuncSig
	Choreography *Choreography
	FuncCtx      parser.IFuncContext
	Body         []Statement
}

type FuncParam struct {
	FuncSig  *FuncSig
	ParamCtx parser.IFuncParamContext
	Name     string
	Type     *types.Type
}

func NewFuncSig(role string, funcSigCtx parser.IFuncSigContext, returnValue types.Value) *FuncSig {
	return &FuncSig{
		FuncSigCtx:  funcSigCtx,
		Name:        funcSigCtx.Ident().GetText(),
		Role:        role,
		Params:      []FuncParam{},
		ReturnValue: returnValue,
	}
}

func (f *FuncSig) AddParam(param parser.IFuncParamContext, paramType *types.Type) *FuncSig {
	f.Params = append(f.Params, FuncParam{
		FuncSig:  f,
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

func (f *Func) Codegen() *jen.Statement {
	result := f.FuncSig.Codegen(true)

	result = result.BlockFunc(func(block *jen.Group) {
		for _, bodyStmt := range f.Body {
			for _, stmt := range bodyStmt.Codegen() {
				block.Add(stmt)
			}
		}
	})

	return result
}

func (f *FuncSig) Codegen(printFuncKeyword bool) *jen.Statement {
	result := jen.Id(fmt.Sprintf("%s_%s", f.Name, f.Role)).
		ParamsFunc(func(params *jen.Group) {
			params.Id("env").Add(jen.Op("*").Qual("tempo/runtime", "Env"))
			for _, param := range f.Params {
				params.Id(param.Name).Add(CodegenType(param.Type.Value()))
			}
		})

	if printFuncKeyword {
		result = jen.Func().Add(result)
	}

	if f.ReturnValue != types.Unit() {
		result = result.Add(CodegenType(f.ReturnValue))
	}

	return result
}
