package types

import (
	"fmt"
	"tempo/misc"
)

type FunctionType struct {
	// returnType Type
	params []*Type
}

func (f *FunctionType) IsSendable() bool {
	return false
}

func (f *FunctionType) ToString() string {
	params := misc.JoinStringsFunc(f.params, ", ", func(param *Type) string { return param.ToString() })
	return fmt.Sprintf("func(%s)", params)
}

func (f *FunctionType) IsValue()    {}
func (f *FunctionType) IsFunction() {}

// func (f *FunctionType) ReturnType() Type
func (f *FunctionType) Params() []*Type {
	return f.params
}

func Function(params []*Type) Value {
	return &FunctionType{
		// returnType: returnType,
		params: params,
	}
}
