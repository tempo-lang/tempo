package compiler

import (
	"chorego/epp"
	"chorego/parser"
	"chorego/type_check"

	"github.com/antlr4-go/antlr/v4"
	"github.com/dave/jennifer/jen"
)

func Compile(input antlr.CharStream) (output string, errors []error) {
	// parse source input
	sourceFile, syntaxErrors := parser.Parse(input)
	if len(syntaxErrors) > 0 {
		errors = syntaxErrors
		return
	}

	// type check ast
	typeErrors := type_check.TypeCheck(sourceFile)
	if len(typeErrors) > 0 {
		for _, err := range typeErrors {
			errors = append(errors, err)
		}
		return
	}

	// endpoint project
	eppFile := epp.EppSourceFile(sourceFile)

	// generate go code
	file := jen.NewFile("choreography")
	eppFile.Codegen(file)

	output = file.GoString()
	return
}
