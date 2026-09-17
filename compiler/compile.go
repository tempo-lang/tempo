// This package bundles together all the different stages of compilation into the single convenient [Compile] function.
package compiler

import (
	"fmt"

	"github.com/tempo-lang/tempo/epp"
	parser "github.com/tempo-lang/tempo/parser/new_parser"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
	"github.com/tempo-lang/tempo/projection/codegen_go"
	"github.com/tempo-lang/tempo/projection/codegen_java"
	"github.com/tempo-lang/tempo/projection/codegen_ts"
	"github.com/tempo-lang/tempo/type_check"

	"github.com/dave/jennifer/jen"
)

type CompilerLanguage string

const (
	LangGo   CompilerLanguage = "go"
	LangTS   CompilerLanguage = "ts"
	LangJS   CompilerLanguage = "js"
	LangJava CompilerLanguage = "java"
)

// Options given to the compiler
type Options struct {
	PackageName string
	Language    CompilerLanguage
	RuntimePath string
}

func DefaultOptions() Options {
	return Options{
		PackageName: "choreography",
		Language:    LangGo,
		RuntimePath: "@tempo-lang/tempo/runtime",
	}
}

// Compile takes Tempo source code and outputs its projection, or all syntax and type errors.
func Compile(source *token.Source, options *Options) (output string, errors []error) {
	// parse source input
	parsed := parser.ParseSource(source)
	sourceFile := parsed.Root
	if len(parsed.Diagnostics) > 0 {
		for _, err := range parsed.Diagnostics {
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
	case LangTS, LangJS:
		tsOpts := codegen_ts.Options{
			DisableTypes: options.Language == LangJS,
			RuntimePath:  options.RuntimePath,
		}
		output = codegen_ts.Codegen(info, eppFile, &tsOpts)
		return
	case LangJava:
		javaOpts := codegen_java.Options{
			Package:   &options.PackageName,
			ClassName: "Choreography",
		}
		output = codegen_java.Codegen(info, eppFile, &javaOpts)
		return
	default:
		panic(fmt.Sprintf("unknown language: %v", options.Language))
	}
}
