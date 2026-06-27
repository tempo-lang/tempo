package new_parser

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

// Precedence levels for binary operators (Pratt parsing)
const (
	PREC_LOWEST     = iota
	PREC_LOGICAL    // && and ||
	PREC_COMPARISON // ==, !=, <, <=, >, >=
	PREC_SUM        // + and -
	PREC_PRODUCT    // *, /, %
	PREC_FIELD      // . (field access)
	PREC_INDEX      // [] (index access)
)

// getPrecedence returns the precedence level for a token type
func getPrecedence(tokenType token.TokenType) int {
	switch tokenType {
	case token.OR, token.AND:
		return PREC_LOGICAL
	case token.EQUAL, token.NOT_EQUAL, token.LESS, token.LESS_EQ, token.GREATER, token.GREATER_EQ:
		return PREC_COMPARISON
	case token.PLUS, token.MINUS:
		return PREC_SUM
	case token.MULTIPLY, token.DIVIDE, token.MODULO:
		return PREC_PRODUCT
	case token.DOT:
		return PREC_FIELD
	case token.LSQUARE:
		return PREC_INDEX
	default:
		return PREC_LOWEST
	}
}

func (p *Parser) ParseExpr() (ast.Expr, bool) {
	return p.parseExprWithPrecedence(PREC_LOWEST)
}

func (p *Parser) parseExprWithPrecedence(precedence int) (ast.Expr, bool) {
	left, needsRecover := p.parsePrimaryExpr()
	if needsRecover {
		return left, true
	}

	// While the next token is a binary operator with precedence > current precedence level
	// Strict inequality combined with passing currentPrec to the recursive call ensures left-associativity
	for precedence < getPrecedence(p.curToken.Type) {
		op := p.readToken()

		// Parse the right-hand side with the same precedence for the current operator
		// Same-precedence operators are not consumed by the recursive call due to strict inequality
		currentPrec := getPrecedence(op.Type)
		right, needsRecover := p.parseExprWithPrecedence(currentPrec)
		if needsRecover {
			return left, true
		}

		// Wrap in BinaryExpr and continue parsing
		left = &ast.BinaryExpr{
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}

	return left, false
}

func (p *Parser) parseLiteralWithRole(lit ast.Literal) (ast.Expr, bool) {
	// Check for optional role annotation: ROLE_AT roleType
	if p.curToken.Type == token.ROLE_AT {
		roleAtToken := p.readToken()
		roleType, needsRecover := p.parseRoleType()
		return &ast.PrimitiveExpr{
			Literal:  lit,
			RoleAt:   roleAtToken,
			RoleType: roleType,
		}, needsRecover
	}

	// No role annotation, return the literal directly as a PrimitiveExpr
	return &ast.PrimitiveExpr{
		Literal:  lit,
		RoleAt:   token.Token{},
		RoleType: nil,
	}, false
}

func (p *Parser) parsePrimaryExpr() (ast.Expr, bool) {
	switch p.curToken.Type {
	case token.FLOAT:
		lit := &ast.FloatLit{FloatToken: p.readToken()}
		return p.parseLiteralWithRole(lit)
	case token.INT:
		lit := &ast.IntLit{IntToken: p.readToken()}
		return p.parseLiteralWithRole(lit)
	case token.STRING:
		lit := &ast.StringLit{StringToken: p.readToken()}
		return p.parseLiteralWithRole(lit)
	case token.TRUE, token.FALSE:
		lit := &ast.BoolLit{BoolToken: p.readToken()}
		return p.parseLiteralWithRole(lit)
	case token.IDENT:
		ident, err := p.parseIdentifierOrAccess()
		if err {
			return nil, true
		}
		return ident, false
	default:
		err := p.errorToken("not an expression")
		return &ast.InvalidExpr{
			ErrorToken:  err,
			PartialExpr: nil,
		}, true
	}
}

// parseIdentifierOrAccess parses an identifier followed by optional field/index access
func (p *Parser) parseIdentifierOrAccess() (ast.Expr, bool) {
	// Start with identifier
	ident, needsRecover := p.parseIdentifier()
	if needsRecover {
		return nil, true
	}

	var expr ast.Expr = ident

	// Handle chained field/index access
	for {
		switch p.curToken.Type {
		case token.DOT:
			// Parse field access: .field
			fieldExpr, needsRecover := p.parseFieldAccess(expr)
			if needsRecover {
				return expr, true // Return what we have so far
			}
			expr = fieldExpr
		case token.LSQUARE:
			// Parse index access: [index]
			indexExpr, needsRecover := p.parseIndexAccess(expr)
			if needsRecover {
				return expr, true // Return what we have so far
			}
			expr = indexExpr
		default:
			return expr, false
		}
	}
}

// parseFieldAccess parses a field access expression: object.field
func (p *Parser) parseFieldAccess(object ast.Expr) (ast.Expr, bool) {
	dotToken := p.assertToken(token.DOT)

	field, needsRecover := p.parseIdentifier()
	if needsRecover {
		return &ast.InvalidExpr{
			ErrorToken: p.errorToken("expected field name after dot"),
		}, true
	}

	return &ast.FieldAccessExpr{
		Object:   object,
		DotToken: dotToken,
		Field:    field,
	}, false
}

// parseIndexAccess parses an index access expression: object[index]
func (p *Parser) parseIndexAccess(object ast.Expr) (ast.Expr, bool) {
	openBracket := p.assertToken(token.LSQUARE)

	index, needsRecover := p.ParseExpr()
	if needsRecover {
		return &ast.InvalidExpr{
			ErrorToken: p.errorToken("expected expression in index"),
		}, true
	}

	if p.curToken.Type != token.RSQUARE {
		errToken := p.errorToken("expected closing bracket")
		return &ast.InvalidExpr{
			ErrorToken: errToken,
		}, true
	}
	closeBracket := p.readToken()

	return &ast.IndexExpr{
		Object:       object,
		OpenBracket:  openBracket,
		Index:        index,
		CloseBracket: closeBracket,
	}, false
}
