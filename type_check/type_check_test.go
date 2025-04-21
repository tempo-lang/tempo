package type_check_test

import (
	"chorego/parser"
	"chorego/type_check"
	"chorego/type_check/type_error"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

type AnalyzerTestData struct {
	name   string
	input  string
	errors []string
}

type AntlrTestErrorListener struct {
	t *testing.T
}

func NewAntlrTestErrorListener(t *testing.T) *AntlrTestErrorListener {
	return &AntlrTestErrorListener{t}
}

// ReportAmbiguity implements antlr.ErrorListener.
func (a *AntlrTestErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	a.t.Error("test failed to parse input")
}

// ReportAttemptingFullContext implements antlr.ErrorListener.
func (a *AntlrTestErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	a.t.Error("test failed to parse input")
}

// ReportContextSensitivity implements antlr.ErrorListener.
func (a *AntlrTestErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, prediction int, configs *antlr.ATNConfigSet) {
	a.t.Error("test failed to parse input")
}

// SyntaxError implements antlr.ErrorListener.
func (a *AntlrTestErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	a.t.Error("test failed to parse input")
}

type TestErrorListener struct {
	errors []type_error.Error
}

// ReportAnalyzerError implements analyzer.ErrorListener.
func (t *TestErrorListener) ReportAnalyzerError(err type_error.Error) {
	t.errors = append(t.errors, err)
}

func NewTestErrorListener() *TestErrorListener {
	return &TestErrorListener{}
}

func (data *AnalyzerTestData) Assert(t *testing.T) {
	t.Run(data.name, func(t *testing.T) {
		t.Parallel()

		input := antlr.NewInputStream(data.input)
		lexer := parser.NewChoregoLexer(input)
		lexer.AddErrorListener(NewAntlrTestErrorListener(t))
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewChoregoParser(stream)
		p.AddErrorListener(NewAntlrTestErrorListener(t))

		f := p.Func_()

		a := type_check.New()
		errorListener := NewTestErrorListener()
		a.ErrorListener = errorListener

		antlr.ParseTreeWalkerDefault.Walk(a, f)

		for i := range min(len(data.errors), len(errorListener.errors)) {
			expected := data.errors[i]
			actual := errorListener.errors[i].Error()
			if expected != actual {
				t.Errorf("error %d did not match, expected '%s', got '%s'.", i, expected, actual)
			}
		}

		if len(data.errors) > len(errorListener.errors) {
			t.Errorf("unexpected few actual errors: %v", data.errors[len(errorListener.errors):])
		}

		if len(errorListener.errors) > len(data.errors) {
			t.Errorf("unexpected extra actual errors, expected %d got %d", len(data.errors), len(errorListener.errors))
		}
	})
}
