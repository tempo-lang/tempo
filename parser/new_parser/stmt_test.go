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
				LHS:         &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 1)},
				AssignToken: makeToken(token.ASSIGN, "=", nil, 1, 3),
				RHS: &ast.PrimitiveExpr{
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
				LHS: &ast.FieldAccessExpr{
					Object:   &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 1)},
					DotToken: makeToken(token.DOT, ".", nil, 1, 2),
					Field:    &ast.Identifier{Token: makeToken(token.IDENT, "y", "y", 1, 3)},
				},
				AssignToken: makeToken(token.ASSIGN, "=", nil, 1, 5),
				RHS: &ast.PrimitiveExpr{
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
				LHS: &ast.IndexExpr{
					Object:      &ast.Identifier{Token: makeToken(token.IDENT, "arr", "arr", 1, 1)},
					OpenBracket: makeToken(token.LSQUARE, "[", nil, 1, 4),
					Index: &ast.PrimitiveExpr{
						Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "0", 0, 1, 5)},
						RoleAt:   token.Token{},
						RoleType: nil,
					},
					CloseBracket: makeToken(token.RSQUARE, "]", nil, 1, 6),
				},
				AssignToken: makeToken(token.ASSIGN, "=", nil, 1, 8),
				RHS: &ast.PrimitiveExpr{
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
		{
			name:  "assignment with chained field and index access",
			input: "x.arr[0].field = 42;",
			expected_stmt: &ast.AssignStmt{
				LHS: &ast.FieldAccessExpr{
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
				AssignToken: makeToken(token.ASSIGN, "=", nil, 1, 16),
				RHS: &ast.PrimitiveExpr{
					Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "42", 42, 1, 18)},
					RoleAt:   token.Token{},
					RoleType: nil,
				},
				SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 20),
			},
			expectError: false,
		},
	}

	runStmtTest(t, tests)
}

func TestParseExprStmt(t *testing.T) {
	// Helper to create a token for testing
	makeToken := func(tokenType token.TokenType, literal string, value any, line, col int) token.Token {
		return token.New(tokenType, literal, token.SourcePos{Line: line, Col: col}, value)
	}

	tests := []stmtTestCase{
		{
			name:  "simple expression statement",
			input: "42;",
			expected_stmt: &ast.ExprStmt{
				Expr: &ast.PrimitiveExpr{
					Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "42", 42, 1, 1)},
					RoleAt:   token.Token{},
					RoleType: nil,
				},
				SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 3),
			},
			expectError: false,
		},
		{
			name:  "binary expression statement",
			input: "a + b;",
			expected_stmt: &ast.ExprStmt{
				Expr: &ast.BinaryExpr{
					Left:     &ast.Identifier{Token: makeToken(token.IDENT, "a", "a", 1, 1)},
					Operator: makeToken(token.PLUS, "+", nil, 1, 3),
					Right:    &ast.Identifier{Token: makeToken(token.IDENT, "b", "b", 1, 5)},
				},
				SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 6),
			},
			expectError: false,
		},
		{
			name:  "field access expression statement",
			input: "x.y;",
			expected_stmt: &ast.ExprStmt{
				Expr: &ast.FieldAccessExpr{
					Object:   &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 1)},
					DotToken: makeToken(token.DOT, ".", nil, 1, 2),
					Field:    &ast.Identifier{Token: makeToken(token.IDENT, "y", "y", 1, 3)},
				},
				SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 4),
			},
			expectError: false,
		},
		{
			name:  "index access expression statement",
			input: "arr[0];",
			expected_stmt: &ast.ExprStmt{
				Expr: &ast.IndexExpr{
					Object:      &ast.Identifier{Token: makeToken(token.IDENT, "arr", "arr", 1, 1)},
					OpenBracket: makeToken(token.LSQUARE, "[", nil, 1, 4),
					Index: &ast.PrimitiveExpr{
						Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "0", 0, 1, 5)},
						RoleAt:   token.Token{},
						RoleType: nil,
					},
					CloseBracket: makeToken(token.RSQUARE, "]", nil, 1, 6),
				},
				SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 7),
			},
			expectError: false,
		},
	}

	runStmtTest(t, tests)
}

