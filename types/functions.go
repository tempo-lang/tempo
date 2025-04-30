package types

import "fmt"

type FunctionType struct {
	// returnType Type
	params []Type
}

func (f *FunctionType) ToString() string {

	params := "("
	for _, param := range f.params {
		params += param.ToString() + ", "
	}
	params = params[:len(params)-2]
	params += ")"

	return fmt.Sprintf("func (%s)", params)
}

func (f *FunctionType) IsType()     {}
func (f *FunctionType) IsFunction() {}

// func (f *FunctionType) ReturnType() Type
func (f *FunctionType) Params() []Type {
	return f.params
}

func Function(params []Type) *FunctionType {
	return &FunctionType{
		// returnType: returnType,
		params: params,
	}
}
