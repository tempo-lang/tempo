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
				Roles: []*ast.Role{{
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
				Roles: []*ast.Role{{
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
				Roles: []*ast.Role{
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
				Roles: []*ast.Role{
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
				Roles: []*ast.Role{
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

			// Check number of roles
			if len(roleType.Roles) != len(tt.expected.Roles) {
				t.Errorf("parseRoleType() Roles length = %d, expected %d", len(roleType.Roles), len(tt.expected.Roles))
				return
			}

			// Check IsShared() method
			if roleType.IsShared() != (roleType.Start.Type == token.LSQUARE) {
				t.Errorf("parseRoleType() IsShared() method not working correctly")
				return
			}

			// For single role case, check that start and end are the same
			if len(tt.expected.Roles) == 1 && tt.expected.Start.Type != token.LSQUARE && tt.expected.Start.Type != token.LPAREN {
				if roleType.Start.Type != roleType.End.Type {
					t.Errorf("parseRoleType() single role: Start.Type = %v, End.Type = %v, expected same", roleType.Start.Type, roleType.End.Type)
				}
			}
		})
	}
}
