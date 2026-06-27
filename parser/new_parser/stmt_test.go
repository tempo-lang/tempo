package new_parser_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tempo-lang/tempo/parser/new_parser"
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

func TestParseLetStmt(t *testing.T) {
	// Helper to create a token for testing
	makeToken := func(tokenType token.TokenType, literal string, value any, line, col int) token.Token {
		return token.New(tokenType, literal, token.SourcePos{Line: line, Col: col}, value)
	}

	tests := []struct {
		name          string
		input         string
		expected_stmt ast.Stmt
		expectError   bool
	}{
		{
			name:  "let declaration with string expression",
			input: "let x = \"hello\";",
			expected_stmt: &ast.LetStmt{
				LetToken: makeToken(token.LET, "let", nil, 1, 1),
				Name:     &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 5)},
				Expr: &ast.PrimitiveExpr{
					Literal:  &ast.StringLit{StringToken: makeToken(token.STRING, "\"hello\"", "hello", 1, 9)},
					RoleAt:   token.Token{},
					RoleType: nil,
				},
				SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 16),
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := new_parser.FromString(test.input)
			actual_stmt, needsRecover := p.ParseStmt()

			if test.expectError {
				if !needsRecover {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if needsRecover {
				t.Fatalf("Got parser error when expecting success: %v", p.ParserErrors())
			}

			if len(p.ParserErrors()) > 0 {
				t.Fatalf("Got %d parser errors: %v", len(p.ParserErrors()), p.ParserErrors())
			}

			if diff := cmp.Diff(test.expected_stmt, actual_stmt); diff != "" {
				t.Errorf("Statement mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}
