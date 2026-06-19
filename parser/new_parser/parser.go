package new_parser

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/lexer"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

// Precedence levels for binary operators (Pratt parsing)
const (
	PREC_LOWEST  = iota
	PREC_SUM     // + and -
	PREC_PRODUCT // * and /
	PREC_PREFIX  // unary operators (future use)
)

type Parser struct {
	l      *lexer.Lexer
	errors []token.Token

	prevToken token.Token
	curToken  token.Token
}

func New(sourceReader io.RuneReader) *Parser {
	p := &Parser{
		l: lexer.New(sourceReader),
	}

	// read current token
	p.readToken()

	return p
}

func FromString(source string) *Parser {
	return New(strings.NewReader(source))
}

func (p *Parser) ParserErrors() []token.Token {
	return p.errors
}

func (p *Parser) readToken() token.Token {
	token := p.curToken
	p.prevToken = p.curToken
	p.curToken = p.l.ReadToken()
	return token
}

func (p *Parser) curTokenIs(tokenTypes ...token.TokenType) bool {
	return slices.Contains(tokenTypes, p.curToken.Type)
}

// assertToken will read the next token and panic if it does not match the expected token type.
// This function should only be used when an earlier check ensures that the token type is correct.
func (p *Parser) assertToken(tokenType token.TokenType) token.Token {
	if p.curToken.Type != tokenType {
		panic(fmt.Sprintf("expected token type '%s' got '%v'", tokenType, p.curToken))
	}
	return p.readToken()
}

func (p *Parser) errorToken(errorMessage string) token.Token {
	errToken := token.Error(
		p.curToken.Literal,
		p.curToken.Pos,
		errorMessage,
	)
	p.errors = append(p.errors, errToken)
	return errToken
}

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

func (p *Parser) expectSemicolon(stmt ast.Stmt) (actualToken token.Token, errorStmt *ast.InvalidStmt) {
	actualToken = p.curToken
	if actualToken.Type != token.SEMICOLON {
		errToken := token.Error(p.prevToken.Literal, p.prevToken.Pos, "missing semicolon")
		p.errors = append(p.errors, errToken)
		errorStmt = &ast.InvalidStmt{
			ErrorToken:  errToken,
			PartialStmt: stmt,
		}
	} else {
		p.readToken() // consume semicolon if successfully found
	}
	return // output variables assigned
}

func (p *Parser) ParseScope() *ast.Scope {
	openToken := p.assertToken(token.LCURLY)

	stmts := []ast.Stmt{}

	for p.curToken.Type != token.RCURLY {
		if p.curToken.Type == token.EOF {
			return &ast.Scope{
				OpenToken:  openToken,
				CloseToken: p.errorToken("unexpected EOF when parsing scope"),
				Stmts:      stmts,
			}
		}

		stmt, needsRecover := p.parseStmt()
		stmts = append(stmts, stmt)

		if needsRecover {
			for !p.curTokenIs(token.RCURLY, token.SEMICOLON, token.EOF) {
				p.readToken()
			}
			switch p.curToken.Type {
			case token.RCURLY:
				break
			case token.SEMICOLON:
				p.readToken()
			case token.EOF:
				return &ast.Scope{
					OpenToken:  openToken,
					CloseToken: p.errorToken("unexpected EOF when parsing scope"),
					Stmts:      stmts,
				}
			}
		}
	}

	closeToken := p.assertToken(token.RCURLY)

	return &ast.Scope{
		OpenToken:  openToken,
		CloseToken: closeToken,
		Stmts:      stmts,
	}
}

func (p *Parser) parseStmt() (ast.Stmt, bool) {
	var stmt ast.Stmt
	switch p.curToken.Type {
	case token.LET:
		letStmt, needsRecovery := p.parseLetStmt()
		stmt = letStmt
		if needsRecovery {
			return letStmt, true
		}
	default:
		return &ast.InvalidStmt{
			ErrorToken: p.errorToken("invalid statement"),
		}, true
	}

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

	expr, needsRecover := p.parseExpr()
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

func (p *Parser) parseIdentifier() (*ast.Identifier, bool) {
	if p.curToken.Type != token.IDENT {
		return nil, true
	}
	return &ast.Identifier{
		Token: p.readToken(),
	}, false
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
