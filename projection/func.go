package projection

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"
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
