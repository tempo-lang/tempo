package main

import (
	"fmt"
	"os"

	"chorego/analyzer"
	"chorego/epp"
	"chorego/parser"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewChoregoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewChoregoParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	function := p.Func_()

	antlr.ParseTreeWalkerDefault.Walk(analyzer.New(), function)

	file := epp.EppFunc(function)
	fmt.Printf("\n\n%#v", file)

}
