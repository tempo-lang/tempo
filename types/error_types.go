package types

import (
	"fmt"
	"tempo/misc"
	"tempo/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Error interface {
	error
	ParserRule() antlr.ParserRuleContext
	IsTypeError()
}

type DuplicateRolesError struct {
	Func           parser.IFuncContext
	DuplicateRoles []parser.IIdentContext
}

func NewDuplicateRolesError(function parser.IFuncContext, duplicateRoles []parser.IIdentContext) Error {
	return &DuplicateRolesError{
		Func:           function,
		DuplicateRoles: duplicateRoles,
	}
}

func (e *DuplicateRolesError) IsTypeError() {}

func (e *DuplicateRolesError) Error() string {
	return fmt.Sprintf("function '%s' has duplicate role '%s'", e.Func.Ident().GetText(), e.DuplicateRoles[0].GetText())
}

func (e *DuplicateRolesError) ParserRule() antlr.ParserRuleContext {
	return e.Func
}

type RolesNotInScopeError struct {
	RoleType     parser.IRoleTypeContext
	UnknownRoles []string
}

func NewRolesNotInScopeError(roleType parser.IRoleTypeContext, unknownRoles []string) Error {
	return &RolesNotInScopeError{
		RoleType:     roleType,
		UnknownRoles: unknownRoles,
	}
}

func (e *RolesNotInScopeError) IsTypeError() {}

func (e *RolesNotInScopeError) Error() string {
	return fmt.Sprintf("roles '%s' are not in scope", misc.JoinStrings(e.UnknownRoles, ","))
}

func (e *RolesNotInScopeError) ParserRule() antlr.ParserRuleContext {
	return e.RoleType
}

type ValueRoleNotInScopeError struct {
	Value             antlr.ParserRuleContext
	ValueRoles        *Roles
	InaccessibleRoles []string
}

func (v *ValueRoleNotInScopeError) Error() string {
	return fmt.Sprintf("value '%s' has roles '%s' that are not in scope", v.Value.GetText(), misc.JoinStrings(v.InaccessibleRoles, ","))
}

func (v *ValueRoleNotInScopeError) IsTypeError() {}

func (v *ValueRoleNotInScopeError) ParserRule() antlr.ParserRuleContext {
	return v.Value
}

func NewValueRoleNotInScopeError(value antlr.ParserRuleContext, valueRoles *Roles, inaccessibleRoles []string) Error {
	return &ValueRoleNotInScopeError{
		Value:             value,
		ValueRoles:        valueRoles,
		InaccessibleRoles: inaccessibleRoles,
	}
}

type ReturnNotAllRolesError struct {
	Return       *parser.StmtReturnContext
	MissignRoles []string
}

func NewReturnNotAllRolesError(ret *parser.StmtReturnContext, missingRoles []string) Error {
	return &ReturnNotAllRolesError{
		Return:       ret,
		MissignRoles: missingRoles,
	}
}

func (e *ReturnNotAllRolesError) Error() string {
	return fmt.Sprintf("return statement is missing roles: %s", misc.JoinStrings(e.MissignRoles, ","))
}

func (e *ReturnNotAllRolesError) IsTypeError() {}

func (e *ReturnNotAllRolesError) ParserRule() antlr.ParserRuleContext {
	return e.Return.Expr()
}

type UnmergableRolesError struct {
	Expr  parser.IExprContext
	Roles []*Roles
}

func (u *UnmergableRolesError) Error() string {
	roles := misc.JoinStringsFunc(u.Roles, ", ", func(role *Roles) string { return role.ToString() })
	return fmt.Sprintf("can not merge roles %s", roles)
}

func (u *UnmergableRolesError) IsTypeError() {}

func (u *UnmergableRolesError) ParserRule() antlr.ParserRuleContext {
	return u.Expr
}

func NewUnmergableRolesError(expr parser.IExprContext, roles []*Roles) Error {
	return &UnmergableRolesError{
		Expr:  expr,
		Roles: roles,
	}
}

