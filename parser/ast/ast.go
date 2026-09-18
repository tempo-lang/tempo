package ast

import (
	"strings"

	"github.com/tempo-lang/tempo/parser/token"
)

type Node interface {
	StartToken() token.Token
	EndToken() token.Token
}

type TokenRange struct{ First, Last token.Token }

func (r *TokenRange) StartToken() token.Token { return r.First }
func (r *TokenRange) EndToken() token.Token   { return r.Last }

// Text returns the source spelling carried by simple nodes. It is primarily
// used in diagnostics, where retaining a parser-runtime context is undesirable.
func Text(n Node) string {
	switch n := n.(type) {
	case *Identifier:
		return n.Value()
	case *IdentAccessExpr:
		text := n.Ident.Value()
		if n.RoleType != nil {
			text += "@" + Text(n.RoleType)
		}
		return text
	case *FieldAccessExpr:
		return Text(n.Object) + "." + n.Field.Value()
	case *NamedType:
		return Text(n.RoleIdent)
	case *RoleIdent:
		text := n.Ident.Value()
		if n.RoleType != nil {
			text += "@" + Text(n.RoleType)
		}
		return text
	case *Role:
		return n.Token.Text
	case *RoleType:
		rs := n.Roles()
		parts := make([]string, len(rs))
		for i, role := range rs {
			parts[i] = role.Token.Text
		}
		if n.IsShared() {
			return "[" + strings.Join(parts, ",") + "]"
		}
		if len(parts) == 1 {
			return parts[0]
		}
		if len(parts) > 1 {
			return "(" + strings.Join(parts, ",") + ")"
		}
	}
	return n.StartToken().Text
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
	Skipped     []token.Ref
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
	// Declarations preserves the lexical top-level order. The typed slices are
	// retained as convenient indexes for consumers.
	Declarations []Node
	Tokens       []token.Token
}

func (f *SourceFile) StartToken() token.Token {
	return f.FirstToken
}

func (f *SourceFile) EndToken() token.Token {
	return f.LastToken
}
