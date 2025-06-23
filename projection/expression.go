package projection

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"
)

// An expression that when evaluated returns a value.
type Expression interface {
	// The type returned after evaluating the expression.
	Type() Type
	// Whether or not the expression performs side effects.
	// In cases where this is false, the compiler might remove the expression, if its result is never used.
	HasSideEffects() bool
	// Whether or not the expression returns a value.
	ReturnsValue() bool
	IsExpression()
}

// An expression that constructs a constant [IntType] value when evaluated.
type ExprInt struct {
	Value int
}

func (e *ExprInt) Type() Type {
	return IntType
}

func (e *ExprInt) ReturnsValue() bool {
	return true
}

func (e *ExprInt) HasSideEffects() bool {
	return false
}

func (e *ExprInt) IsExpression() {}

func NewExprInt(value int) Expression {
	return &ExprInt{
		Value: value,
	}
}

// An expression that constructs a constant [FloatType] value when evaluated.
type ExprFloat struct {
	Value float64
}

func (e *ExprFloat) Type() Type {
	return FloatType
}

func (e *ExprFloat) ReturnsValue() bool {
	return true
}

func (e *ExprFloat) HasSideEffects() bool {
	return false
}

func (e *ExprFloat) IsExpression() {}

func NewExprFloat(value float64) Expression {
	return &ExprFloat{
		Value: value,
	}
}

// An expression that constructs a constant [StringType] value when evaluated.
type ExprString struct {
	Value string
}

func (e *ExprString) Type() Type {
	return StringType
}

func (e *ExprString) ReturnsValue() bool {
	return true
}

func (e *ExprString) HasSideEffects() bool {
	return false
}

func (e *ExprString) IsExpression() {}

func NewExprString(value string) Expression {
	return &ExprString{
		Value: value,
	}
}

// An expression that returns the value of symbol with the given name.
type ExprIdent struct {
	Name    string
	typeVal Type
}

func (e *ExprIdent) Type() Type {
	return e.typeVal
}

func (e *ExprIdent) ReturnsValue() bool {
	return true
}

func (e *ExprIdent) HasSideEffects() bool {
	return false
}

func (e *ExprIdent) IsExpression() {}

func NewExprIdent(name string, typeVal Type) Expression {
	return &ExprIdent{Name: name, typeVal: typeVal}
}

// An expression that constructs a constant [BoolType] value when evaluated.
type ExprBool struct {
	Value bool
}

func (e *ExprBool) Type() Type {
	return BoolType
}

func (e *ExprBool) ReturnsValue() bool {
	return true
}

func (e *ExprBool) HasSideEffects() bool {
	return false
}

func (e *ExprBool) IsExpression() {}

func NewExprBool(value bool) Expression {
	return &ExprBool{Value: value}
}

// An operator used in a binary operation.
type Operator string

const (
	// For [IntType] and [FloatType] values, add the numbers.
	// For [StringType] concatenate the strings.
	OpAdd Operator = "+"
	// For [IntType] and [FloatType] values, subtract the numbers.
	OpSub Operator = "-"
	// For [IntType] and [FloatType] values, multiply the numbers.
	OpMul Operator = "*"
	// For [FloatType] values performs regular division.
	// In the case of [IntType] performs interger division.
	OpDiv Operator = "/"
	// For [FloatType] values performs regular modulus.
	// In the case of [IntType] performs interger modulus.
	OpMod Operator = "%"
	// Performs value based deep equality.
	OpEq Operator = "=="
	// Performs value based deep inequality.
	OpNotEq Operator = "!="
	// For [IntType] and [FloatType] values, evalues whether the first value is less than the second.
	OpLess Operator = "<"
	// For [IntType] and [FloatType] values, evalues whether the first value is greater than or equal to the second.
	OpLessEq Operator = "<="
	// For [IntType] and [FloatType] values, evalues whether the first value is greater than the second.
	OpGreater Operator = ">"
	// For [IntType] and [FloatType] values, evalues whether the first value is greater than or equal to the second.
	OpGreaterEq Operator = ">="
	// For [BoolType] values, eagerly evaluates whether both values are true.
	OpAnd Operator = "&&"
	// For [BoolType] values, eagerly evaluates whether one of the values is false.
	OpOr Operator = "||"
)

