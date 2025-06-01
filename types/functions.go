package types

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
)

type FunctionType struct {
	params       []*Type
	returnType   *Type
	roleSubst    *RoleSubst
	instantiated bool
}

func (f *FunctionType) SubstituteRoles(substMap *RoleSubst) Value {
	if f.instantiated {
		panic("function already instantiated")
	}

	substParams := []*Type{}
	for _, p := range f.params {
		substParams = append(substParams, p.SubstituteRoles(substMap))
	}

	newRoleSubst := NewRoleSubst()
	for _, from := range f.roleSubst.Roles {
		newRoleSubst.AddRole(from, substMap.Map[from])
	}

	return &FunctionType{
		params:       substParams,
		returnType:   f.returnType.SubstituteRoles(substMap),
		roleSubst:    newRoleSubst,
		instantiated: true,
	}
}

func (f *FunctionType) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(f, other); ok {
		return value, true
	}

	g, ok := other.(*FunctionType)
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

	newFunc := &FunctionType{
		params:       newParams,
		returnType:   newReturn,
		roleSubst:    f.roleSubst,
		instantiated: f.instantiated,
	}

	return newFunc, canCoerce
}

func (f *FunctionType) IsSendable() bool {
	return false
}

func (t *FunctionType) IsEquatable() bool {
	return false
}

func (f *FunctionType) ToString() string {
	params := misc.JoinStringsFunc(f.params, ", ", func(param *Type) string { return param.ToString() })
	returnType := ""
	if f.returnType.Value() != Unit() {
		returnType = f.returnType.ToString()
	}
	return fmt.Sprintf("func(%s)%s", params, returnType)
}

func (f *FunctionType) IsValue()    {}
func (f *FunctionType) IsFunction() {}

func (f *FunctionType) Params() []*Type {
	return f.params
}

func (f *FunctionType) ReturnType() *Type {
	return f.returnType
}

func (f *FunctionType) IsInstantiated() bool {
	return f.instantiated
}

func (f *FunctionType) RoleSubstitution() *RoleSubst {
	return f.roleSubst
}

func Function(params []*Type, returnType *Type, roles []string, instantiated bool) Value {
	roleSubst := NewRoleSubst()
	for _, role := range roles {
		roleSubst.AddRole(role, role)
	}

	return &FunctionType{
		params:       params,
		returnType:   returnType,
		roleSubst:    roleSubst,
		instantiated: instantiated,
	}
}
