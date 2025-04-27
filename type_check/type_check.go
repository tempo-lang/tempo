package type_check

import (
	"chorego/parser"
	"chorego/type_check/sym_table"
	"chorego/type_check/type_error"

	"github.com/antlr4-go/antlr/v4"
)

type typeChecker struct {
	parser.BaseChoregoListener
	errorListener ErrorListener

	funcScopeStack []parser.IFuncContext
	symTable       *sym_table.SymTable
}

func TypeCheck(sourceFile parser.ISourceFileContext) []type_error.Error {
	tc := new()

	antlr.ParseTreeWalkerDefault.Walk(tc, sourceFile)

	analyzerErrorListener, ok := tc.errorListener.(*DefaultErrorListener)
	if !ok {
		panic("analyzer error listener was expected to be DefaultErrorListener")
	}

	return analyzerErrorListener.Errors
}

func new() *typeChecker {
	return &typeChecker{
		errorListener:  &DefaultErrorListener{},
		funcScopeStack: []parser.IFuncContext{},
		symTable:       sym_table.New(),
	}
}

func (tc *typeChecker) reportTypeError(err type_error.Error) {
	tc.errorListener.ReportTypeError(err)
}

func (tc *typeChecker) funcScope() parser.IFuncContext {
	if len(tc.funcScopeStack) == 0 {
		return nil
	}
	return tc.funcScopeStack[len(tc.funcScopeStack)-1]
}

func (tc *typeChecker) EnterSourceFile(ctx *parser.SourceFileContext) {
	globalScope := tc.symTable.EnterScope() // enter global scope
	typeErrors := addGlobalSymbols(globalScope, ctx)
	for _, err := range typeErrors {
		tc.reportTypeError(err)
	}
}

func (tc *typeChecker) ExitSourceFile(ctx *parser.SourceFileContext) {
	tc.symTable.ExitScope() // exit global scope
}

func (tc *typeChecker) EnterFunc(ctx *parser.FuncContext) {
	tc.funcScopeStack = append(tc.funcScopeStack, ctx)
	tc.symTable.EnterScope()

	tc.checkFuncDuplicateRoles(ctx)
}

func (tc *typeChecker) ExitFunc(ctx *parser.FuncContext) {
	tc.symTable.ExitScope()
	tc.funcScopeStack = tc.funcScopeStack[:len(tc.funcScopeStack)-1]
}

func (tc *typeChecker) EnterFuncParam(ctx *parser.FuncParamContext) {
	tc.checkFuncParamUnknownRoles(ctx)

	err := tc.symTable.InsertSymbol(sym_table.NewFuncParamSymbol(ctx))
	if err != nil {
		tc.reportTypeError(err)
	}
}

func (tc *typeChecker) EnterValueType(ctx *parser.ValueTypeContext) {
	tc.checkValueType(ctx)
}

func (tc *typeChecker) EnterStmtVarDecl(ctx *parser.StmtVarDeclContext) {
	err := tc.symTable.InsertSymbol(sym_table.NewVariableSymbol(ctx))
	if err != nil {
		tc.reportTypeError(err)
	}
}