func TestParseIfStmt(t *testing.T) {
	// Helper to create a token for testing
	makeToken := func(tokenType token.TokenType, literal string, value any, line, col int) token.Token {
		return token.New(tokenType, literal, token.SourcePos{Line: line, Col: col}, value)
	}

	tests := []stmtTestCase{
		{
			name:  "simple if statement",
			input: "if x > 0 { return 1; }",
			expected_stmt: &ast.IfStmt{
				IfToken: makeToken(token.IF, "if", nil, 1, 1),
				Condition: &ast.BinaryExpr{
					Left:     &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 4)},
					Operator: makeToken(token.GREATER, ">", nil, 1, 6),
					Right: &ast.PrimitiveExpr{
						Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "0", 0, 1, 8)},
						RoleAt:   token.Token{},
						RoleType: nil,
					},
				},
				ThenScope: &ast.Scope{
					OpenToken:  makeToken(token.LCURLY, "{", nil, 1, 10),
					CloseToken: makeToken(token.RCURLY, "}", nil, 1, 22),
					Stmts: []ast.Stmt{
						&ast.ReturnStmt{
							ReturnToken: makeToken(token.RETURN, "return", nil, 1, 12),
							Expr: &ast.PrimitiveExpr{
								Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 19)},
								RoleAt:   token.Token{},
								RoleType: nil,
							},
							SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 20),
						},
					},
				},
				ElseToken: token.Token{},
				ElseScope: nil,
			},
			expectError: false,
		},
		{
			name:  "if with else",
			input: "if x > 0 { return 1; } else { return 0; }",
			expected_stmt: &ast.IfStmt{
				IfToken: makeToken(token.IF, "if", nil, 1, 1),
				Condition: &ast.BinaryExpr{
					Left:     &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 4)},
					Operator: makeToken(token.GREATER, ">", nil, 1, 6),
					Right: &ast.PrimitiveExpr{
						Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "0", 0, 1, 8)},
						RoleAt:   token.Token{},
						RoleType: nil,
					},
				},
				ThenScope: &ast.Scope{
					OpenToken:  makeToken(token.LCURLY, "{", nil, 1, 10),
					CloseToken: makeToken(token.RCURLY, "}", nil, 1, 22),
					Stmts: []ast.Stmt{
						&ast.ReturnStmt{
							ReturnToken: makeToken(token.RETURN, "return", nil, 1, 12),
							Expr: &ast.PrimitiveExpr{
								Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 19)},
								RoleAt:   token.Token{},
								RoleType: nil,
							},
							SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 20),
						},
					},
				},
				ElseToken: makeToken(token.ELSE, "else", nil, 1, 24),
				ElseScope: &ast.Scope{
					OpenToken:  makeToken(token.LCURLY, "{", nil, 1, 29),
					CloseToken: makeToken(token.RCURLY, "}", nil, 1, 41),
					Stmts: []ast.Stmt{
						&ast.ReturnStmt{
							ReturnToken: makeToken(token.RETURN, "return", nil, 1, 31),
							Expr: &ast.PrimitiveExpr{
								Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "0", 0, 1, 38)},
								RoleAt:   token.Token{},
								RoleType: nil,
							},
							SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 39),
						},
					},
				},
			},
			expectError: false,
		},
		{
			name:  "nested if statements",
			input: "if x > 0 { if y > 0 { return 1; } }",
			expected_stmt: &ast.IfStmt{
				IfToken: makeToken(token.IF, "if", nil, 1, 1),
				Condition: &ast.BinaryExpr{
					Left:     &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 4)},
					Operator: makeToken(token.GREATER, ">", nil, 1, 6),
					Right: &ast.PrimitiveExpr{
						Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "0", 0, 1, 8)},
						RoleAt:   token.Token{},
						RoleType: nil,
					},
				},
				ThenScope: &ast.Scope{
					OpenToken:  makeToken(token.LCURLY, "{", nil, 1, 10),
					CloseToken: makeToken(token.RCURLY, "}", nil, 1, 35),
					Stmts: []ast.Stmt{
						&ast.IfStmt{
							IfToken: makeToken(token.IF, "if", nil, 1, 12),
							Condition: &ast.BinaryExpr{
								Left:     &ast.Identifier{Token: makeToken(token.IDENT, "y", "y", 1, 15)},
								Operator: makeToken(token.GREATER, ">", nil, 1, 17),
								Right: &ast.PrimitiveExpr{
									Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "0", 0, 1, 19)},
									RoleAt:   token.Token{},
									RoleType: nil,
								},
							},
							ThenScope: &ast.Scope{
								OpenToken:  makeToken(token.LCURLY, "{", nil, 1, 21),
								CloseToken: makeToken(token.RCURLY, "}", nil, 1, 33),
								Stmts: []ast.Stmt{
									&ast.ReturnStmt{
										ReturnToken: makeToken(token.RETURN, "return", nil, 1, 23),
										Expr: &ast.PrimitiveExpr{
											Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 30)},
											RoleAt:   token.Token{},
											RoleType: nil,
										},
										SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 31),
									},
								},
							},
							ElseToken: token.Token{},
							ElseScope: nil,
						},
					},
				},
				ElseToken: token.Token{},
				ElseScope: nil,
			},
			expectError: false,
		},
		{
			name:        "if without condition",
			input:       "if { return 1; }",
			expectError: true,
		},
		{
			name:        "if without scope",
			input:       "if x > 0 return 1;",
			expectError: true,
		},
	}

	runStmtTest(t, tests)
}

