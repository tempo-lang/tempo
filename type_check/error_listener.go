package type_check

import (
	"github.com/tempo-lang/tempo/type_check/type_error"
)

type ErrorListener interface {
	ReportTypeError(err type_error.Error)
}

type DefaultErrorListener struct {
	Errors []type_error.Error
}

func (e *DefaultErrorListener) ReportTypeError(err type_error.Error) {
	e.Errors = append(e.Errors, err)
}
