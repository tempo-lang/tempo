package main

import (
	"chorego/chorego"
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

	output, errors := chorego.Compile(input)
	if errors != nil {
		for _, err := range errors {
			fmt.Printf("%v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("%s", output)
}
