package type_check_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"tempo/compiler"
	"tempo/misc"
	"tempo/types"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

func TestExamples(t *testing.T) {
	paths, err := filepath.Glob(filepath.Join("testdata", "examples", "*.txt"))
	if err != nil {
		t.Fatal(err)
	}

	for _, path := range paths {
		_, filename := filepath.Split(path)
		testname := filename[:len(filename)-len(filepath.Ext(path))]

		t.Run(testname, func(t *testing.T) {

			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}

			split := strings.SplitN(string(data), "---", 2)
			source := split[0]
			expectedErrors := split[1]

			input := antlr.NewInputStream(source)

			_, compilerErrors := compiler.Compile(input, nil)
			formattedErrors := []string{}
			for _, err := range compilerErrors {
				typeError, ok := err.(types.Error)
				if !ok {
					t.Error("expected only type errors")
					continue
				}

				line := typeError.ParserRule().GetStart().GetLine()
				col := typeError.ParserRule().GetStart().GetColumn() + 1
				formattedErrors = append(formattedErrors, fmt.Sprintf("%d:%d: %s", line, col, typeError.Error()))
			}

			actualErrors := misc.JoinStrings(formattedErrors, "\n")

			expectedErrors = strings.Trim(expectedErrors, " \n")
			actualErrors = strings.Trim(actualErrors, " \n")

			if expectedErrors != actualErrors {
				t.Errorf("Errors did not match expected in %s.\nExpected:\n%s\nActual:\n%s", testname, expectedErrors, actualErrors)
			}

		})
	}
}
