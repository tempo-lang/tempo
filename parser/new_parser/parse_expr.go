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
		ident, err := p.parseIdentifier()
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
