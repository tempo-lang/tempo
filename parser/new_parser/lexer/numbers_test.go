package lexer_test

import (
	"testing"

	"github.com/tempo-lang/tempo/parser/new_parser/lexer"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

func TestReadNumber(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedType  token.TokenType
		expectedValue any
	}{
		{
			name:          "simple int",
			input:         "123",
			expectedType:  token.INT,
			expectedValue: 123,
		},
		{
			name:          "simple float",
			input:         "123.4",
			expectedType:  token.FLOAT,
			expectedValue: 123.4,
		},
		{
			name:          "float starting with dot",
			input:         ".5",
			expectedType:  token.FLOAT,
			expectedValue: 0.5,
		},
		{
			name:          "float ending with dot",
			input:         "1.",
			expectedType:  token.FLOAT,
			expectedValue: 1.0,
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
				t.Errorf("wrong value: got %#v (%T), expected %#v (%T)", tok.Value, tok.Value, tt.expectedValue, tt.expectedValue)
			}

			if tok.Literal != tt.input {
				t.Errorf("wrong literal: got %q, expected %q", tok.Literal, tt.input)
			}
		})
	}
}
