package lexer_test

import (
	"testing"

	"github.com/tempo-lang/tempo/parser/new_parser/lexer"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){}"

	tests :=
		[]struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LCURLY, "{"},
			{token.RCURLY, "}"},
			{token.EOF, ""},
		}

	l := lexer.FromString(input)

	for _, tt := range tests {
		tok := l.ReadToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("wrong token type %s expected %s", tok.Type, tt.expectedType)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("wrong literal %s expected %s", tok.Literal, tt.expectedLiteral)
		}
	}
}