func ParseOperator(binOp *parser.ExprBinOpContext) Operator {
	var operator Operator
	switch {
	case binOp.PLUS() != nil:
		operator = OpAdd
	case binOp.MINUS() != nil:
		operator = OpSub
	case binOp.MULTIPLY() != nil:
		operator = OpMul
	case binOp.DIVIDE() != nil:
		operator = OpDiv
	case binOp.MODULO() != nil:
		operator = OpMod
	case binOp.EQUAL() != nil:
		operator = OpEq
	case binOp.NOT_EQUAL() != nil:
		operator = OpNotEq
	case binOp.LESS() != nil:
		operator = OpLess
	case binOp.LESS_EQ() != nil:
		operator = OpLessEq
	case binOp.GREATER() != nil:
		operator = OpGreater
	case binOp.GREATER_EQ() != nil:
		operator = OpGreaterEq
	case binOp.AND() != nil:
		operator = OpAnd
	case binOp.OR() != nil:
		operator = OpOr
	}
	return operator
}

type ExprBinaryOp struct {
	Operator Operator
	Lhs      Expression
	Rhs      Expression
	typeVal  Type
}

func (e *ExprBinaryOp) Type() Type {
	return e.typeVal
}

func (e *ExprBinaryOp) ReturnsValue() bool {
	return true
}

func (e *ExprBinaryOp) HasSideEffects() bool {
	return e.Lhs != nil && e.Lhs.HasSideEffects() || e.Rhs.HasSideEffects()
}

func (e *ExprBinaryOp) IsExpression() {}

func NewExprBinaryOp(operator Operator, lhs Expression, rhs Expression, typeVal Type) Expression {
	return &ExprBinaryOp{
		Operator: operator,
		Lhs:      lhs,
		Rhs:      rhs,
		typeVal:  typeVal,
	}
}

// Converts the inner expression to a resolved async value.
type ExprAsync struct {
	inner Expression
}

func (e *ExprAsync) Type() Type {
	return NewAsyncType(e.inner.Type())
}

func (e *ExprAsync) ReturnsValue() bool {
	return true
}

func (e *ExprAsync) HasSideEffects() bool {
	return e.Inner().HasSideEffects()
}

func (e *ExprAsync) IsExpression() {}

func (e *ExprAsync) Inner() Expression {
	return e.inner
}

func NewExprAsync(inner Expression) Expression {
	return &ExprAsync{inner: inner}
}

// Waits for the async expression to resolve and returns the resolved value.
type ExprAwait struct {
	expr    Expression
	typeVal Type
}

func (e *ExprAwait) Inner() Expression {
	return e.expr
}

func (e *ExprAwait) Type() Type {
	return e.typeVal
}

func (e *ExprAwait) ReturnsValue() bool {
	return true
}

func (e *ExprAwait) HasSideEffects() bool {
	return e.Inner().HasSideEffects()
}

func (e *ExprAwait) IsExpression() {}

func NewExprAwait(inner Expression, typeVal Type) Expression {
	return &ExprAwait{expr: inner, typeVal: typeVal}
}

// Send the value of the inner expression to a set of recipients.
// The recipient set cannot be empty but can contain the sender.
type ExprSend struct {
	Expr      Expression
	Receivers []string
}

func (e *ExprSend) ReturnsValue() bool {
	return false
}

func (e *ExprSend) HasSideEffects() bool {
	return true
}

func (e *ExprSend) IsExpression() {}

func (e *ExprSend) Type() Type {
	return UnitType()
}

