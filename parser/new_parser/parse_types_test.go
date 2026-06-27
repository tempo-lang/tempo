package new_parser

import (
	"testing"

	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

func TestParseRoleType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.RoleType
		wantErr  bool
	}{
		{
			name:  "single role",
			input: "A",
			expected: &ast.RoleType{
				Start: token.New(token.IDENT, "A", token.SourcePos{Line: 1, Col: 1}, nil),
				End:   token.New(token.IDENT, "A", token.SourcePos{Line: 1, Col: 1}, nil),
				RoleNodes: []*ast.Role{{
					Token: token.New(token.IDENT, "A", token.SourcePos{Line: 1, Col: 1}, nil),
				}},
			},
			wantErr: false,
		},
		{
			name:  "underscore role",
			input: "_",
			expected: &ast.RoleType{
				Start: token.New(token.UNDERSCORE, "_", token.SourcePos{Line: 1, Col: 1}, nil),
				End:   token.New(token.UNDERSCORE, "_", token.SourcePos{Line: 1, Col: 1}, nil),
				RoleNodes: []*ast.Role{{
					Token: token.New(token.UNDERSCORE, "_", token.SourcePos{Line: 1, Col: 1}, nil),
				}},
			},
			wantErr: false,
		},
		{
			name:  "shared role type with two roles",
			input: "[A,B]",
			expected: &ast.RoleType{
				Start: token.New(token.LSQUARE, "[", token.SourcePos{Line: 1, Col: 1}, nil),
				End:   token.New(token.RSQUARE, "]", token.SourcePos{Line: 1, Col: 4}, nil),
				RoleNodes: []*ast.Role{
					{Token: token.New(token.IDENT, "A", token.SourcePos{Line: 1, Col: 2}, nil)},
					{Token: token.New(token.IDENT, "B", token.SourcePos{Line: 1, Col: 4}, nil)},
				},
			},
			wantErr: false,
		},
		{
			name:  "parenthesized role type with two roles",
			input: "(A,B)",
			expected: &ast.RoleType{
				Start: token.New(token.LPAREN, "(", token.SourcePos{Line: 1, Col: 1}, nil),
				End:   token.New(token.RPAREN, ")", token.SourcePos{Line: 1, Col: 4}, nil),
				RoleNodes: []*ast.Role{
					{Token: token.New(token.IDENT, "A", token.SourcePos{Line: 1, Col: 2}, nil)},
					{Token: token.New(token.IDENT, "B", token.SourcePos{Line: 1, Col: 4}, nil)},
				},
			},
			wantErr: false,
		},
		{
			name:  "shared role type with single role",
			input: "[A]",
			expected: &ast.RoleType{
				Start: token.New(token.LSQUARE, "[", token.SourcePos{Line: 1, Col: 1}, nil),
				End:   token.New(token.RSQUARE, "]", token.SourcePos{Line: 1, Col: 3}, nil),
				RoleNodes: []*ast.Role{
					{Token: token.New(token.IDENT, "A", token.SourcePos{Line: 1, Col: 2}, nil)},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := FromString(tt.input)
			roleType, needsRecover := p.parseRoleType()

			if needsRecover != tt.wantErr {
				t.Errorf("parseRoleType() needsRecover = %v, wantErr %v", needsRecover, tt.wantErr)
				return
			}

			if needsRecover {
				return
			}

			if roleType == nil {
				t.Errorf("parseRoleType() returned nil, expected non-nil")
				return
			}

			// Check number of role nodes
			if len(roleType.RoleNodes) != len(tt.expected.RoleNodes) {
				t.Errorf("parseRoleType() RoleNodes length = %d, expected %d", len(roleType.RoleNodes), len(tt.expected.RoleNodes))
				return
			}

			// Check IsShared() method
			if roleType.IsShared() != (roleType.Start.Type == token.LSQUARE) {
				t.Errorf("parseRoleType() IsShared() method not working correctly")
				return
			}

			// For single role case, check that start and end are the same
			if len(tt.expected.RoleNodes) == 1 && tt.expected.Start.Type != token.LSQUARE && tt.expected.Start.Type != token.LPAREN {
				if roleType.Start.Type != roleType.End.Type {
					t.Errorf("parseRoleType() single role: Start.Type = %v, End.Type = %v, expected same", roleType.Start.Type, roleType.End.Type)
				}
			}
		})
	}
}

