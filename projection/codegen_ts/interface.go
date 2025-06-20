package codegen_ts

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenChoreographyInterface(c *projection.ChoreographyInterface) string {
	out := gen.Writeln("// Projection of interface %s", c.Name)

	for _, role := range c.Roles {
		out += gen.GenInterface(c.Interfaces[role])
	}

	out += gen.Writeln("")
	return out
}

func (gen *codegen) GenInterface(inf *projection.Interface) string {
	out := gen.Writeln("export interface %s {", inf.GenName())
	gen.IncIndent()

	for _, method := range inf.Methods {

		params := []string{"env: Env"}
		for _, param := range method.Params {
			if gen.opts.DisableTypes {
				params = append(params, param.Name)
			} else {
				params = append(params, fmt.Sprintf("%s: %s", param.Name, gen.GenType(param.TypeValue)))
			}
		}

		returnType := gen.GenType(method.ReturnValue)

		out += gen.Writeln("%s(%s): Promise<%s>;", method.Name, misc.JoinStrings(params, ", "), returnType)
	}

	gen.DecIndent()
	out += gen.Writeln("}")
	return out
}
