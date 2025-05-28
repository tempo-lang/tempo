package types

import "fmt"

type StructType struct {
	name string
}

func (s *StructType) SubstituteRoles(substMap map[string]string) Value {
	return s
}

func (s *StructType) IsSendable() bool {
	return true
}

func (t *StructType) IsEquatable() bool {
	return true
}

func (s *StructType) IsValue() {}

func (s *StructType) ToString() string {
	return fmt.Sprintf("struct %s", s.name)
}

func NewStructType(name string) Value {
	return &StructType{name: name}
}

func (s *StructType) Name() string {
	return s.name
}
