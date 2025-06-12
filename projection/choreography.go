// This module includes the data structures that make up choreography projections.
// Is is also responsible for generating the projected code based on these data structures.
package projection

import (
	"github.com/tempo-lang/tempo/parser"

	"github.com/dave/jennifer/jen"
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

func (c *Choreography) Codegen(file *jen.File) {
	file.Commentf("Projection of choreography %s", c.Name)

	for _, role := range c.Roles {
		file.Add(c.Funcs[role].Codegen())
	}
}
