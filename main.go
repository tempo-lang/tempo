package main

import (
	"chorego/compiler"
	"chorego/types"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Printf("failed to get file stream: %v\n", err)
		os.Exit(1)
	}

	output, errors := compiler.Compile(input)
	if errors != nil {
		for _, err := range errors {
			if typeErr, ok := err.(types.Error); ok {
				token := typeErr.ParserRule().GetStart()
				fmt.Printf("Type error %d:%d: %s\n", token.GetLine(), token.GetColumn(), typeErr.Error())
			} else {
				fmt.Printf("%v\n", err)
			}
		}
		os.Exit(1)
	}

	fmt.Printf("%s", output)
}
