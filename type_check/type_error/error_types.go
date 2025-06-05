package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type AnnotationType string

const (
	AnnotationTypeHint AnnotationType = "hint"
	AnnotationTypeNote AnnotationType = "note"
)

type Annotation struct {
	Type    AnnotationType
	Message string
}

type RelatedInfo struct {
	Message    string
	ParserRule antlr.ParserRuleContext
}

type Error interface {
	error
	ParserRule() antlr.ParserRuleContext
	IsTypeError()
	RelatedInfo() []RelatedInfo
	Annotations() []Annotation
}

type baseError struct{}

func (*baseError) IsTypeError() {}
func (*baseError) RelatedInfo() []RelatedInfo {
	return nil
}
func (*baseError) Annotations() []Annotation {
	return nil
}

type ValueRoleNotInScope struct {
	baseError
	Value             antlr.ParserRuleContext
	ValueRoles        *types.Roles
	InaccessibleRoles []string
}

func (v *ValueRoleNotInScope) Error() string {
	return fmt.Sprintf("value '%s' contains roles '%s' that are not in scope", v.Value.GetText(), misc.JoinStrings(v.InaccessibleRoles, ","))
}

func (v *ValueRoleNotInScope) ParserRule() antlr.ParserRuleContext {
	return v.Value
}

func NewValueRoleNotInScope(value antlr.ParserRuleContext, valueRoles *types.Roles, inaccessibleRoles []string) Error {
	return &ValueRoleNotInScope{
		Value:             value,
		ValueRoles:        valueRoles,
		InaccessibleRoles: inaccessibleRoles,
	}
}

type UnexpectedSharedType struct {
	baseError
	RoleType parser.IRoleTypeContext
}

func (u *UnexpectedSharedType) Error() string {
	return "shared type not allowed here"
}

func (u *UnexpectedSharedType) ParserRule() antlr.ParserRuleContext {
	return u.RoleType
}

func NewUnexpectedSharedType(roleType parser.IRoleTypeContext) Error {
	return &UnexpectedSharedType{
		RoleType: roleType,
	}
}

type UnknownType struct {
	baseError
	TypeName parser.IIdentContext
}

func NewUnknownType(typeName parser.IIdentContext) Error {
	return &UnknownType{
		TypeName: typeName,
	}
}

func (e *UnknownType) Error() string {
	return fmt.Sprintf("unknown type '%s'", e.TypeName.GetText())
}

func (e *UnknownType) ParserRule() antlr.ParserRuleContext {
	return e.TypeName
}

type ValueMismatch struct {
	baseError
	Expr        parser.IExprContext
	FirstValue  types.Value
	SecondValue types.Value
}

func (t *ValueMismatch) Error() string {
	return fmt.Sprintf("type values '%s' and '%s' do not match", t.FirstValue.ToString(), t.SecondValue.ToString())
}

func (t *ValueMismatch) ParserRule() antlr.ParserRuleContext {
	return t.Expr
}

func NewValueMismatch(expr parser.IExprContext, firstValue types.Value, secondValue types.Value) Error {
	return &ValueMismatch{
		Expr:        expr,
		FirstValue:  firstValue,
		SecondValue: secondValue,
	}
}

type IncompatibleTypes struct {
	baseError
	Expr         parser.IExprContext
	ExprType     *types.Type
	ExpectedType *types.Type
}

func (e *IncompatibleTypes) Error() string {
	return fmt.Sprintf("type '%s' is not compatible with type '%s'", e.ExprType.ToString(), e.ExpectedType.ToString())
}

func (e *IncompatibleTypes) ParserRule() antlr.ParserRuleContext {
	return e.Expr
}

func NewIncompatibleTypes(expr parser.IExprContext, exprType *types.Type, expectedType *types.Type) Error {
	return &IncompatibleTypes{
		Expr:         expr,
		ExprType:     exprType,
		ExpectedType: expectedType,
	}
}

type InvalidValue struct {
	baseError
	Expr          parser.IExprContext
	ActualValue   types.Value
	ExpectedValue types.Value
}

func (i *InvalidValue) Error() string {
	return fmt.Sprintf("invalid value, expected '%s' but found '%s'", i.ExpectedValue.ToString(), i.ActualValue.ToString())
}

func (i *InvalidValue) ParserRule() antlr.ParserRuleContext {
	return i.Expr
}

func NewInvalidValue(expr parser.IExprContext, actualValue types.Value, expectedValue types.Value) Error {
	return &InvalidValue{
		Expr:          expr,
		ActualValue:   actualValue,
		ExpectedValue: expectedValue,
	}
}

type ExpectedAsyncType struct {
	baseError
	Expr parser.IExprContext
	Type *types.Type
}

func (e *ExpectedAsyncType) Error() string {
	return fmt.Sprintf("expected async type '%s'", e.Type.ToString())
}

func (e *ExpectedAsyncType) ParserRule() antlr.ParserRuleContext {
	return e.Expr
}

func NewExpectedAsyncType(expr parser.IExprContext, errType *types.Type) Error {
	return &ExpectedAsyncType{
		Expr: expr,
		Type: errType,
	}
}

type BinOpIncompatibleType struct {
	baseError
	BinOp   *parser.ExprBinOpContext
	Value   types.Value
	Allowed []types.Value
}

func (e *BinOpIncompatibleType) Error() string {
	allowed := make([]string, len(e.Allowed))
	for i, v := range e.Allowed {
		allowed[i] = v.ToString()
	}

	op := projection.ParseOperator(e.BinOp)

	return fmt.Sprintf("value '%s' not allowed for operation '%s', allowed types '%s'", e.Value.ToString(), op, misc.JoinStrings(allowed, ", "))
}

func (e *BinOpIncompatibleType) ParserRule() antlr.ParserRuleContext {
	return e.BinOp
}

func NewBinOpIncompatibleType(binOp *parser.ExprBinOpContext, value types.Value, allowed []types.Value) Error {
	return &BinOpIncompatibleType{
		BinOp:   binOp,
		Value:   value,
		Allowed: allowed,
	}
}

type UnsendableType struct {
	baseError
	Com            *parser.ExprComContext
	UnsendableType *types.Type
}

func (u *UnsendableType) Error() string {
	return fmt.Sprintf("can not send type '%s'", u.UnsendableType.ToString())
}

func (u *UnsendableType) ParserRule() antlr.ParserRuleContext {
	return u.Com.Expr()
}

func NewUnsendableType(com *parser.ExprComContext, unsendableType *types.Type) Error {
	return &UnsendableType{
		Com:            com,
		UnsendableType: unsendableType,
	}
}

type NotDistributedType struct {
	baseError
	typeCtx antlr.ParserRuleContext
}

func NewNotDistributedType(typeCtx antlr.ParserRuleContext) Error {
	return &NotDistributedType{
		typeCtx: typeCtx,
	}
}

func (e *NotDistributedType) Error() string {
	return fmt.Sprintf("type '%s' is not distributed", e.typeCtx.GetText())
}

func (e *NotDistributedType) ParserRule() antlr.ParserRuleContext {
	return e.typeCtx
}

func (e *NotDistributedType) Annotations() []Annotation {
	return []Annotation{{
		Type:    AnnotationTypeHint,
		Message: "change the type a single role, or make it a shared type instead.",
	}}
}
