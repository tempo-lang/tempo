package lexer_test

import (
	"testing"

	"github.com/tempo-lang/tempo/parser/new_parser/lexer"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

func TestReadStringLiteral(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedType  token.TokenType
		expectedValue any
	}{
		{
			name:          "simple string",
			input:         `"hello"`,
			expectedType:  token.STRING,
			expectedValue: "hello",
		},
		{
			name:          "empty string",
			input:         `""`,
			expectedType:  token.STRING,
			expectedValue: "",
		},
		{
			name:          "escaped quote",
			input:         `"say \"hi\""`,
			expectedType:  token.STRING,
			expectedValue: `say \"hi\"`,
		},
		{
			name:          "escaped backslash",
			input:         `"back\\slash"`,
			expectedType:  token.STRING,
			expectedValue: `back\\slash`,
		},
		{
			name:          "escaped newline",
			input:         `"line1\nline2"`,
			expectedType:  token.STRING,
			expectedValue: "line1\nline2",
		},
		{
			name:          "escaped tab",
			input:         `"col1\tcol2"`,
			expectedType:  token.STRING,
			expectedValue: "col1\tcol2",
		},
		{
			name: "escaped newline in source",
			input: `"line1\
line2"`,
			expectedType:  token.STRING,
			expectedValue: "line1line2",
		},
		{
			name:          "mixed escapes",
			input:         `"a\nb\tc\\d\"e"`,
			expectedType:  token.STRING,
			expectedValue: "a\nb\tc\\\\d\\\"e",
		},
		{
			name:          "unclosed string",
			input:         `"unclosed`,
			expectedType:  token.ILLEGAL,
			expectedValue: token.TokenError{Message: "unexpected end of file, expected \""},
		},
		{
			name:          "unclosed string with escapes",
			input:         `"has \n escape`,
			expectedType:  token.ILLEGAL,
			expectedValue: token.TokenError{Message: "unexpected end of file, expected \""},
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
