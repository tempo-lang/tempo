package new_parser_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tempo-lang/tempo/parser/new_parser"
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

func TestParseExpr(t *testing.T) {
	// Helper to create a token for testing
	makeToken := func(tokenType token.TokenType, literal string, value any, line, col int) token.Token {
		return token.New(tokenType, literal, token.SourcePos{Line: line, Col: col}, value)
	}

	tests := []struct {
		name          string
		input         string
		expected_expr ast.Expr
		expectError   bool
	}{
		{
			name:          "identifier",
			input:         "x",
			expected_expr: &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 1)},
			expectError:   false,
		},
		{
			name:  "integer literal",
			input: "42",
			expected_expr: &ast.PrimitiveExpr{
				Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "42", 42, 1, 1)},
				RoleAt:   token.Token{},
				RoleType: nil,
			},
			expectError: false,
		},
		{
			name:  "float literal",
			input: "3.14",
			expected_expr: &ast.PrimitiveExpr{
				Literal:  &ast.FloatLit{FloatToken: makeToken(token.FLOAT, "3.14", 3.14, 1, 1)},
				RoleAt:   token.Token{},
				RoleType: nil,
			},
			expectError: false,
		},
		{
			name:  "string literal",
			input: "\"hello\"",
			expected_expr: &ast.PrimitiveExpr{
				Literal:  &ast.StringLit{StringToken: makeToken(token.STRING, "\"hello\"", "hello", 1, 1)},
				RoleAt:   token.Token{},
				RoleType: nil,
			},
			expectError: false,
		},
		{
			name:  "boolean literal true",
			input: "true",
			expected_expr: &ast.PrimitiveExpr{
				Literal:  &ast.BoolLit{BoolToken: makeToken(token.TRUE, "true", true, 1, 1)},
				RoleAt:   token.Token{},
				RoleType: nil,
			},
			expectError: false,
		},
		{
			name:  "boolean literal false",
			input: "false",
			expected_expr: &ast.PrimitiveExpr{
				Literal:  &ast.BoolLit{BoolToken: makeToken(token.FALSE, "false", false, 1, 1)},
				RoleAt:   token.Token{},
				RoleType: nil,
			},
			expectError: false,
		},
		{
			name:  "binary addition",
			input: "1 + 2",
			expected_expr: &ast.BinaryExpr{
				Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
				Operator: makeToken(token.PLUS, "+", nil, 1, 3),
				Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 5)}},
			},
			expectError: false,
		},
		{
			name:  "binary multiplication",
			input: "a * b",
			expected_expr: &ast.BinaryExpr{
				Left:     &ast.Identifier{Token: makeToken(token.IDENT, "a", "a", 1, 1)},
				Operator: makeToken(token.MULTIPLY, "*", nil, 1, 3),
				Right:    &ast.Identifier{Token: makeToken(token.IDENT, "b", "b", 1, 5)},
			},
			expectError: false,
		},
		{
			name:  "chained addition (left associative)",
			input: "1 + 2 + 3",
			// (1 + 2) + 3
			expected_expr: &ast.BinaryExpr{
				Left: &ast.BinaryExpr{
					Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
					Operator: makeToken(token.PLUS, "+", nil, 1, 3),
					Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 5)}},
				},
				Operator: makeToken(token.PLUS, "+", nil, 1, 7),
				Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "3", 3, 1, 9)}},
			},
			expectError: false,
		},
		{
			name:  "precedence: addition and multiplication",
			input: "1 + 2 * 3",
			// 1 + (2 * 3)
			expected_expr: &ast.BinaryExpr{
				Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
				Operator: makeToken(token.PLUS, "+", nil, 1, 3),
				Right: &ast.BinaryExpr{
					Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 5)}},
					Operator: makeToken(token.MULTIPLY, "*", nil, 1, 7),
					Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "3", 3, 1, 9)}},
				},
			},
			expectError: false,
		},
		{
			name:  "literal with single role annotation",
			input: "42@A",
			expected_expr: &ast.PrimitiveExpr{
				Literal: &ast.IntLit{IntToken: makeToken(token.INT, "42", 42, 1, 1)},
				RoleAt:  makeToken(token.ROLE_AT, "@", nil, 1, 3),
				RoleType: &ast.RoleType{
					Start: makeToken(token.IDENT, "A", "A", 1, 4),
					End:   makeToken(token.IDENT, "A", "A", 1, 4),
					RoleNodes: []*ast.Role{
						{Token: makeToken(token.IDENT, "A", "A", 1, 4)},
					},
				},
			},
			expectError: false,
		},
		{
			name:  "literal with shared role annotation",
			input: "42@[A,B]",
			expected_expr: &ast.PrimitiveExpr{
				Literal: &ast.IntLit{IntToken: makeToken(token.INT, "42", 42, 1, 1)},
				RoleAt:  makeToken(token.ROLE_AT, "@", nil, 1, 3),
				RoleType: &ast.RoleType{
					Start: makeToken(token.LSQUARE, "[", nil, 1, 4),
					End:   makeToken(token.RSQUARE, "]", nil, 1, 8),
					RoleNodes: []*ast.Role{
						{Token: makeToken(token.IDENT, "A", "A", 1, 5)},
						{Token: makeToken(token.IDENT, "B", "B", 1, 7)},
					},
				},
			},
			expectError: false,
		},
		{
			name:  "string literal with role annotation",
			input: "\"hello\"@A",
			expected_expr: &ast.PrimitiveExpr{
				Literal: &ast.StringLit{StringToken: makeToken(token.STRING, "\"hello\"", "hello", 1, 1)},
				RoleAt:  makeToken(token.ROLE_AT, "@", nil, 1, 8),
				RoleType: &ast.RoleType{
					Start: makeToken(token.IDENT, "A", "A", 1, 9),
					End:   makeToken(token.IDENT, "A", "A", 1, 9),
					RoleNodes: []*ast.Role{
						{Token: makeToken(token.IDENT, "A", "A", 1, 9)},
					},
				},
			},
			expectError: false,
		},
		{
			name:  "binary expression with role annotation",
			input: "1@A + 2@B",
			expected_expr: &ast.BinaryExpr{
				Left: &ast.PrimitiveExpr{
					Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)},
					RoleAt:  makeToken(token.ROLE_AT, "@", nil, 1, 2),
					RoleType: &ast.RoleType{
						Start: makeToken(token.IDENT, "A", "A", 1, 3),
						End:   makeToken(token.IDENT, "A", "A", 1, 3),
						RoleNodes: []*ast.Role{
							{Token: makeToken(token.IDENT, "A", "A", 1, 3)},
						},
					},
				},
				Operator: makeToken(token.PLUS, "+", nil, 1, 5),
				Right: &ast.PrimitiveExpr{
					Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 7)},
					RoleAt:  makeToken(token.ROLE_AT, "@", nil, 1, 8),
					RoleType: &ast.RoleType{
						Start: makeToken(token.IDENT, "B", "B", 1, 9),
						End:   makeToken(token.IDENT, "B", "B", 1, 9),
						RoleNodes: []*ast.Role{
							{Token: makeToken(token.IDENT, "B", "B", 1, 9)},
						},
					},
				},
			},
			expectError: false,
		},
		// New binary operator tests
		{
			name:  "binary modulo",
			input: "1 % 2",
			expected_expr: &ast.BinaryExpr{
				Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
				Operator: makeToken(token.MODULO, "%", nil, 1, 3),
				Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 5)}},
			},
			expectError: false,
		},
		{
			name:  "binary equality",
			input: "1 == 2",
			expected_expr: &ast.BinaryExpr{
				Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
				Operator: makeToken(token.EQUAL, "==", nil, 1, 3),
				Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 6)}},
			},
			expectError: false,
		},
		{
			name:  "binary not equal",
			input: "1 != 2",
			expected_expr: &ast.BinaryExpr{
				Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
				Operator: makeToken(token.NOT_EQUAL, "!=", nil, 1, 3),
				Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 6)}},
			},
			expectError: false,
		},
		{
			name:  "binary less than",
			input: "1 < 2",
			expected_expr: &ast.BinaryExpr{
				Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
				Operator: makeToken(token.LESS, "<", nil, 1, 3),
				Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 5)}},
			},
			expectError: false,
		},
		{
			name:  "binary less than or equal",
			input: "1 <= 2",
			expected_expr: &ast.BinaryExpr{
				Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
				Operator: makeToken(token.LESS_EQ, "<=", nil, 1, 3),
				Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 6)}},
			},
			expectError: false,
		},
		{
			name:  "binary greater than",
			input: "1 > 2",
			expected_expr: &ast.BinaryExpr{
				Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
				Operator: makeToken(token.GREATER, ">", nil, 1, 3),
				Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 5)}},
			},
			expectError: false,
		},
		{
			name:  "binary greater than or equal",
			input: "1 >= 2",
			expected_expr: &ast.BinaryExpr{
				Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
				Operator: makeToken(token.GREATER_EQ, ">=", nil, 1, 3),
				Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 6)}},
			},
			expectError: false,
		},
		{
			name:  "binary logical and",
			input: "true && false",
			expected_expr: &ast.BinaryExpr{
				Left:     &ast.PrimitiveExpr{Literal: &ast.BoolLit{BoolToken: makeToken(token.TRUE, "true", true, 1, 1)}},
				Operator: makeToken(token.AND, "&&", nil, 1, 6),
				Right:    &ast.PrimitiveExpr{Literal: &ast.BoolLit{BoolToken: makeToken(token.FALSE, "false", false, 1, 9)}},
			},
			expectError: false,
		},
		{
			name:  "binary logical or",
			input: "true || false",
			expected_expr: &ast.BinaryExpr{
				Left:     &ast.PrimitiveExpr{Literal: &ast.BoolLit{BoolToken: makeToken(token.TRUE, "true", true, 1, 1)}},
				Operator: makeToken(token.OR, "||", nil, 1, 6),
				Right:    &ast.PrimitiveExpr{Literal: &ast.BoolLit{BoolToken: makeToken(token.FALSE, "false", false, 1, 9)}},
			},
			expectError: false,
		},
		// Precedence tests for new operators
		{
			name:  "precedence: product and modulo",
			input: "1 * 2 % 3",
			// (1 * 2) % 3 - left associative at same precedence
			expected_expr: &ast.BinaryExpr{
				Left: &ast.BinaryExpr{
					Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
					Operator: makeToken(token.MULTIPLY, "*", nil, 1, 3),
					Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 5)}},
				},
				Operator: makeToken(token.MODULO, "%", nil, 1, 7),
				Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "3", 3, 1, 9)}},
			},
			expectError: false,
		},
		{
			name:  "precedence: sum and product with modulo",
			input: "1 + 2 * 3 % 4",
			// 1 + ((2 * 3) % 4)
			expected_expr: &ast.BinaryExpr{
				Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
				Operator: makeToken(token.PLUS, "+", nil, 1, 3),
				Right: &ast.BinaryExpr{
					Left: &ast.BinaryExpr{
						Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 5)}},
						Operator: makeToken(token.MULTIPLY, "*", nil, 1, 7),
						Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "3", 3, 1, 9)}},
					},
					Operator: makeToken(token.MODULO, "%", nil, 1, 11),
					Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "4", 4, 1, 13)}},
				},
			},
			expectError: false,
		},
		{
			name:  "precedence: comparison and sum",
			input: "1 + 2 == 3 * 4",
			// (1 + 2) == (3 * 4)
			expected_expr: &ast.BinaryExpr{
				Left: &ast.BinaryExpr{
					Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
					Operator: makeToken(token.PLUS, "+", nil, 1, 3),
					Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 5)}},
				},
				Operator: makeToken(token.EQUAL, "==", nil, 1, 7),
				Right: &ast.BinaryExpr{
					Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "3", 3, 1, 10)}},
					Operator: makeToken(token.MULTIPLY, "*", nil, 1, 12),
					Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "4", 4, 1, 14)}},
				},
			},
			expectError: false,
		},
		{
			name:  "precedence: comparison chaining",
			input: "1 == 2 != 3",
			// (1 == 2) != 3 - left associative at same precedence
			expected_expr: &ast.BinaryExpr{
				Left: &ast.BinaryExpr{
					Left:     &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 1)}},
					Operator: makeToken(token.EQUAL, "==", nil, 1, 3),
					Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "2", 2, 1, 6)}},
				},
				Operator: makeToken(token.NOT_EQUAL, "!=", nil, 1, 8),
				Right:    &ast.PrimitiveExpr{Literal: &ast.IntLit{IntToken: makeToken(token.INT, "3", 3, 1, 11)}},
			},
			expectError: false,
		},
		{
			name:  "precedence: logical operators",
			input: "true && false || true",
			// (true && false) || true - left associative at same precedence
			expected_expr: &ast.BinaryExpr{
				Left: &ast.BinaryExpr{
					Left:     &ast.PrimitiveExpr{Literal: &ast.BoolLit{BoolToken: makeToken(token.TRUE, "true", true, 1, 1)}},
					Operator: makeToken(token.AND, "&&", nil, 1, 6),
					Right:    &ast.PrimitiveExpr{Literal: &ast.BoolLit{BoolToken: makeToken(token.FALSE, "false", false, 1, 9)}},
				},
				Operator: makeToken(token.OR, "||", nil, 1, 15),
				Right:    &ast.PrimitiveExpr{Literal: &ast.BoolLit{BoolToken: makeToken(token.TRUE, "true", true, 1, 18)}},
			},
			expectError: false,
		},
		{
			name:  "precedence: complex mixed operators",
			input: "a * b + c == d && e || f",
			// (((a * b) + c) == d) && e) || f
			expected_expr: &ast.BinaryExpr{
				Left: &ast.BinaryExpr{
					Left: &ast.BinaryExpr{
						Left: &ast.BinaryExpr{
							Left: &ast.BinaryExpr{
								Left:     &ast.Identifier{Token: makeToken(token.IDENT, "a", "a", 1, 1)},
								Operator: makeToken(token.MULTIPLY, "*", nil, 1, 3),
								Right:    &ast.Identifier{Token: makeToken(token.IDENT, "b", "b", 1, 5)},
							},
							Operator: makeToken(token.PLUS, "+", nil, 1, 7),
							Right:    &ast.Identifier{Token: makeToken(token.IDENT, "c", "c", 1, 9)},
						},
						Operator: makeToken(token.EQUAL, "==", nil, 1, 11),
						Right:    &ast.Identifier{Token: makeToken(token.IDENT, "d", "d", 1, 14)},
					},
					Operator: makeToken(token.AND, "&&", nil, 1, 16),
					Right:    &ast.Identifier{Token: makeToken(token.IDENT, "e", "e", 1, 19)},
				},
				Operator: makeToken(token.OR, "||", nil, 1, 21),
				Right:    &ast.Identifier{Token: makeToken(token.IDENT, "f", "f", 1, 24)},
			},
			expectError: false,
		},
		{
			name:  "field access expression",
			input: "x.y",
			expected_expr: &ast.FieldAccessExpr{
				Object:   &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 1)},
				DotToken: makeToken(token.DOT, ".", nil, 1, 2),
				Field:    &ast.Identifier{Token: makeToken(token.IDENT, "y", "y", 1, 3)},
			},
			expectError: false,
		},
		{
			name:  "index access expression",
			input: "arr[0]",
			expected_expr: &ast.IndexExpr{
				Object:      &ast.Identifier{Token: makeToken(token.IDENT, "arr", "arr", 1, 1)},
				OpenBracket: makeToken(token.LSQUARE, "[", nil, 1, 4),
				Index: &ast.PrimitiveExpr{
					Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "0", 0, 1, 5)},
					RoleAt:   token.Token{},
					RoleType: nil,
				},
				CloseBracket: makeToken(token.RSQUARE, "]", nil, 1, 6),
			},
			expectError: false,
		},
		{
			name:  "chained field access",
			input: "x.y.z",
			expected_expr: &ast.FieldAccessExpr{
				Object: &ast.FieldAccessExpr{
					Object:   &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 1)},
					DotToken: makeToken(token.DOT, ".", nil, 1, 2),
					Field:    &ast.Identifier{Token: makeToken(token.IDENT, "y", "y", 1, 3)},
				},
				DotToken: makeToken(token.DOT, ".", nil, 1, 4),
				Field:    &ast.Identifier{Token: makeToken(token.IDENT, "z", "z", 1, 5)},
			},
			expectError: false,
		},
		{
			name:  "chained index access",
			input: "arr[0][1]",
			expected_expr: &ast.IndexExpr{
				Object: &ast.IndexExpr{
					Object:      &ast.Identifier{Token: makeToken(token.IDENT, "arr", "arr", 1, 1)},
					OpenBracket: makeToken(token.LSQUARE, "[", nil, 1, 4),
					Index: &ast.PrimitiveExpr{
						Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "0", 0, 1, 5)},
						RoleAt:   token.Token{},
						RoleType: nil,
					},
					CloseBracket: makeToken(token.RSQUARE, "]", nil, 1, 6),
				},
				OpenBracket: makeToken(token.LSQUARE, "[", nil, 1, 7),
				Index: &ast.PrimitiveExpr{
					Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 8)},
					RoleAt:   token.Token{},
					RoleType: nil,
				},
				CloseBracket: makeToken(token.RSQUARE, "]", nil, 1, 9),
			},
			expectError: false,
		},
		{
			name:  "mixed field and index access",
			input: "x.arr[0].field",
			expected_expr: &ast.FieldAccessExpr{
				Object: &ast.IndexExpr{
					Object: &ast.FieldAccessExpr{
						Object:   &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 1)},
						DotToken: makeToken(token.DOT, ".", nil, 1, 2),
						Field:    &ast.Identifier{Token: makeToken(token.IDENT, "arr", "arr", 1, 3)},
					},
					OpenBracket: makeToken(token.LSQUARE, "[", nil, 1, 6),
					Index: &ast.PrimitiveExpr{
						Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "0", 0, 1, 7)},
						RoleAt:   token.Token{},
						RoleType: nil,
					},
					CloseBracket: makeToken(token.RSQUARE, "]", nil, 1, 8),
				},
				DotToken: makeToken(token.DOT, ".", nil, 1, 9),
				Field:    &ast.Identifier{Token: makeToken(token.IDENT, "field", "field", 1, 10)},
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := new_parser.FromString(test.input)
			actual_expr, needsRecover := p.ParseExpr()

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

			if diff := cmp.Diff(test.expected_expr, actual_expr); diff != "" {
				t.Errorf("Expression mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}
