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

func Function(funcIdent parser.IIdentContext, params []*Type, returnType *Type) Value {
	return &FunctionType{
		funcIdent:  funcIdent,
		params:     params,
		returnType: returnType,
	}
}
