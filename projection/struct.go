package projection

import (
	"fmt"
	"tempo/parser"
	"tempo/types"

	"github.com/dave/jennifer/jen"
)

type ChoreographyStruct struct {
	Name    string
	Roles   []string
	Structs map[string]*Struct
}

func NewChoreographyStruct(name string) *ChoreographyStruct {
	return &ChoreographyStruct{
		Name:    name,
		Roles:   []string{},
		Structs: map[string]*Struct{},
	}
}

func (c *ChoreographyStruct) Codegen(file *jen.File) {
	file.Commentf("Projection of struct %s", c.Name)

	for _, role := range c.Roles {
		file.Add(c.Structs[role].Codegen())
	}
}

type Struct struct {
	StructCtx parser.IStructContext
	Name      string
	Role      string
	Fields    []StructField
}

func (c *ChoreographyStruct) AddStruct(role string, structCtx parser.IStructContext) *Struct {
	c.Roles = append(c.Roles, role)
	c.Structs[role] = &Struct{
		StructCtx: structCtx,
		Name:      structCtx.Ident().GetText(),
		Role:      role,
		Fields:    []StructField{},
	}
	return c.Structs[role]
}

func (s *Struct) AddField(field parser.IStructFieldContext, fieldType types.Value) {
	s.Fields = append(s.Fields, StructField{
		Struct:   s,
		FieldCtx: field,
		Name:     field.Ident().GetText(),
		Type:     fieldType,
	})
}

func (s *Struct) Codegen() *jen.Statement {

	fields := []jen.Code{}

	for _, field := range s.Fields {
		fields = append(fields, field.Codegen())
	}

	return jen.Type().Id(fmt.Sprintf("%s_%s", s.Name, s.Role)).Struct(fields...)
}

type StructField struct {
	Struct   *Struct
	FieldCtx parser.IStructFieldContext
	Name     string
	Type     types.Value
}

func (f *StructField) Codegen() *jen.Statement {
	return jen.Id(f.Name).Add(CodegenType(f.Type))
}

type StructType struct {
	types.StructType
	role string
}

func NewStructType(structType *types.StructType, role string) *StructType {
	return &StructType{
		StructType: *structType,
		role:       role,
	}
}

func (s *StructType) Role() string {
	return s.role
}
