package types

type Builtin interface {
	Type
	IsBuiltin()
}

var builtin_string StringType = StringType{}
var builtin_int IntType = IntType{}
var builtin_float FloatType = FloatType{}
var builtin_bool BoolType = BoolType{}

type StringType struct{}

func (t *StringType) ToString() string {
	return "String"
}
func (t *StringType) IsType()    {}
func (t *StringType) IsBuiltin() {}

func String() *StringType {
	return &builtin_string
}

type IntType struct{}

func (t *IntType) ToString() string {
	return "Int"
}
func (t *IntType) IsType()    {}
func (t *IntType) IsBuiltin() {}

func Int() *IntType {
	return &builtin_int
}

type FloatType struct{}

func (t *FloatType) ToString() string {
	return "Float"
}
func (t *FloatType) IsType()    {}
func (t *FloatType) IsBuiltin() {}

func Float() *FloatType {
	return &builtin_float
}

type BoolType struct{}

func (t *BoolType) ToString() string {
	return "Bool"
}
func (t *BoolType) IsType()    {}
func (t *BoolType) IsBuiltin() {}

func Bool() *BoolType {
	return &builtin_bool
}
