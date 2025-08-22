package codegen_java

import (
	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenChoreographyInterface(c *projection.ChoreographyInterface) string {
	out := gen.Writeln("// Projection of interface `%s`", c.Name)

	for _, role := range c.Roles {
		out += gen.GenInterface(c.Interfaces[role])
	}

	out += gen.Writeln("")
	return out
}

func (gen *codegen) GenInterface(inf *projection.Interface) string {
	out := gen.Writeln("public interface %s {", inf.GenName())
	gen.IncIndent()

	for _, method := range inf.Methods {
		params := gen.GenFuncParams(method.FuncSig)
		returnType := gen.GenType(method.ReturnValue)
		out += gen.Writeln("public %s %s(%s) throws Exception;", returnType, method.Name, params)
	}

	gen.DecIndent()
	out += gen.Writeln("}")
	return out
}
