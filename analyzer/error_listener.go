package analyzer

import (
	"chorego/analyzer/analyzer_error"
	"fmt"
)

type ErrorListener interface {
	ReportAnalyzerError(err analyzer_error.Error)
}

type DefaultErrorListener struct {
	ProducedError bool
}

func (e *DefaultErrorListener) ReportAnalyzerError(err analyzer_error.Error) {
	e.ProducedError = true
	fmt.Println(err.Error())
}
