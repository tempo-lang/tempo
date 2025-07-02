// This package includes structures for all types in Tempo from the choreographic perspective (before endpoint projection).
//
// The [Type] interface is implemented by all types in the language.
package types

import "iter"

func baseCoerceValue(thisValue, otherValue Type) (Type, *bool) {
	fail := false
	success := true

	if !thisValue.Roles().Encompass(otherValue.Roles()) {
		return otherValue, &fail
	}

	if thisValue.IsInvalid() {
		return otherValue, &success
	}

	if otherValue.IsInvalid() {
		return Invalid(), &success
	}

	// plain types can coerce to async types
	if _, isAsync := thisValue.(*AsyncType); !isAsync {
		if otherAsync, otherIsAsync := otherValue.(*AsyncType); otherIsAsync {
			innerCoerce, ok := thisValue.CoerceTo(otherAsync.Inner())
			return Async(innerCoerce), &ok
		}
	}

	return nil, nil
}

type TypeFieldMap map[string]Type

// A common interface for all types on the choreographic level
type Type interface {
	// IsSendable describes whether the type can be communicated using the `->` operator.
	IsSendable() bool
	// IsEquatable describes whether values of this type can be compared with eachother using the `==` operator.
	IsEquatable() bool
	// ToString returns the textual representation of the type.
	// The string is also a valid way to write down the type in the language itself.
	ToString() string
	// SubstituteRoles substitutes all roles in the type (recursively) using the given role substitiution map.
	SubstituteRoles(substMap *RoleSubst) Type
	// ReplaceSharedRoles replaces all roles if the type has shared roles.
	// If the roles are distributed, then this method is a no-op.
	ReplaceSharedRoles(participants []string) Type
	// CoerceTo attempts to coerce this type to the `other` type, and returns the coerced type and whether the coercion was successful.
	CoerceTo(other Type) (Type, bool)
	// Roles returns the roles of the type
	Roles() *Roles
	// IsInvalid returns whether this type is the invalid type.
	IsInvalid() bool
	// Fields returns an iterator over all fields that can be accessed on this type.
	Fields() iter.Seq2[string, Type]
	// Field returns the field with the given name, the returned boolean indicates whether the field was found.
	Field(name string) (Type, bool)
}

type baseType struct{}

func (*baseType) IsInvalid() bool {
	return false
}

func (*baseType) Fields() iter.Seq2[string, Type] {
	return func(yield func(string, Type) bool) {}
}

func (*baseType) Field(name string) (Type, bool) {
	return Invalid(), false
}

type InvalidType struct {
	baseType
}

func (*InvalidType) IsInvalid() bool {
	return true
}

func (t *InvalidType) SubstituteRoles(substMap *RoleSubst) Type {
	return t
}

func (t *InvalidType) ReplaceSharedRoles(participants []string) Type {
	return t
}

func (t *InvalidType) CoerceTo(other Type) (Type, bool) {
	return other, true
}

func (t *InvalidType) Roles() *Roles {
	return EveryoneRole()
}

var invalid_type InvalidType = InvalidType{}

func (t *InvalidType) IsSendable() bool {
	return true
}

func (t *InvalidType) IsEquatable() bool {
	return true
}

func (t *InvalidType) ToString() string {
	return "ERROR"
}

// Invalid returns the special invalid type, if a program has no type errors, then this type will not exist.
// However, this type is introduced when an expression produces a type error.
// The invalid type silences all further type errors, since it can coerce (and be coerced) to all other types.
func Invalid() Type {
	return &invalid_type
}

type UnitType struct {
	baseType
}

func (u *UnitType) SubstituteRoles(substMap *RoleSubst) Type {
	return u
}

func (u *UnitType) ReplaceSharedRoles(participants []string) Type {
	return u
}

func (t *UnitType) CoerceTo(other Type) (Type, bool) {
	return Unit(), other == Unit()
}

func (t *UnitType) Roles() *Roles {
	return EveryoneRole()
}

func (u *UnitType) ToString() string {
	return "()"
}

func (u *UnitType) IsSendable() bool {
	return true
}

func (t *UnitType) IsEquatable() bool {
	return false
}

var unit_value UnitType = UnitType{}

// Unit returns the unit type. It is the type of a call to a function that does not return a value.
// This type cannot be assigned to anything.
func Unit() Type {
	return &unit_value
}
