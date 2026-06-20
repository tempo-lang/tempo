package new_parser

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

// Precedence levels for binary operators (Pratt parsing)
const (
	PREC_LOWEST  = iota
	PREC_SUM     // + and -
	PREC_PRODUCT // * and /
	PREC_PREFIX  // unary operators (future use)
)

// getPrecedence returns the precedence level for a token type
func getPrecedence(tokenType token.TokenType) int {
	switch tokenType {
	case token.PLUS, token.MINUS:
		return PREC_SUM
	case token.MULTIPLY, token.DIVIDE:
		return PREC_PRODUCT
	default:
		return PREC_LOWEST
	}
}

func (p *Parser) parseExpr() (ast.Expr, bool) {
	return p.parseExprWithPrecedence(PREC_LOWEST)
}

func (p *Parser) parseExprWithPrecedence(precedence int) (ast.Expr, bool) {
	left, needsRecover := p.parsePrimaryExpr()
	if needsRecover {
		return left, true
	}

	// While the next token is a binary operator with precedence > current precedence level
	// Using strict inequality ensures left-associativity for operators at the same precedence
	for precedence < getPrecedence(p.curToken.Type) {
		op := p.readToken()

		// Parse the right-hand side with higher precedence for left-associative operators
		right, needsRecover := p.parseExprWithPrecedence(precedence + 1)
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

func (p *Parser) parsePrimaryExpr() (ast.Expr, bool) {
	switch p.curToken.Type {
	case token.FLOAT:
		return &ast.FloatExpr{FloatToken: p.readToken()}, false
	case token.INT:
		return &ast.IntExpr{IntToken: p.readToken()}, false
	case token.IDENT:
		return p.parseIdentifier()
	default:
		err := p.errorToken("not an expression")
		return &ast.InvalidExpr{
			ErrorToken:  err,
			PartialExpr: nil,
		}, true
	}
}
