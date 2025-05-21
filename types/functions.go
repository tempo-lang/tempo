package types

import (
	"fmt"
	"tempo/misc"
)

type FunctionType struct {
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
	return fmt.Sprintf("func(%s)", params)
}

func (f *FunctionType) IsValue()    {}
func (f *FunctionType) IsFunction() {}

func (f *FunctionType) Params() []*Type {
	return f.params
}

func (f *FunctionType) ReturnType() *Type {
	return f.returnType
}

func Function(params []*Type, returnType *Type) Value {
	return &FunctionType{
		params:     params,
		returnType: returnType,
	}
}
