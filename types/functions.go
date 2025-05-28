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
	roleSubst  *RoleSubst
}

func (f *FunctionType) SubstituteRoles(substMap *RoleSubst) Value {
	substParams := []*Type{}
	for _, p := range f.params {
		substParams = append(substParams, p.SubstituteRoles(substMap))
	}

	newRoleSubst := NewRoleSubst()
	for _, from := range f.roleSubst.Roles {
		newRoleSubst.AddRole(from, substMap.Map[from])
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

func (f *FunctionType) RoleSubstitution() *RoleSubst {
	return f.roleSubst
}

func Function(funcIdent parser.IIdentContext, params []*Type, returnType *Type, roles []string) Value {
	roleSubst := NewRoleSubst()
	for _, role := range roles {
		roleSubst.AddRole(role, role)
	}

	return &FunctionType{
		funcIdent:  funcIdent,
		params:     params,
		returnType: returnType,
		roleSubst:  roleSubst,
	}
}
