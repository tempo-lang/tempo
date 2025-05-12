package type_check

import "tempo/types"

type ErrorListener interface {
	ReportTypeError(err types.Error)
}

type DefaultErrorListener struct {
	Errors []types.Error
}

func (e *DefaultErrorListener) ReportTypeError(err types.Error) {
	e.Errors = append(e.Errors, err)
}
