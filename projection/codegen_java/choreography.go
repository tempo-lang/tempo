package codegen_java

import "strings"

import "github.com/tempo-lang/tempo/projection"

func (gen *codegen) GenChoreography(c *projection.Choreography) string {
	var out strings.Builder
	out.WriteString(gen.Writeln(""))
	out.WriteString(gen.Writeln("// Projection of choreography `%s`", c.Name))

	for _, role := range c.Roles {
		out.WriteString(gen.GenFunc(c.Funcs[role]))
	}

	return out.String()
}
