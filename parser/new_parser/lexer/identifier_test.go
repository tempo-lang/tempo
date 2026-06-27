package lexer_test

import (
	"testing"

	"github.com/tempo-lang/tempo/parser/new_parser/lexer"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

func TestReadIdentifier(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedType  token.TokenType
		expectedValue any
	}{
		{
			name:          "simple text",
			input:         "name",
			expectedType:  token.IDENT,
			expectedValue: "name",
		},
		{
			name:          "begin underscore",
			input:         "_test",
			expectedType:  token.IDENT,
			expectedValue: "_test",
		},
		{
			name:          "including numbers",
			input:         "hello123",
			expectedType:  token.IDENT,
			expectedValue: "hello123",
		},
		{
			name:          "underscore",
			input:         "_",
			expectedType:  token.UNDERSCORE,
			expectedValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.FromString(tt.input)
			tok := l.ReadToken()

			if tok.Type != tt.expectedType {
				t.Errorf("wrong token type: got %s, expected %s", tok.Type, tt.expectedType)
			}

			if tok.Value != tt.expectedValue {
				t.Errorf("wrong value: got %q, expected %q", tok.Value, tt.expectedValue)
			}

			if tok.Literal != tt.input {
				t.Errorf("wrong literal: got %q, expected %q", tok.Literal, tt.input)
			}
		})
	}
}
