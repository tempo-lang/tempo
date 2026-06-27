package new_parser

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

// parseRoleType parses a role type which can be:
// - [A,B] (shared role type)
// - (A,B) or A (normal role type)
func (p *Parser) parseRoleType() (*ast.RoleType, bool) {
	var startToken token.Token
	var roles []*ast.Role
	var endToken token.Token

	switch p.curToken.Type {
	case token.LSQUARE:
		// Shared role type: [role1, role2, ...]
		startToken = p.assertToken(token.LSQUARE)

		// Parse first role
		role, needsRecover := p.parseRole()
		roles = append(roles, role)
		if needsRecover {
			return &ast.RoleType{
				Start: startToken,
				End:   token.Token{},
				Roles: roles,
			}, true
		}

		// Parse additional roles separated by commas
		for p.curToken.Type == token.COMMA {
			p.assertToken(token.COMMA)

			role, needsRecover := p.parseRole()
			roles = append(roles, role)
			if needsRecover {
				return &ast.RoleType{
					Start: startToken,
					End:   token.Token{},
					Roles: roles,
				}, true
			}
		}

		// Expect RSQUARE
		if p.curToken.Type != token.RSQUARE {
			return &ast.RoleType{
				Start: startToken,
				End:   p.errorToken("expected ']'"),
				Roles: roles,
			}, true
		}
		endToken = p.assertToken(token.RSQUARE)

	case token.LPAREN:
		// Parenthesized role type: (role1, role2, ...)
		startToken = p.assertToken(token.LPAREN)

		// Parse first role
		role, needsRecover := p.parseRole()
		roles = append(roles, role)
		if needsRecover {
			return &ast.RoleType{
				Start: startToken,
				End:   token.Token{},
				Roles: roles,
			}, true
		}

		// Parse additional roles separated by commas
		for p.curToken.Type == token.COMMA {
			p.assertToken(token.COMMA)

			role, needsRecover := p.parseRole()
			roles = append(roles, role)
			if needsRecover {
				return &ast.RoleType{
					Start: startToken,
					End:   token.Token{},
					Roles: roles,
				}, true
			}
		}

		// Expect RPAREN
		if p.curToken.Type != token.RPAREN {
			return &ast.RoleType{
				Start: startToken,
				End:   p.errorToken("expected ')'"),
				Roles: roles,
			}, true
		}
		endToken = p.assertToken(token.RPAREN)

	default:
		// Single role type: role
		role, needsRecover := p.parseRole()
		roles = append(roles, role)
		// For single role, both start and end are the role's token
		startToken = role.Token
		endToken = role.Token

		if needsRecover {
			return &ast.RoleType{
				Start: startToken,
				End:   endToken,
				Roles: roles,
			}, true
		}
	}

	return &ast.RoleType{
		Start: startToken,
		End:   endToken,
		Roles: roles,
	}, false
}

// parseRole parses a single role which can be an identifier or underscore
func (p *Parser) parseRole() (*ast.Role, bool) {
	if p.curTokenIs(token.IDENT, token.UNDERSCORE) {
		return &ast.Role{
			Token: p.readToken(),
		}, false
	}

	return &ast.Role{
		Token: p.errorToken("expected role (identifier or '_')"),
	}, true
}
