package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type SymbolAlreadyExists struct {
	baseError
	ExistingSymbol parser.IIdentContext
	NewSymbol      parser.IIdentContext
}

func NewSymbolAlreadyExistsError(existing parser.IIdentContext, newSym parser.IIdentContext) Error {
	return &SymbolAlreadyExists{
		ExistingSymbol: existing,
		NewSymbol:      newSym,
	}
}

func (s *SymbolAlreadyExists) Error() string {
	return fmt.Sprintf("symbol `%s` is already declared", s.NewSymbol.GetText())
}

func (e *SymbolAlreadyExists) ParserRule() antlr.ParserRuleContext {
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
	SymName parser.IIdentContext
}

func NewUnknownSymbol(symName parser.IIdentContext) Error {
	return &UnknownSymbol{
		SymName: symName,
	}
}

func (e *UnknownSymbol) Error() string {
	return fmt.Sprintf("value `%s` is undefined or not in scope", e.SymName.GetText())
}

func (e *UnknownSymbol) ParserRule() antlr.ParserRuleContext {
	return e.SymName
}

func (e *UnknownSymbol) Code() ErrorCode {
	return CodeUnknownSymbol
}

type UnassignableSymbol struct {
	baseError
	Assign *parser.StmtAssignContext
	Type   types.Type
}

func (u *UnassignableSymbol) Error() string {
	return fmt.Sprintf("type `%s` is not assignable", u.Type.ToString())
}

func (u *UnassignableSymbol) ParserRule() antlr.ParserRuleContext {
	return u.Assign
}

func (e *UnassignableSymbol) Code() ErrorCode {
	return CodeUnassignableSymbol
}

func NewUnassignableSymbol(assign *parser.StmtAssignContext, symType types.Type) Error {
	return &UnassignableSymbol{
		Assign: assign,
		Type:   symType,
	}
}

type FieldAccessUnknownField struct {
	baseError
	FieldIdent parser.IIdentContext
	BaseType   types.Type
}

func (e *FieldAccessUnknownField) Error() string {
	return fmt.Sprintf("value of type `%s` has not field named `%s`", e.BaseType.ToString(), e.FieldIdent.GetText())
}

func (e *FieldAccessUnknownField) ParserRule() antlr.ParserRuleContext {
	return e.FieldIdent
}

func (e *FieldAccessUnknownField) Code() ErrorCode {
	return CodeFieldAccessUnknownField
}

func NewFieldAccessUnknownField(fieldIdent parser.IIdentContext, baseType types.Type) Error {
	return &FieldAccessUnknownField{
		FieldIdent: fieldIdent,
		BaseType:   baseType,
	}
}

type ExpectedInterfaceType struct {
	baseError
	sym   sym_table.Symbol
	ident parser.IIdentContext
}

func (e *ExpectedInterfaceType) Error() string {
	return fmt.Sprintf("type `%s` is not an interface", e.sym.Type().ToString())
}

func (e *ExpectedInterfaceType) ParserRule() antlr.ParserRuleContext {
	return e.ident
}

func (e *ExpectedInterfaceType) Code() ErrorCode {
	return CodeExpectedInterfaceType
}

func NewExpectedInterfaceType(sym sym_table.Symbol, ident parser.IIdentContext) Error {
	return &ExpectedInterfaceType{
		sym:   sym,
		ident: ident,
	}
}
