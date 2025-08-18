package type_error

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"
)

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

type MissingImplementationMethod struct {
	baseError
	structSym    *sym_table.StructSymbol
	interfaceSym *sym_table.InterfaceSymbol
	methodName   string
}

func (e *MissingImplementationMethod) Error() string {
	return fmt.Sprintf("struct `%s` does not implement method `%s` of interface `%s`", e.structSym.SymbolName(), e.methodName, e.interfaceSym.SymbolName())
}

func (e *MissingImplementationMethod) ParserRule() antlr.ParserRuleContext {
	for _, impl := range e.structSym.StructCtx().StructImplements().AllRoleIdent() {
		if impl.Ident().GetText() == e.interfaceSym.SymbolName() {
			return impl
		}
	}
	panic("interface name should exist")
}

func (e *MissingImplementationMethod) Code() ErrorCode {
	return CodeMissingImplementationMethod
}

func (e *MissingImplementationMethod) RelatedInfo() []RelatedInfo {
	fn, found := e.interfaceSym.Method(e.methodName)
	if !found {
		panic("method should always exist")
	}

	return []RelatedInfo{{
		Message:    "missing method declared here",
		ParserRule: fn.FuncSig().Ident(),
	}}
}

func NewMissingImplementationMethod(structSym *sym_table.StructSymbol, interfaceSym *sym_table.InterfaceSymbol, methodName string) Error {
	return &MissingImplementationMethod{
		structSym:    structSym,
		interfaceSym: interfaceSym,
		methodName:   methodName,
	}
}

type IncompatibleImplementationMethod struct {
	baseError
	structSym    *sym_table.StructSymbol
	interfaceSym *sym_table.InterfaceSymbol
	methodName   string
}

func (e *IncompatibleImplementationMethod) Error() string {
	return fmt.Sprintf(
		"method `%s` does not match the signature required by interface `%s`",
		e.methodName, e.interfaceSym.SymbolName(),
	)
}

func (e *IncompatibleImplementationMethod) ParserRule() antlr.ParserRuleContext {
	fn, found := e.structSym.Method(e.methodName)
	if !found {
		panic("method should always exist")
	}

	return fn.FuncSig()
}

func (e *IncompatibleImplementationMethod) Code() ErrorCode {
	return CodeIncompatibleImplementationMethod
}

func (e *IncompatibleImplementationMethod) RelatedInfo() []RelatedInfo {
	fn, found := e.interfaceSym.Method(e.methodName)
	if !found {
		panic("method should always exist")
	}

	return []RelatedInfo{{
		Message:    "interface method declared here",
		ParserRule: fn.FuncSig().Ident(),
	}}
}

func (e *IncompatibleImplementationMethod) Annotations() []Annotation {
	fn, found := e.interfaceSym.Method(e.methodName)
	if !found {
		panic("method should always exist")
	}

	return []Annotation{{
		Type:    AnnotationTypeHint,
		Message: fmt.Sprintf("change function signature to `%s`", fn.FuncType().ToString()),
	}}
}

func NewIncompatibleImplementationMethod(structSym *sym_table.StructSymbol, interfaceSym *sym_table.InterfaceSymbol, methodName string) Error {
	return &IncompatibleImplementationMethod{
		structSym:    structSym,
		interfaceSym: interfaceSym,
		methodName:   methodName,
	}
}
