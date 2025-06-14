package codegen_ts

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenChoreographyInterface(c *projection.ChoreographyInterface) {
	gen.Writeln("// Projection of interface %s", c.Name)

	for _, role := range c.Roles {
		gen.GenInterface(c.Interfaces[role])
	}

	gen.Writeln("")
}

func (gen *codegen) GenInterface(inf *projection.Interface) {
	gen.Writeln("interface %s {", inf.GenName())
	gen.IncIndent()

	for _, method := range inf.Methods {

		params := []string{}
		for _, param := range method.Params {
			if gen.opts.DisableTypes {
				params = append(params, param.Name)
			} else {
				params = append(params, fmt.Sprintf("%s: %s", param.Name, gen.GenType(param.TypeValue)))
			}
		}

		returnType := gen.GenType(method.ReturnValue)

		gen.Writeln("%s(%s): %s;", method.Name, misc.JoinStrings(params, ", "), returnType)
	}

	gen.DecIndent()
	gen.Writeln("}")
}
