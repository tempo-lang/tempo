package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type Error interface {
	error
	ParserRule() antlr.ParserRuleContext
	IsTypeError()
}

type ValueRoleNotInScopeError struct {
	Value             antlr.ParserRuleContext
	ValueRoles        *types.Roles
	InaccessibleRoles []string
}

func (v *ValueRoleNotInScopeError) Error() string {
	return fmt.Sprintf("value '%s' contains roles '%s' that are not in scope", v.Value.GetText(), misc.JoinStrings(v.InaccessibleRoles, ","))
}

func (v *ValueRoleNotInScopeError) IsTypeError() {}

func (v *ValueRoleNotInScopeError) ParserRule() antlr.ParserRuleContext {
	return v.Value
}

func NewValueRoleNotInScopeError(value antlr.ParserRuleContext, valueRoles *types.Roles, inaccessibleRoles []string) Error {
	return &ValueRoleNotInScopeError{
		Value:             value,
		ValueRoles:        valueRoles,
		InaccessibleRoles: inaccessibleRoles,
	}
}

type UnexpectedSharedType struct {
	RoleType parser.IRoleTypeContext
}

func (u *UnexpectedSharedType) Error() string {
	return "shared type not allowed here"
}

func (u *UnexpectedSharedType) IsTypeError() {}

func (u *UnexpectedSharedType) ParserRule() antlr.ParserRuleContext {
	return u.RoleType
}

func NewUnexpectedSharedTypeError(roleType parser.IRoleTypeContext) Error {
	return &UnexpectedSharedType{
		RoleType: roleType,
	}
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

type ValueMismatchError struct {
	Expr        parser.IExprContext
	FirstValue  types.Value
	SecondValue types.Value
}

func (t *ValueMismatchError) Error() string {
	return fmt.Sprintf("type values '%s' and '%s' do not match", t.FirstValue.ToString(), t.SecondValue.ToString())
}

func (t *ValueMismatchError) IsTypeError() {}

func (t *ValueMismatchError) ParserRule() antlr.ParserRuleContext {
	return t.Expr
}

func NewValueMismatchError(expr parser.IExprContext, firstValue types.Value, secondValue types.Value) Error {
	return &ValueMismatchError{
		Expr:        expr,
		FirstValue:  firstValue,
		SecondValue: secondValue,
	}
}

type IncompatibleTypesError struct {
	Expr         parser.IExprContext
	ExprType     *types.Type
	ExpectedType *types.Type
}

func (e *IncompatibleTypesError) Error() string {
	return fmt.Sprintf("type '%s' is not compatible with type '%s'", e.ExprType.ToString(), e.ExpectedType.ToString())
}

func (e *IncompatibleTypesError) IsTypeError() {}

func (e *IncompatibleTypesError) ParserRule() antlr.ParserRuleContext {
	return e.Expr
}

func NewIncompatibleTypesError(expr parser.IExprContext, exprType *types.Type, expectedType *types.Type) Error {
	return &IncompatibleTypesError{
		Expr:         expr,
		ExprType:     exprType,
		ExpectedType: expectedType,
	}
}

type InvalidValueError struct {
	Expr          parser.IExprContext
	ActualValue   types.Value
	ExpectedValue types.Value
}

func (i *InvalidValueError) Error() string {
	return fmt.Sprintf("invalid value, expected '%s' but found '%s'", i.ExpectedValue.ToString(), i.ActualValue.ToString())
}

func (i *InvalidValueError) IsTypeError() {}

func (i *InvalidValueError) ParserRule() antlr.ParserRuleContext {
	return i.Expr
}

func NewInvalidValueError(expr parser.IExprContext, actualValue types.Value, expectedValue types.Value) Error {
	return &InvalidValueError{
		Expr:          expr,
		ActualValue:   actualValue,
		ExpectedValue: expectedValue,
	}
}

type ExpectedAsyncTypeError struct {
	Expr parser.IExprContext
	Type *types.Type
}

func (e *ExpectedAsyncTypeError) Error() string {
	return fmt.Sprintf("expected async type '%s'", e.Type.ToString())
}

func (e *ExpectedAsyncTypeError) IsTypeError() {}

func (e *ExpectedAsyncTypeError) ParserRule() antlr.ParserRuleContext {
	return e.Expr
}

func NewExpectedAsyncTypeError(expr parser.IExprContext, exprType *types.Type) Error {
	return &ExpectedAsyncTypeError{
		Expr: expr,
		Type: exprType,
	}
}

type UnsendableTypeError struct {
	Com            *parser.ExprComContext
	UnsendableType *types.Type
}

func (u *UnsendableTypeError) Error() string {
	return fmt.Sprintf("can not send type '%s'", u.UnsendableType.ToString())
}

func (u *UnsendableTypeError) IsTypeError() {}

func (u *UnsendableTypeError) ParserRule() antlr.ParserRuleContext {
	return u.Com.Expr()
}

func NewUnsendableTypeError(com *parser.ExprComContext, unsendableType *types.Type) Error {
	return &UnsendableTypeError{
		Com:            com,
		UnsendableType: unsendableType,
	}
}

type NotDistributedTypeError struct {
	typeCtx antlr.ParserRuleContext
}

func NewNotDistributedTypeError(typeCtx antlr.ParserRuleContext) Error {
	return &NotDistributedTypeError{
		typeCtx: typeCtx,
	}
}

func (e *NotDistributedTypeError) Error() string {
	return fmt.Sprintf("type '%s' is not distributed", e.typeCtx.GetText())
}

func (e *NotDistributedTypeError) IsTypeError() {}

func (e *NotDistributedTypeError) ParserRule() antlr.ParserRuleContext {
	return e.typeCtx
}
