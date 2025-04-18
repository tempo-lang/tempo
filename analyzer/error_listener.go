package analyzer

import "fmt"

type ErrorListener interface {
	ReportAnalyzerError(err Error)
}

type DefaultErrorListener struct {
	ProducedError bool
}

func (e *DefaultErrorListener) ReportAnalyzerError(err Error) {
	e.ProducedError = true
	fmt.Println(err.Error())
}
