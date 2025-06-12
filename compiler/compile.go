// This package bundles together all the different stages of compilation into the single convenient [Compile] function.
package compiler

import (
	"github.com/tempo-lang/tempo/epp"
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/type_check"

	"github.com/antlr4-go/antlr/v4"
	"github.com/dave/jennifer/jen"
)

// Options given to the compiler
type Options struct {
	PackageName string
}

// Compile takes Tempo source code and outputs its projection, or all syntax and type errors.
func Compile(input antlr.CharStream, options *Options) (output string, errors []error) {
	// parse source input
	sourceFile, syntaxErrors := parser.Parse(input)
	if len(syntaxErrors) > 0 {
		for _, err := range syntaxErrors {
			errors = append(errors, err)
		}
	}

	// type check ast
	info, typeErrors := type_check.TypeCheck(sourceFile)
	if len(typeErrors) > 0 {
		for _, err := range typeErrors {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return
	}

	// endpoint project
	eppFile := epp.EndpointProject(info, sourceFile)

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
