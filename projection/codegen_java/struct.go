package codegen_java

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenChoreographyStruct(c *projection.ChoreographyStruct) string {
	out := ""
	out += gen.Writeln("// Projection of struct `%s`", c.Name)

	for _, role := range c.Roles {
		out += gen.GenStructDecl(c.Structs[role])
		out += gen.Writeln("")
	}

	out += gen.Writeln("")

	return out
}

func (gen *codegen) GenStructDecl(s *projection.Struct) string {
	out := ""

	impls := misc.JoinStringsFunc(s.Implements, ", ", func(impl projection.Type) string {
		return gen.GenType(impl)
	})

	if len(impls) > 0 {
		impls = fmt.Sprintf("Cloneable, %s", impls)
	} else {
		impls = "Cloneable"
	}

	out += gen.Writeln("public static final class %s_%s implements %s {", s.Name, s.Role, impls)
	gen.IncIndent()

	for _, field := range s.Fields {
		out += gen.Writeln("public %s %s;", gen.GenType(field.Type), field.Name)
	}
	out += gen.Writeln("")

	fields := misc.JoinStringsFunc(s.Fields, ", ", func(field projection.StructField) string {
		return fmt.Sprintf("%s %s", gen.GenType(field.Type), field.Name)
	})

	out += gen.Writeln("public %s_%s(%s) {", s.Name, s.Role, fields)

	gen.IncIndent()
	for _, field := range s.Fields {
		out += gen.Writeln("this.%s = %s;", field.Name, field.Name)
	}
	gen.DecIndent()
	out += gen.Writeln("}")

	out += gen.GenStructMethods(s)

	out += gen.GenStructDefaultMethods(s)

	gen.DecIndent()
	out += gen.Writeln("}")
	return out
}

func (gen *codegen) GenStructMethods(s *projection.Struct) string {
	out := ""

	for _, method := range s.Methods {
		out += gen.Writeln("")
		params := gen.GenFuncParams(method.FuncSig)

		out += gen.Writeln("public %s %s(%s) throws Exception {", gen.GenType(method.FuncSig.ReturnValue), method.FuncSig.Name, params)
		gen.IncIndent()
		for _, stmt := range method.Body {
			out += gen.GenStmt(stmt)
		}
		gen.DecIndent()
		out += gen.Writeln("}")
	}

	return out
}

func (gen *codegen) GenStructDefaultMethods(s *projection.Struct) string {
	out := gen.Writeln("")

	attrs := misc.JoinStringsFunc(s.Fields, ", ", func(field projection.StructField) string {
		return field.Name
	})

	out += gen.Writeln("public %s_%s clone() {", s.Name, s.Role)
	gen.IncIndent()
	out += gen.Writeln("return new %s_%s(%s);", s.Name, s.Role, attrs)
	gen.DecIndent()
	out += gen.Writeln("}")

	return out
}
