package ast

import "github.com/tempo-lang/tempo/parser/new_parser/token"

// Expressions

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

type FloatLit struct {
	FloatToken token.Token
}

func (n *FloatLit) Value() float64 {
	return n.FloatToken.Value.(float64)
}

func (n *FloatLit) StartToken() token.Token {
	return n.FloatToken
}

func (n *FloatLit) EndToken() token.Token {
	return n.FloatToken
}

func (n *FloatLit) literalNode() {}

type IntLit struct {
	IntToken token.Token
}

func (n *IntLit) Value() int {
	return n.IntToken.Value.(int)
}

func (n *IntLit) StartToken() token.Token {
	return n.IntToken
}

func (n *IntLit) EndToken() token.Token {
	return n.IntToken
}

func (n *IntLit) literalNode() {}

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

// StringLit
type StringLit struct {
	StringToken token.Token
	RoleType    *RoleType
}

func (s *StringLit) Value() string {
	return s.StringToken.Value.(string)
}

func (s *StringLit) StartToken() token.Token {
	return s.StringToken
}

func (s *StringLit) EndToken() token.Token {
	if s.RoleType != nil {
		return s.RoleType.EndToken()
	}
	return s.StringToken
}

func (s *StringLit) literalNode() {}

// BoolLit
type BoolLit struct {
	BoolToken token.Token
	RoleType  *RoleType
}

func (b *BoolLit) Value() bool {
	return b.BoolToken.Type == token.TRUE
}

func (b *BoolLit) StartToken() token.Token {
	return b.BoolToken
}

func (b *BoolLit) EndToken() token.Token {
	if b.RoleType != nil {
		return b.RoleType.EndToken()
	}
	return b.BoolToken
}

func (b *BoolLit) literalNode() {}

// ClosureExpr
type ClosureExpr struct {
	ClosureSig *ClosureSig
	Scope      *Scope
}

func (c *ClosureExpr) StartToken() token.Token {
	return c.ClosureSig.StartToken()
}

func (c *ClosureExpr) EndToken() token.Token {
	return c.Scope.EndToken()
}

func (c *ClosureExpr) expressionNode() {}

// ClosureSig
type ClosureSig struct {
	FuncToken   token.Token
	RoleAtToken token.Token
	RoleType    *RoleType
	Params      *FuncParams
	ReturnType  ValueType
}

func (c *ClosureSig) StartToken() token.Token {
	return c.FuncToken
}

func (c *ClosureSig) EndToken() token.Token {
	if c.ReturnType != nil {
		return c.ReturnType.EndToken()
	}
	return c.Params.EndToken()
}

// StructExpr
type StructExpr struct {
	RoleIdent    *RoleIdent
	StructFields *StructFields
}

func (s *StructExpr) StartToken() token.Token {
	return s.RoleIdent.StartToken()
}

func (s *StructExpr) EndToken() token.Token {
	return s.StructFields.EndToken()
}

func (s *StructExpr) expressionNode() {}

// StructFields
type StructFields struct {
	OpenToken  token.Token
	Fields     []*StructFieldExpr
	CloseToken token.Token
}

func (s *StructFields) StartToken() token.Token {
	return s.OpenToken
}

func (s *StructFields) EndToken() token.Token {
	return s.CloseToken
}

type StructFieldExpr struct {
	Name  *Identifier
	Colon token.Token
	Expr  Expr
}

func (f *StructFieldExpr) StartToken() token.Token {
	return f.Name.StartToken()
}

func (f *StructFieldExpr) EndToken() token.Token {
	return f.Expr.EndToken()
}

// CallExpr
type CallExpr struct {
	Function   Expr
	OpenParen  token.Token
	Args       []Expr
	CloseParen token.Token
}

func (c *CallExpr) StartToken() token.Token {
	return c.Function.StartToken()
}

func (c *CallExpr) EndToken() token.Token {
	return c.CloseParen
}

func (c *CallExpr) expressionNode() {}

// FieldAccessExpr
type FieldAccessExpr struct {
	Object   Expr
	DotToken token.Token
	Field    *Identifier
}

func (f *FieldAccessExpr) StartToken() token.Token {
	return f.Object.StartToken()
}

func (f *FieldAccessExpr) EndToken() token.Token {
	return f.Field.EndToken()
}

func (f *FieldAccessExpr) expressionNode() {}

// IndexExpr
type IndexExpr struct {
	Object       Expr
	OpenBracket  token.Token
	Index        Expr
	CloseBracket token.Token
}

func (i *IndexExpr) StartToken() token.Token {
	return i.Object.StartToken()
}

func (i *IndexExpr) EndToken() token.Token {
	return i.CloseBracket
}

func (i *IndexExpr) expressionNode() {}

// ListExpr
type ListExpr struct {
	OpenBracket  token.Token
	Elements     []Expr
	CloseBracket token.Token
}

func (l *ListExpr) StartToken() token.Token {
	return l.OpenBracket
}

func (l *ListExpr) EndToken() token.Token {
	return l.CloseBracket
}

func (l *ListExpr) expressionNode() {}

// IdentAccessExpr
type IdentAccessExpr struct {
	Ident    *Identifier
	RoleAt   token.Token
	RoleType *RoleType
}

func (i *IdentAccessExpr) StartToken() token.Token {
	return i.Ident.StartToken()
}

func (i *IdentAccessExpr) EndToken() token.Token {
	if i.RoleType != nil {
		return i.RoleType.EndToken()
	}
	return i.Ident.EndToken()
}

func (i *IdentAccessExpr) expressionNode() {}

// ComExpr (communication expression: sender -> receiver expr)
type ComExpr struct {
	Sender   *RoleType
	ComToken token.Token
	Receiver *RoleType
	Expr     Expr
}

func (c *ComExpr) StartToken() token.Token {
	return c.Sender.StartToken()
}

func (c *ComExpr) EndToken() token.Token {
	return c.Expr.EndToken()
}

func (c *ComExpr) expressionNode() {}

// AwaitExpr
type AwaitExpr struct {
	AwaitToken token.Token
	Expr       Expr
}

func (a *AwaitExpr) StartToken() token.Token {
	return a.AwaitToken
}

func (a *AwaitExpr) EndToken() token.Token {
	return a.Expr.EndToken()
}

func (a *AwaitExpr) expressionNode() {}

// GroupExpr (parenthesized expression)
type GroupExpr struct {
	OpenParen  token.Token
	Expr       Expr
	CloseParen token.Token
}

func (g *GroupExpr) StartToken() token.Token {
	return g.OpenParen
}

func (g *GroupExpr) EndToken() token.Token {
	return g.CloseParen
}

func (g *GroupExpr) expressionNode() {}

// PrimitiveExpr (literal with optional role type)
type PrimitiveExpr struct {
	Literal  Literal
	RoleAt   token.Token
	RoleType *RoleType
}

func (p *PrimitiveExpr) StartToken() token.Token {
	return p.Literal.StartToken()
}

func (p *PrimitiveExpr) EndToken() token.Token {
	if p.RoleType != nil {
		return p.RoleType.EndToken()
	}
	return p.Literal.EndToken()
}

func (p *PrimitiveExpr) expressionNode() {}
