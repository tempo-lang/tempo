package sym_table

import (
	"chorego/parser"
	"chorego/types"
)

type Symbol interface {
	SymbolName() string
	Ident() parser.IIdentContext
	Type() *types.Type
}

type FuncSymbol struct {
	Func     parser.IFuncContext
	funcType *types.Type
}

func (f *FuncSymbol) Ident() parser.IIdentContext {
	return f.Func.Ident()
}

func (f *FuncSymbol) SymbolName() string {
	return f.Func.Ident().GetText()
}

func (f *FuncSymbol) Type() *types.Type {
	return f.funcType
}

func NewFuncSymbol(fn parser.IFuncContext, funcType *types.Type) *FuncSymbol {
	return &FuncSymbol{Func: fn, funcType: funcType}
}

type FuncParamSymbol struct {
	Param     parser.IFuncParamContext
	paramType *types.Type
}

func NewFuncParamSymbol(param parser.IFuncParamContext, paramType *types.Type) *FuncParamSymbol {
	return &FuncParamSymbol{Param: param, paramType: paramType}
}

func (param *FuncParamSymbol) SymbolName() string {
	return param.Param.Ident().GetText()
}

func (param *FuncParamSymbol) Ident() parser.IIdentContext {
	return param.Param.Ident()
}

func (f *FuncParamSymbol) Type() *types.Type {
	return f.paramType
}

type VariableSymbol struct {
	Decl    parser.IStmtVarDeclContext
	varType *types.Type
}

func NewVariableSymbol(decl parser.IStmtVarDeclContext, varType *types.Type) *VariableSymbol {
	return &VariableSymbol{Decl: decl, varType: varType}
}

func (v *VariableSymbol) SymbolName() string {
	return v.Decl.Ident().GetText()
}

func (v *VariableSymbol) Ident() parser.IIdentContext {
	return v.Decl.Ident()
}

func (v *VariableSymbol) Type() *types.Type {
	return v.varType
}
