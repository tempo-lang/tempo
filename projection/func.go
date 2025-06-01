package projection

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"

	"github.com/dave/jennifer/jen"
)

type FuncSig struct {
	FuncSigCtx  parser.IFuncSigContext
	Name        string
	Role        string
	Params      []FuncParam
	ReturnValue Type
}

type Func struct {
	*FuncSig
	Choreography *Choreography
	FuncCtx      parser.IFuncContext
	Body         []Statement
}

type FuncParam struct {
	FuncSig   *FuncSig
	ParamCtx  parser.IFuncParamContext
	Name      string
	TypeValue Type
}

func NewFuncSig(role string, funcSigCtx parser.IFuncSigContext, returnValue Type) *FuncSig {
	return &FuncSig{
		FuncSigCtx:  funcSigCtx,
		Name:        funcSigCtx.Ident().GetText(),
		Role:        role,
		Params:      []FuncParam{},
		ReturnValue: returnValue,
	}
}

func (f *FuncSig) AddParam(param parser.IFuncParamContext, paramType Type) *FuncSig {
	f.Params = append(f.Params, FuncParam{
		FuncSig:   f,
		ParamCtx:  param,
		Name:      param.Ident().GetText(),
		TypeValue: paramType,
	})
	return f
}

func (f *Func) AddStmt(stmt ...Statement) *Func {
	f.Body = append(f.Body, stmt...)
	return f
}

func (f *Func) Codegen() *jen.Statement {
	result := f.FuncSig.Codegen(false)

	result = result.BlockFunc(func(block *jen.Group) {
		for _, bodyStmt := range f.Body {
			for _, stmt := range bodyStmt.Codegen() {
				block.Add(stmt)
			}
		}
	})

	return result
}

func (f *FuncSig) Codegen(isInterfaceMethod bool) *jen.Statement {
	params := []jen.Code{
		jen.Id("env").Add(jen.Op("*").Qual("github.com/tempo-lang/tempo/runtime", "Env")),
	}
	for _, param := range f.Params {
		params = append(params, jen.Id(param.Name).Add(param.TypeValue.Codegen()))
	}

	var result *jen.Statement
	if isInterfaceMethod {
		result = jen.Id(f.Name).Params(params...)
	} else {
		result = jen.Func().Id(fmt.Sprintf("%s_%s", f.Name, f.Role)).Params(params...)
	}

	if f.ReturnValue != UnitType() {
		result = result.Add(f.ReturnValue.Codegen())
	}

	return result
}

type FunctionType struct {
	types.FunctionType
	Params     []Type
	ReturnType Type
}

func NewFunctionType(funcType *types.FunctionType, params []Type, returnType Type) *FunctionType {
	return &FunctionType{
		FunctionType: *funcType,
		Params:       params,
		ReturnType:   returnType,
	}
}

func (c *FunctionType) IsType() {}

func (f *FunctionType) Codegen() jen.Code {
	params := []jen.Code{
		jen.Op("*").Qual("github.com/tempo-lang/tempo/runtime", "Env"),
	}

	for _, param := range f.Params {
		params = append(params, param.Codegen())
	}

	result := jen.Func().Params(params...)

	if f.ReturnType != UnitType() {
		result = result.Add(f.ReturnType.Codegen())
	}

	return result
}
