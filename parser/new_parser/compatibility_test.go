package new_parser

import (
	"os"
	"path/filepath"
	"testing"
)

// These fixtures were classified as accepted by both parsers in the final
// pre-cutover comparison. They are now the checked-in compatibility oracle.
func TestCompatibilityCorpus(t *testing.T) {
	paths, err := filepath.Glob(filepath.Join("..", "..", "projection", "testdata", "examples", "*.tempo"))
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) == 0 {
		t.Fatal("compatibility corpus is empty")
	}
	for _, path := range paths {
		path := path
		t.Run(filepath.Base(path), func(t *testing.T) {
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			result := Parse(string(data))
			if len(result.Diagnostics) > 0 {
				t.Fatalf("expected accepted fixture:\n%s", FormatDiagnostics(result.Diagnostics))
			}
		})
	}
}

func TestCompatibilityRejectedCorpus(t *testing.T) {
	fixtures := []string{"let x = 1;", "func@(A f() {}", "struct T { x: Int }", "func f(){ let x = ; }"}
	for _, source := range fixtures {
		if len(Parse(source).Diagnostics) == 0 {
			t.Errorf("expected rejected fixture: %q", source)
		}
	}
}
