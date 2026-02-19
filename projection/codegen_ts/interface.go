package codegen_ts

import (
	"strings"

	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenChoreographyInterface(c *projection.ChoreographyInterface) string {
	var out strings.Builder
	out.WriteString(gen.Writeln("// Projection of interface `%s`", c.Name))

	for _, role := range c.Roles {
		out.WriteString(gen.GenInterface(c.Interfaces[role]))
	}

	out.WriteString(gen.Writeln(""))
	return out.String()
}

func (gen *codegen) GenInterface(inf *projection.Interface) string {
	var out strings.Builder
	out.WriteString(gen.Writeln("export interface %s {", inf.InterfaceName()))
	gen.IncIndent()

	for _, method := range inf.Methods {
		params := gen.GenFuncParams(method.FuncSig)
		returnType := gen.GenType(method.ReturnValue)
		out.WriteString(gen.Writeln("%s(%s): Promise<%s>;", method.Name, params, returnType))
	}

	gen.DecIndent()
	out.WriteString(gen.Writeln("}"))
	return out.String()
}
