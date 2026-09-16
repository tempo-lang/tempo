package new_parser

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

func (p *Parser) parseValueType(stop TokenSet) ast.ValueType {
	switch p.current().Kind {
	case token.ASYNC:
		k := p.advance()
		return &ast.AsyncType{AsyncKeyword: k, Inner: p.parseValueType(stop)}
	case token.LSQUARE:
		o := p.advance()
		inner := p.parseValueType(stop.With(token.RSQUARE))
		c := p.expect(token.RSQUARE, stop)
		return &ast.ListType{OpenBracket: o, Inner: inner, CloseBracket: c}
	case token.IDENT:
		id := p.parseIdentifier()
		r := &ast.RoleIdent{Ident: id}
		if p.current().Kind == token.ROLE_AT {
			r.RoleAt = p.advance()
			r.RoleType = p.parseRoleType(stop)
		}
		return &ast.NamedType{RoleIdent: r}
	default:
		cur := p.current()
		p.add(Diagnostic{Code: CodeUnexpectedToken, Message: "expected type", PrimarySpan: cur.Span, Expected: []token.Kind{token.ASYNC, token.LSQUARE, token.IDENT, token.FUNC}, Found: cur.Kind})
		sk := p.skipUntil(stop)
		return &ast.InvalidType{Token: token.Missing(token.Ref(len(p.tokens)), token.IDENT, cur.Span.Start, p.source), Skipped: sk}
	}
}
func ParseType(s string) ProductionResult {
	p := FromString(s)
	return p.production(p.parseValueType(TokenSet{token.EOF}))
}
func ParseRoleType(s string) ProductionResult {
	p := FromString(s)
	return p.production(p.parseRoleType(TokenSet{token.EOF}))
}
func (p *Parser) parseRoleType(caller TokenSet) *ast.RoleType {
	var open token.Token
	var closeKind token.Kind
	switch p.current().Kind {
	case token.LSQUARE:
		open = p.advance()
		closeKind = token.RSQUARE
	case token.LPAREN:
		open = p.advance()
		closeKind = token.RPAREN
	default:
		r := p.parseRole(caller)
		return &ast.RoleType{Start: r.Token, End: r.Token, RoleNodes: []*ast.Role{r}}
	}
	r := &ast.RoleType{Start: open}
	stops := caller.Union(TokenSet{token.COMMA, closeKind, token.EOF})
	if p.current().Kind == closeKind {
		r.RoleNodes = append(r.RoleNodes, p.parseRole(stops))
	}
	for p.current().Kind != closeKind && p.current().Kind != token.EOF {
		pos := p.cursor.Position()
		r.RoleNodes = append(r.RoleNodes, p.parseRole(stops))
		if p.current().Kind == token.COMMA {
			p.advance()
			continue
		}
		if p.cursor.Position() == pos {
			break
		}
		if p.current().Kind != closeKind {
			break
		}
	}
	r.End = p.expect(closeKind, caller)
	return r
}
func (p *Parser) parseRole(stop TokenSet) *ast.Role {
	if p.curTokenIs(token.IDENT, token.UNDERSCORE) {
		return &ast.Role{Token: p.advance()}
	}
	cur := p.current()
	p.add(Diagnostic{Code: CodeExpectedRole, Message: "expected role", PrimarySpan: cur.Span, Expected: []token.Kind{token.IDENT, token.UNDERSCORE}, Found: cur.Kind})
	skipped := p.skipUntil(stop)
	return &ast.Role{Token: token.Missing(token.Ref(len(p.tokens)), token.IDENT, cur.Span.Start, p.source), Invalid: true, Skipped: skipped}
}
