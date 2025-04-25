package misc

import (
	"errors"

	"github.com/antlr4-go/antlr/v4"
)

type ErrorListener struct {
	Errors []error
}

// ReportAmbiguity implements antlr.ErrorListener.
func (t *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	t.Errors = append(t.Errors, errors.New("ambiguity error"))
}

// ReportAttemptingFullContext implements antlr.ErrorListener.
func (t *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	t.Errors = append(t.Errors, errors.New("attempting full context error"))
}

// ReportContextSensitivity implements antlr.ErrorListener.
func (t *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, prediction int, configs *antlr.ATNConfigSet) {
	t.Errors = append(t.Errors, errors.New("context sensitivity error"))
}

// SyntaxError implements antlr.ErrorListener.
func (t *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	t.Errors = append(t.Errors, errors.New("syntax error"))
}
