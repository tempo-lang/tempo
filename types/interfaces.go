package types

import (
	"fmt"
	"tempo/parser"
)

type InterfaceType struct {
	ident parser.IIdentContext
}

func (s *InterfaceType) SubstituteRoles(substMap map[string]string) Value {
	return s
}

func (s *InterfaceType) IsSendable() bool {
	return false
}

func (t *InterfaceType) IsEquatable() bool {
	return true
}

func (s *InterfaceType) IsValue() {}

func (s *InterfaceType) ToString() string {
	return fmt.Sprintf("interface %s", s.ident.GetText())
}

func NewInterfaceType(ident parser.IIdentContext) Value {
	return &InterfaceType{ident: ident}
}

func (s *InterfaceType) Name() string {
	return s.ident.GetText()
}

func (s *InterfaceType) Ident() parser.IIdentContext {
	return s.ident
}
