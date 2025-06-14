package codegen_go

import (
	"github.com/dave/jennifer/jen"
	"github.com/tempo-lang/tempo/projection"
)

func GenChoreography(file *jen.File, c *projection.Choreography) {
	file.Commentf("Projection of choreography %s", c.Name)

	for _, role := range c.Roles {
		file.Add(GenFunc(c.Funcs[role]))
	}
}