func NewExprSend(expr Expression, receivers []string) Expression {
	return &ExprSend{
		Expr:      expr,
		Receivers: receivers,
	}
}

// Receive a value from the sender and returns an async value that resolved when the value is received.
// The sender and receiver may be the same process.
type ExprRecv struct {
	RecvType Type
	Sender   string
}

func (e *ExprRecv) ReturnsValue() bool {
	return true
}

func (e *ExprRecv) HasSideEffects() bool {
	return true
}

func (e *ExprRecv) IsExpression() {}

func (e *ExprRecv) Type() Type {
	return NewAsyncType(e.RecvType)
}

func NewExprRecv(recvType Type, sender string) Expression {
	return &ExprRecv{
		RecvType: recvType,
		Sender:   sender,
	}
}

type ExprCallFunc struct {
	FuncExpr   Expression
	FuncRole   string
	Args       []Expression
	ReturnType Type
	RoleSubs   *types.RoleSubst
}

func (e *ExprCallFunc) Type() Type {
	return e.ReturnType
}

func (e *ExprCallFunc) ReturnsValue() bool {
	return e.ReturnType != UnitType()
}

func (e *ExprCallFunc) HasSideEffects() bool {
	return true
}

func (e *ExprCallFunc) IsExpression() {}

func NewExprCallFunc(funcExpr Expression, funcRole string, args []Expression, returnType Type, roleSubs *types.RoleSubst) Expression {
	return &ExprCallFunc{
		FuncExpr:   funcExpr,
		FuncRole:   funcRole,
		Args:       args,
		ReturnType: returnType,
		RoleSubs:   roleSubs,
	}
}

type ExprCallClosure struct {
	ClosureExpr Expression
	Role        string
	Args        []Expression
	ReturnType  Type
}

func (e *ExprCallClosure) Type() Type {
	return e.ReturnType
}

func (e *ExprCallClosure) ReturnsValue() bool {
	return e.ReturnType != UnitType()
}

func (e *ExprCallClosure) HasSideEffects() bool {
	return true
}

func (e *ExprCallClosure) IsExpression() {}

func NewExprCallClosure(closureExpr Expression, role string, args []Expression, returnType Type) Expression {
	return &ExprCallClosure{
		ClosureExpr: closureExpr,
		Role:        role,
		Args:        args,
		ReturnType:  returnType,
	}
}

// Construct an instance of the struct with the given name.
type ExprStruct struct {
	StructName string
	StructRole string
	FieldNames []string
	Fields     map[string]Expression
	StructType *StructType
}

func (e *ExprStruct) Type() Type {
	return e.StructType
}

func (e *ExprStruct) ReturnsValue() bool {
	return true
}

func (e *ExprStruct) HasSideEffects() bool {
	for _, expr := range e.Fields {
		if expr.HasSideEffects() {
			return true
		}
	}
	return false
}

func (e *ExprStruct) IsExpression() {}

func NewExprStruct(structName, structRole string, fieldNames []string, fields map[string]Expression, typ *StructType) Expression {
	return &ExprStruct{
		StructName: structName,
		StructRole: structRole,
		FieldNames: fieldNames,
		Fields:     fields,
		StructType: typ,
	}
}

// Access a field on an inner expression.
// For [StructType], each struct field can be accessed.
type ExprFieldAccess struct {
	BaseExpr  Expression
	FieldName string
	FieldType Type
}

func (e *ExprFieldAccess) Type() Type {
	return e.FieldType
}

func (e *ExprFieldAccess) ReturnsValue() bool {
	return true
}

func (e *ExprFieldAccess) HasSideEffects() bool {
	return e.BaseExpr.HasSideEffects()
}

func (e *ExprFieldAccess) IsExpression() {}

func NewExprFieldAccess(structExpr Expression, fieldName string, typ Type) Expression {
	return &ExprFieldAccess{
		BaseExpr:  structExpr,
		FieldName: fieldName,
		FieldType: typ,
	}
}
