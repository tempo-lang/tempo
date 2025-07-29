package codegen_ts

import "github.com/tempo-lang/tempo/projection"

func (gen *codegen) GenChoreography(c *projection.Choreography) string {
	out := gen.Writeln("// Projection of choreography `%s`", c.Name)

	for _, role := range c.Roles {
		out += gen.GenFunc(c.Funcs[role])
	}

	out += gen.Writeln("")
	return out
}
