package codegen_go

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/tempo-lang/tempo/projection"
)

func GenChoreographyInterface(file *jen.File, c *projection.ChoreographyInterface) {
	file.Commentf("Projection of interface %s", c.Name)

	for _, role := range c.Roles {
		file.Add(GenInterface(c.Interfaces[role]))
	}
}

func GenInterface(inf *projection.Interface) *jen.Statement {
	methods := []jen.Code{}

	for _, method := range inf.Methods {
		methods = append(methods, GenFuncSig(method.FuncSig, true))
	}

	return jen.Type().Id(inf.GenName()).Interface(methods...)
}

func GenInterfaceType(t *projection.InterfaceType) jen.Code {
	return jen.Id(fmt.Sprintf("%s_%s", t.Name(), t.Role()))
}
