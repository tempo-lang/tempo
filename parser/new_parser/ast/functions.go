package ast

import "github.com/tempo-lang/tempo/parser/new_parser/token"

// Function

type Func struct {
	FuncSig *FuncSig
	Scope   *Scope
}

func (f *Func) StartToken() token.Token {
	return f.FuncSig.StartToken()
}

func (f *Func) EndToken() token.Token {
	return f.Scope.EndToken()
}

type FuncSig struct {
	FuncToken  token.Token // the func keyword
	RoleType   *RoleType
	Name       *Identifier
	Params     *FuncParams
	ReturnType ValueType
}

func (s *FuncSig) StartToken() token.Token {
	return s.FuncToken
}

func (s *FuncSig) EndToken() token.Token {
	if s.ReturnType != nil {
		return s.ReturnType.EndToken()
	}
	return s.Params.EndToken()
}

type FuncParams struct {
	OpenParen  token.Token
	Params     []*FuncParam
	CloseParen token.Token
}

func (p *FuncParams) StartToken() token.Token {
	return p.OpenParen
}

func (p *FuncParams) EndToken() token.Token {
	return p.CloseParen
}

type FuncParam struct {
	Name *Identifier
	Type ValueType
}

func (p *FuncParam) StartToken() token.Token {
	return p.Name.StartToken()
}

func (p *FuncParam) EndToken() token.Token {
	return p.Type.EndToken()
}

type Scope struct {
	OpenToken  token.Token
	CloseToken token.Token
	Stmts      []Stmt
}

func (s *Scope) StartToken() token.Token {
	return s.OpenToken
}

func (s *Scope) EndToken() token.Token {
	return s.CloseToken
}
