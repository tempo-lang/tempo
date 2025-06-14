// This package bundles together all the different stages of compilation into the single convenient [Compile] function.
package compiler

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/tempo-lang/tempo/epp"
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/projection/codegen_go"
	"github.com/tempo-lang/tempo/projection/codegen_ts"
	"github.com/tempo-lang/tempo/type_check"

	"github.com/antlr4-go/antlr/v4"
	"github.com/dave/jennifer/jen"
)

type CompilerLanguage string

const (
	LangGo = "go"
	LangTS = "ts"
)

// Options given to the compiler
type Options struct {
	PackageName string
	Language    CompilerLanguage
}

func DefaultOptions() Options {
	return Options{
		PackageName: "choreography",
		Language:    LangGo,
	}
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

	if options == nil {
		defaultOpts := DefaultOptions()
		options = &defaultOpts
	}

	if options.PackageName == "" {
		options.PackageName = "choreography"
	}

	if options.Language == "" {
		options.Language = LangGo
	}

	// generate go code
	switch options.Language {
	case LangGo:
		file := jen.NewFile(options.PackageName)
		codegen_go.GenSourceFile(file, eppFile)

		output = file.GoString()
		return
	case LangTS:
		b := strings.Builder{}
		buf := bufio.NewWriter(&b)
		codegen_ts.Codegen(buf, eppFile, nil)
		buf.Flush()
		output = b.String()
		return
	default:
		panic(fmt.Sprintf("unknown language: %v", options.Language))
	}
}
