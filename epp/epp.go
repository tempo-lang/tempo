package epp

import (
	"chorego/misc"
	"chorego/parser"
	"chorego/type_check"
	"slices"

	"github.com/antlr4-go/antlr/v4"
	"github.com/dave/jennifer/jen"
)

func EndpointProject(input antlr.CharStream) (output string, errors []error) {
	// input, _ := antlr.NewFileStream(os.Args[1])
	errorListener := misc.ErrorListener{}

	// lexer
	lexer := parser.NewChoregoLexer(input)
	lexer.AddErrorListener(&errorListener)

	// parser
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewChoregoParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	p.AddErrorListener(&errorListener)

	// parse program
	sourceFile := p.SourceFile()

	if len(errorListener.Errors) > 0 {
		errors = errorListener.Errors
		return
	}

	typeErrors := type_check.TypeCheck(sourceFile)
	if len(typeErrors) > 0 {
		for _, err := range typeErrors {
			errors = append(errors, err)
		}
		return
	}

	eppFile := EppSourceFile(sourceFile)

	file := jen.NewFile("choreography")
	eppFile.Codegen(file)

	output = file.GoString()
	return
}

func ValueExistsAtRole(value parser.IValueTypeContext, roleName string) bool {
	var roles []parser.IIdentContext
	if roleTypeNormal := value.RoleType().RoleTypeNormal(); roleTypeNormal != nil {
		roles = roleTypeNormal.AllIdent()
	}
	if roleTypeShared := value.RoleType().RoleTypeShared(); roleTypeShared != nil {
		roles = roleTypeShared.AllIdent()
	}

	containsRole := slices.ContainsFunc(roles, func(role parser.IIdentContext) bool {
		return role.GetText() == roleName
	})
	return containsRole
}
