package new_parser_test

import (
	"testing"

	"github.com/tempo-lang/tempo/parser/new_parser"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "empty block",
			input: "{}",
		},
		{
			name:  "let decl block",
			input: "{let x = x;}",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := new_parser.FromString(test.input)
			p.ParseScope()

			if len(p.ParserErrors()) > 0 {
				t.Fatalf("Got %d parser errors: %v", len(p.ParserErrors()), p.ParserErrors())
			}
		})
	}

}
