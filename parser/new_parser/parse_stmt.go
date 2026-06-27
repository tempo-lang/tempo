package new_parser

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

func (p *Parser) ParseStmt() (ast.Stmt, bool) {
	var stmt ast.Stmt
	switch p.curToken.Type {
	case token.LET:
		letStmt, needsRecovery := p.parseLetStmt()
		stmt = letStmt
		if needsRecovery {
			return letStmt, true
		}
	default:
		return &ast.InvalidStmt{
			ErrorToken: p.errorToken("invalid statement"),
		}, true
	}

	return stmt, false
}

func (p *Parser) parseLetStmt() (ast.Stmt, bool) {
	stmt := &ast.LetStmt{
		LetToken: p.assertToken(token.LET),
	}

	ident, needsRecover := p.parseIdentifier()
	stmt.Name = ident
	if needsRecover {
		return &ast.InvalidStmt{
			ErrorToken:  p.errorToken("expected identifier when parsing let statement"),
			PartialStmt: stmt,
		}, true
	}

	if p.curToken.Type != token.ASSIGN {
		return &ast.InvalidStmt{
			ErrorToken:  p.errorToken("expected `=` when parsing let statement"),
			PartialStmt: stmt,
		}, true
	}
	p.assertToken(token.ASSIGN)

	expr, needsRecover := p.ParseExpr()
	stmt.Expr = expr
	if needsRecover {
		return &ast.InvalidStmt{
			ErrorToken:  expr.StartToken(),
			PartialStmt: stmt,
		}, true
	}

	semi, errStmt := p.expectSemicolon(stmt)
	if errStmt != nil {
		return errStmt, true
	}
	stmt.SemiToken = semi

	return stmt, false
}
