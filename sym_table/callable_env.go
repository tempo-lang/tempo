package sym_table

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"
)

// CallableEnv describes an environment that can be called.
// Such as a [FuncSymbol] or a [ClosureEnv].
type CallableEnv interface {
	ReturnCtx() parser.IValueTypeContext
	Scope() *Scope
	Params() []*FuncParamSymbol
	AddParam(param *FuncParamSymbol)
	ReturnType() types.Type
}

type ClosureEnv struct {
	scope      *Scope
	params     []*FuncParamSymbol
	returnType types.Type
	returnCtx  parser.IValueTypeContext
}

// ReturnCtx implements CallableEnv.
func (f *ClosureEnv) ReturnCtx() parser.IValueTypeContext {
	return f.returnCtx
}

// Params implements CallableEnv.
func (f *ClosureEnv) Params() []*FuncParamSymbol {
	return f.params
}

// AddParam implements CallableEnv.
func (f *ClosureEnv) AddParam(param *FuncParamSymbol) {
	f.params = append(f.params, param)
}

// ReturnType implements CallableEnv.
func (c *ClosureEnv) ReturnType() types.Type {
	return c.returnType
}

// Scope implements CallableEnv.
func (c *ClosureEnv) Scope() *Scope {
	return c.scope
}

// NewClosureEnv constructs a new closure callable environment.
func NewClosureEnv(scope *Scope, returnType types.Type, returnCtx parser.IValueTypeContext) CallableEnv {
	return &ClosureEnv{
		scope:      scope,
		params:     []*FuncParamSymbol{},
		returnType: returnType,
		returnCtx:  returnCtx,
	}
}
