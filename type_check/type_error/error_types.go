package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type ErrorCode int

const (
	CodeValueRoleNotInScope ErrorCode = iota + 1
	CodeUnexpectedSharedType
	CodeUnknownType
	CodeValueMismatch
	CodeIncompatibleTypes
	CodeInvalidValue
	CodeExpectedAsyncType
	CodeBinOpIncompatibleType
	CodeUnsendableType
	CodeNotDistributedType
	CodeDuplicateRoles
	CodeRolesNotInScope
	CodeUnmergableRoles
	CodeSharedRoleSingleParticipant
	CodeSymbolAlreadyExists
	CodeUnknownSymbol
	CodeUnassignableSymbol
	CodeExpectedStructType
	CodeUnexpectedStructField
	CodeMissingStructField
	CodeStructWrongRoleCount
	CodeFieldAccessUnknownField
	CodeFieldAccessUnexpectedType
	CodeInvalidAssignType
	CodeReturnNotAllRoles
	CodeInvalidNumber
	CodeComNonLocalSender
	CodeComValueNotAtSender
	CodeUnequatableType
	CodeStructNotInitialized
	CodeCallNonFunction
	CodeCallWrongArgCount
	CodeInstantiateNonFunction
	CodeFunctionNotInstantiated
	CodeFunctionMissingReturn
	CodeReturnValueMissing
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
	Code() ErrorCode
}

type baseError struct{}

func (*baseError) IsTypeError() {}
func (*baseError) RelatedInfo() []RelatedInfo {
	return nil
}
func (*baseError) Annotations() []Annotation {
	return nil
}

func formatList(singular, plural string, items []string, combine string) string {
	if len(items) == 1 {
		if singular != "" {
			singular += " "
		}

		return fmt.Sprintf("%s`%s`", singular, items[0])
	}

	if plural != "" {
		plural += " "
	}

	formattedItems := misc.JoinStringsFunc(items[:len(items)-1], ", ", func(item string) string {
		return fmt.Sprintf("`%s`", item)
	})

	return fmt.Sprintf("%s%s %s `%s`", plural, formattedItems, combine, items[len(items)-1])
}

func toBe[T any](items []T) string {
	if len(items) == 1 {
		return "is"
	} else {
		return "are"
	}
}

type ValueRoleNotInScope struct {
	baseError
	Value             antlr.ParserRuleContext
	ValueRoles        *types.Roles
	InaccessibleRoles []string
}

func (v *ValueRoleNotInScope) Error() string {
	roles := formatList("role", "roles", v.InaccessibleRoles, "and")
	return fmt.Sprintf("value `%s` contains %s that %s not in scope", v.Value.GetText(), roles, toBe(v.InaccessibleRoles))
}

func (v *ValueRoleNotInScope) ParserRule() antlr.ParserRuleContext {
	return v.Value
}

func (e *ValueRoleNotInScope) Code() ErrorCode {
	return CodeValueRoleNotInScope
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
	return "shared type is not allowed here"
}

func (u *UnexpectedSharedType) ParserRule() antlr.ParserRuleContext {
	return u.RoleType
}

func (e *UnexpectedSharedType) Code() ErrorCode {
	return CodeUnexpectedSharedType
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
	return fmt.Sprintf("type name `%s` is undefined or not in scope", e.TypeName.GetText())
}

func (e *UnknownType) ParserRule() antlr.ParserRuleContext {
	return e.TypeName
}

func (e *UnknownType) Code() ErrorCode {
	return CodeUnknownType
}

type ValueMismatch struct {
	baseError
	Expr        parser.IExprContext
	FirstValue  types.Value
	SecondValue types.Value
}

func (t *ValueMismatch) Error() string {
	return fmt.Sprintf("types `%s` and `%s` are not compatible", t.FirstValue.ToString(), t.SecondValue.ToString())
}

func (t *ValueMismatch) ParserRule() antlr.ParserRuleContext {
	return t.Expr
}

func (e *ValueMismatch) Code() ErrorCode {
	return CodeValueMismatch
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
	return fmt.Sprintf("type `%s` is not compatible with type `%s`", e.ExprType.ToString(), e.ExpectedType.ToString())
}

func (e *IncompatibleTypes) ParserRule() antlr.ParserRuleContext {
	return e.Expr
}

func (e *IncompatibleTypes) Code() ErrorCode {
	return CodeIncompatibleTypes
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
	return fmt.Sprintf("type `%s` does not match expected type `%s`", i.ActualValue.ToString(), i.ExpectedValue.ToString())
}

func (i *InvalidValue) ParserRule() antlr.ParserRuleContext {
	return i.Expr
}

func (e *InvalidValue) Code() ErrorCode {
	return CodeInvalidValue
}

func NewInvalidValue(expr parser.IExprContext, actualValue types.Value, expectedValue types.Value) Error {
	return &InvalidValue{
		Expr:          expr,
		ActualValue:   actualValue,
		ExpectedValue: expectedValue,
	}
}

type AwaitNonAsyncType struct {
	baseError
	Expr parser.IExprContext
	Type *types.Type
}

func (e *AwaitNonAsyncType) Error() string {
	return fmt.Sprintf("cannot await type `%s` since it is not an async type", e.Type.ToString())
}

func (e *AwaitNonAsyncType) ParserRule() antlr.ParserRuleContext {
	return e.Expr
}

func (e *AwaitNonAsyncType) Code() ErrorCode {
	return CodeExpectedAsyncType
}

func NewAwaitNonAsyncType(expr parser.IExprContext, errType *types.Type) Error {
	return &AwaitNonAsyncType{
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
	op := projection.ParseOperator(e.BinOp)
	return fmt.Sprintf("operation `%s` cannot be performed on `%s` types", op, e.Value.ToString())
}

func (e *BinOpIncompatibleType) Annotations() []Annotation {
	allowed := make([]string, len(e.Allowed))
	for i, v := range e.Allowed {
		allowed[i] = fmt.Sprintf("`%s`", v.ToString())
	}

	var msg string
	if len(allowed) == 1 {
		msg = fmt.Sprintf("only type %s is allowed.", allowed[0])
	} else {
		msg = fmt.Sprintf("allowed types are %s or %s.", misc.JoinStrings(allowed[:len(allowed)-1], ", "), allowed[len(allowed)-1])
	}

	return []Annotation{{
		Type:    AnnotationTypeNote,
		Message: msg,
	}}
}

func (e *BinOpIncompatibleType) ParserRule() antlr.ParserRuleContext {
	return e.BinOp
}

func (e *BinOpIncompatibleType) Code() ErrorCode {
	return CodeBinOpIncompatibleType
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
	return fmt.Sprintf("can not send values of type `%s`", u.UnsendableType.ToString())
}

func (u *UnsendableType) ParserRule() antlr.ParserRuleContext {
	return u.Com.Expr()
}

func (e *UnsendableType) Code() ErrorCode {
	return CodeUnsendableType
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
	return fmt.Sprintf("type `%s` cannot be distributed", e.typeCtx.GetText())
}

func (e *NotDistributedType) ParserRule() antlr.ParserRuleContext {
	return e.typeCtx
}

func (e *NotDistributedType) Code() ErrorCode {
	return CodeNotDistributedType
}

func (e *NotDistributedType) Annotations() []Annotation {
	return []Annotation{{
		Type:    AnnotationTypeHint,
		Message: "change the type to be a single or a shared role instead.",
	}}
}
