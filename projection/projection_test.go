package projection_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tempo-lang/tempo/compiler"

	"github.com/andreyvit/diff"
	"github.com/antlr4-go/antlr/v4"
)

// TestExamples finds all examples located in the testdata directory/examples,
// compiles them and verifies that they match the expected outputs.
func TestExamples(t *testing.T) {

	paths, err := filepath.Glob(filepath.Join("testdata", "examples", "*.tempo"))
	if err != nil {
		t.Fatal(err)
	}

	for _, path := range paths {
		_, filename := filepath.Split(path)
		testname := filename[:len(filename)-len(filepath.Ext(path))]

		t.Run(testname+".go", func(t *testing.T) {
			t.Parallel()

			input, err := antlr.NewFileStream(path)
			if err != nil {
				t.Fatal("error reading source file:", err)
			}

			output, compilerErrors := compiler.Compile(input, nil)
			if len(compilerErrors) > 0 {
				errorsFormatted := ""
				for _, err := range compilerErrors {
					errorsFormatted = fmt.Sprintf("%s\n%s", errorsFormatted, err)
				}
				t.Fatalf("source produced errors: %s", errorsFormatted)
			}

			// Generate output if not already exists
			outputPath := filepath.Join("testdata", "examples", testname)
			compareOrWriteOutput(t, outputPath, "go", output)
		})

		t.Run(testname+".ts", func(t *testing.T) {
			t.Parallel()

			input, err := antlr.NewFileStream(path)
			if err != nil {
				t.Fatal("error reading source file:", err)
			}

			options := compiler.DefaultOptions()
			options.Language = compiler.LangTS
			options.RuntimePath = "../../../typescript/runtime.ts"

			output, compilerErrors := compiler.Compile(input, &options)
			if len(compilerErrors) > 0 {
				errorsFormatted := ""
				for _, err := range compilerErrors {
					errorsFormatted = fmt.Sprintf("%s\n%s", errorsFormatted, err)
				}
				t.Fatalf("source produced errors: %s", errorsFormatted)
			}

			// Generate output if not already exists
			outputPath := filepath.Join("testdata", "examples", testname)
			compareOrWriteOutput(t, outputPath, "ts", output)
		})
	}

}

func compareOrWriteOutput(t *testing.T, filename string, extension string, output string) {
	// Generate output if not already exists
	outputPath := filename + "." + extension
	if _, err := os.Stat(outputPath); errors.Is(err, os.ErrNotExist) {
		newPath := filename + "_new." + extension
		if err := os.WriteFile(newPath, []byte(output), 0655); err != nil {
			t.Fatalf("error writing output: %v", err)
		}
		t.Fatal("no output file, generated one")
	}

	outputFile, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("error reading output file: %v", outputFile)
	}

	whitespace := " \r\n\t"
	trimmedOutput := strings.Trim(output, whitespace)
	trimmedOutputFile := strings.Trim(string(outputFile), whitespace)

	if trimmedOutput != trimmedOutputFile {
		t.Fatalf("projected code differes from output file %s.%s:\n%s\n", filename, extension, diff.LineDiff(trimmedOutputFile, trimmedOutput))
	}
}
