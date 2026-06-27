package new_parser_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tempo-lang/tempo/parser/new_parser"
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

// stmtTestCase is a shared test case structure for statement parsing tests
type stmtTestCase struct {
	name          string
	input         string
	expected_stmt ast.Stmt
	expectError   bool
}

// runStmtTest is a shared test function for statement parsing
func runStmtTest(t *testing.T, tests []stmtTestCase) {
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

func TestParseLetStmt(t *testing.T) {
	// Helper to create a token for testing
	makeToken := func(tokenType token.TokenType, literal string, value any, line, col int) token.Token {
		return token.New(tokenType, literal, token.SourcePos{Line: line, Col: col}, value)
	}

	tests := []stmtTestCase{
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

	runStmtTest(t, tests)
}

func TestParseReturnStmt(t *testing.T) {
	// Helper to create a token for testing
	makeToken := func(tokenType token.TokenType, literal string, value any, line, col int) token.Token {
		return token.New(tokenType, literal, token.SourcePos{Line: line, Col: col}, value)
	}

	tests := []stmtTestCase{
		{
			name:  "return with expression",
			input: "return 42;",
			expected_stmt: &ast.ReturnStmt{
				ReturnToken: makeToken(token.RETURN, "return", nil, 1, 1),
				Expr: &ast.PrimitiveExpr{
					Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "42", 42, 1, 8)},
					RoleAt:   token.Token{},
					RoleType: nil,
				},
				SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 10),
			},
			expectError: false,
		},
		{
			name:  "return without expression",
			input: "return;",
			expected_stmt: &ast.ReturnStmt{
				ReturnToken: makeToken(token.RETURN, "return", nil, 1, 1),
				Expr:        nil,
				SemiToken:   makeToken(token.SEMICOLON, ";", nil, 1, 7),
			},
			expectError: false,
		},
		{
			name:        "return without semicolon",
			input:       "return 42",
			expectError: true,
		},
	}

	runStmtTest(t, tests)
}

func TestParseAssignStmt(t *testing.T) {
	// Helper to create a token for testing
	makeToken := func(tokenType token.TokenType, literal string, value any, line, col int) token.Token {
		return token.New(tokenType, literal, token.SourcePos{Line: line, Col: col}, value)
	}

	tests := []stmtTestCase{
		{
			name:  "simple assignment",
			input: "x = 42;",
			expected_stmt: &ast.AssignStmt{
				AssignExpr: &ast.AssignExpr{
					Ident:      &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 1)},
					Specifiers: nil,
				},
				AssignToken: makeToken(token.ASSIGN, "=", nil, 1, 3),
				Expr: &ast.PrimitiveExpr{
					Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "42", 42, 1, 5)},
					RoleAt:   token.Token{},
					RoleType: nil,
				},
				SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 7),
			},
			expectError: false,
		},
		{
			name:  "assignment with field access",
			input: "x.y = 100;",
			expected_stmt: &ast.AssignStmt{
				AssignExpr: &ast.AssignExpr{
					Ident: &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 1)},
					Specifiers: []ast.AssignSpecifier{
						&ast.AssignFieldSpecifier{
							DotToken: makeToken(token.DOT, ".", nil, 1, 2),
							Ident:    &ast.Identifier{Token: makeToken(token.IDENT, "y", "y", 1, 3)},
						},
					},
				},
				AssignToken: makeToken(token.ASSIGN, "=", nil, 1, 5),
				Expr: &ast.PrimitiveExpr{
					Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "100", 100, 1, 7)},
					RoleAt:   token.Token{},
					RoleType: nil,
				},
				SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 10),
			},
			expectError: false,
		},
		{
			name:  "assignment with index access",
			input: "arr[0] = 5;",
			expected_stmt: &ast.AssignStmt{
				AssignExpr: &ast.AssignExpr{
					Ident: &ast.Identifier{Token: makeToken(token.IDENT, "arr", "arr", 1, 1)},
					Specifiers: []ast.AssignSpecifier{
						&ast.AssignIndexSpecifier{
							OpenBracket: makeToken(token.LSQUARE, "[", nil, 1, 4),
							IndexExpr: &ast.PrimitiveExpr{
								Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "0", 0, 1, 5)},
								RoleAt:   token.Token{},
								RoleType: nil,
							},
							CloseBracket: makeToken(token.RSQUARE, "]", nil, 1, 6),
						},
					},
				},
				AssignToken: makeToken(token.ASSIGN, "=", nil, 1, 8),
				Expr: &ast.PrimitiveExpr{
					Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "5", 5, 1, 10)},
					RoleAt:   token.Token{},
					RoleType: nil,
				},
				SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 11),
			},
			expectError: false,
		},
		{
			name:        "assignment without semicolon",
			input:       "x = 42",
			expectError: true,
		},
		{
			name:        "assignment without equals",
			input:       "x 42;",
			expectError: true,
		},
	}

	runStmtTest(t, tests)
}
