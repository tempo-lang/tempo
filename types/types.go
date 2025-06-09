package types

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

func MergeTypes(a, b Type) (Type, bool) {
	if a.Roles().IsDistributedRole() && !b.Roles().IsDistributedRole() {
		return Invalid(), false
	}

	if b.Roles().IsDistributedRole() && !a.Roles().IsDistributedRole() {
		return Invalid(), false
	}

	intersectRoles, ok := RoleIntersect(a.Roles(), b.Roles())
	if !ok {
		return Invalid(), false
	}

	intersectParticipants := intersectRoles.participants

	newA := a.ReplaceSharedRoles(intersectParticipants)
	newB := b.ReplaceSharedRoles(intersectParticipants)

	if result, ok := newA.CoerceTo(newB); ok {
		return result, true
	}

	if result, ok := newB.CoerceTo(newA); ok {
		return result, true
	}

	return Invalid(), false
}

type Type interface {
	IsSendable() bool
	IsEquatable() bool
	ToString() string
	IsValue()
	SubstituteRoles(substMap *RoleSubst) Type
	ReplaceSharedRoles(participants []string) Type
	CoerceTo(other Type) (Type, bool)
	Roles() *Roles
	IsInvalid() bool
}

type baseType struct{}

func (*baseType) IsValue() {}

func (*baseType) IsInvalid() bool {
	return false
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

func Unit() Type {
	return &unit_value
}
