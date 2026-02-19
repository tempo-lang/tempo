package codegen_java

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
	out.WriteString(gen.Writeln("public interface %s {", inf.InterfaceName()))
	gen.IncIndent()

	for _, method := range inf.Methods {
		params := gen.GenFuncParams(method.FuncSig)
		returnType := gen.GenType(method.ReturnValue)
		out.WriteString(gen.Writeln("public %s %s(%s) throws Exception;", returnType, method.Name, params))
	}

	gen.DecIndent()
	out.WriteString(gen.Writeln("}"))
	return out.String()
}
