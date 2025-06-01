package types

type Builtin interface {
	Value
	IsBuiltin()
}

type baseBuiltin struct{}

func (b *baseBuiltin) IsValue()   {}
func (b *baseBuiltin) IsBuiltin() {}

func (b *baseBuiltin) IsSendable() bool {
	return true
}

func (b *baseBuiltin) IsEquatable() bool {
	return true
}

var builtin_string StringType = StringType{}
var builtin_int IntType = IntType{}
var builtin_float FloatType = FloatType{}
var builtin_bool BoolType = BoolType{}

type StringType struct {
	baseBuiltin
}

func (t *StringType) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(t, other); ok {
		return value, true
	}

	if other == String() {
		return String(), true
	}
	return Invalid().Value(), false
}

func (t *StringType) SubstituteRoles(substMap *RoleSubst) Value {
	return t
}

func (t *StringType) ToString() string {
	return "String"
}

func String() Value {
	return &builtin_string
}

type IntType struct {
	baseBuiltin
}

func (t *IntType) SubstituteRoles(substMap *RoleSubst) Value {
	return t
}

func (t *IntType) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(t, other); ok {
		return value, true
	}

	if other == Int() {
		return Int(), true
	}
	return Invalid().Value(), false
}

func (t *IntType) ToString() string {
	return "Int"
}

func Int() Value {
	return &builtin_int
}

type FloatType struct {
	baseBuiltin
}

func (t *FloatType) SubstituteRoles(substMap *RoleSubst) Value {
	return t
}

func (t *FloatType) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(t, other); ok {
		return value, true
	}

	if other == Float() {
		return Float(), true
	}
	return Invalid().Value(), false
}

func (t *FloatType) ToString() string {
	return "Float"
}

func Float() Value {
	return &builtin_float
}

type BoolType struct {
	baseBuiltin
}

func (t *BoolType) SubstituteRoles(substMap *RoleSubst) Value {
	return t
}

func (t *BoolType) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(t, other); ok {
		return value, true
	}

	if other == Bool() {
		return Bool(), true
	}
	return Invalid().Value(), false
}

func (t *BoolType) ToString() string {
	return "Bool"
}

func Bool() Value {
	return &builtin_bool
}
