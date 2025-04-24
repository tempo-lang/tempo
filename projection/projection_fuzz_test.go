package projection_test

import (
	"chorego/epp"
	"chorego/misc"
	"chorego/parser"
	"chorego/type_check"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/dave/jennifer/jen"

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
		terminateErrorListener := misc.TerminateErrorListener{}
		input := antlr.NewInputStream(source)
		lexer := parser.NewChoregoLexer(input)
		lexer.AddErrorListener(&terminateErrorListener)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewChoregoParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.AddErrorListener(&terminateErrorListener)

		// parse program
		function := p.Func_()

		if terminateErrorListener.ProducedError {
			return
		}

		a := type_check.New()
		antlr.ParseTreeWalkerDefault.Walk(a, function)

		analyzerErrorListener := a.ErrorListener.(*type_check.DefaultErrorListener)

		if analyzerErrorListener.ProducedError {
			return
		}

		choreography := epp.EppFunc(function)
		file := jen.NewFile("choreography")
		choreography.Codegen(file)

		output := file.GoString()

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
