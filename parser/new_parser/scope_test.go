package new_parser_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tempo-lang/tempo/parser/new_parser"
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

func TestScope(t *testing.T) {
	// Helper to create a token for testing
	makeToken := func(tokenType token.TokenType, literal string, value any, line, col int) token.Token {
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
				OpenToken:  makeToken(token.LCURLY, "{", nil, 1, 1),
				CloseToken: makeToken(token.RCURLY, "}", nil, 1, 2),
				Stmts:      []ast.Stmt{},
			},
		},
		{
			name:  "two let statements",
			input: "{let x = 1; let y = 2;}",
			expected_ast: &ast.Scope{
				OpenToken:  makeToken(token.LCURLY, "{", nil, 1, 1),
				CloseToken: makeToken(token.RCURLY, "}", nil, 1, 23),
				Stmts: []ast.Stmt{
					&ast.LetStmt{
						LetToken: makeToken(token.LET, "let", nil, 1, 2),
						Name:     &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 6)},
						Expr: &ast.PrimitiveExpr{
							Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 10)},
							RoleAt:   token.Token{},
							RoleType: nil,
						},
						SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 11),
					},
					&ast.LetStmt{
						LetToken: makeToken(token.LET, "let", nil, 1, 13),
						Name:     &ast.Identifier{Token: makeToken(token.IDENT, "y", "y", 1, 17)},
						Expr: &ast.PrimitiveExpr{
							Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 21)},
							RoleAt:   token.Token{},
							RoleType: nil,
						},
						SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 22),
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

func TestParseScopeWithErrors(t *testing.T) {
	// Helper to create an expected error token
	makeError := func(literal string, line, col int, message string) token.Token {
		return token.Error(literal, token.SourcePos{Line: line, Col: col}, message)
	}

	tests := []struct {
		name           string
		input          string
		expectedErrors []token.Token
	}{
		{
			name:  "missing semicolon",
			input: "{let x = 1}",
			expectedErrors: []token.Token{
				makeError("1", 1, 10, "missing semicolon"),
			},
		},
		{
			name:  "missing equals in let",
			input: "{let x 1;}",
			expectedErrors: []token.Token{
				makeError("1", 1, 8, "expected `=` when parsing let statement"),
			},
		},
		{
			name:  "invalid expression in let",
			input: "{let x = +;}",
			expectedErrors: []token.Token{
				makeError("+", 1, 10, "not an expression"),
			},
		},
		{
			name:  "unclosed scope",
			input: "{let x = 1;",
			expectedErrors: []token.Token{
				makeError("", 1, 11, "unexpected EOF when parsing scope"),
			},
		},
		{
			name:  "missing identifier in let",
			input: "{let = 1;}",
			expectedErrors: []token.Token{
				makeError("=", 1, 6, "expected identifier when parsing let statement"),
			},
		},
		{
			name:  "empty let name",
			input: "{let x = ; let y = 10;}",
			expectedErrors: []token.Token{
				makeError(";", 1, 10, "not an expression"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := new_parser.FromString(test.input)
			p.ParseScope()

			actualErrors := p.ParserErrors()
			if diff := cmp.Diff(test.expectedErrors, actualErrors); diff != "" {
				t.Errorf("ParserErrors mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}