type UnexpectedSharedType struct {
	Ident parser.IIdentContext
	Type  *Type
}

func (u *UnexpectedSharedType) Error() string {
	return fmt.Sprintf("symbol '%s' is not allowed to have shared type", u.Ident.GetText())
}

func (u *UnexpectedSharedType) IsTypeError() {}

func (u *UnexpectedSharedType) ParserRule() antlr.ParserRuleContext {
	return u.Ident
}

func NewUnexpectedSharedTypeError(ident parser.IIdentContext, identType *Type) Error {
	return &UnexpectedSharedType{
		Ident: ident,
		Type:  identType,
	}
}

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

type InvalidDeclTypeError struct {
	DeclToken parser.IValueTypeContext
	DeclType  *Type
	ExprToken parser.IExprContext
	ExprType  *Type
}

func NewInvalidDeclTypeError(declToken parser.IValueTypeContext, declType *Type, exprToken parser.IExprContext, exprType *Type) Error {
	return &InvalidDeclTypeError{
		DeclToken: declToken,
		DeclType:  declType,
		ExprToken: exprToken,
		ExprType:  exprType,
	}
}

func (e *InvalidDeclTypeError) Error() string {
	return fmt.Sprintf("invalid declaration type, expected %s found %s", e.DeclType.ToString(), e.ExprType.ToString())
}

func (e *InvalidDeclTypeError) IsTypeError() {}

func (e *InvalidDeclTypeError) ParserRule() antlr.ParserRuleContext {
	return e.ExprToken
}

type InvalidAssignTypeError struct {
	Assign   *parser.StmtAssignContext
	VarType  *Type
	ExprType *Type
}

func (i *InvalidAssignTypeError) Error() string {
	return fmt.Sprintf("invalid assignment type, expected %s found %s", i.VarType.ToString(), i.ExprType.ToString())
}

func (i *InvalidAssignTypeError) IsTypeError() {}

func (i *InvalidAssignTypeError) ParserRule() antlr.ParserRuleContext {
	return i.Assign.Expr()
}

func NewInvalidAssignTypeError(assign *parser.StmtAssignContext, varType *Type, exprType *Type) Error {
	return &InvalidAssignTypeError{
		Assign:   assign,
		VarType:  varType,
		ExprType: exprType,
	}
}

type ValueMismatchError struct {
	Expr        parser.IExprContext
	FirstValue  Value
	SecondValue Value
}

func (t *ValueMismatchError) Error() string {
	return fmt.Sprintf("type values '%s' and '%s' do not match", t.FirstValue.ToString(), t.SecondValue.ToString())
}

func (t *ValueMismatchError) IsTypeError() {}

func (t *ValueMismatchError) ParserRule() antlr.ParserRuleContext {
	return t.Expr
}

func NewValueMismatchError(expr parser.IExprContext, firstValue Value, secondValue Value) Error {
	return &ValueMismatchError{
		Expr:        expr,
		FirstValue:  firstValue,
		SecondValue: secondValue,
	}
}

type IncompatibleTypesError struct {
	Expr         parser.IExprContext
	ExprType     *Type
	ExpectedType *Type
}

func (e *IncompatibleTypesError) Error() string {
	return fmt.Sprintf("type '%s' is not compatible with type '%s'", e.ExprType.ToString(), e.ExpectedType.ToString())
}

func (e *IncompatibleTypesError) IsTypeError() {}

func (e *IncompatibleTypesError) ParserRule() antlr.ParserRuleContext {
	return e.Expr
}

func NewIncompatibleTypesError(expr parser.IExprContext, exprType *Type, expectedType *Type) Error {
	return &IncompatibleTypesError{
		Expr:         expr,
		ExprType:     exprType,
		ExpectedType: expectedType,
	}
}

type InvalidValueError struct {
	Expr          parser.IExprContext
	ActualValue   Value
	ExpectedValue Value
}

func (i *InvalidValueError) Error() string {
	return fmt.Sprintf("invalid value, expected '%s' but found '%s'", i.ExpectedValue.ToString(), i.ActualValue.ToString())
}

