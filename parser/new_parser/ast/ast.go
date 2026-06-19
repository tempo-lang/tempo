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

type InvalidExpr struct {
	ErrorToken  token.Token
	PartialExpr Expr
}

func (e *InvalidExpr) StartToken() token.Token {
	return e.ErrorToken
}

func (e *InvalidExpr) EndToken() token.Token {
	return e.ErrorToken
}

func (e *InvalidExpr) expressionNode() {}

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

type FloatExpr struct {
	FloatToken token.Token
}

func (n *FloatExpr) Value() float64 {
	return n.FloatToken.Value.(float64)
}

func (n *FloatExpr) StartToken() token.Token {
	return n.FloatToken
}

func (n *FloatExpr) EndToken() token.Token {
	return n.FloatToken
}

func (n *FloatExpr) expressionNode() {}

type IntExpr struct {
	IntToken token.Token
}

func (n *IntExpr) Value() int {
	return n.IntToken.Value.(int)
}

func (n *IntExpr) StartToken() token.Token {
	return n.IntToken
}

func (n *IntExpr) EndToken() token.Token {
	return n.IntToken
}

func (n *IntExpr) expressionNode() {}

// BinaryExpr represents a binary operation like a + b, x * y, etc.
type BinaryExpr struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (e *BinaryExpr) StartToken() token.Token {
	return e.Left.StartToken()
}

func (e *BinaryExpr) EndToken() token.Token {
	return e.Right.EndToken()
}

func (e *BinaryExpr) expressionNode() {}
