package types

import "fmt"

type BuiltinType string

type Builtin interface {
	Type
	IsBuiltin()
}

const (
	BuiltinString BuiltinType = "String"
	BuiltinInt    BuiltinType = "Int"
	BuiltinFloat  BuiltinType = "Float"
	BuiltinBool   BuiltinType = "Bool"
)

func BuiltinKind(t Type) BuiltinType {
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
	baseType
	participants []string
}

func newBaseBuiltin(participants []string) baseBuiltin {
	if len(participants) == 1 && participants[0] == "" {
		participants = []string{}
	}

	return baseBuiltin{
		participants: participants,
	}
}

func (*baseBuiltin) IsBuiltin() {}

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

func (t *StringType) CoerceTo(other Type) (Type, bool) {
	if value, ok := baseCoerceValue(t, other); ok != nil {
		return value, *ok
	}

	if _, ok := other.(*StringType); ok {
		return other, true
	}

	return Invalid(), false
}

func (t *StringType) SubstituteRoles(substMap *RoleSubst) Type {
	return String(t.baseBuiltin.substituteParticipants(substMap))
}

func (t *StringType) ReplaceSharedRoles(participants []string) Type {
	return String(participants)
}

func (t *StringType) ToString() string {
	return t.formatType(BuiltinString)
}

func String(participants []string) Type {
	return &StringType{
		baseBuiltin: newBaseBuiltin(participants),
	}
}

type IntType struct {
	baseBuiltin
}

func (t *IntType) SubstituteRoles(substMap *RoleSubst) Type {
	return Int(t.baseBuiltin.substituteParticipants(substMap))
}

func (t *IntType) ReplaceSharedRoles(participants []string) Type {
	return Int(participants)
}

func (t *IntType) CoerceTo(other Type) (Type, bool) {
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

func Int(participants []string) Type {
	return &IntType{
		baseBuiltin: newBaseBuiltin(participants),
	}
}

type FloatType struct {
	baseBuiltin
}

func (t *FloatType) SubstituteRoles(substMap *RoleSubst) Type {
	return Float(t.baseBuiltin.substituteParticipants(substMap))
}

func (t *FloatType) ReplaceSharedRoles(participants []string) Type {
	return Float(participants)
}

func (t *FloatType) CoerceTo(other Type) (Type, bool) {
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

func Float(participants []string) Type {
	return &FloatType{
		baseBuiltin: newBaseBuiltin(participants),
	}
}

type BoolType struct {
	baseBuiltin
}

func (t *BoolType) SubstituteRoles(substMap *RoleSubst) Type {
	return Bool(t.baseBuiltin.substituteParticipants(substMap))
}

func (t *BoolType) ReplaceSharedRoles(participants []string) Type {
	return Bool(participants)
}

func (t *BoolType) CoerceTo(other Type) (Type, bool) {
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

func Bool(participants []string) Type {
	return &BoolType{
		baseBuiltin: newBaseBuiltin(participants),
	}
}
