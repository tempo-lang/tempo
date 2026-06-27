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
	case token.RETURN:
		returnStmt, needsRecovery := p.parseReturnStmt()
		stmt = returnStmt
		if needsRecovery {
			return returnStmt, true
		}
	default:
		exprStmt, needsRecovery := p.parseExprOrAssignStmt()
		stmt = exprStmt
		if needsRecovery {
			return exprStmt, true
		}
	}

	return stmt, false
}

func (p *Parser) parseReturnStmt() (ast.Stmt, bool) {
	stmt := &ast.ReturnStmt{
		ReturnToken: p.assertToken(token.RETURN),
	}

	// The expression is optional
	if p.curToken.Type != token.SEMICOLON {
		expr, needsRecover := p.ParseExpr()
		if needsRecover {
			return &ast.InvalidStmt{
				ErrorToken:  p.errorToken("expected expression when parsing return statement"),
				PartialStmt: stmt,
			}, true
		}
		stmt.Expr = expr
	}

	semi, errStmt := p.expectSemicolon(stmt)
	if errStmt != nil {
		return errStmt, true
	}
	stmt.SemiToken = semi

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

func (p *Parser) parseExprOrAssignStmt() (ast.Stmt, bool) {
	// Parse the left-hand side expression
	lhs, needsRecover := p.ParseExpr()
	if needsRecover {
		return &ast.InvalidStmt{
			ErrorToken: p.errorToken("expected expression"),
		}, true
	}

	// Check if this is an assignment
	if p.curToken.Type == token.ASSIGN {
		assignToken := p.assertToken(token.ASSIGN)

		// Parse the right-hand side expression
		rhs, needsRecover := p.ParseExpr()
		if needsRecover {
			return &ast.InvalidStmt{
				ErrorToken: p.errorToken("expected expression after assignment"),
				PartialStmt: &ast.AssignStmt{
					LHS:         lhs,
					AssignToken: assignToken,
				},
			}, true
		}

		// Expect semicolon
		semi, errStmt := p.expectSemicolon(&ast.AssignStmt{
			LHS:         lhs,
			AssignToken: assignToken,
			RHS:         rhs,
		})
		if errStmt != nil {
			return errStmt, true
		}

		return &ast.AssignStmt{
			LHS:         lhs,
			AssignToken: assignToken,
			RHS:         rhs,
			SemiToken:   semi,
		}, false
	} else {
		// This is an expression statement
		semi, errStmt := p.expectSemicolon(&ast.ExprStmt{
			Expr: lhs,
		})
		if errStmt != nil {
			return errStmt, true
		}

		return &ast.ExprStmt{
			Expr:      lhs,
			SemiToken: semi,
		}, false
	}
}

func (p *Parser) parseExprStmt() (ast.Stmt, bool) {
	// Parse the expression
	expr, needsRecover := p.ParseExpr()
	if needsRecover {
		return &ast.InvalidStmt{
			ErrorToken: p.errorToken("expected expression when parsing expression statement"),
		}, true
	}

	// Expect semicolon
	semi, errStmt := p.expectSemicolon(&ast.ExprStmt{
		Expr: expr,
	})
	if errStmt != nil {
		return errStmt, true
	}

	return &ast.ExprStmt{
		Expr:      expr,
		SemiToken: semi,
	}, false
}
