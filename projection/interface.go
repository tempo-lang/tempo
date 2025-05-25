package projection

import (
	"fmt"
	"tempo/parser"
	"tempo/types"

	"github.com/dave/jennifer/jen"
)

type ChoreographyInterface struct {
	Name       string
	Roles      []string
	Interfaces map[string]*Interface
}

type Interface struct {
	InterfaceCtx parser.IInterfaceContext
	Name         string
	Role         string
	Methods      []*InterfaceMethod
}

type InterfaceMethod struct {
	*FuncSig
	MethodCtx parser.IInterfaceMethodContext
}

func NewChoreographyInterface(name string) *ChoreographyInterface {
	return &ChoreographyInterface{
		Name:       name,
		Roles:      []string{},
		Interfaces: make(map[string]*Interface),
	}
}

func (inf *ChoreographyInterface) AddInterface(role string, ctx parser.IInterfaceContext) *Interface {
	result := &Interface{
		Name:         inf.Name,
		Role:         role,
		InterfaceCtx: ctx,
		Methods:      []*InterfaceMethod{},
	}

	inf.Roles = append(inf.Roles, role)
	inf.Interfaces[role] = result

	return result
}

func (inf *Interface) AddMethod(sig *FuncSig, ctx parser.IInterfaceMethodContext) *InterfaceMethod {
	method := &InterfaceMethod{
		FuncSig:   sig,
		MethodCtx: ctx,
	}

	inf.Methods = append(inf.Methods, method)

	return method
}

func (c *ChoreographyInterface) Codegen(file *jen.File) {
	file.Commentf("Projection of interface %s", c.Name)

	for _, role := range c.Roles {
		file.Add(c.Interfaces[role].Codegen())
	}
}

func (inf *Interface) Codegen() *jen.Statement {
	methods := []jen.Code{}

	for _, method := range inf.Methods {
		methods = append(methods, method.Codegen(true))
	}

	return jen.Type().Id(fmt.Sprintf("%s_%s", inf.Name, inf.Role)).Interface(methods...)
}

type InterfaceType struct {
	types.InterfaceType
	role string
}

func NewInterfaceType(interfaceType *types.InterfaceType, role string) *InterfaceType {
	return &InterfaceType{
		InterfaceType: *interfaceType,
		role:          role,
	}
}

func (i *InterfaceType) Role() string {
	return i.role
}
