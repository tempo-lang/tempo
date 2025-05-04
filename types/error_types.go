package types

import (
	"chorego/misc"
	"chorego/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Error interface {
	error
	ParserRule() antlr.ParserRuleContext
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

func (e *DuplicateRolesError) ParserRule() antlr.ParserRuleContext {
	return e.Func
}

type UnknownRoleError struct {
	ValueType    parser.IValueTypeContext
	UnknownRoles []string
}

func NewUnknownRoleError(valueType parser.IValueTypeContext, unknownRoles []string) Error {
	return &UnknownRoleError{
		ValueType:    valueType,
		UnknownRoles: unknownRoles,
	}
}

func (e *UnknownRoleError) IsTypeError() {}

func (e *UnknownRoleError) Error() string {
	return fmt.Sprintf("unknown roles [%s] in '%s'", misc.JoinStrings(e.UnknownRoles, ","), e.ValueType.GetText())
}

func (e *UnknownRoleError) ParserRule() antlr.ParserRuleContext {
	return e.ValueType
}

type SymbolAlreadyExists struct {
	ExistingSymbol parser.IIdentContext
	NewSymbol      parser.IIdentContext
}

func NewSymbolAlreadyExistsError(existing parser.IIdentContext, newSym parser.IIdentContext) Error {
	return &SymbolAlreadyExists{
		ExistingSymbol: existing,
		NewSymbol:      newSym,
	}
}

func (s *SymbolAlreadyExists) IsTypeError() {}

func (s *SymbolAlreadyExists) Error() string {
	return fmt.Sprintf("symbol '%s' already declared", s.NewSymbol.GetText())
}

func (e *SymbolAlreadyExists) ParserRule() antlr.ParserRuleContext {
	return e.NewSymbol
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

func (e *UnknownTypeError) ParserRule() antlr.ParserRuleContext {
	return e.TypeName
}

type UnknownSymbolError struct {
	SymName parser.IIdentContext
}

func NewUnknownSymbolError(symName parser.IIdentContext) Error {
	return &UnknownSymbolError{
		SymName: symName,
	}
}

func (e *UnknownSymbolError) Error() string {
	return fmt.Sprintf("unknown symbol '%s'", e.SymName.GetText())
}

func (e *UnknownSymbolError) IsTypeError() {}

func (e *UnknownSymbolError) ParserRule() antlr.ParserRuleContext {
	return e.SymName
}

type InvalidDeclTypeError struct {
	DeclToken parser.IValueTypeContext
	DeclType  *Type
	ExprToken parser.IExprContext
	ExprType  *Type
}

func NewInvalidDeclTypeError(declToken parser.IValueTypeContext, declType *Type, exprToken parser.IExprContext, exprType *Type) Error {
	return &InvalidDeclTypeError{
		DeclToken: declToken,
		DeclType:  declType,
		ExprToken: exprToken,
		ExprType:  exprType,
	}
}

func (e *InvalidDeclTypeError) Error() string {
	return fmt.Sprintf("invalid declaration type, expected %s found %s", e.DeclType.ToString(), e.ExprType.ToString())
}

func (e *InvalidDeclTypeError) IsTypeError() {}

func (e *InvalidDeclTypeError) ParserRule() antlr.ParserRuleContext {
	return e.ExprToken
}

type TypeMismatchError struct {
	Expr       parser.IExprContext
	FirstType  *Type
	SecondType *Type
}

func (t *TypeMismatchError) Error() string {
	return fmt.Sprintf("types %s and %s do not match", t.FirstType.ToString(), t.SecondType.ToString())
}

func (t *TypeMismatchError) IsTypeError() {}

func (t *TypeMismatchError) ParserRule() antlr.ParserRuleContext {
	return t.Expr
}

func NewTypeMismatchError(expr parser.IExprContext, firstType *Type, secondType *Type) Error {
	return &TypeMismatchError{
		Expr:       expr,
		FirstType:  firstType,
		SecondType: secondType,
	}
}

type InvalidFunctionError struct {
	Func        parser.IFuncContext
	ParamErrors [][]Error
}

func NewInvalidFuncError(fn parser.IFuncContext, params [][]Error) Error {
	return &InvalidFunctionError{
		Func:        fn,
		ParamErrors: params,
	}
}

func (e *InvalidFunctionError) Error() string {
	paramErrors := ""
	for i := 0; i < len(e.ParamErrors); i++ {
		if len(e.ParamErrors[i]) > 0 {
			paramErrors += fmt.Sprintf("param %d: %v, ", i+1, e.ParamErrors[i])
		}
	}
	paramErrors = paramErrors[:len(paramErrors)-2]

	return fmt.Sprintf("invalid function type for '%s': %s", e.Func.Ident().GetText(), paramErrors)
}

func (i *InvalidFunctionError) IsTypeError() {}

func (e *InvalidFunctionError) ParserRule() antlr.ParserRuleContext {
	return e.Func
}

type InvalidNumberError struct {
	Num parser.IExprContext
}

func NewInvalidNumberError(num parser.IExprContext) Error {
	return &InvalidNumberError{
		Num: num,
	}
}

func (i *InvalidNumberError) Error() string {
	return fmt.Sprintf("invalid number '%s'", i.Num.GetText())
}

func (i *InvalidNumberError) IsTypeError() {}

func (i *InvalidNumberError) ParserRule() antlr.ParserRuleContext {
	return i.Num
}
