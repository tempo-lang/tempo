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
			name:  "let decl block",
			input: "{let x = x;}",
			expected_ast: &ast.Scope{
				OpenToken:  makeToken(token.LCURLY, "{", nil, 1, 1),
				CloseToken: makeToken(token.RCURLY, "}", nil, 1, 12),
				Stmts: []ast.Stmt{
					&ast.LetStmt{
						LetToken:  makeToken(token.LET, "let", nil, 1, 2),
						Name:      &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 6)},
						Expr:      &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 10)},
						SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 11),
					},
				},
			},
		},
		{
			name:  "let with binary addition",
			input: "{let x = 1 + 2;}",
			expected_ast: &ast.Scope{
				OpenToken:  makeToken(token.LCURLY, "{", nil, 1, 1),
				CloseToken: makeToken(token.RCURLY, "}", nil, 1, 16),
				Stmts: []ast.Stmt{
					&ast.LetStmt{
						LetToken: makeToken(token.LET, "let", nil, 1, 2),
						Name:     &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 6)},
						Expr: &ast.BinaryExpr{
							Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 10)}},
							Operator: makeToken(token.PLUS, "+", nil, 1, 12),
							Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 14)}},
						},
						SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 15),
					},
				},
			},
		},
		{
			name:  "let with binary multiplication",
			input: "{let x = a * b;}",
			expected_ast: &ast.Scope{
				OpenToken:  makeToken(token.LCURLY, "{", nil, 1, 1),
				CloseToken: makeToken(token.RCURLY, "}", nil, 1, 16),
				Stmts: []ast.Stmt{
					&ast.LetStmt{
						LetToken: makeToken(token.LET, "let", nil, 1, 2),
						Name:     &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 6)},
						Expr: &ast.BinaryExpr{
							Left:     &ast.Identifier{Token: makeToken(token.IDENT, "a", "a", 1, 10)},
							Operator: makeToken(token.MULTIPLY, "*", nil, 1, 12),
							Right:    &ast.Identifier{Token: makeToken(token.IDENT, "b", "b", 1, 14)},
						},
						SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 15),
					},
				},
			},
		},
		{
			name:  "let with chained addition (left associative)",
			input: "{let x = 1 + 2 + 3;}",
			expected_ast: &ast.Scope{
				OpenToken:  makeToken(token.LCURLY, "{", nil, 1, 1),
				CloseToken: makeToken(token.RCURLY, "}", nil, 1, 20),
				Stmts: []ast.Stmt{
					&ast.LetStmt{
						LetToken: makeToken(token.LET, "let", nil, 1, 2),
						Name:     &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 6)},
						// (1 + 2) + 3
						Expr: &ast.BinaryExpr{
							Left: &ast.BinaryExpr{
								Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 10)}},
								Operator: makeToken(token.PLUS, "+", nil, 1, 12),
								Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 14)}},
							},
							Operator: makeToken(token.PLUS, "+", nil, 1, 16),
							Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "3", 3, 1, 18)}},
						},
						SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 19),
					},
				},
			},
		},
		{
			name:  "let with precedence: addition and multiplication",
			input: "{let x = 1 + 2 * 3;}",
			expected_ast: &ast.Scope{
				OpenToken:  makeToken(token.LCURLY, "{", nil, 1, 1),
				CloseToken: makeToken(token.RCURLY, "}", nil, 1, 20),
				Stmts: []ast.Stmt{
					&ast.LetStmt{
						LetToken: makeToken(token.LET, "let", nil, 1, 2),
						Name:     &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 6)},
						// 1 + (2 * 3)
						Expr: &ast.BinaryExpr{
							Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 10)}},
							Operator: makeToken(token.PLUS, "+", nil, 1, 12),
							Right: &ast.BinaryExpr{
								Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 14)}},
								Operator: makeToken(token.MULTIPLY, "*", nil, 1, 16),
								Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "3", 3, 1, 18)}},
							},
						},
						SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 19),
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
