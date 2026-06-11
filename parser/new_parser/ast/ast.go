package ast

import "github.com/tempo-lang/tempo/parser/new_parser/token"

type Node interface {
	StartToken() token.Token
	EndToken() token.Token
}

type Stmt interface {
	Node
	statementNode()
}

type Expr interface {
	Node
	expressionNode()
}

// type SourceFile struct {
// 	Functions []Function
// }

// // Function

// type Function struct {
// 	FuncSig FuncSig
// 	Scope   Scope
// }

// type FuncSig struct {
// 	FuncToken token.Token // the func keyword
// 	RoleType  RoleType
// }

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

// // Types

// type RoleType struct {
// 	Token token.Token
// }

type InvalidStmt struct {
	ErrorToken  token.Token
	PartialStmt Stmt
}

func (s *InvalidStmt) StartToken() token.Token {
	return s.ErrorToken
}

func (s *InvalidStmt) EndToken() token.Token {
	return s.ErrorToken
}

func (s *InvalidStmt) statementNode() {}

type LetStmt struct {
	LetToken  token.Token
	Name      *Identifier
	Expr      Expr
	SemiToken token.Token
}

func (s *LetStmt) StartToken() token.Token {
	return s.LetToken
}

func (s *LetStmt) EndToken() token.Token {
	return s.SemiToken
}

func (s *LetStmt) statementNode() {}

type Identifier struct {
	Token token.Token
}

func (i *Identifier) Value() string {
	return i.Token.Value.(string)
}

func (i *Identifier) StartToken() token.Token {
	return i.Token
}

func (i *Identifier) EndToken() token.Token {
	return i.Token
}

func (i *Identifier) expressionNode() {}
