package epp

import (
	"chorego/parser"
	"chorego/process"
)

func EppFunc(function parser.IFuncContext) process.Network {
	func_role := function.Role_type_normal()
	func_name := function.Ident().GetText()

	network := process.NewNetwork()

	for _, role := range func_role.AllIdent() {
		roleName := role.ID().GetText()
		proc := network.Process(roleName)
		proc.AddFunc(func_name)
	}

	return network
}