func (i *InvalidValueError) IsTypeError() {}

func (i *InvalidValueError) ParserRule() antlr.ParserRuleContext {
	return i.Expr
}

func NewInvalidValueError(expr parser.IExprContext, actualValue Value, expectedValue Value) Error {
	return &InvalidValueError{
		Expr:          expr,
		ActualValue:   actualValue,
		ExpectedValue: expectedValue,
	}
}

type InvalidFunctionError struct {
	Func        parser.IFuncContext
	ParamErrors [][]Error
}

func NewInvalidFuncError(fn parser.IFuncContext, params [][]Error) Error {
	return &InvalidFunctionError{
		Func:        fn,
		ParamErrors: params,
	}
}

func (e *InvalidFunctionError) Error() string {
	paramErrors := ""
	for i := 0; i < len(e.ParamErrors); i++ {
		if len(e.ParamErrors[i]) > 0 {
			paramErrors += fmt.Sprintf("param %d: %v, ", i+1, e.ParamErrors[i])
		}
	}
	paramErrors = paramErrors[:len(paramErrors)-2]

	return fmt.Sprintf("invalid function type for '%s': %s", e.Func.Ident().GetText(), paramErrors)
}

func (i *InvalidFunctionError) IsTypeError() {}

func (e *InvalidFunctionError) ParserRule() antlr.ParserRuleContext {
	return e.Func
}

type InvalidNumberError struct {
	Num parser.IExprContext
}

func NewInvalidNumberError(num parser.IExprContext) Error {
	return &InvalidNumberError{
		Num: num,
	}
}

func (i *InvalidNumberError) Error() string {
	return fmt.Sprintf("invalid number '%s'", i.Num.GetText())
}

func (i *InvalidNumberError) IsTypeError() {}

func (i *InvalidNumberError) ParserRule() antlr.ParserRuleContext {
	return i.Num
}

type UnassignableSymbolError struct {
	Assign *parser.StmtAssignContext
	Type   *Type
}

func (u *UnassignableSymbolError) Error() string {
	return fmt.Sprintf("can not assign value to '%s' of type '%s'", u.Assign.Ident().GetText(), u.Type.ToString())
}

func (u *UnassignableSymbolError) IsTypeError() {}

func (u *UnassignableSymbolError) ParserRule() antlr.ParserRuleContext {
	return u.Assign
}

func NewUnassignableSymbolError(assign *parser.StmtAssignContext, symType *Type) Error {
	return &UnassignableSymbolError{
		Assign: assign,
		Type:   symType,
	}
}

type ExpectedAsyncTypeError struct {
	Expr parser.IExprContext
	Type *Type
}

func (e *ExpectedAsyncTypeError) Error() string {
	return fmt.Sprintf("expected async type '%s'", e.Type.ToString())
}

func (e *ExpectedAsyncTypeError) IsTypeError() {}

func (e *ExpectedAsyncTypeError) ParserRule() antlr.ParserRuleContext {
	return e.Expr
}

func NewExpectedAsyncTypeError(expr parser.IExprContext, exprType *Type) Error {
	return &ExpectedAsyncTypeError{
		Expr: expr,
		Type: exprType,
	}
}

type ComSharedTypeError struct {
	Com       *parser.ExprComContext
	InnerType *Type
}

func (e *ComSharedTypeError) Error() string {
	return fmt.Sprintf("can not communicate shared type '%s'", e.InnerType.ToString())
}

func (e *ComSharedTypeError) IsTypeError() {}

func (e *ComSharedTypeError) ParserRule() antlr.ParserRuleContext {
	return e.Com
}

func NewComSharedTypeError(com *parser.ExprComContext, innerType *Type) Error {
	return &ComSharedTypeError{
		Com:       com,
		InnerType: innerType,
	}
}

type ComDistributedTypeError struct {
	Com       *parser.ExprComContext
	InnerType *Type
}

func (e *ComDistributedTypeError) Error() string {
	return fmt.Sprintf("can not communicate distributed type '%s'", e.InnerType.ToString())
}

