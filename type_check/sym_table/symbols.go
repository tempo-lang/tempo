package sym_table

import "chorego/parser"

type Symbol interface {
	SymbolName() string
	Ident() parser.IIdentContext
}

type FuncSymbol struct {
	Func parser.IFuncContext
}

func (f *FuncSymbol) Ident() parser.IIdentContext {
	return f.Func.Ident()
}

func (f *FuncSymbol) SymbolName() string {
	return f.Func.Ident().GetText()
}

func NewFuncSymbol(fn parser.IFuncContext) *FuncSymbol {
	return &FuncSymbol{Func: fn}
}

type FuncParamSymbol struct {
	Param parser.IFuncParamContext
}

func NewFuncParamSymbol(param parser.IFuncParamContext) *FuncParamSymbol {
	return &FuncParamSymbol{Param: param}
}

func (param *FuncParamSymbol) SymbolName() string {
	return param.Param.Ident().GetText()
}

func (param *FuncParamSymbol) Ident() parser.IIdentContext {
	return param.Param.Ident()
}
