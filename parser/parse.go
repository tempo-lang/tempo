package parser

import (
	"fmt"
	"tempo/misc"

	"github.com/antlr4-go/antlr/v4"
)

func Parse(input antlr.CharStream) (sourceFile ISourceFileContext, errors []error) {
	// input, _ := antlr.NewFileStream(os.Args[1])
	errorListener := misc.ErrorListener{}

	// lexer
	lexer := NewTempoLexer(input)
	lexer.AddErrorListener(&errorListener)

	// parser
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewTempoParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	p.AddErrorListener(&errorListener)

	// parse program
	sourceFile = p.SourceFile()

	if len(errorListener.Errors) > 0 {
		errors = errorListener.Errors
	}

	return
}

func RoleTypeAllIdents(ctx IRoleTypeContext) []IIdentContext {
	switch ctx := ctx.(type) {
	case *RoleTypeNormalContext:
		return ctx.AllIdent()
	case *RoleTypeSharedContext:
		return ctx.AllIdent()
	}
	panic(fmt.Sprintf("unexpected IRoleTypeContext: %#v", ctx))
}
