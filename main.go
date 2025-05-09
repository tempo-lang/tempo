package main

import (
	"chorego/compiler"
	"chorego/types"
	"fmt"
	"os"
	"path"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
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
			if typeErr, ok := err.(types.Error); ok {
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
