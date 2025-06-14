package codegen_ts

import "github.com/tempo-lang/tempo/projection"

func (gen *codegen) GenChoreography(c *projection.Choreography) {
	gen.Writeln("// Projection of choreography %s", c.Name)

	for _, role := range c.Roles {
		gen.GenFunc(c.Funcs[role])
	}

	gen.Writeln("")
}
