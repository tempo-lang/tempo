package projection_test

import (
	"chorego/compiler"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/antlr4-go/antlr/v4"
)

func TestExamples(t *testing.T) {

	paths, err := filepath.Glob(filepath.Join("testdata", "examples", "*.chorego"))
	if err != nil {
		t.Fatal(err)
	}

	for _, path := range paths {
		_, filename := filepath.Split(path)
		testname := filename[:len(filename)-len(filepath.Ext(path))]

		t.Run(testname, func(t *testing.T) {
			input, err := antlr.NewFileStream(path)
			if err != nil {
				t.Fatal("error reading source file:", err)
			}

			output, compilerErrors := compiler.Compile(input)
			if len(compilerErrors) > 0 {
				errorsFormatted := ""
				for _, err := range compilerErrors {
					errorsFormatted = fmt.Sprintf("%s\n%s", errorsFormatted, err)
				}
				t.Fatalf("source produced errors: %s", errorsFormatted)
			}

			outputPath := filepath.Join("testdata", "examples", testname+".out")

			// If output path does not exist already
			if _, err := os.Stat(outputPath); errors.Is(err, os.ErrNotExist) {
				newPath := filepath.Join("testdata", "examples", testname+"_new.out")
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
				t.Fatalf("projected code differes from output file:\n%s\n", diff.LineDiff(trimmedOutput, trimmedOutputFile))
			}

		})
	}

}
