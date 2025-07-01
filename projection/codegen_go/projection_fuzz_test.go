package codegen_go_test

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/tempo-lang/tempo/compiler"
	"github.com/tempo-lang/tempo/projection/codegen_go"

	"github.com/antlr4-go/antlr/v4"

	"go/ast"
	"go/importer"
	goparser "go/parser"
	"go/token"
	"go/types"
)

type runtimeImporter struct {
	pkg             *types.Package
	defaultImporter types.Importer
}

// Import implements types.Importer.
func (r *runtimeImporter) Import(path string) (*types.Package, error) {
	if path == codegen_go.RUNTIME_PATH {
		return r.pkg, nil
	}
	// return nil, fmt.Errorf("failed to import %s", path)
	return r.defaultImporter.Import(path)
}

func newRuntimeImporter() (*runtimeImporter, error) {

	runtimeSource, err := os.ReadFile("../../runtime/runtime.go")
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()
	parsedAST, err := goparser.ParseFile(fset, "", runtimeSource, goparser.AllErrors)
	if err != nil {
		return nil, err
	}

	conf := types.Config{}
	conf.Importer = importer.ForCompiler(fset, "gc", nil)

	pkg, err := conf.Check(codegen_go.RUNTIME_PATH, fset, []*ast.File{parsedAST}, nil)
	if err != nil {
		return nil, err
	}

	return &runtimeImporter{
		pkg: pkg,
	}, nil
}

func FuzzProjection(f *testing.F) {
	f.Add("func@(X,Y,Z) foo() {}")
	f.Add("func@(A,B) foo(snd: Int@(A,B)) {}")
	f.Add("func@(A,B,C) foo(fst: Float@A, snd: String@(A,B)) {}")
	f.Add(`
		func@(A,B,C) foo(fst: Int@A, snd: Int@(A,B)) {
			let x: Int@(B,C) = 42;
			let y: Int@(A,B) = 12;
		}
	`)

	// Add valid examples
	paths, err := filepath.Glob(filepath.Join("..", "testdata", "examples", "*.tempo"))
	if err != nil {
		f.Fatal(err)
	}
	for _, path := range paths {
		dat, err := os.ReadFile(path)
		if err != nil {
			f.Fatal(err)
		}
		f.Add(string(dat))
	}

	// Add type checker examples
	paths, err = filepath.Glob(filepath.Join("..", "..", "type_check", "testdata", "examples", "*.txt"))
	if err != nil {
		f.Fatal(err)
	}
	for _, path := range paths {
		dat, err := os.ReadFile(path)
		if err != nil {
			f.Fatal(err)
		}
		source := strings.SplitN(string(dat), "---", 2)[0]
		f.Add(source)
	}

	importer, err := newRuntimeImporter()
	if err != nil {
		f.Fatalf("Failed to make runtime importer: %v", err)
	}

	f.Fuzz(func(t *testing.T, source string) {
		input := antlr.NewInputStream(source)
		output, errors := compiler.Compile(input, nil)
		if len(errors) > 0 {
			return
		}

		fset := token.NewFileSet()
		parsedAST, err := goparser.ParseFile(fset, "", output, goparser.AllErrors)
		if err != nil {
			t.Fatalf("Invalid generated go code: %v", err)
		}

		conf := types.Config{}
		conf.Importer = importer

		// allow these Go errors for now
		ignoredGoErrors := []string{
			"invalid operation: division by zero",
			"(overflows)",
		}

		_, err = conf.Check("choreography", fset, []*ast.File{parsedAST}, nil)
		if err != nil {
			isIgnoredError := slices.ContainsFunc(ignoredGoErrors, func(e string) bool {
				return strings.Contains(err.Error(), e)
			})

			if !isIgnoredError {
				t.Fatalf("Go code type error: %v\n\nINPUT:\n%s\n\nPROJECTION:\n%s", err, source, output)
			}
		}
	})
}
