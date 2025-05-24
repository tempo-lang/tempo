package type_check

import (
	"tempo/parser"
	"tempo/sym_table"
	"tempo/type_check/type_error"
	"tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type typeChecker struct {
	*antlr.BaseParseTreeVisitor
	errorListener ErrorListener

	currentScope *sym_table.Scope
	info         *Info
}

func TypeCheck(sourceFile parser.ISourceFileContext) (*Info, []type_error.Error) {
	tc := new()

	// check that tc implements visitor
	var visitor parser.TempoVisitor = tc
	_ = visitor

	sourceFile.Accept(tc)

	analyzerErrorListener, ok := tc.errorListener.(*DefaultErrorListener)
	if !ok {
		panic("analyzer error listener was expected to be DefaultErrorListener")
	}

	return tc.info, analyzerErrorListener.Errors
}

func new() *typeChecker {
	return &typeChecker{
		errorListener: &DefaultErrorListener{},
		currentScope:  nil,
		info:          newInfo(),
	}
}

func (tc *typeChecker) reportError(err type_error.Error) {
	tc.errorListener.ReportTypeError(err)
}

func (tc *typeChecker) VisitSourceFile(ctx *parser.SourceFileContext) (result any) {
	tc.info.GlobalScope = sym_table.NewScope(ctx.GetStart(), ctx.GetStop(), nil, nil)
	tc.currentScope = tc.info.GlobalScope

	tc.addGlobalSymbols(ctx)

	for _, st := range ctx.AllStruct_() {
		st.Accept(tc)
	}

	for _, fn := range ctx.AllFunc_() {
		fn.Accept(tc)
	}

	tc.currentScope = tc.currentScope.Parent()
	return
}

func (tc *typeChecker) VisitFunc(ctx *parser.FuncContext) any {
	// functions are already resolved by addGlobalSymbols
	sym, found := tc.info.Symbols[ctx.FuncSig().GetName()].(*sym_table.FuncSymbol)
	if !found {
		// was not found if function has parser errors
		return nil
	}

	tc.currentScope = sym.Scope()

	if sym.Type().Roles().IsSharedRole() {
		tc.reportError(type_error.NewUnexpectedSharedTypeError(ctx.FuncSig().GetName(), sym.Type()))
	}

	ctx.FuncSig().Accept(tc)

	// nil if parser error
	if ctx.Scope() != nil {
		for _, stmt := range ctx.Scope().AllStmt() {
			stmt.Accept(tc)
		}
	}

	tc.currentScope = tc.currentScope.Parent()
	return nil
}

func (tc *typeChecker) VisitFuncSig(ctx *parser.FuncSigContext) any {

	tc.checkFuncDuplicateRoles(ctx)
	ctx.FuncParamList().Accept(tc)

	return nil
}

func (tc *typeChecker) VisitFuncParam(ctx *parser.FuncParamContext) any {
	paramType := tc.visitValueType(ctx.ValueType())
	paramSym := sym_table.NewFuncParamSymbol(ctx, tc.currentScope, paramType)
	tc.insertSymbol(paramSym)

	funcSym := tc.currentScope.GetFunc()
	funcSym.AddParam(paramSym.(*sym_table.FuncParamSymbol))

	return nil
}

func (tc *typeChecker) VisitFuncParamList(ctx *parser.FuncParamListContext) any {
	for _, param := range ctx.AllFuncParam() {
		param.Accept(tc)
	}
	return nil
}

func (tc *typeChecker) VisitScope(ctx *parser.ScopeContext) any {
	return nil
}

func (tc *typeChecker) visitValueType(ctx parser.IValueTypeContext) *types.Type {
	if ctx == nil {
		return types.Invalid()
	}

	valType, err := tc.parseValueType(ctx)
	if err != nil {
		tc.reportError(err)
		return types.Invalid()
	}

	return valType
}

func (tc *typeChecker) VisitValueType(ctx *parser.ValueTypeContext) any {
	return nil
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
