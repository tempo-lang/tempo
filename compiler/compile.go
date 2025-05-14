package compiler

import (
	"tempo/epp"
	"tempo/parser"
	"tempo/type_check"

	"github.com/antlr4-go/antlr/v4"
	"github.com/dave/jennifer/jen"
)

type Options struct {
	PackageName string
}

func Compile(input antlr.CharStream, options *Options) (output string, errors []error) {
	// parse source input
	sourceFile, syntaxErrors := parser.Parse(input)
	if len(syntaxErrors) > 0 {
		for _, err := range syntaxErrors {
			errors = append(errors, err)
		}
		return
	}

	// type check ast
	info, typeErrors := type_check.TypeCheck(sourceFile)
	if len(typeErrors) > 0 {
		for _, err := range typeErrors {
			errors = append(errors, err)
		}
		return
	}

	// endpoint project
	eppFile := epp.EppSourceFile(info, sourceFile)

	// generate go code
	packageName := "choreography"
	if options != nil && options.PackageName != "" {
		packageName = options.PackageName
	}

	file := jen.NewFile(packageName)
	eppFile.Codegen(file)

	output = file.GoString()
	return
}
