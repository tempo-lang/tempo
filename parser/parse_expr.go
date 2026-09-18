package parser

import (
	"github.com/tempo-lang/tempo/parser/ast"
	"github.com/tempo-lang/tempo/parser/token"
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
	return k == token.FLOAT || k == token.INT || k == token.STRING || k == token.TRUE || k == token.FALSE || k == token.IDENT || k == token.UNDERSCORE || k == token.LPAREN || k == token.LSQUARE || k == token.AWAIT || k == token.FUNC
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
					if p.current().Kind == token.RPAREN {
						p.trailingComma(token.RPAREN)
						break
					}
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
		if p.roleTypeFollowedByCom() {
			return p.parseComExpr(stop)
		}
		ri := p.parseRoleIdent(stop.Union(TokenSet{token.LCURLY, token.DOT, token.LSQUARE, token.LPAREN}))
		// A brace following an identifier is also the boundary between an
		// if/while condition and its scope. In that context only commit to a
		// struct literal when the body visibly starts with `name:`.
		if p.current().Kind == token.LCURLY && (!stop.Contains(token.LCURLY) || (p.peek(1).Kind == token.IDENT && p.peek(2).Kind == token.COLON)) {
			return p.parseStructExpr(ri, stop)
		}
		return &ast.IdentAccessExpr{Ident: ri.Ident, RoleAt: ri.RoleAt, RoleType: ri.RoleType}
	case token.UNDERSCORE:
		if p.roleTypeFollowedByCom() {
			return p.parseComExpr(stop)
		}
		fallthrough
	case token.FUNC:
		if cur.Kind == token.FUNC {
			sig := p.parseClosureSig()
			return &ast.ClosureExpr{ClosureSig: sig, Scope: p.parseScope()}
		}
		fallthrough
	case token.AWAIT:
		a := p.advance()
		return &ast.AwaitExpr{AwaitToken: a, Expr: p.parsePrecedence(stop, precPostfix-1)}
	case token.LPAREN:
		if p.roleTypeFollowedByCom() {
			return p.parseComExpr(stop)
		}
		o := p.advance()
		e := p.parseExpr(stop.Union(TokenSet{token.RPAREN, token.COMMA, token.SEMICOLON, token.RSQUARE, token.RCURLY, token.EOF}))
		c := p.expect(token.RPAREN, stop)
		return &ast.GroupExpr{OpenParen: o, Expr: e, CloseParen: c}
	case token.LSQUARE:
		if p.roleTypeFollowedByCom() {
			return p.parseComExpr(stop)
		}
		o := p.advance()
		var es []ast.Expr
		for p.current().Kind != token.RSQUARE && p.current().Kind != token.EOF {
			pos := p.cursor.Position()
			es = append(es, p.parseExpr(TokenSet{token.COMMA, token.RSQUARE, token.EOF}))
			if p.current().Kind == token.COMMA {
				p.advance()
				if p.current().Kind == token.RSQUARE {
					p.trailingComma(token.RSQUARE)
					break
				}
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

func (p *Parser) roleTypeFollowedByCom() bool {
	i := p.cursor.Position()
	k := p.peek(0).Kind
	if k == token.IDENT || k == token.UNDERSCORE {
		return p.peek(1).Kind == token.COM
	}
	var close token.Kind
	switch k {
	case token.LPAREN:
		close = token.RPAREN
	case token.LSQUARE:
		close = token.RSQUARE
	default:
		return false
	}
	depth := 0
	for n := 0; i+n < len(p.tokens); n++ {
		q := p.peek(n).Kind
		if q == k {
			depth++
		}
		if q == close {
			depth--
			if depth == 0 {
				return p.peek(n+1).Kind == token.COM
			}
		}
		if q == token.EOF {
			return false
		}
	}
	return false
}

func (p *Parser) parseComExpr(stop TokenSet) ast.Expr {
	s := p.parseRoleType(TokenSet{token.COM})
	c := p.expect(token.COM, TokenSet{token.IDENT, token.UNDERSCORE, token.LPAREN, token.LSQUARE})
	r := p.parseRoleType(stop)
	return &ast.ComExpr{Sender: s, ComToken: c, Receiver: r, Expr: p.parsePrecedence(stop, precLowest)}
}

func (p *Parser) parseStructExpr(ri *ast.RoleIdent, stop TokenSet) ast.Expr {
	f := &ast.StructFields{OpenToken: p.advance()}
	for p.current().Kind != token.RCURLY && p.current().Kind != token.EOF {
		n := p.parseIdentifier()
		c := p.expect(token.COLON, TokenSet{token.COMMA, token.RCURLY}.Union(stop))
		e := p.parseExpr(TokenSet{token.COMMA, token.RCURLY, token.EOF})
		f.Fields = append(f.Fields, &ast.StructFieldExpr{Name: n, Colon: c, Expr: e})
		if p.current().Kind != token.COMMA {
			break
		}
		p.advance()
		if p.current().Kind == token.RCURLY {
			p.trailingComma(token.RCURLY)
			break
		}
	}
	f.CloseToken = p.expect(token.RCURLY, stop)
	return &ast.StructExpr{RoleIdent: ri, StructFields: f}
}

func (p *Parser) primitive(l ast.Literal, stop TokenSet) ast.Expr {
	r := &ast.PrimitiveExpr{Literal: l}
	if p.current().Kind == token.ROLE_AT {
		r.RoleAt = p.advance()
		r.RoleType = p.parseRoleType(stop)
	}
	return r
}
