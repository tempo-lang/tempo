package type_check_test

import (
	"chorego/compiler"
	"chorego/misc"
	"go/types"
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
	errors []types.Error
}

// ReportTypeError implements analyzer.ErrorListener.
func (t *TestErrorListener) ReportTypeError(err types.Error) {
	t.errors = append(t.errors, err)
}

func NewTestErrorListener() *TestErrorListener {
	return &TestErrorListener{}
}

func (data *AnalyzerTestData) Assert(t *testing.T) {
	t.Run(data.name, func(t *testing.T) {

		input := antlr.NewInputStream(data.input)
		_, errors := compiler.Compile(input)

		for i := range min(len(data.errors), len(errors)) {
			expected := data.errors[i]
			actual := errors[i].Error()
			if expected != actual {
				t.Errorf("error %d did not match, expected '%s', got '%s'.", i, expected, actual)
			}
		}

		if len(data.errors) > len(errors) {
			t.Errorf("unexpected few actual errors: %v", data.errors[len(errors):])
		}

		if len(errors) > len(data.errors) {
			t.Errorf("unexpected extra actual errors, expected %d got %d:\n- %s", len(data.errors), len(errors), misc.JoinStrings(errors, "\n- "))
		}
	})
}
