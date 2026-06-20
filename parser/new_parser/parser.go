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

func (p *Parser) parseIdentifier() (*ast.Identifier, bool) {
	if p.curToken.Type != token.IDENT {
		return nil, true
	}
	return &ast.Identifier{
		Token: p.readToken(),
	}, false
}
