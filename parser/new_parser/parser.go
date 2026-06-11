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

func (p *Parser) ParseScope() *ast.Scope {
	openToken := p.assertToken(token.LCURLY)

	stmts := []ast.Stmt{}

	for p.curToken.Type != token.RCURLY {
		stmt, needsRecover := p.parseStmt()
		stmts = append(stmts, stmt)

		if needsRecover {
			for !p.curTokenIs(token.RCURLY, token.SEMICOLON, token.EOF) {
				p.readToken()
			}
			if p.curTokenIs(token.SEMICOLON) {
				p.readToken()
			}
			if p.curTokenIs(token.EOF) {
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

	if p.curToken.Type != token.SEMICOLON {
		return &ast.InvalidStmt{
			ErrorToken:  token.Error(p.prevToken.Literal, p.prevToken.Pos, "missing semicolon"),
			PartialStmt: stmt,
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

	expr, needsRecover := p.parseExpr()
	stmt.Expr = expr
	if needsRecover {
		return &ast.InvalidStmt{
			ErrorToken:  expr.StartToken(),
			PartialStmt: stmt,
		}, true
	}

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
	return nil, false
}
