package new_parser_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tempo-lang/tempo/parser/new_parser"
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

func TestParser(t *testing.T) {
	// Helper to create a token for testing
	makeToken := func(tokenType token.TokenType, literal string, line, col int) token.Token {
		// For identifiers, the value is the literal string
		var value any
		if tokenType == token.IDENT {
			value = literal
		}
		return token.New(tokenType, literal, token.SourcePos{Line: line, Col: col}, value)
	}

	tests := []struct {
		name         string
		input        string
		expected_ast *ast.Scope
	}{
		{
			name:  "empty block",
			input: "{}",
			expected_ast: &ast.Scope{
				OpenToken:  makeToken(token.LCURLY, "{", 1, 1),
				CloseToken: makeToken(token.RCURLY, "}", 1, 2),
				Stmts:      []ast.Stmt{},
			},
		},
		{
			name:  "let decl block",
			input: "{let x = x;}",
			expected_ast: &ast.Scope{
				OpenToken:  makeToken(token.LCURLY, "{", 1, 1),
				CloseToken: makeToken(token.RCURLY, "}", 1, 12),
				Stmts: []ast.Stmt{
					&ast.LetStmt{
						LetToken:  makeToken(token.LET, "let", 1, 2),
						Name:      &ast.Identifier{Token: makeToken(token.IDENT, "x", 1, 6)},
						Expr:      &ast.Identifier{Token: makeToken(token.IDENT, "x", 1, 10)},
						SemiToken: makeToken(token.SEMICOLON, ";", 1, 11),
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := new_parser.FromString(test.input)
			actual_ast := p.ParseScope()

			if len(p.ParserErrors()) > 0 {
				t.Fatalf("Got %d parser errors: %v", len(p.ParserErrors()), p.ParserErrors())
			}

			if diff := cmp.Diff(test.expected_ast, actual_ast); diff != "" {
				t.Errorf("AST mismatch (-expected +actual):\n%s", diff)
			}
		})
	}

}
