package type_check_test

import "testing"

func TestFuncParams(t *testing.T) {

	tests := [...]AnalyzerTestData{
		{
			name:  "unknown role",
			input: "func@(A,B) foo(value: Int@C) {}",
			errors: []string{
				"invalid function type for 'foo': param 1: [unknown roles [C] in 'Int@C']",
			},
		},
		{
			name:  "duplicate param names",
			input: "func@(A,B) foo(value: Int@A, value: Int@B) {}",
			errors: []string{
				"symbol 'value' already declared",
			},
		},
	}

	for _, test := range tests {
		test.Assert(t)
	}
}
