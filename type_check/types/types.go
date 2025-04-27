package types

type Type interface {
	ToString() string
	IsType()
}

type InvalidType struct{}

var invalid_type InvalidType = InvalidType{}

func (t *InvalidType) ToString() string {
	return "ERROR"
}

func (t *InvalidType) IsType() {}

func Invalid() *InvalidType {
	return &invalid_type
}
