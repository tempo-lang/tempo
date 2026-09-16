package new_parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	antlrparser "github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
)

func antlrAccepts(source string) bool {
	_, diagnostics := antlrparser.Parse(antlr.NewInputStream(source))
	return len(diagnostics) == 0
}

func handwrittenAccepts(source string) bool { return len(Parse(source).Diagnostics) == 0 }

func TestRepositoryCorpusCompatibility(t *testing.T) {
	patterns := []string{"../../examples/*/*.tempo", "../../projection/testdata/examples/*.tempo"}
	for _, pattern := range patterns {
		paths, err := filepath.Glob(pattern)
		if err != nil {
			t.Fatal(err)
		}
		for _, path := range paths {
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if a, h := antlrAccepts(string(data)), handwrittenAccepts(string(data)); a != h {
				t.Errorf("acceptance disagreement for %s: ANTLR=%v handwritten=%v\n%s", path, a, h, FormatDiagnostics(Parse(string(data)).Diagnostics))
			}
			r := Parse(string(data))
			if len(r.Diagnostics) == 0 {
				if err := ast.Validate(r.Root, r.Tokens, r.Source.Len()); err != nil {
					t.Errorf("invalid AST for %s: %v", path, err)
				}
			}
		}
	}
}

func TestGrammarAlternativeCompatibility(t *testing.T) {
	valid := []string{
		"", "func f() {}", "func@[A,B] f(x: async [T@A]) func@A(T) U {}",
		"struct@[A,B] S implements I@A, J@B { x: T@A; func m() { return; } }",
		"interface@A I { func f(x: T@A) U@A; }",
		"func f() { let x: [T@A] = S@A{x: 1@A}; let y = func@A(x: T@A) T@A { return x; }; A -> B y(1).x[0]; }",
	}
	invalid := []string{"let x = 1;", "func f( {}", "struct S implements I, {}", "struct S { x T; }", "interface I { func f() }", "func f() { return }", "func f() { let x = [1,]; }", "func f() { f(1,); }", "func f() { x@A = 1; }", "func f() {} junk"}
	for _, source := range append(valid, invalid...) {
		if a, h := antlrAccepts(source), handwrittenAccepts(source); a != h {
			t.Errorf("%q: ANTLR=%v handwritten=%v; %s", source, a, h, FormatDiagnostics(Parse(source).Diagnostics))
		}
	}
}

func FuzzAcceptanceParity(f *testing.F) {
	for _, seed := range []string{"", "func f() {}", "struct S { x: T; }", "interface I { func f(); }"} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, source string) {
		if a, h := antlrAccepts(source), handwrittenAccepts(source); a != h {
			t.Fatalf("ANTLR=%v handwritten=%v", a, h)
		}
	})
}
