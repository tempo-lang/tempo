package new_parser

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

const (
	precLowest = iota
	precLogical
	precComparison
	precSum
	precProduct
	precPostfix
)

func precedence(k token.Kind) int {
	switch k {
	case token.OR, token.AND:
		return precLogical
	case token.EQUAL, token.NOT_EQUAL, token.LESS, token.LESS_EQ, token.GREATER, token.GREATER_EQ:
		return precComparison
	case token.PLUS, token.MINUS:
		return precSum
	case token.MULTIPLY, token.DIVIDE, token.MODULO:
		return precProduct
	}
	return precLowest
}
func exprStart(k token.Kind) bool {
	return k == token.FLOAT || k == token.INT || k == token.STRING || k == token.TRUE || k == token.FALSE || k == token.IDENT || k == token.LPAREN || k == token.LSQUARE || k == token.AWAIT
}
func (p *Parser) ParseExpr() (ast.Expr, bool) {
	before := len(p.diagnostics)
	e := p.parseExpr(TokenSet{token.EOF})
	p.expectEOF()
	return e, len(p.diagnostics) > before
}
func (p *Parser) parseExpr(stop TokenSet) ast.Expr { return p.parsePrecedence(stop, precLowest) }
func (p *Parser) parsePrecedence(stop TokenSet, min int) ast.Expr {
	left := p.parsePrefix(stop)
	for !stop.Contains(p.current().Kind) {
		switch p.current().Kind {
		case token.DOT:
			dot := p.advance()
			field := p.parseIdentifier()
			left = &ast.FieldAccessExpr{Object: left, DotToken: dot, Field: field}
			continue
		case token.LSQUARE:
			open := p.advance()
			idx := p.parseExpr(stop.Union(TokenSet{token.RSQUARE, token.COMMA, token.SEMICOLON, token.RPAREN, token.RCURLY, token.EOF}))
			close := p.expect(token.RSQUARE, stop.Union(TokenSet{token.COMMA, token.SEMICOLON, token.RPAREN, token.RCURLY, token.EOF}))
			left = &ast.IndexExpr{Object: left, OpenBracket: open, Index: idx, CloseBracket: close}
			continue
		case token.LPAREN:
			open := p.advance()
			var args []ast.Expr
			for p.current().Kind != token.RPAREN && p.current().Kind != token.EOF {
				pos := p.cursor.Position()
				args = append(args, p.parseExpr(TokenSet{token.COMMA, token.RPAREN, token.EOF}))
				if p.current().Kind == token.COMMA {
					p.advance()
					continue
				}
				if p.cursor.Position() == pos {
					break
				}
				break
			}
			close := p.expect(token.RPAREN, stop)
			left = &ast.CallExpr{Function: left, OpenParen: open, Args: args, CloseParen: close}
			continue
		}
		prec := precedence(p.current().Kind)
		if prec == precLowest || prec <= min {
			break
		}
		op := p.advance()
		right := p.parsePrecedence(stop, prec)
		left = &ast.BinaryExpr{Left: left, Operator: op, Right: right}
	}
	return left
}
func (p *Parser) parsePrefix(stop TokenSet) ast.Expr {
	cur := p.current()
	switch cur.Kind {
	case token.FLOAT:
		lit := &ast.FloatLit{FloatToken: p.advance()}
		return p.primitive(lit, stop)
	case token.INT:
		lit := &ast.IntLit{IntToken: p.advance()}
		return p.primitive(lit, stop)
	case token.STRING:
		lit := &ast.StringLit{StringToken: p.advance()}
		return p.primitive(lit, stop)
	case token.TRUE, token.FALSE:
		lit := &ast.BoolLit{BoolToken: p.advance()}
		return p.primitive(lit, stop)
	case token.IDENT:
		return p.parseIdentifier()
	case token.AWAIT:
		a := p.advance()
		return &ast.AwaitExpr{AwaitToken: a, Expr: p.parsePrecedence(stop, precPostfix-1)}
	case token.LPAREN:
		o := p.advance()
		e := p.parseExpr(stop.Union(TokenSet{token.RPAREN, token.COMMA, token.SEMICOLON, token.RSQUARE, token.RCURLY, token.EOF}))
		c := p.expect(token.RPAREN, stop)
		return &ast.GroupExpr{OpenParen: o, Expr: e, CloseParen: c}
	case token.LSQUARE:
		o := p.advance()
		var es []ast.Expr
		for p.current().Kind != token.RSQUARE && p.current().Kind != token.EOF {
			pos := p.cursor.Position()
			es = append(es, p.parseExpr(TokenSet{token.COMMA, token.RSQUARE, token.EOF}))
			if p.current().Kind == token.COMMA {
				p.advance()
				continue
			}
			if p.cursor.Position() == pos {
				break
			}
			break
		}
		c := p.expect(token.RSQUARE, stop)
		return &ast.ListExpr{OpenBracket: o, Elements: es, CloseBracket: c}
	default:
		p.add(Diagnostic{Code: CodeExpectedExpression, Message: "expected expression", PrimarySpan: cur.Span, Expected: []token.Kind{token.IDENT, token.INT, token.FLOAT, token.STRING, token.TRUE, token.FALSE, token.LPAREN, token.LSQUARE}, Found: cur.Kind})
		invalid := &ast.InvalidExpr{ErrorToken: token.Missing(token.Ref(len(p.tokens)), token.BadToken, cur.Span.Start, p.source)}
		if !stop.Contains(cur.Kind) && cur.Kind != token.EOF && closing(cur.Kind) == false {
			invalid.Skipped = []token.Ref{cur.Index}
			p.advance()
		}
		return invalid
	}
}
func (p *Parser) primitive(l ast.Literal, stop TokenSet) ast.Expr {
	r := &ast.PrimitiveExpr{Literal: l}
	if p.current().Kind == token.ROLE_AT {
		r.RoleAt = p.advance()
		r.RoleType = p.parseRoleType(stop)
	}
	return r
}
