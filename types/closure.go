package types

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
)

type ClosureType struct {
	baseType
	params       []Type
	returnType   Type
	participants []string
}

func (f *ClosureType) CoerceTo(other Type) (Type, bool) {
	if value, ok := baseCoerceValue(f, other); ok != nil {
		return value, *ok
	}

	g, ok := other.(*ClosureType)
	if !ok {
		return Invalid(), false
	}

	if len(f.params) != len(g.params) {
		return Invalid(), false
	}

	canCoerce := true
	newParams := []Type{}
	for i := range f.params {
		if newParam, ok := f.params[i].CoerceTo(g.params[i]); ok {
			newParams = append(newParams, newParam)
		} else {
			newParams = append(newParams, Invalid())
			canCoerce = false
		}
	}

	newReturn, ok := f.returnType.CoerceTo(g.returnType)
	if !ok {
		canCoerce = false
	}

	return Closure(newParams, newReturn, other.Roles().participants), canCoerce
}

func (c *ClosureType) IsEquatable() bool {
	return false
}

func (c *ClosureType) IsSendable() bool {
	return false
}

func (c *ClosureType) SubstituteRoles(substMap *RoleSubst) Type {
	substParams := []Type{}
	for _, p := range c.params {
		substParams = append(substParams, p.SubstituteRoles(substMap))
	}

	return Closure(
		substParams,
		c.returnType.SubstituteRoles(substMap),
		c.Roles().SubstituteRoles(substMap).participants,
	)
}

func (c *ClosureType) ReplaceSharedRoles(participants []string) Type {
	return c
}

func (c *ClosureType) Roles() *Roles {
	return NewRole(c.participants, false)
}

func (c *ClosureType) ToString() string {
	params := misc.JoinStringsFunc(c.params, ", ", func(param Type) string { return param.ToString() })
	returnType := ""
	if c.returnType != Unit() {
		returnType = c.returnType.ToString()
	}
	return fmt.Sprintf("func@%s(%s)%s", c.Roles().ToString(), params, returnType)
}

func Closure(params []Type, returnType Type, participants []string) *ClosureType {
	return &ClosureType{
		params:       params,
		returnType:   returnType,
		participants: participants,
	}
}

func (c *ClosureType) Params() []Type {
	return c.params
}

func (c *ClosureType) ReturnType() Type {
	return c.returnType
}
