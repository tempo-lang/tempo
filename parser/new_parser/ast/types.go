package ast

import "github.com/tempo-lang/tempo/parser/new_parser/token"

// Types

type RoleType struct {
	Start token.Token
	End   token.Token
	// RoleNodes contains all parsed roles (including errors)
	// Use [RoleType.Roles] to get all successfully parsed roles.
	RoleNodes []*Role
}

func (r *RoleType) StartToken() token.Token {
	return r.Start
}

func (r *RoleType) EndToken() token.Token {
	return r.End
}

func (r *RoleType) IsShared() bool {
	return r.Start.Type == token.LSQUARE
}

// Error returns the error token in the role type, or nil if there are no errors.
func (r *RoleType) Error() *token.Token {
	if r.Start.Type == token.ILLEGAL {
		return &r.Start
	}
	for _, role := range r.RoleNodes {
		if role.Token.Type == token.ILLEGAL {
			return &role.Token
		}
	}
	if r.End.Type == token.ILLEGAL {
		return &r.End
	}
	return nil
}

// Roles returns all successfully parsed roles.
func (r *RoleType) Roles() (roles []*Role) {
	for _, role := range r.RoleNodes {
		if role.Token.Type != token.ILLEGAL {
			roles = append(roles, role)
		}
	}
	return
}

type Role struct {
	// ident or _
	Token token.Token
}

func (r *Role) StartToken() token.Token {
	return r.Token
}

func (r *Role) EndToken() token.Token {
	return r.Token
}

// AsyncType
type AsyncType struct {
	AsyncKeyword token.Token
	Inner        ValueType
}

func (a *AsyncType) StartToken() token.Token {
	return a.AsyncKeyword
}

func (a *AsyncType) EndToken() token.Token {
	return a.Inner.EndToken()
}

func (a *AsyncType) valueTypeNode() {}

// ListType
type ListType struct {
	OpenBracket  token.Token
	Inner        ValueType
	CloseBracket token.Token
}

func (l *ListType) StartToken() token.Token {
	return l.OpenBracket
}

func (l *ListType) EndToken() token.Token {
	return l.CloseBracket
}

func (l *ListType) valueTypeNode() {}

// ClosureType
type ClosureType struct {
	FuncToken   token.Token
	RoleAtToken token.Token
	RoleType    *RoleType
	Params      *ClosureTypeParams
	ReturnType  ValueType
}

func (c *ClosureType) StartToken() token.Token {
	return c.FuncToken
}

func (c *ClosureType) EndToken() token.Token {
	if c.ReturnType != nil {
		return c.ReturnType.EndToken()
	}
	return c.Params.EndToken()
}

func (c *ClosureType) valueTypeNode() {}

// ClosureTypeParams
type ClosureTypeParams struct {
	OpenParen  token.Token
	Params     []ValueType
	CloseParen token.Token
}

func (c *ClosureTypeParams) StartToken() token.Token {
	return c.OpenParen
}

func (c *ClosureTypeParams) EndToken() token.Token {
	return c.CloseParen
}

// NamedType
type NamedType struct {
	RoleIdent *RoleIdent
}

func (n *NamedType) StartToken() token.Token {
	return n.RoleIdent.StartToken()
}

func (n *NamedType) EndToken() token.Token {
	return n.RoleIdent.EndToken()
}

func (n *NamedType) valueTypeNode() {}

// RoleIdent
type RoleIdent struct {
	Ident    *Identifier
	RoleAt   token.Token // optional
	RoleType *RoleType   // optional
}

func (r *RoleIdent) StartToken() token.Token {
	return r.Ident.StartToken()
}

func (r *RoleIdent) EndToken() token.Token {
	if r.RoleType != nil {
		return r.RoleType.EndToken()
	}
	return r.Ident.EndToken()
}