func TestParseWhileStmt(t *testing.T) {
	// Helper to create a token for testing
	makeToken := func(tokenType token.TokenType, literal string, value any, line, col int) token.Token {
		return token.New(tokenType, literal, token.SourcePos{Line: line, Col: col}, value)
	}

	tests := []stmtTestCase{
		{
			name:  "simple while statement",
			input: "while x > 0 { x = x - 1; }",
			expected_stmt: &ast.WhileStmt{
				WhileKeyword: makeToken(token.WHILE, "while", nil, 1, 1),
				Condition: &ast.BinaryExpr{
					Left:     &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 7)},
					Operator: makeToken(token.GREATER, ">", nil, 1, 9),
					Right: &ast.PrimitiveExpr{
						Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "0", 0, 1, 11)},
						RoleAt:   token.Token{},
						RoleType: nil,
					},
				},
				Scope: &ast.Scope{
					OpenToken:  makeToken(token.LCURLY, "{", nil, 1, 13),
					CloseToken: makeToken(token.RCURLY, "}", nil, 1, 26),
					Stmts: []ast.Stmt{
						&ast.AssignStmt{
							LHS:         &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 15)},
							AssignToken: makeToken(token.ASSIGN, "=", nil, 1, 17),
							RHS: &ast.BinaryExpr{
								Left:     &ast.Identifier{Token: makeToken(token.IDENT, "x", "x", 1, 19)},
								Operator: makeToken(token.MINUS, "-", nil, 1, 21),
								Right: &ast.PrimitiveExpr{
									Literal:  &ast.IntLit{IntToken: makeToken(token.INT, "1", 1, 1, 23)},
									RoleAt:   token.Token{},
									RoleType: nil,
								},
							},
							SemiToken: makeToken(token.SEMICOLON, ";", nil, 1, 24),
						},
					},
				},
			},
			expectError: false,
		},
		{
			name:        "while without condition",
			input:       "while { x = x - 1; }",
			expectError: true,
		},
		{
			name:        "while without scope",
			input:       "while x > 0 x = x - 1;",
			expectError: true,
		},
	}

	runStmtTest(t, tests)
}
