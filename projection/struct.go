package projection

import (
	"fmt"

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
	StructCtx  parser.IStructContext
	Name       string
	Role       string
	Fields     []StructField
	Methods    []*StructMethod
	Implements []Type
}

func (c *ChoreographyStruct) AddStruct(role string, structCtx parser.IStructContext) *Struct {
	c.Roles = append(c.Roles, role)
	c.Structs[role] = &Struct{
		StructCtx:  structCtx,
		Name:       structCtx.Ident().GetText(),
		Role:       role,
		Fields:     []StructField{},
		Implements: []Type{},
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

func (s *Struct) AddMethod(sig *FuncSig, funcCtx parser.IFuncContext) *StructMethod {
	method := &StructMethod{
		Struct:  s,
		FuncSig: sig,
		FuncCtx: funcCtx,
		Body:    []Statement{},
	}

	s.Methods = append(s.Methods, method)

	return method
}

func (s *Struct) AddImplements(infType Type) {
	s.Implements = append(s.Implements, infType)
}

func (s *Struct) StructName() string {
	if s.Role == "" {
		return s.Name
	} else {
		return fmt.Sprintf("%s_%s", s.Name, s.Role)
	}
}

type StructField struct {
	Struct   *Struct
	FieldCtx parser.IStructFieldContext
	Name     string
	Type     Type
}

type StructMethod struct {
	Struct  *Struct
	FuncSig *FuncSig
	FuncCtx parser.IFuncContext
	Body    []Statement
}

func (m *StructMethod) AddStmt(stmt ...Statement) *StructMethod {
	m.Body = append(m.Body, stmt...)
	return m
}

// A record that combines multiple values with potentially different types, identified by names.
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

func (s *StructType) StructName() string {
	if s.role == "" {
		return s.Name()
	} else {
		return fmt.Sprintf("%s_%s", s.Name(), s.Role())
	}
}
