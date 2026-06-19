package new_parser_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tempo-lang/tempo/parser/new_parser"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

func TestParserErrors(t *testing.T) {
	// Helper to create an expected error token
	makeError := func(literal string, line, col int, message string) token.Token {
		return token.Error(literal, token.SourcePos{Line: line, Col: col}, message)
	}

	tests := []struct {
		name           string
		input          string
		expectedErrors []token.Token
	}{
		{
			name:  "missing semicolon",
			input: "{let x = 1}",
			expectedErrors: []token.Token{
				makeError("1", 1, 10, "missing semicolon"),
			},
		},
		{
			name:  "missing equals in let",
			input: "{let x 1;}",
			expectedErrors: []token.Token{
				makeError("1", 1, 8, "expected `=` when parsing let statement"),
			},
		},
		{
			name:  "number as statement",
			input: "{123;}",
			expectedErrors: []token.Token{
				makeError("123", 1, 2, "invalid statement"),
			},
		},
		{
			name:  "invalid expression in let",
			input: "{let x = +;}",
			expectedErrors: []token.Token{
				makeError("+", 1, 10, "not an expression"),
			},
		},
		{
			name:  "unclosed scope",
			input: "{let x = 1;",
			expectedErrors: []token.Token{
				makeError("", 1, 11, "unexpected EOF when parsing scope"),
			},
		},
		{
			name:  "missing identifier in let",
			input: "{let = 1;}",
			expectedErrors: []token.Token{
				makeError("=", 1, 6, "expected identifier when parsing let statement"),
			},
		},
		{
			name:  "empty let name",
			input: "{let x = ; let y = 10;}",
			expectedErrors: []token.Token{
				makeError(";", 1, 10, "not an expression"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := new_parser.FromString(test.input)
			p.ParseScope()

			actualErrors := p.ParserErrors()
			if diff := cmp.Diff(test.expectedErrors, actualErrors); diff != "" {
				t.Errorf("ParserErrors mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}
