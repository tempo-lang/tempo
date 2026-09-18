// This package exposes the [TypeCheck] function which is responsible for type checking Tempo programs.
package type_check

import (
	"github.com/tempo-lang/tempo/parser/ast"
	"github.com/tempo-lang/tempo/parser/token"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

type typeChecker struct {
	errors []type_error.Error

	currentScope    *sym_table.Scope
	info            *Info
	currentTypeHint types.Type
}

func nodeSpan(n ast.Node) token.Span {
	return token.Span{Start: n.StartToken().Span.Start, End: n.EndToken().Span.End}
}

// TypeCheck takes a parsed AST and returns an [Info] object and list of type errors.
// If the list of errors is empty, then the input program is valid.
func TypeCheck(sourceFile *ast.SourceFile) (*Info, []type_error.Error) {
	tc := new()
	tc.checkSourceFile(sourceFile)

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

func (tc *typeChecker) checkSourceFile(ctx *ast.SourceFile) {
	span := token.Span{Start: ctx.StartToken().Span.Start, End: ctx.EndToken().Span.End}
	tc.info.GlobalScope = sym_table.NewScope(span, nil, nil)
	tc.currentScope = tc.info.GlobalScope
	tc.populateGlobalSymbols()

	tc.addGlobalSymbols(ctx)

	for _, inf := range ctx.Interfaces {
		tc.visitInterface(inf)
	}

	for _, st := range ctx.Structs {
		tc.visitStruct(st)
	}

	for _, fn := range ctx.Functions {
		tc.visitFunc(fn)
	}

	tc.currentScope = tc.currentScope.Parent()
}

func (tc *typeChecker) visitScope(ctx *ast.Scope) bool {
	returnsValue := false
	for _, stmt := range ctx.Stmts {
		returnsValue = returnsValue || tc.visitStmt(stmt)
	}
	return returnsValue
}
