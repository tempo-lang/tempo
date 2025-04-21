package type_check

import (
	"chorego/type_check/type_error"
	"fmt"
)

type ErrorListener interface {
	ReportAnalyzerError(err type_error.Error)
}

type DefaultErrorListener struct {
	ProducedError bool
}

func (e *DefaultErrorListener) ReportAnalyzerError(err type_error.Error) {
	e.ProducedError = true
	fmt.Println(err.Error())
}
