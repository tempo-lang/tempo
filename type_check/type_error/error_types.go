package type_error

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/types"

	"github.com/tempo-lang/tempo/parser/ast"
)

type ErrorCode int

const (
	CodeValueRoleNotInScope ErrorCode = iota + 1
	CodeUnexpectedSharedType
	CodeUndefinedType
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
	_ // CodeFieldAccessUnexpectedType
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
	CodeNestedAsync
	CodeUnknownType
	CodeIndexWrongBaseType
	CodeIndexRoleNotEncompassBase
	CodeAssignUnitValue
	CodeExpectedInterfaceType
	CodeMissingImplementationMethod
	CodeIncompatibleImplementationMethod
	CodeMissingRoles
	CodeDuplicateStructField
	CodeTypeNotAnExpression
	CodeIncompatibleTypeCast
	CodeUnexpectedHiddenRoles
	CodeHiddenExpression
	CodeHiddenTypeSignature
	CodeIncompleteFunction
	CodeHiddenStructField
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
	ParserRule ast.Node
}

type CodeAction struct {
	Title     string
	Range     ast.Node
	NewSource string
}

type Error interface {
	error
	ParserRule() ast.Node
	IsTypeError()
	RelatedInfo() []RelatedInfo
	Annotations() []Annotation
	CodeAction() *CodeAction
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
func (*baseError) CodeAction() *CodeAction {
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

func amount(number int, singular, plural string) string {
	if number == 1 {
		return fmt.Sprintf("%d %s", number, singular)
	} else {
		return fmt.Sprintf("%d %s", number, plural)
	}
}

type ValueRoleNotInScope struct {
	baseError
	Value             ast.Node
	ValueRoles        *types.Roles
	InaccessibleRoles []string
}

func (v *ValueRoleNotInScope) Error() string {
	roles := formatList("role", "roles", v.InaccessibleRoles, "and")
	return fmt.Sprintf("value `%s` contains %s that %s not in scope", ast.Text(v.Value), roles, toBe(v.InaccessibleRoles))
}

func (v *ValueRoleNotInScope) ParserRule() ast.Node {
	return v.Value
}

func (e *ValueRoleNotInScope) Code() ErrorCode {
	return CodeValueRoleNotInScope
}

func NewValueRoleNotInScope(value ast.Node, valueRoles *types.Roles, inaccessibleRoles []string) Error {
	return &ValueRoleNotInScope{
		Value:             value,
		ValueRoles:        valueRoles,
		InaccessibleRoles: inaccessibleRoles,
	}
}

type UnexpectedSharedType struct {
	baseError
	RoleType *ast.RoleType
}

func (u *UnexpectedSharedType) Error() string {
	return "shared type is not allowed here"
}

func (u *UnexpectedSharedType) CodeAction() *CodeAction {
	roles := misc.JoinStringsFunc(u.RoleType.Roles(), ",", func(role *ast.Role) string {
		return role.Token.Text
	})

	return &CodeAction{
		Title:     "Change to distributed type",
		Range:     u.RoleType,
		NewSource: fmt.Sprintf("(%s)", roles),
	}
}

func (u *UnexpectedSharedType) ParserRule() ast.Node {
	return u.RoleType
}

func (e *UnexpectedSharedType) Code() ErrorCode {
	return CodeUnexpectedSharedType
}

func NewUnexpectedSharedType(roleType *ast.RoleType) Error {
	return &UnexpectedSharedType{
		RoleType: roleType,
	}
}

type UndefinedType struct {
	baseError
	TypeName *ast.Identifier
}

func NewUndefinedType(typeName *ast.Identifier) Error {
	return &UndefinedType{
		TypeName: typeName,
	}
}

func (e *UndefinedType) Error() string {
	return fmt.Sprintf("type name `%s` is undefined or not in scope", e.TypeName.Value())
}

func (e *UndefinedType) ParserRule() ast.Node {
	return e.TypeName
}

func (e *UndefinedType) Code() ErrorCode {
	return CodeUndefinedType
}

type UnknownType struct {
	baseError
	Expr ast.Expr
}

func NewUnknownType(expr ast.Expr) Error {
	return &UnknownType{
		Expr: expr,
	}
}

func (e *UnknownType) Error() string {
	return "cannot determine the type of this expression"
}

func (e *UnknownType) Annotations() []Annotation {
	return []Annotation{{
		Type:    AnnotationTypeHint,
		Message: "give an explicit type in the declaration of the variable",
	}}
}

func (e *UnknownType) ParserRule() ast.Node {
	return e.Expr
}

func (e *UnknownType) Code() ErrorCode {
	return CodeUnknownType
}

type HiddenTypeSignature struct {
	baseError
	ValueType ast.ValueType
	Type      types.Type
}

func NewHiddenTypeSignature(valueType ast.ValueType, typeSig types.Type) Error {
	return &HiddenTypeSignature{
		ValueType: valueType,
		Type:      typeSig,
	}
}

func (e *HiddenTypeSignature) Error() string {
	return fmt.Sprintf("hidden type signature `%s`", e.Type.ToString())
}

func (e *HiddenTypeSignature) Annotations() []Annotation {
	return []Annotation{
		{
			Type:    AnnotationTypeNote,
			Message: "all the roles in the type are hidden",
		},
	}
}

func (e *HiddenTypeSignature) ParserRule() ast.Node {
	return e.ValueType
}

func (e *HiddenTypeSignature) Code() ErrorCode {
	return CodeHiddenTypeSignature
}

type ValueMismatch struct {
	baseError
	Expr        ast.Expr
	FirstValue  types.Type
	SecondValue types.Type
}

func (t *ValueMismatch) Error() string {
	return fmt.Sprintf("types `%s` and `%s` are not compatible", t.FirstValue.ToString(), t.SecondValue.ToString())
}

func (t *ValueMismatch) ParserRule() ast.Node {
	return t.Expr
}

func (e *ValueMismatch) Code() ErrorCode {
	return CodeValueMismatch
}

func NewValueMismatch(expr ast.Expr, firstValue types.Type, secondValue types.Type) Error {
	return &ValueMismatch{
		Expr:        expr,
		FirstValue:  firstValue,
		SecondValue: secondValue,
	}
}

type IncompatibleTypes struct {
	baseError
	Expr         ast.Expr
	ExprType     types.Type
	ExpectedType types.Type
}

func (e *IncompatibleTypes) Error() string {
	return fmt.Sprintf("type `%s` is not compatible with type `%s`", e.ExprType.ToString(), e.ExpectedType.ToString())
}

func (e *IncompatibleTypes) ParserRule() ast.Node {
	return e.Expr
}

func (e *IncompatibleTypes) Code() ErrorCode {
	return CodeIncompatibleTypes
}

func NewIncompatibleTypes(expr ast.Expr, exprType types.Type, expectedType types.Type) Error {
	return &IncompatibleTypes{
		Expr:         expr,
		ExprType:     exprType,
		ExpectedType: expectedType,
	}
}

type InvalidValue struct {
	baseError
	Expr          ast.Expr
	ActualValue   types.Type
	ExpectedValue types.Type
}

func (i *InvalidValue) Error() string {
	return fmt.Sprintf("type `%s` does not match expected type `%s`", i.ActualValue.ToString(), i.ExpectedValue.ToString())
}

func (i *InvalidValue) ParserRule() ast.Node {
	return i.Expr
}

func (e *InvalidValue) Code() ErrorCode {
	return CodeInvalidValue
}

func NewInvalidValue(expr ast.Expr, actualValue types.Type, expectedValue types.Type) Error {
	return &InvalidValue{
		Expr:          expr,
		ActualValue:   actualValue,
		ExpectedValue: expectedValue,
	}
}

type AwaitNonAsyncType struct {
	baseError
	Expr *ast.AwaitExpr
	Type types.Type
}

func (e *AwaitNonAsyncType) Error() string {
	return fmt.Sprintf("cannot await type `%s` since it is not an async type", e.Type.ToString())
}

func (e *AwaitNonAsyncType) ParserRule() ast.Node {
	return e.Expr
}

func (e *AwaitNonAsyncType) Code() ErrorCode {
	return CodeExpectedAsyncType
}

func NewAwaitNonAsyncType(expr *ast.AwaitExpr, errType types.Type) Error {
	return &AwaitNonAsyncType{
		Expr: expr,
		Type: errType,
	}
}

type BinOpIncompatibleType struct {
	baseError
	BinOp   *ast.BinaryExpr
	Value   types.Type
	Allowed []types.BuiltinType
}

func (e *BinOpIncompatibleType) Error() string {
	op := e.BinOp.Operator.Text
	return fmt.Sprintf("operation `%s` cannot be performed on `%s` types", op, e.Value.ToString())
}

func (e *BinOpIncompatibleType) Annotations() []Annotation {
	allowed := make([]string, len(e.Allowed))
	for i, v := range e.Allowed {
		allowed[i] = fmt.Sprintf("`%s`", v)
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

func (e *BinOpIncompatibleType) ParserRule() ast.Node {
	return e.BinOp
}

func (e *BinOpIncompatibleType) Code() ErrorCode {
	return CodeBinOpIncompatibleType
}

func NewBinOpIncompatibleType(binOp *ast.BinaryExpr, value types.Type, allowed []types.BuiltinType) Error {
	return &BinOpIncompatibleType{
		BinOp:   binOp,
		Value:   value,
		Allowed: allowed,
	}
}

type UnsendableType struct {
	baseError
	Com            *ast.ComExpr
	UnsendableType types.Type
}

func (u *UnsendableType) Error() string {
	return fmt.Sprintf("can not send value of type `%s`", u.UnsendableType.ToString())
}

func (u *UnsendableType) ParserRule() ast.Node {
	return u.Com.Expr
}

func (e *UnsendableType) Code() ErrorCode {
	return CodeUnsendableType
}

func (u *UnsendableType) Annotations() []Annotation {
	result := []Annotation{}

	if _, isStruct := u.UnsendableType.(*types.StructType); isStruct {
		result = append(result, Annotation{
			Type:    AnnotationTypeNote,
			Message: "struct contains unsendable fields",
		})
	}

	return result
}

func NewUnsendableType(com *ast.ComExpr, unsendableType types.Type) Error {
	return &UnsendableType{
		Com:            com,
		UnsendableType: unsendableType,
	}
}

type NotDistributedType struct {
	baseError
	typeCtx ast.Node
}

func NewNotDistributedType(typeCtx ast.Node) Error {
	return &NotDistributedType{
		typeCtx: typeCtx,
	}
}

func (e *NotDistributedType) Error() string {
	return fmt.Sprintf("type `%s` cannot be distributed", ast.Text(e.typeCtx))
}

func (e *NotDistributedType) ParserRule() ast.Node {
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

type NestedAsync struct {
	baseError
	AsyncCtx *ast.AsyncType
}

func (e *NestedAsync) Error() string {
	return "nested async types are not allowed"
}

func (e *NestedAsync) ParserRule() ast.Node {
	return e.AsyncCtx
}

func (e *NestedAsync) Code() ErrorCode {
	return CodeNestedAsync
}

func NewNestedAsync(asyncCtx *ast.AsyncType) Error {
	return &NestedAsync{
		AsyncCtx: asyncCtx,
	}
}

type IndexWrongBaseType struct {
	baseError
	BaseRule ast.Node
	BaseType types.Type
}

func (e *IndexWrongBaseType) Error() string {
	return fmt.Sprintf("cannot index value of type `%s`", e.BaseType.ToString())
}

func (e *IndexWrongBaseType) ParserRule() ast.Node {
	return e.BaseRule
}

func (e *IndexWrongBaseType) Code() ErrorCode {
	return CodeIndexWrongBaseType
}

func NewIndexWrongBaseType(baseRule ast.Node, baseType types.Type) Error {
	return &IndexWrongBaseType{
		BaseRule: baseRule,
		BaseType: baseType,
	}
}

type IndexRoleNotEncompassBase struct {
	baseError
	IndexExpr  *ast.IndexExpr
	BaseType   types.Type
	IndexRoles *types.Roles
}

func (e *IndexRoleNotEncompassBase) Error() string {

	var doesDo string
	if e.IndexRoles.IsLocalRole() {
		doesDo = "does"
	} else {
		doesDo = "do"
	}

	roles := formatList("role", "roles", e.IndexRoles.Participants(), "and")

	return fmt.Sprintf("index %s %s not encompass roles in base type `%s`", roles, doesDo, e.BaseType.ToString())
}

func (e *IndexRoleNotEncompassBase) ParserRule() ast.Node {
	return e.IndexExpr.Index
}

func (e *IndexRoleNotEncompassBase) Code() ErrorCode {
	return CodeIndexRoleNotEncompassBase
}

func NewIndexRoleNotEncompassBase(indexExpr *ast.IndexExpr, baseType types.Type, indexRoles *types.Roles) Error {
	return &IndexRoleNotEncompassBase{
		IndexExpr:  indexExpr,
		BaseType:   baseType,
		IndexRoles: indexRoles,
	}
}
