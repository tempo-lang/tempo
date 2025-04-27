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

type VariableSymbol struct {
	Decl parser.IStmtVarDeclContext
}

func NewVariableSymbol(decl parser.IStmtVarDeclContext) *VariableSymbol {
	return &VariableSymbol{Decl: decl}
}

func (v *VariableSymbol) SymbolName() string {
	return v.Decl.Ident().GetText()
}

func (v *VariableSymbol) Ident() parser.IIdentContext {
	return v.Decl.Ident()
}
