package ast

import "github.com/tempo-lang/tempo/parser/new_parser/token"

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
