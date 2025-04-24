package main

import (
	"fmt"
	"os"

	"chorego/epp"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Printf("failed to get file stream: %v\n", err)
		os.Exit(1)
	}

	output, errors := epp.EndpointProject(input)
	if errors != nil {
		for _, err := range errors {
			fmt.Printf("%v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("%s", output)
}
