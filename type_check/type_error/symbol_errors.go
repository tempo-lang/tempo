package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"

	"github.com/tempo-lang/tempo/parser/new_parser/ast"
)

type SymbolAlreadyExists struct {
	baseError
	ExistingSymbol *ast.Identifier
	NewSymbol      *ast.Identifier
}

func NewSymbolAlreadyExistsError(existing *ast.Identifier, newSym *ast.Identifier) Error {
	return &SymbolAlreadyExists{
		ExistingSymbol: existing,
		NewSymbol:      newSym,
	}
}

func (s *SymbolAlreadyExists) Error() string {
	return fmt.Sprintf("symbol `%s` is already declared", s.NewSymbol.Value())
}

func (e *SymbolAlreadyExists) ParserRule() ast.Node {
	return e.NewSymbol
}

func (e *SymbolAlreadyExists) Code() ErrorCode {
	return CodeSymbolAlreadyExists
}

func (e *SymbolAlreadyExists) RelatedInfo() []RelatedInfo {
	return []RelatedInfo{{
		Message:    "symbol is first declared here",
		ParserRule: e.ExistingSymbol,
	}}
}

type UnknownSymbol struct {
	baseError
	SymName *ast.Identifier
}

func NewUnknownSymbol(symName *ast.Identifier) Error {
	return &UnknownSymbol{
		SymName: symName,
	}
}

func (e *UnknownSymbol) Error() string {
	return fmt.Sprintf("value `%s` is undefined or not in scope", e.SymName.Value())
}

func (e *UnknownSymbol) ParserRule() ast.Node {
	return e.SymName
}

func (e *UnknownSymbol) Code() ErrorCode {
	return CodeUnknownSymbol
}

type UnassignableSymbol struct {
	baseError
	Assign *ast.AssignStmt
	Type   types.Type
}

func (u *UnassignableSymbol) Error() string {
	return fmt.Sprintf("type `%s` is not assignable", u.Type.ToString())
}

func (u *UnassignableSymbol) ParserRule() ast.Node {
	return u.Assign
}

func (e *UnassignableSymbol) Code() ErrorCode {
	return CodeUnassignableSymbol
}

func NewUnassignableSymbol(assign *ast.AssignStmt, symType types.Type) Error {
	return &UnassignableSymbol{
		Assign: assign,
		Type:   symType,
	}
}

type FieldAccessUnknownField struct {
	baseError
	FieldIdent *ast.Identifier
	BaseType   types.Type
}

func (e *FieldAccessUnknownField) Error() string {
	return fmt.Sprintf("value of type `%s` has not field named `%s`", e.BaseType.ToString(), e.FieldIdent.Value())
}

func (e *FieldAccessUnknownField) ParserRule() ast.Node {
	return e.FieldIdent
}

func (e *FieldAccessUnknownField) Code() ErrorCode {
	return CodeFieldAccessUnknownField
}

func NewFieldAccessUnknownField(fieldIdent *ast.Identifier, baseType types.Type) Error {
	return &FieldAccessUnknownField{
		FieldIdent: fieldIdent,
		BaseType:   baseType,
	}
}

type ExpectedInterfaceType struct {
	baseError
	sym   sym_table.Symbol
	ident *ast.Identifier
}

func (e *ExpectedInterfaceType) Error() string {
	return fmt.Sprintf("type `%s` is not an interface", e.sym.Type().ToString())
}

func (e *ExpectedInterfaceType) ParserRule() ast.Node {
	return e.ident
}

func (e *ExpectedInterfaceType) Code() ErrorCode {
	return CodeExpectedInterfaceType
}

func NewExpectedInterfaceType(sym sym_table.Symbol, ident *ast.Identifier) Error {
	return &ExpectedInterfaceType{
		sym:   sym,
		ident: ident,
	}
}
