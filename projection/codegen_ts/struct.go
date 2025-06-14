package codegen_ts

import (
	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenChoreographyStruct(c *projection.ChoreographyStruct) {
	gen.Writeln("// Projection of struct %s", c.Name)

	for _, role := range c.Roles {
		gen.GenStruct(c.Structs[role])
	}

	gen.Writeln("")
}

func (gen *codegen) GenStruct(s *projection.Struct) {
	gen.Writeln("export type %s_%s = {", s.Name, s.Role)
	gen.IncIndent()

	for _, field := range s.Fields {
		gen.Writeln("%s: %s;", field.Name, gen.GenType(field.Type))
	}

	gen.DecIndent()
	gen.Writeln("}")
}
