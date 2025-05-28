package types

import (
	"fmt"
	"tempo/misc"
	"tempo/parser"
)

type FunctionType struct {
	funcIdent  parser.IIdentContext
	params     []*Type
	returnType *Type
	roleSubst  map[string]string
}

func (f *FunctionType) SubstituteRoles(substMap map[string]string) Value {
	substParams := []*Type{}
	for _, p := range f.params {
		substParams = append(substParams, p.SubstituteRoles(substMap))
	}

	newRoleSubst := map[string]string{}
	for from, to := range f.roleSubst {
		newRoleSubst[from] = substMap[to]
	}

	return &FunctionType{
		funcIdent:  f.funcIdent,
		params:     substParams,
		returnType: f.returnType.SubstituteRoles(substMap),
		roleSubst:  newRoleSubst,
	}
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

func (f *FunctionType) FuncIdent() parser.IIdentContext {
	return f.funcIdent
}

func (f *FunctionType) RoleSubstitution() map[string]string {
	return f.roleSubst
}

func Function(funcIdent parser.IIdentContext, params []*Type, returnType *Type, roles []string) Value {
	roleSubst := map[string]string{}
	for _, role := range roles {
		roleSubst[role] = role
	}

	return &FunctionType{
		funcIdent:  funcIdent,
		params:     params,
		returnType: returnType,
		roleSubst:  roleSubst,
	}
}
