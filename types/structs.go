package types

import "fmt"

type StructType struct {
	name string
}

// IsSendable implements Value.
func (s *StructType) IsSendable() bool {
	return true
}

func (t *StructType) IsEquatable() bool {
	return true
}

// IsValue implements Value.
func (s *StructType) IsValue() {}

// ToString implements Value.
func (s *StructType) ToString() string {
	return fmt.Sprintf("struct %s", s.name)
}

func NewStructType(name string) Value {
	return &StructType{name: name}
}

func (s *StructType) Name() string {
	return s.name
}
