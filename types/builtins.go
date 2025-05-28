package types

type Builtin interface {
	Value
	IsBuiltin()
}

var builtin_string StringType = StringType{}
var builtin_int IntType = IntType{}
var builtin_float FloatType = FloatType{}
var builtin_bool BoolType = BoolType{}

type StringType struct{}

func (t *StringType) SubstituteRoles(substMap *RoleSubst) Value {
	return t
}

func (t *StringType) IsSendable() bool {
	return true
}

func (t *StringType) IsEquatable() bool {
	return true
}

func (t *StringType) ToString() string {
	return "String"
}
func (t *StringType) IsValue()   {}
func (t *StringType) IsBuiltin() {}

func String() Value {
	return &builtin_string
}

type IntType struct{}

func (t *IntType) SubstituteRoles(substMap *RoleSubst) Value {
	return t
}

func (t *IntType) IsSendable() bool {
	return true
}

func (t *IntType) IsEquatable() bool {
	return true
}

func (t *IntType) ToString() string {
	return "Int"
}
func (t *IntType) IsValue()   {}
func (t *IntType) IsBuiltin() {}

func Int() Value {
	return &builtin_int
}

type FloatType struct{}

func (t *FloatType) SubstituteRoles(substMap *RoleSubst) Value {
	return t
}

func (t *FloatType) IsEquatable() bool {
	return true
}

func (t *FloatType) IsSendable() bool {
	return true
}

func (t *FloatType) ToString() string {
	return "Float"
}
func (t *FloatType) IsValue()   {}
func (t *FloatType) IsBuiltin() {}

func Float() Value {
	return &builtin_float
}

type BoolType struct{}

func (t *BoolType) SubstituteRoles(substMap *RoleSubst) Value {
	return t
}

func (t *BoolType) IsSendable() bool {
	return true
}

func (t *BoolType) IsEquatable() bool {
	return true
}

func (t *BoolType) ToString() string {
	return "Bool"
}
func (t *BoolType) IsValue()   {}
func (t *BoolType) IsBuiltin() {}

func Bool() Value {
	return &builtin_bool
}
