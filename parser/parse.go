package parser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func Parse(input antlr.CharStream) (sourceFile ISourceFileContext, errors []*SyntaxError) {
	// input, _ := antlr.NewFileStream(os.Args[1])
	errorListener := ErrorListener{}

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

type SyntaxError struct {
	line   int
	column int
	msg    string
}

func newSyntaxError(line int, column int, msg string) *SyntaxError {
	return &SyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	}
}

func (e *SyntaxError) Line() int {
	return e.line
}

func (e *SyntaxError) Column() int {
	return e.column
}

func (e *SyntaxError) Message() string {
	return e.msg
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("syntax error: %d:%d: %s", e.line, e.column, e.msg)
}

type ErrorListener struct {
	Errors []*SyntaxError
}

// ReportAmbiguity implements antlr.ErrorListener.
func (t *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// t.Errors = append(t.Errors, errors.New("ambiguity error"))
	fmt.Printf("ambiguity error")
}

// ReportAttemptingFullContext implements antlr.ErrorListener.
func (t *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// t.Errors = append(t.Errors, errors.New("attempting full context error"))
	fmt.Printf("context error")
}

// ReportContextSensitivity implements antlr.ErrorListener.
func (t *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, prediction int, configs *antlr.ATNConfigSet) {
	// t.Errors = append(t.Errors, errors.New("context sensitivity error"))
	fmt.Printf("sensitivity error")
}

// SyntaxError implements antlr.ErrorListener.
func (t *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	t.Errors = append(
		t.Errors,
		newSyntaxError(line, column, msg),
	)
}