func TestParseRoleTypeRecovery(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectError    bool
		expectRoleType bool
		checkFn        func(t *testing.T, roleType *ast.RoleType, p *Parser)
	}{
		{
			name:           "unclosed shared role type",
			input:          "[A,B",
			expectError:    true,
			expectRoleType: true,
			checkFn: func(t *testing.T, roleType *ast.RoleType, p *Parser) {
				if roleType == nil {
					t.Error("expected roleType to be non-nil for error recovery")
					return
				}
				if roleType.Start.Type != token.LSQUARE {
					t.Errorf("expected Start.Type to be LSQUARE, got %v", roleType.Start.Type)
				}
				if roleType.End.Type != token.ILLEGAL {
					t.Errorf("expected End.Type to be ILLEGAL (error token), got %v", roleType.End.Type)
				}
				if len(roleType.RoleNodes) != 2 {
					t.Errorf("expected 2 role nodes, got %d", len(roleType.RoleNodes))
				}
				// Check that parser has errors
				if len(p.ParserErrors()) == 0 {
					t.Error("expected parser to have errors")
				}
				// Use the Error() helper method
				if roleType.Error() == nil {
					t.Error("expected roleType.Error() to return non-nil")
				}
			},
		},
		{
			name:           "unclosed parenthesized role type",
			input:          "(A,B",
			expectError:    true,
			expectRoleType: true,
			checkFn: func(t *testing.T, roleType *ast.RoleType, p *Parser) {
				if roleType == nil {
					t.Error("expected roleType to be non-nil for error recovery")
					return
				}
				if roleType.Start.Type != token.LPAREN {
					t.Errorf("expected Start.Type to be LPAREN, got %v", roleType.Start.Type)
				}
				if roleType.End.Type != token.ILLEGAL {
					t.Errorf("expected End.Type to be ILLEGAL (error token), got %v", roleType.End.Type)
				}
				if len(roleType.RoleNodes) != 2 {
					t.Errorf("expected 2 role nodes, got %d", len(roleType.RoleNodes))
				}
				// Check that parser has errors
				if len(p.ParserErrors()) == 0 {
					t.Error("expected parser to have errors")
				}
				// Use the Error() helper method
				if roleType.Error() == nil {
					t.Error("expected roleType.Error() to return non-nil")
				}
			},
		},
		{
			name:           "invalid role in shared type",
			input:          "[A,123]",
			expectError:    true,
			expectRoleType: true,
			checkFn: func(t *testing.T, roleType *ast.RoleType, p *Parser) {
				if roleType == nil {
					t.Error("expected roleType to be non-nil for error recovery")
					return
				}
				if roleType.Start.Type != token.LSQUARE {
					t.Errorf("expected Start.Type to be LSQUARE, got %v", roleType.Start.Type)
				}
				// Should have at least one valid role (A) and one error role (123)
				if len(roleType.RoleNodes) != 2 {
					t.Errorf("expected 2 role nodes, got %d", len(roleType.RoleNodes))
				}
				// The second role should have an error token
				if roleType.RoleNodes[1].Token.Type != token.ILLEGAL {
					t.Errorf("expected second role to have error token, got %v", roleType.RoleNodes[1].Token.Type)
				}
				// Check that parser has errors
				if len(p.ParserErrors()) == 0 {
					t.Error("expected parser to have errors")
				}
				// Use the Error() helper method
				if roleType.Error() == nil {
					t.Error("expected roleType.Error() to return non-nil")
				}
				// Use the Roles() helper method to get valid roles
				validRoles := roleType.Roles()
				if len(validRoles) != 1 {
					t.Errorf("expected 1 valid role, got %d", len(validRoles))
				}
			},
		},
		{
			name:           "invalid single role",
			input:          "123",
			expectError:    true,
			expectRoleType: true,
			checkFn: func(t *testing.T, roleType *ast.RoleType, p *Parser) {
				if roleType == nil {
					t.Error("expected roleType to be non-nil for error recovery")
					return
				}
				// Should have one role with error token
				if len(roleType.RoleNodes) != 1 {
					t.Errorf("expected 1 role node, got %d", len(roleType.RoleNodes))
				}
				if roleType.RoleNodes[0].Token.Type != token.ILLEGAL {
					t.Errorf("expected role to have error token, got %v", roleType.RoleNodes[0].Token.Type)
				}
				// Check that parser has errors
				if len(p.ParserErrors()) == 0 {
					t.Error("expected parser to have errors")
				}
				// Use the Error() helper method
				if roleType.Error() == nil {
					t.Error("expected roleType.Error() to return non-nil")
				}
				// Use the Roles() helper method to get valid roles
				validRoles := roleType.Roles()
				if len(validRoles) != 0 {
					t.Errorf("expected 0 valid roles, got %d", len(validRoles))
				}
			},
		},
		{
			name:           "empty shared role type",
			input:          "[]",
			expectError:    true,
			expectRoleType: true,
			checkFn: func(t *testing.T, roleType *ast.RoleType, p *Parser) {
				if roleType == nil {
					t.Error("expected roleType to be non-nil for error recovery")
					return
				}
				if roleType.Start.Type != token.LSQUARE {
					t.Errorf("expected Start.Type to be LSQUARE, got %v", roleType.Start.Type)
				}
				// For empty brackets, End should be empty token since we return early on error
				if roleType.End.Type != token.ILLEGAL && roleType.End.Type != "" {
					t.Errorf("expected End.Type to be ILLEGAL or empty, got %v", roleType.End.Type)
				}
				// Should have one error role (empty)
				if len(roleType.RoleNodes) != 1 {
					t.Errorf("expected 1 role node, got %d", len(roleType.RoleNodes))
				}
				if roleType.RoleNodes[0].Token.Type != token.ILLEGAL {
					t.Errorf("expected role to have error token, got %v", roleType.RoleNodes[0].Token.Type)
				}
				// Check that parser has errors
				if len(p.ParserErrors()) == 0 {
					t.Error("expected parser to have errors")
				}
				// Use the Error() helper method
				if roleType.Error() == nil {
					t.Error("expected roleType.Error() to return non-nil")
				}
				// Use the Roles() helper method to get valid roles
				validRoles := roleType.Roles()
				if len(validRoles) != 0 {
					t.Errorf("expected 0 valid roles, got %d", len(validRoles))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := FromString(tt.input)
			roleType, needsRecover := p.parseRoleType()

			if needsRecover != tt.expectError {
				t.Errorf("parseRoleType() needsRecover = %v, expectError %v", needsRecover, tt.expectError)
				return
			}

			if tt.expectRoleType && roleType == nil {
				t.Error("expected roleType to be non-nil for error recovery")
				return
			}

			if tt.checkFn != nil {
				tt.checkFn(t, roleType, p)
			}
		})
	}
}
