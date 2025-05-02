package type_check_test

import "testing"

func TestExpressions(t *testing.T) {

	tests := [...]AnalyzerTestData{
		{
			name: "constant number expression overflow",
			input: `
			  func@(A) overflow() {
				  let x: Int@A = 99999999999999999999;
				}
			`,
			errors: []string{
				"invalid number '99999999999999999999'",
			},
		},
	}

	for _, test := range tests {
		test.Assert(t)
	}
}
