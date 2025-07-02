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
	ident      parser.IIdentContext
	structType *types.StructType
}

func (e *UnexpectedStructField) Error() string {
	return fmt.Sprintf("unexpected field `%s` in struct `%s`", e.ident.GetText(), e.structType.Name())
}

func (e *UnexpectedStructField) ParserRule() antlr.ParserRuleContext {
	return e.ident
}

func (e *UnexpectedStructField) Code() ErrorCode {
	return CodeUnexpectedStructField
}

func NewUnexpectedStructField(ident parser.IIdentContext, structType *types.StructType) Error {
	return &UnexpectedStructField{
		ident:      ident,
		structType: structType,
	}
}

type MissingStructField struct {
	baseError
	expr       *parser.ExprStructContext
	field      string
	structType *types.StructType
}

func (e *MissingStructField) Error() string {
	return fmt.Sprintf("missing field `%s` in struct `%s`", e.field, e.structType.Name())
}

func (e *MissingStructField) ParserRule() antlr.ParserRuleContext {
	return e.expr
}

func (e *MissingStructField) Code() ErrorCode {
	return CodeMissingStructField
}

func NewMissingStructField(expr *parser.ExprStructContext, field string, structType *types.StructType) Error {
	return &MissingStructField{
		expr:       expr,
		field:      field,
		structType: structType,
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
