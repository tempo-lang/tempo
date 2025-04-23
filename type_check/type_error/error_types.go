package type_error

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

type UnknownRoleError struct {
	Func        parser.IFuncContext
	UnknownRole parser.IIdentContext
}

func NewUnknownRoleError(function parser.IFuncContext, unknownRole parser.IIdentContext) *UnknownRoleError {
	return &UnknownRoleError{
		Func:        function,
		UnknownRole: unknownRole,
	}
}

func (e *UnknownRoleError) analyzerError() {}

func (e *UnknownRoleError) Error() string {
	return fmt.Sprintf("unknown role '%s' in function '%s'", e.UnknownRole.GetText(), e.Func.Ident().GetText())
}

type SymbolAlreadyExists struct {
	ExistingSymbol parser.IIdentContext
	NewSymbol      parser.IIdentContext
}

func (s *SymbolAlreadyExists) analyzerError() {}

func (s *SymbolAlreadyExists) Error() string {
	return fmt.Sprintf("symbol '%s' already declared", s.NewSymbol.GetText())
}

func NewSymbolAlreadyExistsError(existing parser.IIdentContext, newSym parser.IIdentContext) *SymbolAlreadyExists {
	return &SymbolAlreadyExists{
		ExistingSymbol: existing,
		NewSymbol:      newSym,
	}
}

type UnknownTypeError struct {
	TypeName parser.IIdentContext
}

func NewUnknownTypeError(typeName parser.IIdentContext) *UnknownTypeError {
	return &UnknownTypeError{
		TypeName: typeName,
	}
}

func (e *UnknownTypeError) analyzerError() {}

func (e *UnknownTypeError) Error() string {
	return fmt.Sprintf("unknown type '%s'", e.TypeName.GetText())
}
