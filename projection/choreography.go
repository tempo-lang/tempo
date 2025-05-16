package projection

import (
	"tempo/parser"
	"tempo/types"

	"github.com/dave/jennifer/jen"
)

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

func (c *Choreography) AddFunc(role string, funcCtx parser.IFuncContext, returnValue types.Value) *Func {
	c.Roles = append(c.Roles, role)
	c.Funcs[role] = &Func{
		Choreography: c,
		FuncCtx:      funcCtx,
		Name:         funcCtx.Ident().GetText(),
		Role:         role,
		ReturnValue:  returnValue,
		Body:         []Statement{},
	}
	return c.Funcs[role]
}

func (c *Choreography) Codegen(file *jen.File) {
	file.Commentf("Projection of choreography %s", c.Name)

	for _, role := range c.Roles {
		file.Add(c.Funcs[role].Codegen())
	}
}
