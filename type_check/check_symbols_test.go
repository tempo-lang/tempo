package type_check_test

import "testing"

func TestSymbols(t *testing.T) {

	tests := [...]AnalyzerTestData{
		{
			name: "duplicate function names",
			input: `
			  func@(A,B) foo(value: Int@A) {}
				func@(C,D) foo(value: Int@C) {}
			`,
			errors: []string{
				"symbol 'foo' already declared",
			},
		},
	}

	for _, test := range tests {
		test.Assert(t)
	}
}
