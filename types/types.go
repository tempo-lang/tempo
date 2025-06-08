package types

func baseCoerceValue(thisValue, otherValue Value) (Value, bool) {
	if !thisValue.Roles().Encompass(otherValue.Roles()) {
		return otherValue, false
	}

	if thisValue.IsInvalid() {
		return otherValue, true
	}

	if otherValue.IsInvalid() {
		return Invalid(), true
	}

	// plain types can coerce to async types
	if _, isAsync := thisValue.(*Async); !isAsync {
		if otherAsync, otherIsAsync := otherValue.(*Async); otherIsAsync {
			innerCoerce, ok := thisValue.CoerceTo(otherAsync.Inner())
			return NewAsync(innerCoerce), ok
		}
	}

	return nil, false
}

// func (t *Type) ToString() string {

// 	switch value := t.value.(type) {
// 	case *ClosureType:
// 		params := misc.JoinStringsFunc(value.params, ", ", func(param *Type) string { return param.ToString() })
// 		returnType := ""
// 		if value.returnType.Value() != Unit() {
// 			returnType = value.returnType.ToString()
// 		}
// 		return fmt.Sprintf("func@%s(%s)%s", t.roles.ToString(), params, returnType)
// 	case *FunctionType:
// 		params := misc.JoinStringsFunc(value.params, ", ", func(param *Type) string { return param.ToString() })
// 		returnType := ""
// 		if value.returnType.Value() != Unit() {
// 			returnType = value.returnType.ToString()
// 		}
// 		return fmt.Sprintf("func@%s %s(%s)%s", t.roles.ToString(), value.NameIdent().GetText(), params, returnType)
// 	case *StructType:
// 		return fmt.Sprintf("struct@%s %s", t.roles.ToString(), value.Name())
// 	case *InterfaceType:
// 		return fmt.Sprintf("interface@%s %s", t.roles.ToString(), value.Name())
// 	}

// 	return fmt.Sprintf("%s@%s", t.value.ToString(), t.roles.ToString())
// }

type Value interface {
	IsSendable() bool
	IsEquatable() bool
	ToString() string
	IsValue()
	SubstituteRoles(substMap *RoleSubst) Value
	ReplaceSharedRoles(participants []string) Value
	CoerceTo(other Value) (Value, bool)
	Roles() *Roles
	IsInvalid() bool
}

type baseValue struct{}

func (*baseValue) IsValue() {}

func (*baseValue) IsInvalid() bool {
	return false
}

type InvalidValue struct {
	baseValue
}

func (*InvalidValue) IsInvalid() bool {
	return true
}

func (t *InvalidValue) SubstituteRoles(substMap *RoleSubst) Value {
	return t
}

func (t *InvalidValue) ReplaceSharedRoles(participants []string) Value {
	return t
}

func (t *InvalidValue) CoerceTo(other Value) (Value, bool) {
	return other, true
}

func (t *InvalidValue) Roles() *Roles {
	return EveryoneRole()
}

var invalid_type InvalidValue = InvalidValue{}

func (t *InvalidValue) IsSendable() bool {
	return true
}

func (t *InvalidValue) IsEquatable() bool {
	return true
}

func (t *InvalidValue) ToString() string {
	return "ERROR"
}

func Invalid() Value {
	return &invalid_type
}

type UnitValue struct {
	baseValue
}

func (u *UnitValue) SubstituteRoles(substMap *RoleSubst) Value {
	return u
}

func (u *UnitValue) ReplaceSharedRoles(participants []string) Value {
	return u
}

func (t *UnitValue) CoerceTo(other Value) (Value, bool) {
	return Unit(), other == Unit()
}

func (t *UnitValue) Roles() *Roles {
	return EveryoneRole()
}

func (u *UnitValue) ToString() string {
	return "()"
}

func (u *UnitValue) IsSendable() bool {
	return true
}

func (t *UnitValue) IsEquatable() bool {
	return false
}

var unit_value UnitValue = UnitValue{}

func Unit() Value {
	return &unit_value
}
