package ast

import "github.com/tempo-lang/tempo/parser/new_parser/token"

// Interface
type Interface struct {
	InterfaceKeyword token.Token
	RoleType         *RoleType
	Name             *Identifier
	Methods          *InterfaceMethodsList
}

func (i *Interface) StartToken() token.Token {
	return i.InterfaceKeyword
}

func (i *Interface) EndToken() token.Token {
	return i.Methods.EndToken()
}

type InterfaceMethodsList struct {
	OpenToken  token.Token
	Methods    []*InterfaceMethod
	CloseToken token.Token
}

func (m *InterfaceMethodsList) StartToken() token.Token {
	return m.OpenToken
}

func (m *InterfaceMethodsList) EndToken() token.Token {
	return m.CloseToken
}

type InterfaceMethod struct {
	FuncSig   *FuncSig
	SemiToken token.Token
}

func (m *InterfaceMethod) StartToken() token.Token {
	return m.FuncSig.StartToken()
}

func (m *InterfaceMethod) EndToken() token.Token {
	return m.SemiToken
}

func (m *InterfaceMethod) statementNode() {}
