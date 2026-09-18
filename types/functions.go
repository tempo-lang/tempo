package types

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/parser/ast"
)

type FunctionType struct {
	baseType
	fnSig      *ast.FuncSig
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
		f.fnSig,
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
		return f.ToClosure().CoerceTo(closure)
	}

	g, ok := other.(*FunctionType)
	if !ok {
		return Invalid(), false
	}

	if f.NameIdent().Value() != g.NameIdent().Value() {
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

	newFunc := Function(f.fnSig, newParams, newReturn, f.roles)

	return newFunc, canCoerce
}

func (f *FunctionType) Roles() *Roles {
	return f.roles
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

	if f.Roles().IsUnnamedRole() {
		return fmt.Sprintf("func %s(%s)%s", f.NameIdent().Value(), params, returnType)
	} else {
		return fmt.Sprintf("func@%s %s(%s)%s", f.Roles().ToString(), f.NameIdent().Value(), params, returnType)
	}
}

// FormatFunctionSig returns a valid function signature for this function type
func (f *FunctionType) FormatFunctionSig() string {
	allParams := f.fnSig.Params.Params
	params := []string{}
	for i, paramType := range f.params {
		paramName := allParams[i].Name.Value()
		params = append(params, fmt.Sprintf("%s: %s", paramName, paramType.ToString()))
	}

	paramsStr := misc.JoinStrings(params, ", ")

	returnType := ""
	if f.returnType != Unit() {
		returnType = " " + f.returnType.ToString()
	}

	if f.Roles().IsUnnamedRole() {
		return fmt.Sprintf("func %s(%s)%s", f.NameIdent().Value(), paramsStr, returnType)
	} else {
		return fmt.Sprintf("func@%s %s(%s)%s", f.Roles().ToString(), f.NameIdent().Value(), paramsStr, returnType)
	}
}

func (f *FunctionType) Params() []Type {
	return f.params
}

func (f *FunctionType) ReturnType() Type {
	return f.returnType
}

func (f *FunctionType) NameIdent() *ast.Identifier {
	return f.fnSig.Name
}

func (f *FunctionType) FuncSig() *ast.FuncSig {
	return f.fnSig
}

func (f *FunctionType) ToClosure() Type {
	return Closure(f.Params(), f.ReturnType(), f.Roles())
}

func Function(fnSig *ast.FuncSig, params []Type, returnType Type, roles *Roles) Type {
	return &FunctionType{
		fnSig:      fnSig,
		params:     params,
		returnType: returnType,
		roles:      roles,
	}
}
