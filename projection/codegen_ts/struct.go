package codegen_ts

import (
	"fmt"
	"strings"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenChoreographyStruct(c *projection.ChoreographyStruct) string {
	var out strings.Builder
	out.WriteString(gen.Writeln("// Projection of struct `%s`", c.Name))

	for _, role := range c.Roles {
		out.WriteString(gen.GenStructAttrs(c.Structs[role]))
		out.WriteString(gen.GenStructDecl(c.Structs[role]))
		out.WriteString(gen.Writeln(""))
	}

	out.WriteString(gen.Writeln(""))

	return out.String()
}

func structAttrs(s *projection.Struct) string {
	return fmt.Sprintf("%s_attrs", s.StructName())
}

func (gen *codegen) GenStructAttrs(s *projection.Struct) string {
	if gen.opts.DisableTypes {
		return ""
	}

	var out strings.Builder

	out.WriteString(gen.Writeln("export interface %s {", structAttrs(s)))
	gen.IncIndent()

	for _, field := range s.Fields {
		out.WriteString(gen.Writeln("%s: %s;", field.Name, gen.GenType(field.Type)))
	}

	gen.DecIndent()
	out.WriteString(gen.Writeln("}"))

	return out.String()
}

func (gen *codegen) GenStructDecl(s *projection.Struct) string {
	var out strings.Builder

	if gen.opts.DisableTypes {
		out.WriteString(gen.Writeln("export class %s {", s.StructName()))
	} else {
		impls := []string{structAttrs(s)}
		for _, impl := range s.Implements {
			impls = append(impls, gen.GenType(impl))
		}

		out.WriteString(gen.Writeln("export class %s implements %s {", s.StructName(), misc.JoinStrings(impls, ", ")))
	}
	gen.IncIndent()

	if !gen.opts.DisableTypes && len(s.Fields) > 0 {
		for _, field := range s.Fields {
			out.WriteString(gen.Writeln("%s: %s;", field.Name, gen.GenType(field.Type)))
		}
		out.WriteString(gen.Writeln(""))
	}

	fields := misc.JoinStringsFunc(s.Fields, ", ", func(field projection.StructField) string {
		return field.Name
	})

	if gen.opts.DisableTypes {
		out.WriteString(gen.Writeln("constructor({ %s }) {", fields))
	} else {
		out.WriteString(gen.Writeln("constructor({ %s }: %s) {", fields, structAttrs(s)))
	}

	gen.IncIndent()
	for _, field := range s.Fields {
		out.WriteString(gen.Writeln("this.%s = %s;", field.Name, field.Name))
	}
	gen.DecIndent()
	out.WriteString(gen.Writeln("}"))

	out.WriteString(gen.GenStructMethods(s))

	gen.DecIndent()
	out.WriteString(gen.Writeln("}"))
	return out.String()
}

func (gen *codegen) GenStructMethods(s *projection.Struct) string {
	var out strings.Builder

	for _, method := range s.Methods {
		out.WriteString(gen.Writeln(""))
		params := gen.GenFuncParams(method.FuncSig)
		fnSig := fmt.Sprintf("async %s(%s)", method.FuncSig.Name, params)
		if method.FuncSig.ReturnValue != projection.UnitType() && !gen.opts.DisableTypes {
			fnSig += fmt.Sprintf(": Promise<%s>", gen.GenType(method.FuncSig.ReturnValue))
		}
		out.WriteString(gen.Writeln("%s {", fnSig))
		gen.IncIndent()
		for _, stmt := range method.Body {
			out.WriteString(gen.GenStmt(stmt))
		}
		gen.DecIndent()
		out.WriteString(gen.Writeln("}"))
	}

	return out.String()
}
