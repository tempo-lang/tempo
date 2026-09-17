package projection

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/types"
)

type ChoreographyInterface struct {
	Name       string
	Roles      []string
	Interfaces map[string]*Interface
}

type Interface struct {
	InterfaceCtx *ast.Interface
	Name         string
	Role         string
	Methods      []*InterfaceMethod
}

type InterfaceMethod struct {
	*FuncSig
	MethodCtx *ast.InterfaceMethod
}

func NewChoreographyInterface(name string) *ChoreographyInterface {
	return &ChoreographyInterface{
		Name:       name,
		Roles:      []string{},
		Interfaces: make(map[string]*Interface),
	}
}

func (inf *ChoreographyInterface) AddInterface(role string, ctx *ast.Interface) *Interface {
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

func (inf *Interface) AddMethod(sig *FuncSig, ctx *ast.InterfaceMethod) *InterfaceMethod {
	method := &InterfaceMethod{
		FuncSig:   sig,
		MethodCtx: ctx,
	}

	inf.Methods = append(inf.Methods, method)

	return method
}

func (inf *Interface) InterfaceName() string {
	if inf.Role == "" {
		return inf.Name
	} else {
		return fmt.Sprintf("%s_%s", inf.Name, inf.Role)
	}
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

func (c *InterfaceType) IsType() {}

func (i *InterfaceType) Role() string {
	return i.role
}

func (inf *InterfaceType) InterfaceName() string {
	if inf.Role() == "" {
		return inf.Name()
	} else {
		return fmt.Sprintf("%s_%s", inf.Name(), inf.Role())
	}
}
