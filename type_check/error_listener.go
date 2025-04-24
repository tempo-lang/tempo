package type_check

import (
	"chorego/type_check/type_error"
)

type ErrorListener interface {
	ReportAnalyzerError(err type_error.Error)
}

type DefaultErrorListener struct {
	Errors []type_error.Error
}

func (e *DefaultErrorListener) ReportAnalyzerError(err type_error.Error) {
	e.Errors = append(e.Errors, err)
}
