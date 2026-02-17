package compiler_test

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tempo-lang/tempo/compiler"
	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/type_check/type_error"
)

// FuzzExtendedCompile runs the compiler with a large corpus to look for crashes.
func FuzzExtendedCompile(f *testing.F) {
	samples := loadSamples(f)

	f.Add("") // empty input
	for _, sample := range samples {
		for i := range len(sample) - 1 {
			prefix := sample[0 : i+1]
			f.Add(prefix)
		}
	}

	opts := []*compiler.Options{
		misc.ToPtr(compiler.DefaultOptions()),
		{Language: compiler.LangJS},
		{Language: compiler.LangTS},
		{Language: compiler.LangJava},
	}

	f.Fuzz(func(t *testing.T, source string) {
		for _, opt := range opts {
			input_stream := antlr.NewInputStream(source)
			_, type_errors := compiler.Compile(input_stream, opt)
			for _, err := range type_errors {
				if type_err, ok := err.(type_error.Error); ok {
					type_error.FormatError(io.Discard, input_stream, "test_sample", type_err, false)
					type_error.FormatError(io.Discard, input_stream, "test_sample", type_err, true)
				}
			}
		}
	})
}

func loadSamples(t testing.TB) []string {
	results := []string{}

	// Add valid examples
	paths, err := filepath.Glob(filepath.Join("..", "projection", "testdata", "examples", "*.tempo"))
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range paths {
		dat, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		results = append(results, (string(dat)))
	}

	// Add type checker examples
	paths, err = filepath.Glob(filepath.Join("..", "type_check", "testdata", "examples", "*.txt"))
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range paths {
		dat, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		source := strings.SplitN(string(dat), "---", 2)[0]
		results = append(results, source)
	}

	return results
}
