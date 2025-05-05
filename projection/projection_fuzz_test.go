package projection_test

import (
	"chorego/compiler"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/antlr4-go/antlr/v4"

	"go/ast"
	goparser "go/parser"
	"go/token"
	"go/types"
)

type runtimeImporter struct {
	pkg *types.Package
}

// Import implements types.Importer.
func (r *runtimeImporter) Import(path string) (*types.Package, error) {
	if path == "chorego/runtime" {
		return r.pkg, nil
	}
	return nil, fmt.Errorf("failed to import %s", path)
}

func newRuntimeImporter() (*runtimeImporter, error) {

	runtimeSource, err := os.ReadFile("../runtime/runtime.go")
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()
	parsedAST, err := goparser.ParseFile(fset, "", runtimeSource, goparser.AllErrors)
	if err != nil {
		return nil, err
	}

	conf := types.Config{}
	pkg, err := conf.Check("chorego/runtime", fset, []*ast.File{parsedAST}, nil)
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
	paths, err := filepath.Glob(filepath.Join("testdata", "examples", "*.chorego"))
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
	paths, err = filepath.Glob(filepath.Join("..", "type_check", "testdata", "examples", "*.txt"))
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
		timeout := time.After(3 * time.Second)
		result := make(chan error)

		go func() {
			input := antlr.NewInputStream(source)
			output, errors := compiler.Compile(input)
			if len(errors) > 0 {
				result <- nil
				return
			}

			fset := token.NewFileSet()
			parsedAST, err := goparser.ParseFile(fset, "", output, goparser.AllErrors)
			if err != nil {
				result <- fmt.Errorf("Invalid generated go code: %v", err)
				return
			}

			conf := types.Config{}
			conf.Importer = importer

			_, err = conf.Check("choreography", fset, []*ast.File{parsedAST}, nil)
			if err != nil {
				result <- fmt.Errorf("Go code type error: %v\n\nINPUT:\n%s\n\nPROJECTION:\n%s", err, source, output)
				return
			}

			result <- nil
		}()

		select {
		case <-timeout:
			t.Fatal("Test didn't finish in time")
		case err := <-result:
			if err != nil {
				t.Fatal(err)
			}
		}

	})
}
