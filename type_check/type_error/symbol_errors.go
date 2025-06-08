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
	Type   types.Value
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

func NewUnassignableSymbol(assign *parser.StmtAssignContext, symType types.Value) Error {
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
	return fmt.Sprintf("type `%s` is not a struct", e.sym.Type().ToString())
}

func (e *ExpectedStructType) ParserRule() antlr.ParserRuleContext {
	return e.expr
}

func (e *ExpectedStructType) Code() ErrorCode {
	return CodeExpectedStructType
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
	return fmt.Sprintf("unexpected field `%s` in struct `%s`", e.ident.GetText(), e.sym.SymbolName())
}

func (e *UnexpectedStructField) ParserRule() antlr.ParserRuleContext {
	return e.ident
}

func (e *UnexpectedStructField) Code() ErrorCode {
	return CodeUnexpectedStructField
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
	return fmt.Sprintf("missing field `%s` in struct `%s`", e.field.SymbolName(), e.field.Struct().SymbolName())
}

func (e *MissingStructField) ParserRule() antlr.ParserRuleContext {
	return e.expr
}

func (e *MissingStructField) Code() ErrorCode {
	return CodeMissingStructField
}

func NewMissingStructField(expr *parser.ExprStructContext, field *sym_table.StructFieldSymbol) Error {
	return &MissingStructField{
		expr:  expr,
		field: field,
	}
}

type StructWrongRoleCount struct {
	baseError
	sym        sym_table.Symbol
	roleType   parser.IRoleTypeContext
	parsedRole *types.Roles
}

func (e *StructWrongRoleCount) Error() string {
	return fmt.Sprintf("wrong number of roles in `%s`", e.sym.SymbolName())
}

func (e *StructWrongRoleCount) ParserRule() antlr.ParserRuleContext {
	return e.roleType
}

func (e *StructWrongRoleCount) Code() ErrorCode {
	return CodeStructWrongRoleCount
}

func NewWrongRoleCount(sym sym_table.Symbol, roleType parser.IRoleTypeContext, parsedRole *types.Roles) Error {
	return &StructWrongRoleCount{
		sym:        sym,
		roleType:   roleType,
		parsedRole: parsedRole,
	}
}

func (e *StructWrongRoleCount) Annotations() []Annotation {
	expectedRoles := amount(len(e.sym.Type().Roles().Participants()), "role", "roles")
	actualRoles := amount(len(e.parsedRole.Participants()), "role", "roles")
	wasWere := "were"
	if len(e.parsedRole.Participants()) == 1 {
		wasWere = "was"
	}

	return []Annotation{{
		Type:    AnnotationTypeHint,
		Message: fmt.Sprintf("type `%s` expects %s, but %s %s found", e.sym.Type().ToString(), expectedRoles, actualRoles, wasWere),
	}}
}

func (e *StructWrongRoleCount) RelatedInfo() []RelatedInfo {
	return []RelatedInfo{{
		Message:    "type declared here",
		ParserRule: e.sym.Ident(),
	}}
}

type FieldAccessUnknownField struct {
	baseError
	expr *parser.ExprFieldAccessContext
	sym  sym_table.Symbol
}

func (e *FieldAccessUnknownField) Error() string {
	return fmt.Sprintf("unknown field `%s` in `%s`", e.expr.Ident().GetText(), e.sym.SymbolName())
}

func (e *FieldAccessUnknownField) ParserRule() antlr.ParserRuleContext {
	return e.expr.Ident()
}

func (e *FieldAccessUnknownField) Code() ErrorCode {
	return CodeFieldAccessUnknownField
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
	objectType  types.Value
}

func (e *FieldAccessUnexpectedType) Error() string {
	return fmt.Sprintf("cannot access field of type `%s`", e.objectType.ToString())
}

func (e *FieldAccessUnexpectedType) IsTypeError() {}

func (e *FieldAccessUnexpectedType) ParserRule() antlr.ParserRuleContext {
	return e.fieldAccess.Ident()
}

func (e *FieldAccessUnexpectedType) Code() ErrorCode {
	return CodeFieldAccessUnexpectedType
}

func NewFieldAccessUnexpectedType(expr *parser.ExprFieldAccessContext, objectType types.Value) Error {
	return &FieldAccessUnexpectedType{
		fieldAccess: expr,
		objectType:  objectType,
	}
}
