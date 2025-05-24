package types

import "fmt"

type InterfaceType struct {
	name string
}

func (s *InterfaceType) IsSendable() bool {
	return false
}

func (t *InterfaceType) IsEquatable() bool {
	return true
}

func (s *InterfaceType) IsValue() {}

func (s *InterfaceType) ToString() string {
	return fmt.Sprintf("interface %s", s.name)
}

func NewInterfaceType(name string) Value {
	return &InterfaceType{name: name}
}

func (s *InterfaceType) Name() string {
	return s.name
}
