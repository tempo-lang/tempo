package analyzer

import "fmt"

type ErrorListener interface {
	ReportAnalyzerError(err Error)
}

type DefaultErrorListener struct{}

func (e *DefaultErrorListener) ReportAnalyzerError(err Error) {
	fmt.Println(err.Error())
}
