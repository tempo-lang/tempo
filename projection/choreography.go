// This module includes the data structures that make up choreography projections.
// The data structures are used for generating the projected code.
//
// The [SourceFile] is contains the root of a choreography projection.
package projection

import (
	"github.com/tempo-lang/tempo/parser"
)

// Choreography represents a Tempo choreography from the perspective of each projected role.
type Choreography struct {
	Name  string
	Roles []string
	Funcs map[string]*Func
}

func NewChoreography(name string) *Choreography {
	return &Choreography{
		Name:  name,
		Roles: []string{},
		Funcs: map[string]*Func{},
	}
}

func (c *Choreography) AddFunc(sig *FuncSig, funcCtx parser.IFuncContext) *Func {
	c.Roles = append(c.Roles, sig.Role)
	c.Funcs[sig.Role] = &Func{
		FuncSig:      sig,
		Choreography: c,
		FuncCtx:      funcCtx,
		Body:         []Statement{},
	}
	return c.Funcs[sig.Role]
}
