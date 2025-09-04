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
	}

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

	out += gen.Writeln("public static final class %s implements %s {", s.StructName(), impls)
	gen.IncIndent()

	if len(s.Fields) > 0 {
		for _, field := range s.Fields {
			out += gen.Writeln("public %s %s;", gen.GenType(field.Type), field.Name)
		}
		out += gen.Writeln("")
	}

	fields := misc.JoinStringsFunc(s.Fields, ", ", func(field projection.StructField) string {
		return fmt.Sprintf("%s %s", gen.GenType(field.Type), field.Name)
	})

	out += gen.Writeln("public %s(%s) {", s.StructName(), fields)

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
	out := ""

	out += gen.GenStructDefaultMethodToString(s)
	out += gen.GenStructDefaultMethodEquals(s)
	out += gen.GenStructDefaultMethodClone(s)

	return out
}

func (gen *codegen) GenStructDefaultMethodToString(s *projection.Struct) string {
	out := gen.Writeln("")

	attrs := misc.JoinStringsFunc(s.Fields, ", ", func(field projection.StructField) string {
		return fmt.Sprintf("%s=\"+this.%s+\"", field.Name, field.Name)
	})

	out += gen.Writeln("@Override")
	out += gen.Writeln("public String toString() {")
	gen.IncIndent()
	out += gen.Writeln("return \"%s[%s]\";", s.StructName(), attrs)
	gen.DecIndent()
	out += gen.Writeln("}")

	return out
}

func (gen *codegen) GenStructDefaultMethodEquals(s *projection.Struct) string {
	out := gen.Writeln("")

	out += gen.Writeln("@Override")
	out += gen.Writeln("public boolean equals(Object o) {")
	gen.IncIndent()
	out += gen.Writeln("if (this == o) return true;")
	out += gen.Writeln("if (o == null) return false;")
	out += gen.Writeln("if (getClass() != o.getClass()) return false;")

	if len(s.Fields) > 0 {
		gen.AddImport(javaPkgObjects)
		attrs := misc.JoinStringsFunc(s.Fields, " && ", func(field projection.StructField) string {
			return fmt.Sprintf("Objects.equals(this.%s, oo.%s)", field.Name, field.Name)
		})

		out += gen.Writeln("%s oo = (%s) o;", s.StructName(), s.StructName())
		out += gen.Writeln("return %s;", attrs)
	} else {
		out += gen.Writeln("return true;")
	}

	gen.DecIndent()
	out += gen.Writeln("}")

	return out
}

func (gen *codegen) GenStructDefaultMethodClone(s *projection.Struct) string {
	out := gen.Writeln("")

	attrs := misc.JoinStringsFunc(s.Fields, ", ", func(field projection.StructField) string {
		return field.Name
	})

	out += gen.Writeln("@Override")
	out += gen.Writeln("public %s clone() {", s.StructName())
	gen.IncIndent()
	out += gen.Writeln("return new %s(%s);", s.StructName(), attrs)
	gen.DecIndent()
	out += gen.Writeln("}")

	return out
}
