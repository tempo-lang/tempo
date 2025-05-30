package projection

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"

	"github.com/dave/jennifer/jen"
)

type Expression interface {
	Codegen() jen.Code
	Type() types.Value
	HasSideEffects() bool
	ReturnsValue() bool
	IsExpression()
}

type ExprInt struct {
	Value int
}

func (e *ExprInt) Codegen() jen.Code {
	return jen.Lit(e.Value)
}

func (e *ExprInt) Type() types.Value {
	return types.Int()
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

type ExprString struct {
	Value string
}

func (e *ExprString) Codegen() jen.Code {
	return jen.Lit(e.Value)
}

func (e *ExprString) Type() types.Value {
	return types.String()
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

type ExprIdent struct {
	Name    string
	typeVal types.Value
}

func (e *ExprIdent) Codegen() jen.Code {
	return jen.Id(e.Name)
}

func (e *ExprIdent) Type() types.Value {
	return e.typeVal
}

func (e *ExprIdent) ReturnsValue() bool {
	return true
}

func (e *ExprIdent) HasSideEffects() bool {
	return false
}

func (e *ExprIdent) IsExpression() {}

func NewExprIdent(name string, typeVal types.Value) Expression {
	return &ExprIdent{Name: name, typeVal: typeVal}
}

type ExprBool struct {
	Value bool
}

func (e *ExprBool) Codegen() jen.Code {
	if e.Value {
		return jen.True()
	} else {
		return jen.False()
	}
}

func (e *ExprBool) Type() types.Value {
	return types.Bool()
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

type Operator string

const (
	OpAdd       Operator = "+"
	OpSub       Operator = "-"
	OpMul       Operator = "*"
	OpDiv       Operator = "/"
	OpMod       Operator = "%"
	OpEq        Operator = "=="
	OpNotEq     Operator = "!="
	OpLess      Operator = "<"
	OpLessEq    Operator = "<="
	OpGreater   Operator = ">"
	OpGreaterEq Operator = ">="
	OpAnd       Operator = "&&"
	OpOr        Operator = "||"
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
	operator Operator
	lhs      Expression
	rhs      Expression
	typeVal  types.Value
}

func (e *ExprBinaryOp) Codegen() jen.Code {
	return jen.Add(e.lhs.Codegen()).Op(string(e.operator)).Add(e.rhs.Codegen())
}

func (e *ExprBinaryOp) Type() types.Value {
	return e.typeVal
}

func (e *ExprBinaryOp) ReturnsValue() bool {
	return true
}

func (e *ExprBinaryOp) HasSideEffects() bool {
	return e.lhs != nil && e.lhs.HasSideEffects() || e.rhs.HasSideEffects()
}

func (e *ExprBinaryOp) IsExpression() {}

func NewExprBinaryOp(operator Operator, lhs Expression, rhs Expression, typeVal types.Value) Expression {
	return &ExprBinaryOp{
		operator: operator,
		lhs:      lhs,
		rhs:      rhs,
		typeVal:  typeVal,
	}
}

type ExprAsync struct {
	inner Expression
}

func (e *ExprAsync) Codegen() jen.Code {
	return RuntimeFunc("FixedAsync").Call(e.inner.Codegen())
}

func (e *ExprAsync) Type() types.Value {
	return types.NewAsync(e.inner.Type())
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

type ExprAwait struct {
	expr    Expression
	typeVal types.Value
}

func (e *ExprAwait) Inner() Expression {
	return e.expr
}

func (e *ExprAwait) Codegen() jen.Code {
	return RuntimeFunc("GetAsync").Call(e.Inner().Codegen())
}

func (e *ExprAwait) Type() types.Value {
	return e.typeVal
}

func (e *ExprAwait) ReturnsValue() bool {
	return true
}

func (e *ExprAwait) HasSideEffects() bool {
	return e.Inner().HasSideEffects()
}

func (e *ExprAwait) IsExpression() {}

func NewExprAwait(inner Expression, typeVal types.Value) Expression {
	return &ExprAwait{expr: inner, typeVal: typeVal}
}

type ExprSend struct {
	expr      Expression
	receivers []string
}

func (e *ExprSend) Codegen() jen.Code {
	args := []jen.Code{
		jen.Id("env"),
		e.expr.Codegen(),
	}

	for _, role := range e.receivers {
		args = append(args, jen.Lit(role))
	}

	return RuntimeFunc("Send").Call(args...)
}

func (e *ExprSend) ReturnsValue() bool {
	return false
}

func (e *ExprSend) HasSideEffects() bool {
	return true
}

func (e *ExprSend) IsExpression() {}

func (e *ExprSend) Type() types.Value {
	return types.Unit()
}

func NewExprSend(expr Expression, receivers []string) Expression {
	return &ExprSend{
		expr:      expr,
		receivers: receivers,
	}
}

type ExprRecv struct {
	recvType types.Value
	sender   string
}

func (e *ExprRecv) Codegen() jen.Code {
	recvType := CodegenType(e.recvType)
	return RuntimeFunc("Recv").Types(recvType).Call(jen.Id("env"), jen.Lit(e.sender))
}

func (e *ExprRecv) ReturnsValue() bool {
	return true
}

func (e *ExprRecv) HasSideEffects() bool {
	return true
}

func (e *ExprRecv) IsExpression() {}

func (e *ExprRecv) Type() types.Value {
	return types.NewAsync(e.recvType)
}

func NewExprRecv(recvType types.Value, sender string) Expression {
	return &ExprRecv{
		recvType: recvType,
		sender:   sender,
	}
}

type ExprCall struct {
	FuncExpr   Expression
	FuncRole   string
	Args       []Expression
	ReturnType types.Value
	RoleSubs   *types.RoleSubst
}

func (e *ExprCall) Codegen() jen.Code {
	args := []jen.Code{}

	roleSub := []jen.Code{}
	for _, to := range e.RoleSubs.Roles {
		from := e.RoleSubs.Subst(to)
		roleSub = append(roleSub, jen.Lit(from), jen.Lit(to))
	}

	args = append(args, jen.Id("env").Dot("Subst").Call(roleSub...))

	for _, arg := range e.Args {
		args = append(args, arg.Codegen())
	}

	return jen.Add(e.FuncExpr.Codegen()).Call(args...)
	// return jen.Id(fmt.Sprintf("%s_%s", e.FuncName, e.FuncRole)).Call(args...)
}

func (e *ExprCall) Type() types.Value {
	return e.ReturnType
}

func (e *ExprCall) ReturnsValue() bool {
	return e.ReturnType != types.Unit()
}

func (e *ExprCall) HasSideEffects() bool {
	return true
}

func (e *ExprCall) IsExpression() {}

func NewExprCall(funcExpr Expression, funcRole string, args []Expression, returnType types.Value, roleSubs *types.RoleSubst) Expression {
	return &ExprCall{
		FuncExpr:   funcExpr,
		FuncRole:   funcRole,
		Args:       args,
		ReturnType: returnType,
		RoleSubs:   roleSubs,
	}
}

type ExprStruct struct {
	StructName string
	StructRole string
	FieldNames []string
	Fields     map[string]Expression
	StructType *StructType
}

func (e *ExprStruct) Codegen() jen.Code {
	fields := jen.Dict{}

	for _, fieldName := range e.FieldNames {
		field := e.Fields[fieldName]
		expr := field.Codegen()
		fields[jen.Id(fieldName)] = expr
	}

	name := fmt.Sprintf("%s_%s", e.StructName, e.StructRole)

	return jen.Id(name).Values(fields)
}

func (e *ExprStruct) Type() types.Value {
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

type ExprFieldAccess struct {
	BaseExpr  Expression
	FieldName string
	FieldType types.Value
}

func (e *ExprFieldAccess) Codegen() jen.Code {
	return jen.Add(e.BaseExpr.Codegen()).Dot(e.FieldName)
}

func (e *ExprFieldAccess) Type() types.Value {
	return e.FieldType
}

func (e *ExprFieldAccess) ReturnsValue() bool {
	return true
}

func (e *ExprFieldAccess) HasSideEffects() bool {
	return e.BaseExpr.HasSideEffects()
}

func (e *ExprFieldAccess) IsExpression() {}

func NewExprFieldAccess(structExpr Expression, fieldName string, typ types.Value) Expression {
	return &ExprFieldAccess{
		BaseExpr:  structExpr,
		FieldName: fieldName,
		FieldType: typ,
	}
}
