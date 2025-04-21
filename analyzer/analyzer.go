package analyzer

import (
	"chorego/analyzer/sym_table"
	"chorego/parser"
)

type AnalyzerListener struct {
	parser.BaseChoregoListener
	ErrorListener ErrorListener

	funcScopeStack []parser.IFuncContext
	symTable       *sym_table.SymTable
}

func New() *AnalyzerListener {
	return &AnalyzerListener{
		ErrorListener:  &DefaultErrorListener{},
		funcScopeStack: []parser.IFuncContext{},
		symTable:       sym_table.New(),
	}
}

func (a *AnalyzerListener) funcScope() parser.IFuncContext {
	if len(a.funcScopeStack) == 0 {
		return nil
	}
	return a.funcScopeStack[len(a.funcScopeStack)-1]
}

func (a *AnalyzerListener) EnterFunc(ctx *parser.FuncContext) {
	a.funcScopeStack = append(a.funcScopeStack, ctx)
	a.symTable.EnterScope()

	a.checkFuncDuplicateRoles(ctx)
}

func (a *AnalyzerListener) ExitFunc(ctx *parser.FuncContext) {
	a.symTable.ExitScope()
	a.funcScopeStack = a.funcScopeStack[:len(a.funcScopeStack)-1]
}

func (a *AnalyzerListener) EnterFunc_param(ctx *parser.Func_paramContext) {
	a.checkFuncParamUnknownRoles(ctx)

	err := a.symTable.InsertSymbol(sym_table.NewFuncParamSymbol(ctx))
	if err != nil {
		a.ErrorListener.ReportAnalyzerError(err)
	}
}
