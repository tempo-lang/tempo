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
		{
			name: "duplicate variable names",
			input: `
			  func@A foo() {
				  let x: Int@A = 1;
					let x: Int@A = 2;
				}
			`,
			errors: []string{
				"symbol 'x' already declared",
			},
		},
		{
			name: "duplicate variable and param names",
			input: `
			  func@A foo(p: Int@A) {
				  let p: Int@A = 1;
				}
			`,
			errors: []string{
				"symbol 'p' already declared",
			},
		},
	}

	for _, test := range tests {
		test.Assert(t)
	}
}
