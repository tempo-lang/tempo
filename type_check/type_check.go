package type_check

import (
	"chorego/parser"
	"chorego/type_check/sym_table"
)

type TypeChecker struct {
	parser.BaseChoregoListener
	ErrorListener ErrorListener

	funcScopeStack []parser.IFuncContext
	symTable       *sym_table.SymTable
}

func New() *TypeChecker {
	return &TypeChecker{
		ErrorListener:  &DefaultErrorListener{},
		funcScopeStack: []parser.IFuncContext{},
		symTable:       sym_table.New(),
	}
}

func (tc *TypeChecker) funcScope() parser.IFuncContext {
	if len(tc.funcScopeStack) == 0 {
		return nil
	}
	return tc.funcScopeStack[len(tc.funcScopeStack)-1]
}

func (tc *TypeChecker) EnterFunc(ctx *parser.FuncContext) {
	tc.funcScopeStack = append(tc.funcScopeStack, ctx)
	tc.symTable.EnterScope()

	tc.checkFuncDuplicateRoles(ctx)
}

func (tc *TypeChecker) ExitFunc(ctx *parser.FuncContext) {
	tc.symTable.ExitScope()
	tc.funcScopeStack = tc.funcScopeStack[:len(tc.funcScopeStack)-1]
}

func (tc *TypeChecker) EnterFuncParam(ctx *parser.FuncParamContext) {
	tc.checkFuncParamUnknownRoles(ctx)

	err := tc.symTable.InsertSymbol(sym_table.NewFuncParamSymbol(ctx))
	if err != nil {
		tc.ErrorListener.ReportAnalyzerError(err)
	}
}

func (tc *TypeChecker) EnterValueType(ctx *parser.ValueTypeContext) {
	tc.checkValueType(ctx)
}
