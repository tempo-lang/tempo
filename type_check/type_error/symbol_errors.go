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
	return fmt.Sprintf("symbol '%s' already declared", s.NewSymbol.GetText())
}

func (e *SymbolAlreadyExists) ParserRule() antlr.ParserRuleContext {
	return e.NewSymbol
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
	return fmt.Sprintf("unknown symbol '%s'", e.SymName.GetText())
}

func (e *UnknownSymbol) ParserRule() antlr.ParserRuleContext {
	return e.SymName
}

type UnassignableSymbol struct {
	baseError
	Assign *parser.StmtAssignContext
	Type   *types.Type
}

func (u *UnassignableSymbol) Error() string {
	return fmt.Sprintf("can not assign value to '%s' of type '%s'", u.Assign.Ident().GetText(), u.Type.ToString())
}

func (u *UnassignableSymbol) ParserRule() antlr.ParserRuleContext {
	return u.Assign
}

func NewUnassignableSymbol(assign *parser.StmtAssignContext, symType *types.Type) Error {
	return &UnassignableSymbol{
		Assign: assign,
		Type:   symType,
	}
}

type ExpectedStructType struct {
	baseError
	sym  sym_table.Symbol
	expr *parser.ExprStructContext
}

func (e *ExpectedStructType) Error() string {
	return fmt.Sprintf("type '%s' is not a struct", e.sym.Type().ToString())
}

func (e *ExpectedStructType) ParserRule() antlr.ParserRuleContext {
	return e.expr
}

func NewExpectedStructType(sym sym_table.Symbol, expr *parser.ExprStructContext) Error {
	return &ExpectedStructType{
		sym:  sym,
		expr: expr,
	}
}

type UnexpectedStructField struct {
	baseError
	ident parser.IIdentContext
	sym   *sym_table.StructSymbol
}

func (e *UnexpectedStructField) Error() string {
	return fmt.Sprintf("unexpected field '%s' in struct '%s'", e.ident.GetText(), e.sym.SymbolName())
}

func (e *UnexpectedStructField) ParserRule() antlr.ParserRuleContext {
	return e.ident
}

func NewUnexpectedStructField(ident parser.IIdentContext, sym *sym_table.StructSymbol) Error {
	return &UnexpectedStructField{
		ident: ident,
		sym:   sym,
	}
}

type MissingStructField struct {
	baseError
	expr  *parser.ExprStructContext
	field *sym_table.StructFieldSymbol
}

func (e *MissingStructField) Error() string {
	return fmt.Sprintf("missing field '%s' in struct '%s'", e.field.SymbolName(), e.field.Struct().SymbolName())
}

func (e *MissingStructField) ParserRule() antlr.ParserRuleContext {
	return e.expr
}

func NewMissingStructField(expr *parser.ExprStructContext, field *sym_table.StructFieldSymbol) Error {
	return &MissingStructField{
		expr:  expr,
		field: field,
	}
}

type StructWrongRoleCount struct {
	baseError
	roleType parser.IRoleTypeContext
	sym      *sym_table.StructSymbol
}

func (e *StructWrongRoleCount) Error() string {
	return fmt.Sprintf("wrong number of roles for struct '%s'", e.sym.SymbolName())
}

func (e *StructWrongRoleCount) ParserRule() antlr.ParserRuleContext {
	return e.roleType
}

func NewStructWrongRoleCount(roleType parser.IRoleTypeContext, sym *sym_table.StructSymbol) Error {
	return &StructWrongRoleCount{
		roleType: roleType,
		sym:      sym,
	}
}

type FieldAccessUnknownField struct {
	baseError
	expr *parser.ExprFieldAccessContext
	sym  sym_table.Symbol
}

func (e *FieldAccessUnknownField) Error() string {
	return fmt.Sprintf("unknown field '%s' in '%s'", e.expr.Ident().GetText(), e.sym.SymbolName())
}

func (e *FieldAccessUnknownField) ParserRule() antlr.ParserRuleContext {
	return e.expr.Ident()
}

func NewFieldAccessUnknownField(expr *parser.ExprFieldAccessContext, sym sym_table.Symbol) Error {
	return &FieldAccessUnknownField{
		expr: expr,
		sym:  sym,
	}
}

type FieldAccessUnexpectedType struct {
	baseError
	fieldAccess *parser.ExprFieldAccessContext
	objectType  *types.Type
}

func (e *FieldAccessUnexpectedType) Error() string {
	return fmt.Sprintf("cannot access field of type '%s'", e.objectType.ToString())
}

func (e *FieldAccessUnexpectedType) IsTypeError() {}

func (e *FieldAccessUnexpectedType) ParserRule() antlr.ParserRuleContext {
	return e.fieldAccess.Ident()
}

func NewFieldAccessUnexpectedType(expr *parser.ExprFieldAccessContext, objectType *types.Type) Error {
	return &FieldAccessUnexpectedType{
		fieldAccess: expr,
		objectType:  objectType,
	}
}
