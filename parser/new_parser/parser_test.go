package new_parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tempo-lang/tempo/parser/new_parser/ast"
)

const exampleSeparator = "---"

// TestExamples verifies the S-expression for every parser example. A new
// example may initially contain only Tempo source; its expectation is then
// generated in place and the test fails so that the output must be inspected
// before the next test run.
func TestExamples(t *testing.T) {
	paths, err := filepath.Glob(filepath.Join("testdata", "examples", "*.txt"))
	if err != nil {
		t.Fatal(err)
	}

	for _, path := range paths {
		path := path
		name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}

			source, expected, hasExpectation := splitExample(string(data))
			result := Parse(source)
			actual := strings.TrimSpace((ast.SExprGenerator{Indent: "  "}).Generate(result.Root, result.Tokens))

			if !hasExpectation {
				if err := writeExampleExpectation(path, source, actual); err != nil {
					t.Fatal(err)
				}
				t.Fatalf("missing S-expression expectation; generated one in %s — inspect it and rerun the test", path)
			}

			expected = strings.TrimSpace(expected)
			if expected != actual {
				t.Errorf("S-expression did not match expected output.\nExpected:\n%s\nActual:\n%s", expected, actual)
			}
		})
	}
}

// splitExample recognizes only a separator occupying an entire line. This
// allows three dashes to occur in Tempo source without accidentally splitting
// the fixture.
func splitExample(data string) (source, expected string, ok bool) {
	for lineStart := 0; lineStart <= len(data); {
		relativeLineEnd := strings.IndexByte(data[lineStart:], '\n')
		lineEnd := len(data)
		nextLine := len(data) + 1
		if relativeLineEnd >= 0 {
			lineEnd = lineStart + relativeLineEnd
			nextLine = lineEnd + 1
		}

		line := strings.TrimSuffix(data[lineStart:lineEnd], "\r")
		if line == exampleSeparator {
			if nextLine > len(data) {
				nextLine = len(data)
			}
			return data[:lineStart], data[nextLine:], true
		}
		if nextLine > len(data) {
			break
		}
		lineStart = nextLine
	}
	return data, "", false
}

func writeExampleExpectation(path, source, actual string) error {
	var output strings.Builder
	output.WriteString(source)
	if source != "" && !strings.HasSuffix(source, "\n") {
		output.WriteByte('\n')
	}
	output.WriteString(exampleSeparator)
	output.WriteByte('\n')
	output.WriteString(actual)
	output.WriteByte('\n')
	return os.WriteFile(path, []byte(output.String()), 0o644)
}

func TestSplitExample(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		source   string
		expected string
		ok       bool
	}{
		{name: "separator", data: "let x = 1;\n---\n(let x)", source: "let x = 1;\n", expected: "(let x)", ok: true},
		{name: "CRLF separator", data: "let x = 1;\r\n---\r\n(let x)\r\n", source: "let x = 1;\r\n", expected: "(let x)\r\n", ok: true},
		{name: "source only", data: "let x = \"---\";", source: "let x = \"---\";"},
		{name: "dashes within line", data: "// --- not a separator\n", source: "// --- not a separator\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source, expected, ok := splitExample(test.data)
			if source != test.source || expected != test.expected || ok != test.ok {
				t.Fatalf("splitExample() = (%q, %q, %v), want (%q, %q, %v)", source, expected, ok, test.source, test.expected, test.ok)
			}
		})
	}
}

func TestWriteExampleExpectation(t *testing.T) {
	path := filepath.Join(t.TempDir(), "example.txt")
	if err := writeExampleExpectation(path, "let x = 1;", "(source-file)"); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	want := "let x = 1;\n---\n(source-file)\n"
	if string(data) != want {
		t.Fatalf("generated fixture = %q, want %q", data, want)
	}
}
