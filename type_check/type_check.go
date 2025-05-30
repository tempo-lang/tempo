package type_check

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"

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

	for _, inf := range ctx.AllInterface_() {
		inf.Accept(tc)
	}

	for _, fn := range ctx.AllFunc_() {
		fn.Accept(tc)
	}

	tc.currentScope = tc.currentScope.Parent()
	return
}

func (tc *typeChecker) VisitScope(ctx *parser.ScopeContext) any {
	returnsValue := false
	for _, stmt := range ctx.AllStmt() {
		result := stmt.Accept(tc)
		returnsValue = returnsValue || result == true
	}
	return returnsValue
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

	if err := tc.checkDuplicateRoles(ctx.RoleType(), valType.Roles()); err != nil {
		tc.reportError(err)
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
