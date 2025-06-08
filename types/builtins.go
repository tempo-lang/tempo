package types

import "fmt"

type BuiltinType string

type Builtin interface {
	Value
	IsBuiltin()
}

const (
	BuiltinString BuiltinType = "String"
	BuiltinInt    BuiltinType = "Int"
	BuiltinFloat  BuiltinType = "Float"
	BuiltinBool   BuiltinType = "Bool"
)

func BuiltinKind(t Value) BuiltinType {
	switch t.(type) {
	case *BoolType:
		return BuiltinBool
	case *FloatType:
		return BuiltinFloat
	case *IntType:
		return BuiltinInt
	case *StringType:
		return BuiltinString
	}
	return ""
}

type baseBuiltin struct {
	baseValue
	participants []string
}

func (*baseBuiltin) IsBuiltin() {}

func (*baseBuiltin) IsSendable() bool {
	return true
}

func (*baseBuiltin) IsEquatable() bool {
	return true
}

func (b *baseBuiltin) Roles() *Roles {
	return NewRole(b.participants, true)
}

func (b *baseBuiltin) substituteParticipants(substMap *RoleSubst) []string {
	return b.Roles().SubstituteRoles(substMap).participants
}

func (b *baseBuiltin) formatType(t BuiltinType) string {
	if len(b.participants) == 0 {
		return string(t)
	} else {
		return fmt.Sprintf("%s@%s", t, b.Roles().ToString())
	}
}

type StringType struct {
	baseBuiltin
}

func (t *StringType) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(t, other); ok != nil {
		return value, *ok
	}

	if _, ok := other.(*StringType); ok {
		return other, true
	}

	return Invalid(), false
}

func (t *StringType) SubstituteRoles(substMap *RoleSubst) Value {
	return String(t.baseBuiltin.substituteParticipants(substMap))
}

func (t *StringType) ReplaceSharedRoles(participants []string) Value {
	return String(participants)
}

func (t *StringType) ToString() string {
	return t.formatType(BuiltinString)
}

func String(participants []string) Value {
	return &StringType{
		baseBuiltin: baseBuiltin{
			participants: participants,
		},
	}
}

type IntType struct {
	baseBuiltin
}

func (t *IntType) SubstituteRoles(substMap *RoleSubst) Value {
	return Int(t.baseBuiltin.substituteParticipants(substMap))
}

func (t *IntType) ReplaceSharedRoles(participants []string) Value {
	return Int(participants)
}

func (t *IntType) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(t, other); ok != nil {
		return value, *ok
	}

	if _, ok := other.(*IntType); ok {
		return other, true
	}

	return Invalid(), false
}

func (t *IntType) ToString() string {
	return t.formatType(BuiltinInt)
}

func Int(participants []string) Value {
	return &IntType{
		baseBuiltin: baseBuiltin{
			participants: participants,
		},
	}
}

type FloatType struct {
	baseBuiltin
}

func (t *FloatType) SubstituteRoles(substMap *RoleSubst) Value {
	return Float(t.baseBuiltin.substituteParticipants(substMap))
}

func (t *FloatType) ReplaceSharedRoles(participants []string) Value {
	return Float(participants)
}

func (t *FloatType) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(t, other); ok != nil {
		return value, *ok
	}

	if _, ok := other.(*FloatType); ok {
		return other, true
	}

	return Invalid(), false
}

func (t *FloatType) ToString() string {
	return t.formatType(BuiltinFloat)
}

func Float(participants []string) Value {
	return &FloatType{
		baseBuiltin: baseBuiltin{
			participants: participants,
		},
	}
}

type BoolType struct {
	baseBuiltin
}

func (t *BoolType) SubstituteRoles(substMap *RoleSubst) Value {
	return Bool(t.baseBuiltin.substituteParticipants(substMap))
}

func (t *BoolType) ReplaceSharedRoles(participants []string) Value {
	return Bool(participants)
}

func (t *BoolType) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(t, other); ok != nil {
		return value, *ok
	}

	if _, ok := other.(*BoolType); ok {
		return other, true
	}

	return Invalid(), false
}

func (t *BoolType) ToString() string {
	return t.formatType(BuiltinBool)
}

func Bool(participants []string) Value {
	return &BoolType{
		baseBuiltin: baseBuiltin{
			participants: participants,
		},
	}
}
