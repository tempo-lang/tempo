package codegen_go

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/tempo-lang/tempo/projection"
)

func GenChoreographyStruct(file *jen.File, c *projection.ChoreographyStruct) {
	file.Commentf("Projection of struct %s", c.Name)

	for _, role := range c.Roles {
		file.Add(GenStruct(c.Structs[role]))
	}
}

func GenStruct(s *projection.Struct) *jen.Statement {

	fields := []jen.Code{}

	for _, field := range s.Fields {
		fields = append(fields, GenStructField(&field))
	}

	return jen.Type().Id(fmt.Sprintf("%s_%s", s.Name, s.Role)).Struct(fields...)
}

func GenStructType(s *projection.StructType) jen.Code {
	return jen.Id(fmt.Sprintf("%s_%s", s.Name(), s.Role()))
}

func GenStructField(f *projection.StructField) *jen.Statement {
	return jen.Id(f.Name).Add(GenType(f.Type)).Tag(map[string]string{"json": f.Name})
}
