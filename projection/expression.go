package projection

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"

	"github.com/dave/jennifer/jen"
)

type Expression interface {
	Codegen() jen.Code
	Type() Type
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

func (e *ExprInt) Type() Type {
	return &BuiltinType{Type: types.Int(nil)}
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

type ExprFloat struct {
	Value float64
}

func (e *ExprFloat) Codegen() jen.Code {
	return jen.Lit(e.Value)
}

func (e *ExprFloat) Type() Type {
	return &BuiltinType{Type: types.Float(nil)}
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

type ExprString struct {
	Value string
}

func (e *ExprString) Codegen() jen.Code {
	return jen.Lit(e.Value)
}

func (e *ExprString) Type() Type {
	return &BuiltinType{Type: types.String(nil)}
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
	typeVal Type
}

func (e *ExprIdent) Codegen() jen.Code {
	return jen.Id(e.Name)
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

func (e *ExprBool) Type() Type {
	return &BuiltinType{Type: types.Bool(nil)}
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
	typeVal  Type
}

func (e *ExprBinaryOp) Codegen() jen.Code {
	return jen.Add(e.lhs.Codegen()).Op(string(e.operator)).Add(e.rhs.Codegen())
}

func (e *ExprBinaryOp) Type() Type {
	return e.typeVal
}

func (e *ExprBinaryOp) ReturnsValue() bool {
	return true
}

func (e *ExprBinaryOp) HasSideEffects() bool {
	return e.lhs != nil && e.lhs.HasSideEffects() || e.rhs.HasSideEffects()
}

func (e *ExprBinaryOp) IsExpression() {}

func NewExprBinaryOp(operator Operator, lhs Expression, rhs Expression, typeVal Type) Expression {
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

type ExprAwait struct {
	expr    Expression
	typeVal Type
}

func (e *ExprAwait) Inner() Expression {
	return e.expr
}

func (e *ExprAwait) Codegen() jen.Code {
	return RuntimeFunc("GetAsync").Call(e.Inner().Codegen())
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

func (e *ExprSend) Type() Type {
	return UnitType()
}

func NewExprSend(expr Expression, receivers []string) Expression {
	return &ExprSend{
		expr:      expr,
		receivers: receivers,
	}
}

type ExprRecv struct {
	recvType Type
	sender   string
}

func (e *ExprRecv) Codegen() jen.Code {
	recvType := e.recvType.Codegen()
	return RuntimeFunc("Recv").Types(recvType).Call(jen.Id("env"), jen.Lit(e.sender))
}

func (e *ExprRecv) ReturnsValue() bool {
	return true
}

func (e *ExprRecv) HasSideEffects() bool {
	return true
}

func (e *ExprRecv) IsExpression() {}

func (e *ExprRecv) Type() Type {
	return NewAsyncType(e.recvType)
}

func NewExprRecv(recvType Type, sender string) Expression {
	return &ExprRecv{
		recvType: recvType,
		sender:   sender,
	}
}

type ExprCallFunc struct {
	FuncExpr   Expression
	FuncRole   string
	Args       []Expression
	ReturnType Type
	RoleSubs   *types.RoleSubst
}

func (e *ExprCallFunc) Codegen() jen.Code {
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

func (e *ExprCallClosure) Codegen() jen.Code {
	args := []jen.Code{}

	for _, arg := range e.Args {
		args = append(args, arg.Codegen())
	}

	return jen.Add(e.ClosureExpr.Codegen()).Call(args...)
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

type ExprFieldAccess struct {
	BaseExpr  Expression
	FieldName string
	FieldType Type
}

func (e *ExprFieldAccess) Codegen() jen.Code {
	return jen.Add(e.BaseExpr.Codegen()).Dot(e.FieldName)
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
