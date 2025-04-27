package type_check_test

import (
	"testing"
)

func TestFuncDuplicateRoles(t *testing.T) {

	tests := [...]AnalyzerTestData{
		{
			name:   "no duplicate roles",
			input:  "func@(A,B,C) foo() {}",
			errors: []string{},
		},
		{
			name:  "single duplicate role",
			input: "func@(A,A) foo() {}",
			errors: []string{
				"function 'foo' has duplicate role 'A'",
			},
		},
		{
			name:  "multiple duplicate roles",
			input: "func@(A,B,A,B) foo() {}",
			errors: []string{
				"function 'foo' has duplicate role 'A'",
				"function 'foo' has duplicate role 'B'",
			},
		},
		{
			name:  "duplicate role one between",
			input: "func@(B,A,B) foo(){}",
			errors: []string{
				"function 'foo' has duplicate role 'B'",
			},
		},
	}

	for _, test := range tests {
		test.Assert(t)
	}
}
