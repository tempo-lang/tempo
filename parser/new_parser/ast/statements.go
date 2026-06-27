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
	LHS         Expr
	AssignToken token.Token
	RHS         Expr
	SemiToken   token.Token
}

func (s *AssignStmt) StartToken() token.Token {
	return s.LHS.StartToken()
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
