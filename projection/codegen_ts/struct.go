package codegen_ts

import (
	"fmt"

	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenChoreographyStruct(c *projection.ChoreographyStruct) string {
	out := gen.Writeln("// Projection of struct `%s`", c.Name)

	hasMethods := false
	for _, role := range c.Roles {
		out += gen.GenStructDecl(c.Structs[role])

		if len(c.Structs[role].Methods) > 0 {
			hasMethods = true
		}
	}

	out += gen.Writeln("")

	if hasMethods {
		out += gen.Writeln("// Implementation of struct `%s`", c.Name)

		for _, role := range c.Roles {
			out += gen.GenStructMethods(c.Structs[role])
		}

		out += gen.Writeln("")
	}

	return out
}

func (gen *codegen) GenStructDecl(s *projection.Struct) string {
	out := gen.Writeln("export type %s_%s = {", s.Name, s.Role)
	gen.IncIndent()

	for _, field := range s.Fields {
		out += gen.Writeln("%s: %s;", field.Name, gen.GenType(field.Type))
	}

	gen.DecIndent()
	out += gen.Writeln("}")
	return out
}

func (gen *codegen) GenStructMethods(s *projection.Struct) string {
	out := ""

	out += gen.Writeln("function %s_%s_methods(self: %s_%s) {", s.Name, s.Role, s.Name, s.Role)
	gen.IncIndent()
	out += gen.Writeln("return {")
	gen.IncIndent()
	for _, method := range s.Methods {
		params := gen.GenFuncParams(method.FuncSig)
		fnSig := fmt.Sprintf("async %s(%s)", method.FuncSig.Name, params)
		if method.FuncSig.ReturnValue != projection.UnitType() && !gen.opts.DisableTypes {
			fnSig += fmt.Sprintf(": Promise<%s>", gen.GenType(method.FuncSig.ReturnValue))
		}
		out += gen.Writeln("%s {", fnSig)
		gen.IncIndent()
		for _, stmt := range method.Body {
			out += gen.GenStmt(stmt)
		}
		gen.DecIndent()
		out += gen.Writeln("},")
	}
	gen.DecIndent()
	out += gen.Writeln("};")
	gen.DecIndent()
	out += gen.Writeln("}")

	return out
}
