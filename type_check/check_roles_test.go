package type_check_test

import "testing"

func TestValueRoles(t *testing.T) {

	tests := [...]AnalyzerTestData{
		{
			name: "unknown role in var decl",
			input: `
			  func@(A) foo() {
				  let x: Int@B = 32;
				}
			`,
			errors: []string{
				"unknown role 'B'",
			},
		},
		{
			name: "wrong role in assignment",
			input: `
			  func@(A,B) foo() {
				  let x: Int@A = 32;
					let y: Int@B = x;
				}
			`,
			errors: []string{
				"type mismatch, expected Int@(B) found Int@(A)",
			},
		},
	}

	for _, test := range tests {
		test.Assert(t)
	}
}
