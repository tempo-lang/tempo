package projection

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"
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

func (s *Struct) AddField(field parser.IStructFieldContext, fieldType Type) {
	s.Fields = append(s.Fields, StructField{
		Struct:   s,
		FieldCtx: field,
		Name:     field.Ident().GetText(),
		Type:     fieldType,
	})
}

type StructField struct {
	Struct   *Struct
	FieldCtx parser.IStructFieldContext
	Name     string
	Type     Type
}

type StructType struct {
	types.StructType
	role string
}

func (c *StructType) IsType() {}

func NewStructType(structType *types.StructType, role string) *StructType {
	return &StructType{
		StructType: *structType,
		role:       role,
	}
}

func (s *StructType) Role() string {
	return s.role
}
