package codegen_ts

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenChoreographyStruct(c *projection.ChoreographyStruct) string {
	out := ""
	out += gen.Writeln("// Projection of struct `%s`", c.Name)

	for _, role := range c.Roles {
		out += gen.GenStructAttrs(c.Structs[role])
		out += gen.GenStructDecl(c.Structs[role])
		out += gen.Writeln("")
	}

	out += gen.Writeln("")

	return out
}

func structAttrs(s *projection.Struct) string {
	return fmt.Sprintf("%s_attrs", s.StructName())
}

func (gen *codegen) GenStructAttrs(s *projection.Struct) string {
	if gen.opts.DisableTypes {
		return ""
	}

	out := ""

	out += gen.Writeln("export interface %s {", structAttrs(s))
	gen.IncIndent()

	for _, field := range s.Fields {
		out += gen.Writeln("%s: %s;", field.Name, gen.GenType(field.Type))
	}

	gen.DecIndent()
	out += gen.Writeln("}")

	return out
}

func (gen *codegen) GenStructDecl(s *projection.Struct) string {
	out := ""

	if gen.opts.DisableTypes {
		out += gen.Writeln("export class %s {", s.StructName())
	} else {
		impls := []string{structAttrs(s)}
		for _, impl := range s.Implements {
			impls = append(impls, gen.GenType(impl))
		}

		out += gen.Writeln("export class %s implements %s {", s.StructName(), misc.JoinStrings(impls, ", "))
	}
	gen.IncIndent()

	if !gen.opts.DisableTypes {
		for _, field := range s.Fields {
			out += gen.Writeln("%s: %s;", field.Name, gen.GenType(field.Type))
		}
		out += gen.Writeln("")
	}

	fields := misc.JoinStringsFunc(s.Fields, ", ", func(field projection.StructField) string {
		return field.Name
	})

	if gen.opts.DisableTypes {
		out += gen.Writeln("constructor({ %s }) {", fields)
	} else {
		out += gen.Writeln("constructor({ %s }: %s) {", fields, structAttrs(s))
	}

	gen.IncIndent()
	for _, field := range s.Fields {
		out += gen.Writeln("this.%s = %s;", field.Name, field.Name)
	}
	gen.DecIndent()
	out += gen.Writeln("}")

	out += gen.GenStructMethods(s)

	gen.DecIndent()
	out += gen.Writeln("}")
	return out
}

func (gen *codegen) GenStructMethods(s *projection.Struct) string {
	out := ""

	for _, method := range s.Methods {
		out += gen.Writeln("")
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
		out += gen.Writeln("}")
	}

	return out
}
