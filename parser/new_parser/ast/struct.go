package ast

import "github.com/tempo-lang/tempo/parser/new_parser/token"

// Struct
type Struct struct {
	StructToken token.Token
	RoleType    *RoleType
	Name        *Identifier
	Implements  []*RoleIdent
	Body        *StructBody
}

func (s *Struct) StartToken() token.Token {
	return s.StructToken
}

func (s *Struct) EndToken() token.Token {
	return s.Body.EndToken()
}

type StructBody struct {
	OpenToken  token.Token
	Fields     []*StructField
	Functions  []*Func
	CloseToken token.Token
}

func (b *StructBody) StartToken() token.Token {
	return b.OpenToken
}

func (b *StructBody) EndToken() token.Token {
	return b.CloseToken
}

type StructField struct {
	Name      *Identifier
	Colon     token.Token
	Type      ValueType
	SemiToken token.Token
}

func (f *StructField) StartToken() token.Token {
	return f.Name.StartToken()
}

func (f *StructField) EndToken() token.Token {
	return f.SemiToken
}
