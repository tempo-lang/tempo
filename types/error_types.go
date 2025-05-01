package types

import (
	"chorego/parser"
	"fmt"
)

type Error interface {
	error
	IsTypeError()
}

type DuplicateRolesError struct {
	Func           parser.IFuncContext
	DuplicateRoles []parser.IIdentContext
}

func NewDuplicateRolesError(function parser.IFuncContext, duplicateRoles []parser.IIdentContext) Error {
	return &DuplicateRolesError{
		Func:           function,
		DuplicateRoles: duplicateRoles,
	}
}

func (e *DuplicateRolesError) IsTypeError() {}

func (e *DuplicateRolesError) Error() string {
	return fmt.Sprintf("function '%s' has duplicate role '%s'", e.Func.Ident().GetText(), e.DuplicateRoles[0].GetText())
}

type UnknownRoleError struct {
	Func        parser.IFuncContext
	UnknownRole parser.IIdentContext
}

func NewUnknownRoleError(function parser.IFuncContext, unknownRole parser.IIdentContext) Error {
	return &UnknownRoleError{
		Func:        function,
		UnknownRole: unknownRole,
	}
}

func (e *UnknownRoleError) IsTypeError() {}

func (e *UnknownRoleError) Error() string {
	return fmt.Sprintf("unknown role '%s' in function '%s'", e.UnknownRole.GetText(), e.Func.Ident().GetText())
}

type SymbolAlreadyExists struct {
	ExistingSymbol parser.IIdentContext
	NewSymbol      parser.IIdentContext
}

func (s *SymbolAlreadyExists) IsTypeError() {}

func (s *SymbolAlreadyExists) Error() string {
	return fmt.Sprintf("symbol '%s' already declared", s.NewSymbol.GetText())
}

func NewSymbolAlreadyExistsError(existing parser.IIdentContext, newSym parser.IIdentContext) Error {
	return &SymbolAlreadyExists{
		ExistingSymbol: existing,
		NewSymbol:      newSym,
	}
}

type UnknownTypeError struct {
	TypeName parser.IIdentContext
}

func NewUnknownTypeError(typeName parser.IIdentContext) Error {
	return &UnknownTypeError{
		TypeName: typeName,
	}
}

func (e *UnknownTypeError) IsTypeError() {}

func (e *UnknownTypeError) Error() string {
	return fmt.Sprintf("unknown type '%s'", e.TypeName.GetText())
}

type UnknownSymbolError struct {
	SymName parser.IIdentContext
}

func (e *UnknownSymbolError) Error() string {
	return fmt.Sprintf("unknown symbol '%s'", e.SymName.GetText())
}

func (e *UnknownSymbolError) IsTypeError() {}

func NewUnknownSymbolError(symName parser.IIdentContext) Error {
	return &UnknownSymbolError{
		SymName: symName,
	}
}

type TypeMismatchError struct {
	DeclToken parser.IValueTypeContext
	DeclType  *Type
	ExprToken parser.IExpressionContext
	ExprType  *Type
}

func (e *TypeMismatchError) Error() string {
	return fmt.Sprintf("type mismatch, expected %s found %s", e.DeclType.ToString(), e.ExprType.ToString())
}

func (e *TypeMismatchError) IsTypeError() {}

func NewTypeMismatchError(declToken parser.IValueTypeContext, declType *Type, exprToken parser.IExpressionContext, exprType *Type) Error {
	return &TypeMismatchError{
		DeclToken: declToken,
		DeclType:  declType,
		ExprToken: exprToken,
		ExprType:  exprType,
	}
}

type InvalidFunctionError struct {
	Func        parser.IFuncContext
	ParamErrors map[int][]Error
}

func NewInvalidFuncError(fn parser.IFuncContext, params map[int][]Error) Error {
	return &InvalidFunctionError{
		Func:        fn,
		ParamErrors: params,
	}
}

func (i *InvalidFunctionError) Error() string {
	return fmt.Sprintf("invalid function type for '%s'", i.Func.Ident().GetText())
}

func (i *InvalidFunctionError) IsTypeError() {}
