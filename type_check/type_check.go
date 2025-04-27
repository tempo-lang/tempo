package type_check

import (
	"chorego/parser"
	"chorego/type_check/sym_table"
	"chorego/type_check/type_error"

	"github.com/antlr4-go/antlr/v4"
)

type typeChecker struct {
	*antlr.BaseParseTreeVisitor
	errorListener ErrorListener

	funcScopeStack []parser.IFuncContext
	symTable       *sym_table.SymTable
}

func TypeCheck(sourceFile parser.ISourceFileContext) []type_error.Error {
	tc := new()

	// check that tc implements visitor
	var visitor parser.ChoregoVisitor = tc
	_ = visitor

	sourceFile.Accept(tc)

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

func (tc *typeChecker) reportError(err type_error.Error) {
	tc.errorListener.ReportTypeError(err)
}

func (tc *typeChecker) funcScope() parser.IFuncContext {
	if len(tc.funcScopeStack) == 0 {
		return nil
	}
	return tc.funcScopeStack[len(tc.funcScopeStack)-1]
}

func (tc *typeChecker) VisitSourceFile(ctx *parser.SourceFileContext) (result any) {
	globalScope := tc.symTable.EnterScope() // enter global scope

	typeErrors := addGlobalSymbols(globalScope, ctx)
	for _, err := range typeErrors {
		tc.reportError(err)
	}

	for _, fn := range ctx.AllFunc_() {
		fn.Accept(tc)
	}

	tc.symTable.ExitScope() // exit global scope
	return
}

func (tc *typeChecker) VisitFunc(ctx *parser.FuncContext) any {
	tc.funcScopeStack = append(tc.funcScopeStack, ctx)
	tc.symTable.EnterScope()

	tc.checkFuncDuplicateRoles(ctx)

	ctx.FuncParamList().Accept(tc)

	for _, stmt := range ctx.Scope().AllStatement() {
		stmt.Accept(tc)
	}

	tc.symTable.ExitScope()
	tc.funcScopeStack = tc.funcScopeStack[:len(tc.funcScopeStack)-1]
	return nil
}

func (tc *typeChecker) VisitFuncParam(ctx *parser.FuncParamContext) any {
	tc.checkFuncParamUnknownRoles(ctx)

	err := tc.symTable.InsertSymbol(sym_table.NewFuncParamSymbol(ctx))
	if err != nil {
		tc.reportError(err)
	}

	return tc.VisitChildren(ctx)
}

func (tc *typeChecker) VisitFuncParamList(ctx *parser.FuncParamListContext) any {
	for _, param := range ctx.AllFuncParam() {
		param.Accept(tc)
	}
	return nil
}

func (tc *typeChecker) VisitScope(ctx *parser.ScopeContext) any {
	tc.symTable.EnterScope()

	for _, stmt := range ctx.AllStatement() {
		stmt.Accept(tc)
	}

	tc.symTable.ExitScope()
	return nil
}

func (tc *typeChecker) VisitValueType(ctx *parser.ValueTypeContext) any {
	valType, err := ParseValueType(ctx)
	if err != nil {
		tc.reportError(err)
		return nil
	}

	return valType
}

func (tc *typeChecker) VisitIdent(ctx *parser.IdentContext) any {
	return nil
}

func (tc *typeChecker) VisitRoleType(ctx *parser.RoleTypeContext) any {
	return nil
}

func (tc *typeChecker) VisitRoleTypeNormal(ctx *parser.RoleTypeNormalContext) any {
	return nil
}

func (tc *typeChecker) VisitRoleTypeShared(ctx *parser.RoleTypeSharedContext) any {
	return nil
}
