// This package exposes the [TypeCheck] function which is responsible for type checking Tempo programs.
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
	errors []type_error.Error

	currentScope    *sym_table.Scope
	info            *Info
	currentTypeHint types.Type
}

// TypeCheck takes a parsed AST and returns an [Info] object and list of type errors.
// If the list of errors is empty, then the input program is valid.
func TypeCheck(sourceFile parser.ISourceFileContext) (*Info, []type_error.Error) {
	tc := new()

	// check that tc implements visitor
	var visitor parser.TempoVisitor = tc
	_ = visitor

	sourceFile.Accept(tc)

	return tc.info, tc.errors
}

func new() *typeChecker {
	return &typeChecker{
		errors:          []type_error.Error{},
		currentScope:    nil,
		currentTypeHint: nil,
		info:            newInfo(),
	}
}

func (tc *typeChecker) reportError(err ...type_error.Error) {
	tc.errors = append(tc.errors, err...)
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

func (tc *typeChecker) VisitIdent(ctx *parser.IdentContext) any {
	return nil
}

func (tc *typeChecker) VisitRoleIdent(ctx *parser.RoleIdentContext) any {
	return nil
}
