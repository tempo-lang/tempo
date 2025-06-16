package codegen_ts

import (
	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenChoreographyStruct(c *projection.ChoreographyStruct) string {
	out := gen.Writeln("// Projection of struct %s", c.Name)

	for _, role := range c.Roles {
		out += gen.GenStruct(c.Structs[role])
	}

	out += gen.Writeln("")
	return out
}

func (gen *codegen) GenStruct(s *projection.Struct) string {
	out := gen.Writeln("export type %s_%s = {", s.Name, s.Role)
	gen.IncIndent()

	for _, field := range s.Fields {
		out += gen.Writeln("%s: %s;", field.Name, gen.GenType(field.Type))
	}

	gen.DecIndent()
	out += gen.Writeln("}")
	return out
}
