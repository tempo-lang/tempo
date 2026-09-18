package parser

import (
	"github.com/tempo-lang/tempo/parser/ast"
	"github.com/tempo-lang/tempo/parser/token"
)

var declarationStart = TokenSet{token.FUNC, token.STRUCT, token.INTERFACE, token.EOF}

func (p *Parser) parseFuncSig() *ast.FuncSig {
	s := &ast.FuncSig{FuncToken: p.expect(token.FUNC, TokenSet{token.ROLE_AT, token.IDENT})}
	if p.current().Kind == token.ROLE_AT {
		p.advance()
		s.RoleType = p.parseRoleType(TokenSet{token.IDENT})
	}
	s.Name = p.parseIdentifier()
	s.Params = p.parseFuncParams()
	if valueTypeStart(p.current().Kind) {
		s.ReturnType = p.parseValueType(TokenSet{token.LCURLY, token.SEMICOLON, token.EOF})
	}
	return s
}
func (p *Parser) parseClosureSig() *ast.ClosureSig {
	s := &ast.ClosureSig{FuncToken: p.expect(token.FUNC, TokenSet{token.ROLE_AT}), RoleAtToken: p.expect(token.ROLE_AT, TokenSet{token.IDENT, token.UNDERSCORE, token.LSQUARE, token.LPAREN})}
	s.RoleType = p.parseRoleType(TokenSet{token.LPAREN})
	s.Params = p.parseFuncParams()
	if valueTypeStart(p.current().Kind) {
		s.ReturnType = p.parseValueType(TokenSet{token.LCURLY, token.EOF})
	}
	return s
}
func (p *Parser) parseFuncParams() *ast.FuncParams {
	r := &ast.FuncParams{OpenParen: p.expect(token.LPAREN, TokenSet{token.RPAREN, token.LCURLY})}
	for p.current().Kind != token.RPAREN && p.current().Kind != token.EOF {
		n := p.parseIdentifier()
		p.expect(token.COLON, TokenSet{token.ASYNC, token.LSQUARE, token.FUNC, token.IDENT})
		typ := p.parseValueType(TokenSet{token.COMMA, token.RPAREN, token.EOF})
		r.Params = append(r.Params, &ast.FuncParam{Name: n, Type: typ})
		if p.current().Kind != token.COMMA {
			break
		}
		p.advance()
		if p.current().Kind == token.RPAREN {
			p.trailingComma(token.RPAREN)
			break
		}
	}
	r.CloseParen = p.expect(token.RPAREN, TokenSet{token.LCURLY, token.SEMICOLON, token.ASYNC, token.LSQUARE, token.FUNC, token.IDENT})
	return r
}
func (p *Parser) parseFunc() *ast.Func {
	return &ast.Func{FuncSig: p.parseFuncSig(), Scope: p.parseScope()}
}

func (p *Parser) parseStruct() *ast.Struct {
	s := &ast.Struct{StructToken: p.advance()}
	if p.current().Kind == token.ROLE_AT {
		p.advance()
		s.RoleType = p.parseRoleType(TokenSet{token.IDENT})
	}
	s.Name = p.parseIdentifier()
	if p.current().Kind == token.IMPLEMENTS {
		p.advance()
		for {
			s.Implements = append(s.Implements, p.parseRoleIdent(TokenSet{token.COMMA, token.LCURLY}))
			if p.current().Kind != token.COMMA {
				break
			}
			p.advance()
			if p.current().Kind == token.LCURLY {
				p.trailingComma(token.LCURLY)
				break
			}
		}
	}
	s.Body = p.parseStructBody()
	return s
}
func (p *Parser) parseStructBody() *ast.StructBody {
	b := &ast.StructBody{OpenToken: p.expect(token.LCURLY, TokenSet{token.IDENT, token.FUNC, token.RCURLY})}
	for p.current().Kind != token.RCURLY && p.current().Kind != token.EOF {
		pos := p.cursor.Position()
		if p.current().Kind == token.FUNC {
			f := p.parseFunc()
			b.Functions = append(b.Functions, f)
			b.Members = append(b.Members, f)
		} else if p.current().Kind == token.IDENT {
			n := p.parseIdentifier()
			c := p.expect(token.COLON, TokenSet{token.ASYNC, token.LSQUARE, token.FUNC, token.IDENT})
			typ := p.parseValueType(TokenSet{token.SEMICOLON, token.RCURLY, token.FUNC, token.IDENT})
			semi := p.semicolon()
			f := &ast.StructField{Name: n, Colon: c, Type: typ, SemiToken: semi}
			b.Fields = append(b.Fields, f)
			b.Members = append(b.Members, f)
		} else {
			p.add(Diagnostic{Code: CodeUnexpectedToken, Message: "expected struct member", PrimarySpan: p.current().Span, Expected: []token.Kind{token.IDENT, token.FUNC}, Found: p.current().Kind})
			p.skipUntil(TokenSet{token.IDENT, token.FUNC, token.RCURLY})
		}
		if p.cursor.Position() == pos {
			p.advance()
		}
	}
	b.CloseToken = p.expect(token.RCURLY, declarationStart)
	return b
}
func (p *Parser) parseInterface() *ast.Interface {
	i := &ast.Interface{InterfaceKeyword: p.advance()}
	if p.current().Kind == token.ROLE_AT {
		p.advance()
		i.RoleType = p.parseRoleType(TokenSet{token.IDENT})
	}
	i.Name = p.parseIdentifier()
	m := &ast.InterfaceMethodsList{OpenToken: p.expect(token.LCURLY, TokenSet{token.FUNC, token.RCURLY})}
	for p.current().Kind != token.RCURLY && p.current().Kind != token.EOF {
		if p.current().Kind != token.FUNC {
			p.add(Diagnostic{Code: CodeUnexpectedToken, Message: "expected interface method", PrimarySpan: p.current().Span, Expected: []token.Kind{token.FUNC}, Found: p.current().Kind})
			p.skipUntil(TokenSet{token.FUNC, token.RCURLY})
			continue
		}
		sig := p.parseFuncSig()
		m.Methods = append(m.Methods, &ast.InterfaceMethod{FuncSig: sig, SemiToken: p.semicolon()})
	}
	m.CloseToken = p.expect(token.RCURLY, declarationStart)
	i.Methods = m
	return i
}
