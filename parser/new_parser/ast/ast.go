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

// Statements

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
	LetToken    token.Token
	Name        *Identifier
	Colon       token.Token
	Type        ValueType
	AssignToken token.Token
	Expr        Expr
	SemiToken   token.Token
}

func (s *LetStmt) StartToken() token.Token {
	return s.LetToken
}

func (s *LetStmt) EndToken() token.Token {
	return s.SemiToken
}

func (s *LetStmt) statementNode() {}

type IfStmt struct {
	IfToken   token.Token
	Condition Expr
	ThenScope *Scope
	ElseToken token.Token
	ElseScope *Scope
}

func (s *IfStmt) StartToken() token.Token {
	return s.IfToken
}

func (s *IfStmt) EndToken() token.Token {
	if s.ElseScope != nil {
		return s.ElseScope.EndToken()
	}
	return s.ThenScope.EndToken()
}

func (s *IfStmt) statementNode() {}

type WhileStmt struct {
	WhileKeyword token.Token
	Condition    Expr
	Scope        *Scope
}

func (s *WhileStmt) StartToken() token.Token {
	return s.WhileKeyword
}

func (s *WhileStmt) EndToken() token.Token {
	return s.Scope.EndToken()
}

func (s *WhileStmt) statementNode() {}

type ReturnStmt struct {
	ReturnToken token.Token
	Expr        Expr
	SemiToken   token.Token
}

func (s *ReturnStmt) StartToken() token.Token {
	return s.ReturnToken
}

func (s *ReturnStmt) EndToken() token.Token {
	return s.SemiToken
}

func (s *ReturnStmt) statementNode() {}

type AssignStmt struct {
	AssignExpr  *AssignExpr
	AssignToken token.Token
	Expr        Expr
	SemiToken   token.Token
}

func (s *AssignStmt) StartToken() token.Token {
	return s.AssignExpr.StartToken()
}

func (s *AssignStmt) EndToken() token.Token {
	return s.SemiToken
}

func (s *AssignStmt) statementNode() {}

type ExprStmt struct {
	Expr      Expr
	SemiToken token.Token
}

func (s *ExprStmt) StartToken() token.Token {
	return s.Expr.StartToken()
}

func (s *ExprStmt) EndToken() token.Token {
	return s.SemiToken
}

func (s *ExprStmt) statementNode() {}

// AssignExpr
type AssignExpr struct {
	Ident      *Identifier
	Specifiers []AssignSpecifier
}

func (e *AssignExpr) StartToken() token.Token {
	return e.Ident.StartToken()
}

func (e *AssignExpr) EndToken() token.Token {
	if len(e.Specifiers) > 0 {
		return e.Specifiers[len(e.Specifiers)-1].EndToken()
	}
	return e.Ident.EndToken()
}

// AssignSpecifier is an interface for the different types of assignment specifiers
type AssignSpecifier interface {
	Node
	assignSpecifierNode()
}

// AssignFieldSpecifier represents a field access specifier (e.g., .field)
type AssignFieldSpecifier struct {
	DotToken token.Token
	Ident    *Identifier
}

func (s *AssignFieldSpecifier) StartToken() token.Token {
	return s.DotToken
}

func (s *AssignFieldSpecifier) EndToken() token.Token {
	return s.Ident.EndToken()
}

func (s *AssignFieldSpecifier) assignSpecifierNode() {}

// AssignIndexSpecifier represents an index access specifier (e.g., [index])
type AssignIndexSpecifier struct {
	OpenBracket  token.Token
	IndexExpr    Expr
	CloseBracket token.Token
}

func (s *AssignIndexSpecifier) StartToken() token.Token {
	return s.OpenBracket
}

func (s *AssignIndexSpecifier) EndToken() token.Token {
	return s.CloseBracket
}

func (s *AssignIndexSpecifier) assignSpecifierNode() {}

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
