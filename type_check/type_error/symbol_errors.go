package type_error

import (
	"fmt"
	"tempo/parser"
	"tempo/sym_table"
	"tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

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

type UnassignableSymbolError struct {
	Assign *parser.StmtAssignContext
	Type   *types.Type
}

func (u *UnassignableSymbolError) Error() string {
	return fmt.Sprintf("can not assign value to '%s' of type '%s'", u.Assign.Ident().GetText(), u.Type.ToString())
}

func (u *UnassignableSymbolError) IsTypeError() {}

func (u *UnassignableSymbolError) ParserRule() antlr.ParserRuleContext {
	return u.Assign
}

func NewUnassignableSymbolError(assign *parser.StmtAssignContext, symType *types.Type) Error {
	return &UnassignableSymbolError{
		Assign: assign,
		Type:   symType,
	}
}

type ExpectedStructTypeError struct {
	sym  sym_table.Symbol
	expr *parser.ExprStructContext
}

func (e *ExpectedStructTypeError) Error() string {
	return fmt.Sprintf("type '%s' is not a struct", e.sym.Type().ToString())
}

func (e *ExpectedStructTypeError) IsTypeError() {}

func (e *ExpectedStructTypeError) ParserRule() antlr.ParserRuleContext {
	return e.expr
}

func NewExpectedStructTypeError(sym sym_table.Symbol, expr *parser.ExprStructContext) Error {
	return &ExpectedStructTypeError{
		sym:  sym,
		expr: expr,
	}
}

type UnexpectedStructFieldError struct {
	ident parser.IIdentContext
	sym   *sym_table.StructSymbol
}

func (e *UnexpectedStructFieldError) Error() string {
	return fmt.Sprintf("unexpected field '%s' in struct '%s'", e.ident.GetText(), e.sym.SymbolName())
}

func (e *UnexpectedStructFieldError) IsTypeError() {}

func (e *UnexpectedStructFieldError) ParserRule() antlr.ParserRuleContext {
	return e.ident
}

func NewUnexpectedStructFieldError(ident parser.IIdentContext, sym *sym_table.StructSymbol) Error {
	return &UnexpectedStructFieldError{
		ident: ident,
		sym:   sym,
	}
}

type MissingStructFieldError struct {
	expr  *parser.ExprStructContext
	field *sym_table.StructFieldSymbol
}

func (e *MissingStructFieldError) Error() string {
	return fmt.Sprintf("missing field '%s' in struct '%s'", e.field.SymbolName(), e.field.Struct().SymbolName())
}

func (e *MissingStructFieldError) IsTypeError() {}

func (e *MissingStructFieldError) ParserRule() antlr.ParserRuleContext {
	return e.expr
}

func NewMissingStructFieldError(expr *parser.ExprStructContext, field *sym_table.StructFieldSymbol) Error {
	return &MissingStructFieldError{
		expr:  expr,
		field: field,
	}
}

type StructWrongRoleCountError struct {
	roleType parser.IRoleTypeContext
	sym      *sym_table.StructSymbol
}

func (e *StructWrongRoleCountError) Error() string {
	return fmt.Sprintf("wrong number of roles for struct '%s'", e.sym.SymbolName())
}

func (e *StructWrongRoleCountError) IsTypeError() {}

func (e *StructWrongRoleCountError) ParserRule() antlr.ParserRuleContext {
	return e.roleType
}

func NewStructWrongRoleCountError(roleType parser.IRoleTypeContext, sym *sym_table.StructSymbol) Error {
	return &StructWrongRoleCountError{
		roleType: roleType,
		sym:      sym,
	}
}
