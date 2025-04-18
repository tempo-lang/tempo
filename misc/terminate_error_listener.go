package misc

import "github.com/antlr4-go/antlr/v4"

type TerminateErrorListener struct {
	ProducedError bool
}

// ReportAmbiguity implements antlr.ErrorListener.
func (t *TerminateErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	t.ProducedError = true
}

// ReportAttemptingFullContext implements antlr.ErrorListener.
func (t *TerminateErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	t.ProducedError = true
}

// ReportContextSensitivity implements antlr.ErrorListener.
func (t *TerminateErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, prediction int, configs *antlr.ATNConfigSet) {
	t.ProducedError = true
}

// SyntaxError implements antlr.ErrorListener.
func (t *TerminateErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	t.ProducedError = true
}