func (e *ComDistributedTypeError) IsTypeError() {}

func (e *ComDistributedTypeError) ParserRule() antlr.ParserRuleContext {
	return e.Com
}

func NewComDistributedTypeError(com *parser.ExprComContext, innerType *Type) Error {
	return &ComDistributedTypeError{
		Com:       com,
		InnerType: innerType,
	}
}

type UnsendableTypeError struct {
	Com            *parser.ExprComContext
	UnsendableType *Type
}

func (u *UnsendableTypeError) Error() string {
	return fmt.Sprintf("can not send type '%s'", u.UnsendableType.ToString())
}

func (u *UnsendableTypeError) IsTypeError() {}

func (u *UnsendableTypeError) ParserRule() antlr.ParserRuleContext {
	return u.Com.Expr()
}

func NewUnsendableTypeError(com *parser.ExprComContext, unsendableType *Type) Error {
	return &UnsendableTypeError{
		Com:            com,
		UnsendableType: unsendableType,
	}
}

type ComValueNotAtSenderError struct {
	Com      *parser.ExprComContext
	ExprType *Type
}

func (c *ComValueNotAtSenderError) Error() string {
	sender := parser.RoleTypeAllIdents(c.Com.RoleType(0))[0]
	return fmt.Sprintf("value of type '%s' is not present at sender '%s'", c.ExprType.ToString(), sender.GetText())
}

func (c *ComValueNotAtSenderError) IsTypeError() {}

func (c *ComValueNotAtSenderError) ParserRule() antlr.ParserRuleContext {
	return c.Com.Expr()
}

func NewComValueNotAtSenderError(com *parser.ExprComContext, exprType *Type) Error {
	return &ComValueNotAtSenderError{
		Com:      com,
		ExprType: exprType,
	}
}

type NotDistributedTypeError struct {
	valueType parser.IValueTypeContext
}

func NewNotDistributedTypeError(valueType parser.IValueTypeContext) Error {
	return &NotDistributedTypeError{
		valueType: valueType,
	}
}

func (e *NotDistributedTypeError) Error() string {
	return fmt.Sprintf("type '%s' is not distributed", e.valueType.GetText())
}

func (e *NotDistributedTypeError) IsTypeError() {}

func (e *NotDistributedTypeError) ParserRule() antlr.ParserRuleContext {
	return e.valueType
}

type CallNonFunctionError struct {
	callExpr *parser.ExprCallContext
	symType  *Type
}

func NewCallNonFunctionError(callExpr *parser.ExprCallContext, symType *Type) Error {
	return &CallNonFunctionError{
		callExpr: callExpr,
		symType:  symType,
	}
}

func (e *CallNonFunctionError) Error() string {
	return fmt.Sprintf("call to non-function '%s'", e.symType.ToString())
}

func (e *CallNonFunctionError) IsTypeError() {}

func (e *CallNonFunctionError) ParserRule() antlr.ParserRuleContext {
	return e.callExpr
}

type CallWrongArgCountError struct {
	callExpr *parser.ExprCallContext
}

func NewCallWrongArgCountError(callExpr *parser.ExprCallContext) Error {
	return &CallWrongArgCountError{
		callExpr: callExpr,
	}
}

func (e *CallWrongArgCountError) Error() string {
	return "call has wrong number of arguments"
}

func (e *CallWrongArgCountError) IsTypeError() {}

func (e *CallWrongArgCountError) ParserRule() antlr.ParserRuleContext {
	return e.callExpr
}

type CallWrongRoleCountError struct {
	callExpr *parser.ExprCallContext
}

func NewCallWrongRoleCountError(callExpr *parser.ExprCallContext) Error {
	return &CallWrongRoleCountError{
		callExpr: callExpr,
	}
}

func (e *CallWrongRoleCountError) Error() string {
	return "call has wrong number of roles"
}

func (e *CallWrongRoleCountError) IsTypeError() {}

func (e *CallWrongRoleCountError) ParserRule() antlr.ParserRuleContext {
	return e.callExpr
}
