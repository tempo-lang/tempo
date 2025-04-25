package projection_test

import (
	"chorego/chorego"
	"testing"

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

	f.Fuzz(func(t *testing.T, source string) {
		input := antlr.NewInputStream(source)
		output, errors := chorego.Compile(input)
		if len(errors) > 0 {
			return
		}

		fset := token.NewFileSet()
		parsedAST, err := goparser.ParseFile(fset, "", output, goparser.AllErrors)
		if err != nil {
			t.Errorf("Invalid generated go code: %v", err)
		}

		conf := types.Config{}
		_, err = conf.Check("choreography", fset, []*ast.File{parsedAST}, nil)
		if err != nil {
			t.Errorf("Go code type error: %v\n\nINPUT:\n%s\n\nPROJECTION:\n%s", err, source, output)
		}
	})
}
