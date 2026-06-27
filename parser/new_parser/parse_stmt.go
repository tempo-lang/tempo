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
		// Check if this could be an assignment statement
		if p.curToken.Type == token.IDENT {
			assignStmt, needsRecovery := p.parseAssignStmt()
			stmt = assignStmt
			if needsRecovery {
				return assignStmt, true
			}
		} else {
			// Try to parse as an expression statement
			exprStmt, needsRecovery := p.parseExprStmt()
			if needsRecovery {
				return exprStmt, true
			}
			stmt = exprStmt
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

func (p *Parser) parseAssignStmt() (ast.Stmt, bool) {
	// Parse the assign expression (identifier with optional specifiers)
	assignExpr, needsRecover := p.parseAssignExpr()
	if needsRecover {
		return &ast.InvalidStmt{
			ErrorToken: p.errorToken("expected assignment expression"),
		}, true
	}

	// Check for assignment token
	if p.curToken.Type != token.ASSIGN {
		return &ast.InvalidStmt{
			ErrorToken: p.errorToken("expected `=` in assignment statement"),
			PartialStmt: &ast.AssignStmt{
				AssignExpr: assignExpr,
			},
		}, true
	}
	assignToken := p.assertToken(token.ASSIGN)

	// Parse the right-hand side expression
	expr, needsRecover := p.ParseExpr()
	if needsRecover {
		return &ast.InvalidStmt{
			ErrorToken: expr.StartToken(),
			PartialStmt: &ast.AssignStmt{
				AssignExpr:  assignExpr,
				AssignToken: assignToken,
			},
		}, true
	}

	// Expect semicolon
	semi, errStmt := p.expectSemicolon(&ast.AssignStmt{
		AssignExpr:  assignExpr,
		AssignToken: assignToken,
		Expr:        expr,
	})
	if errStmt != nil {
		return errStmt, true
	}

	return &ast.AssignStmt{
		AssignExpr:  assignExpr,
		AssignToken: assignToken,
		Expr:        expr,
		SemiToken:   semi,
	}, false
}

func (p *Parser) parseAssignExpr() (*ast.AssignExpr, bool) {
	// Start with an identifier
	ident, needsRecover := p.parseIdentifier()
	if needsRecover {
		return nil, true
	}

	// Parse optional specifiers
	var specifiers []ast.AssignSpecifier
	for p.curToken.Type == token.DOT || p.curToken.Type == token.LSQUARE {
		specifier, needsRecover := p.parseAssignSpecifier()
		if needsRecover {
			return nil, true
		}
		specifiers = append(specifiers, specifier)
	}

	return &ast.AssignExpr{
		Ident:      ident,
		Specifiers: specifiers,
	}, false
}

func (p *Parser) parseAssignSpecifier() (ast.AssignSpecifier, bool) {
	switch p.curToken.Type {
	case token.DOT:
		// Field access specifier: .field
		dotToken := p.assertToken(token.DOT)
		ident, needsRecover := p.parseIdentifier()
		if needsRecover {
			return nil, true
		}
		return &ast.AssignFieldSpecifier{
			DotToken: dotToken,
			Ident:    ident,
		}, false
	case token.LSQUARE:
		// Index access specifier: [expr]
		openBracket := p.assertToken(token.LSQUARE)
		indexExpr, needsRecover := p.ParseExpr()
		if needsRecover {
			return nil, true
		}
		if p.curToken.Type != token.RSQUARE {
			p.errorToken("expected `]` in index specifier")
			return nil, true
		}
		closeBracket := p.assertToken(token.RSQUARE)
		return &ast.AssignIndexSpecifier{
			OpenBracket:  openBracket,
			IndexExpr:    indexExpr,
			CloseBracket: closeBracket,
		}, false
	default:
		return nil, true
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
