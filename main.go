package main

import (
	"fmt"
	"os"

	"chorego/analyzer"
	"chorego/epp"
	"chorego/misc"
	"chorego/parser"

	"github.com/antlr4-go/antlr/v4"
	"github.com/dave/jennifer/jen"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewChoregoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewChoregoParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	terminateErrorListener := misc.TerminateErrorListener{}
	p.AddErrorListener(&terminateErrorListener)

	// parse program
	function := p.Func_()

	if terminateErrorListener.ProducedError {
		os.Exit(1)
	}

	a := analyzer.New()
	antlr.ParseTreeWalkerDefault.Walk(a, function)

	analyzerErrorListener, ok := a.ErrorListener.(*analyzer.DefaultErrorListener)
	if !ok {
		fmt.Println("Analyzer error listener was expected to be DefaultErrorListener")
		os.Exit(1)
	}

	if analyzerErrorListener.ProducedError {
		os.Exit(1)
	}

	network := epp.EppFunc(function)

	file := jen.NewFile("choreography")
	network.Codegen(file)

	err := file.Render(os.Stdout)
	if err != nil {
		fmt.Printf("Failed to render file: %v\n", err)
		os.Exit(1)
	}
}
