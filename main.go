package main

import (
	"fmt"
	"os"
	"path"

	"github.com/tempo-lang/tempo/compiler"
	"github.com/tempo-lang/tempo/lsp"
	"github.com/tempo-lang/tempo/type_check/type_error"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	if os.Args[1] == "lsp" {
		lsp.StartServer()
		return
	}

	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Printf("failed to get file stream: %v\n", err)
		os.Exit(1)
	}

	filename := path.Base(os.Args[1])
	options := compiler.Options{
		PackageName: filename[0 : len(filename)-len(path.Ext(filename))],
	}

	output, errors := compiler.Compile(input, &options)
	if errors != nil {
		for _, err := range errors {
			if typeErr, ok := err.(type_error.Error); ok {
				token := typeErr.ParserRule().GetStart()
				fmt.Printf("Type error %d:%d: %s\n", token.GetLine(), token.GetColumn()+1, typeErr.Error())
			} else {
				fmt.Printf("%v\n", err)
			}
		}
		os.Exit(1)
	}

	fmt.Printf("%s", output)
}
