package analyzer

import (
	"chorego/parser"
	"fmt"
)

type Error interface {
	error
	analyzerError()
}

type DuplicateRolesError struct {
	Func           parser.IFuncContext
	DuplicateRoles []parser.IIdentContext
}

func NewDuplicateRolesError(function parser.IFuncContext, duplicateRoles []parser.IIdentContext) *DuplicateRolesError {
	return &DuplicateRolesError{
		Func:           function,
		DuplicateRoles: duplicateRoles,
	}
}

func (e *DuplicateRolesError) analyzerError() {}

func (e *DuplicateRolesError) Error() string {
	return fmt.Sprintf("function '%s' has duplicate role '%s'", e.Func.Ident().GetText(), e.DuplicateRoles[0].GetText())
}
