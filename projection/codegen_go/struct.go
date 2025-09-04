package codegen_go

import (
	"github.com/dave/jennifer/jen"
	"github.com/tempo-lang/tempo/projection"
)

func GenChoreographyStruct(file *jen.File, c *projection.ChoreographyStruct) {
	file.Commentf("Projection of struct `%s`", c.Name)

	hasMethods := false
	for _, role := range c.Roles {
		file.Add(GenStructDecl(c.Structs[role]))

		if len(c.Structs[role].Methods) > 0 {
			hasMethods = true
		}
	}

	if hasMethods {
		file.Commentf("Implementation of struct `%s`", c.Name)
		for _, role := range c.Roles {
			for _, stmt := range GenStructMethods(c.Structs[role]) {
				file.Add(stmt)
			}
		}
	}
}

func GenStructDecl(s *projection.Struct) *jen.Statement {

	fields := []jen.Code{}

	for _, field := range s.Fields {
		fields = append(fields, GenStructField(&field))
	}

	return jen.Type().Id(s.StructName()).Struct(fields...)
}

func GenStructType(s *projection.StructType) jen.Code {
	return jen.Id(s.StructName())
}

func GenStructField(f *projection.StructField) *jen.Statement {
	return jen.Id(f.Name).Add(GenType(f.Type)).Tag(map[string]string{"json": f.Name})
}

func GenStructMethods(s *projection.Struct) []*jen.Statement {
	result := []*jen.Statement{}

	for _, method := range s.Methods {
		params := FuncSigParams(method.FuncSig)

		fn := jen.Func().Params(jen.Id("self").Id(s.StructName())).Id(method.FuncSig.Name).Params(params...)

		if method.FuncSig.ReturnValue != projection.UnitType() {
			fn = fn.Add(GenType(method.FuncSig.ReturnValue))
		}

		fn = fn.BlockFunc(func(block *jen.Group) {
			for _, bodyStmt := range method.Body {
				for _, stmt := range GenStatement(bodyStmt) {
					block.Add(stmt)
				}
			}
		})

		result = append(result, fn)
	}

	return result
}
