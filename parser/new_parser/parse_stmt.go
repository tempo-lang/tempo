package new_parser

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

var stmtBoundaries = TokenSet{
	token.SEMICOLON, token.RCURLY, token.EOF,
	token.LET, token.RETURN, token.IF, token.WHILE,
	token.FLOAT, token.INT, token.STRING, token.TRUE, token.FALSE,
	token.IDENT, token.LPAREN, token.LSQUARE, token.AWAIT,
}

var exprStmtEnd = TokenSet{token.SEMICOLON, token.RCURLY, token.EOF, token.LET, token.RETURN, token.IF, token.WHILE}

func isStmtStart(k token.Kind) bool {
	return k == token.LET || k == token.RETURN || k == token.IF || k == token.WHILE || exprStart(k)
}
func (p *Parser) ParseStmt() (ast.Stmt, bool) {
	before := len(p.diagnostics)
	s := p.parseStmt()
	p.expectEOF()
	return s, len(p.diagnostics) > before
}
func (p *Parser) parseStmt() ast.Stmt {
	switch p.current().Kind {
	case token.LET:
		return p.parseLet()
	case token.RETURN:
		return p.parseReturn()
	case token.IF:
		return p.parseIf()
	case token.WHILE:
		return p.parseWhile()
	default:
		return p.parseExprStmt()
	}
}
func (p *Parser) semicolon() token.Token {
	if p.current().Kind == token.SEMICOLON {
		return p.advance()
	}
	anchor := p.previous.Span
	at := anchor.End
	p.add(Diagnostic{Code: CodeMissingToken, Message: "missing semicolon", PrimarySpan: anchor, Expected: []token.Kind{token.SEMICOLON}, Found: p.current().Kind, Fixes: []TextEdit{{Span: Span{Start: at, End: at}, NewText: ";"}}})
	return token.Missing(token.Ref(len(p.tokens)), token.SEMICOLON, at, p.source)
}
func (p *Parser) parseLet() ast.Stmt {
	s := &ast.LetStmt{LetToken: p.advance()}
	s.Name = p.parseIdentifier()
	if p.current().Kind == token.COLON {
		s.Colon = p.advance()
		s.Type = p.parseValueType(exprStmtEnd.With(token.ASSIGN))
	}
	s.AssignToken = p.expect(token.ASSIGN, stmtBoundaries)
	s.Expr = p.parseExpr(exprStmtEnd)
	s.SemiToken = p.semicolon()
	return s
}
func (p *Parser) parseReturn() ast.Stmt {
	s := &ast.ReturnStmt{ReturnToken: p.advance()}
	if p.current().Kind != token.SEMICOLON && p.current().Kind != token.RCURLY && p.current().Kind != token.EOF {
		s.Expr = p.parseExpr(exprStmtEnd)
	}
	s.SemiToken = p.semicolon()
	return s
}
func (p *Parser) parseIf() ast.Stmt {
	s := &ast.IfStmt{IfToken: p.advance()}
	stops := exprStmtEnd.With(token.LCURLY)
	s.Condition = p.parseExpr(stops)
	s.ThenScope = p.parseScope()
	if p.current().Kind == token.ELSE {
		s.ElseToken = p.advance()
		s.ElseScope = p.parseScope()
	}
	return s
}
func (p *Parser) parseWhile() ast.Stmt {
	s := &ast.WhileStmt{WhileKeyword: p.advance()}
	s.Condition = p.parseExpr(exprStmtEnd.With(token.LCURLY))
	s.Scope = p.parseScope()
	return s
}
func assignable(e ast.Expr) bool {
	switch x := e.(type) {
	case *ast.Identifier:
		return true
	case *ast.IdentAccessExpr:
		return x.RoleType == nil
	case *ast.FieldAccessExpr, *ast.IndexExpr:
		return assignableRoot(e)
	}
	return false
}
func assignableRoot(e ast.Expr) bool {
	switch x := e.(type) {
	case *ast.FieldAccessExpr:
		return assignableRoot(x.Object)
	case *ast.IndexExpr:
		return assignableRoot(x.Object)
	case *ast.IdentAccessExpr:
		return x.RoleType == nil
	case *ast.Identifier:
		return true
	}
	return false
}
func (p *Parser) parseExprStmt() ast.Stmt {
	lhs := p.parseExpr(exprStmtEnd.With(token.ASSIGN))
	if p.current().Kind == token.ASSIGN {
		op := p.advance()
		if !assignable(lhs) {
			p.add(Diagnostic{Code: CodeInvalidAssignmentTarget, Message: "invalid assignment target", PrimarySpan: lhs.StartToken().Span, Found: op.Kind})
		}
		rhs := p.parseExpr(exprStmtEnd)
		s := &ast.AssignStmt{LHS: lhs, AssignToken: op, RHS: rhs}
		s.SemiToken = p.semicolon()
		return s
	}
	s := &ast.ExprStmt{Expr: lhs}
	s.SemiToken = p.semicolon()
	return s
}
func (p *Parser) ParseScope() (*ast.Scope, bool) {
	before := len(p.diagnostics)
	s := p.parseScope()
	p.expectEOF()
	return s, len(p.diagnostics) > before
}
func (p *Parser) parseScope() *ast.Scope {
	s := &ast.Scope{OpenToken: p.expect(token.LCURLY, stmtBoundaries)}
	for p.current().Kind != token.RCURLY && p.current().Kind != token.EOF {
		pos := p.cursor.Position()
		if !isStmtStart(p.current().Kind) {
			p.skipUntil(stmtBoundaries.Union(TokenSet{token.RCURLY}))
			if p.current().Kind == token.SEMICOLON {
				p.advance()
			}
		} else {
			s.Stmts = append(s.Stmts, p.parseStmt())
		}
		if p.cursor.Position() == pos {
			p.add(Diagnostic{Code: CodeInternalRecovery, Message: "parser made no progress", PrimarySpan: p.current().Span, Found: p.current().Kind})
			if p.current().Kind == token.EOF || p.current().Kind == token.RCURLY {
				break
			}
			p.advance()
		}
	}
	s.CloseToken = p.expect(token.RCURLY, TokenSet{token.EOF})
	return s
}
