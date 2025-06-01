package types

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
)

type ClosureType struct {
	params     []*Type
	returnType *Type
}

func (f *ClosureType) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(f, other); ok {
		return value, true
	}

	g, ok := other.(*ClosureType)
	if !ok {
		return Invalid().Value(), false
	}

	if len(f.params) != len(g.params) {
		return Invalid().Value(), false
	}

	canCoerce := true
	newParams := []*Type{}
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

	return Closure(newParams, newReturn), canCoerce
}

func (c *ClosureType) IsEquatable() bool {
	return false
}

func (c *ClosureType) IsSendable() bool {
	return false
}

func (c *ClosureType) IsValue() {}

func (c *ClosureType) SubstituteRoles(substMap *RoleSubst) Value {
	substParams := []*Type{}
	for _, p := range c.params {
		substParams = append(substParams, p.SubstituteRoles(substMap))
	}

	return Closure(
		substParams,
		c.returnType.SubstituteRoles(substMap),
	)
}

func (c *ClosureType) ToString() string {
	params := misc.JoinStringsFunc(c.params, ", ", func(param *Type) string { return param.ToString() })
	returnType := ""
	if c.returnType.Value() != Unit() {
		returnType = c.returnType.ToString()
	}
	return fmt.Sprintf("CLOSURE func(%s)%s", params, returnType)
}

func Closure(params []*Type, returnType *Type) *ClosureType {
	return &ClosureType{
		params:     params,
		returnType: returnType,
	}
}

func (c *ClosureType) Params() []*Type {
	return c.params
}

func (c *ClosureType) ReturnType() *Type {
	return c.returnType
}
