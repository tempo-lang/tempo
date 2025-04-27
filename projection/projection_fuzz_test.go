package projection_test

import (
	"chorego/compiler"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/antlr4-go/antlr/v4"

	"go/ast"
	goparser "go/parser"
	"go/token"
	"go/types"
)

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
