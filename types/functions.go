package types

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/parser"
)

type FunctionType struct {
	baseType
	ident      parser.IIdentContext
	params     []Type
	returnType Type
	roles      *Roles
}

func (f *FunctionType) SubstituteRoles(substMap *RoleSubst) Type {
	substParams := []Type{}
	for _, p := range f.params {
		substParams = append(substParams, p.SubstituteRoles(substMap))
	}

	return Function(
		f.ident,
		substParams,
		f.ReturnType().SubstituteRoles(substMap),
		f.Roles().SubstituteRoles(substMap),
	)
}

func (f *FunctionType) ReplaceSharedRoles(participants []string) Type {
	return f
}

func (f *FunctionType) CoerceTo(other Type) (Type, bool) {
	if value, ok := baseCoerceValue(f, other); ok != nil {
		return value, *ok
	}

	if closure, ok := other.(*ClosureType); ok {
		thisClosure := Closure(f.Params(), f.ReturnType(), f.Roles().Participants())
		return thisClosure.CoerceTo(closure)
	}

	g, ok := other.(*FunctionType)
	if !ok {
		return Invalid(), false
	}

	if f.ident.GetText() != g.ident.GetText() {
		return Invalid(), false
	}

	if len(f.params) != len(g.params) {
		return Invalid(), false
	}

	if !f.Roles().Equals(g.Roles()) {
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

	newFunc := Function(f.ident, newParams, newReturn, f.roles)

	return newFunc, canCoerce
}

func (f *FunctionType) Roles() *Roles {
	return f.roles
}

func (f *FunctionType) IsSendable() bool {
	return false
}

func (t *FunctionType) IsEquatable() bool {
	return false
}

func (f *FunctionType) ToString() string {
	params := misc.JoinStringsFunc(f.params, ", ", func(param Type) string { return param.ToString() })
	returnType := ""
	if f.returnType != Unit() {
		returnType = " " + f.returnType.ToString()
	}
	return fmt.Sprintf("func@%s %s(%s)%s", f.Roles().ToString(), f.ident.GetText(), params, returnType)
}

func (f *FunctionType) Params() []Type {
	return f.params
}

func (f *FunctionType) ReturnType() Type {
	return f.returnType
}

func (f *FunctionType) NameIdent() parser.IIdentContext {
	return f.ident
}

func Function(ident parser.IIdentContext, params []Type, returnType Type, roles *Roles) Type {
	return &FunctionType{
		ident:      ident,
		params:     params,
		returnType: returnType,
		roles:      roles,
	}
}
