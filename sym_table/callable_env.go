package sym_table

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"
)

type CallableEnv interface {
	ReturnCtx() parser.IValueTypeContext
	Scope() *Scope
	Params() []*FuncParamSymbol
	AddParam(param *FuncParamSymbol)
	ReturnType() types.Value
}

type ClosureEnv struct {
	scope      *Scope
	params     []*FuncParamSymbol
	returnType types.Value
	returnCtx  parser.IValueTypeContext
}

// ReturnCtx implements CallableEnv.
func (f *ClosureEnv) ReturnCtx() parser.IValueTypeContext {
	return f.returnCtx
}

func (f *ClosureEnv) Params() []*FuncParamSymbol {
	return f.params
}

func (f *ClosureEnv) AddParam(param *FuncParamSymbol) {
	f.params = append(f.params, param)
}

// ReturnType implements CallableEnv.
func (c *ClosureEnv) ReturnType() types.Value {
	return c.returnType
}

// Scope implements CallableEnv.
func (c *ClosureEnv) Scope() *Scope {
	return c.scope
}

func NewClosureEnv(scope *Scope, returnType types.Value, returnCtx parser.IValueTypeContext) CallableEnv {
	return &ClosureEnv{
		scope:      scope,
		params:     []*FuncParamSymbol{},
		returnType: returnType,
		returnCtx:  returnCtx,
	}
}
