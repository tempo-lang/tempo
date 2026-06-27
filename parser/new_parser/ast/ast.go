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

type Literal interface {
	Node
	literalNode()
}

type ValueType interface {
	Node
	valueTypeNode()
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

// Source File

type SourceFile struct {
	FirstToken token.Token
	LastToken  token.Token

	Functions  []*Func
	Structs    []*Struct
	Interfaces []*Interface
}

func (f *SourceFile) StartToken() token.Token {
	return f.FirstToken
}

func (f *SourceFile) EndToken() token.Token {
	return f.LastToken
}
