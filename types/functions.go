package types

import (
	"fmt"
	"slices"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/parser"
)

type FunctionType struct {
	ident      parser.IIdentContext
	params     []*Type
	returnType *Type
	roles      []string
}

func (f *FunctionType) SubstituteRoles(substMap *RoleSubst) Value {
	substParams := []*Type{}
	for _, p := range f.params {
		substParams = append(substParams, p.SubstituteRoles(substMap))
	}

	newRoles := []string{}
	for _, r := range f.roles {
		newRoles = append(newRoles, substMap.Subst(r))
	}

	return Function(
		f.ident,
		substParams,
		f.returnType.SubstituteRoles(substMap),
		newRoles,
	)
}

func (f *FunctionType) CoerceTo(other Value) (Value, bool) {
	if value, ok := baseCoerceValue(f, other); ok {
		return value, true
	}

	if closure, ok := other.(*ClosureType); ok {
		thisClosure := Closure(f.Params(), f.ReturnType())
		return thisClosure.CoerceTo(closure)
	}

	g, ok := other.(*FunctionType)
	if !ok {
		return Invalid().Value(), false
	}

	if f.ident != g.ident {
		return Invalid().Value(), false
	}

	if len(f.params) != len(g.params) {
		return Invalid().Value(), false
	}

	if !slices.Equal(f.roles, g.roles) {
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

	newFunc := Function(f.ident, newParams, newReturn, f.roles)

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
	return fmt.Sprintf("func %s(%s)%s", f.ident.GetText(), params, returnType)
}

func (f *FunctionType) IsValue()    {}
func (f *FunctionType) IsFunction() {}

func (f *FunctionType) Params() []*Type {
	return f.params
}

func (f *FunctionType) ReturnType() *Type {
	return f.returnType
}

func (f *FunctionType) NameIdent() parser.IIdentContext {
	return f.ident
}

func (f *FunctionType) Roles() *Roles {
	return NewRole(f.roles, false)
}

func Function(ident parser.IIdentContext, params []*Type, returnType *Type, roles []string) Value {
	return &FunctionType{
		ident:      ident,
		params:     params,
		returnType: returnType,
		roles:      roles,
	}
}
