package epp

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/projection"
)

func (epp *epp) eppSourceFile(sourceFile *ast.SourceFile) *projection.SourceFile {
	result := projection.NewSourceFile()

	for _, inf := range sourceFile.Interfaces {
		eppInf := epp.eppInterface(inf)
		result.AddInterface(eppInf)
	}

	for _, st := range sourceFile.Structs {
		eppSt := epp.eppStruct(st)
		result.AddStruct(eppSt)
	}

	for _, fn := range sourceFile.Functions {
		chor := epp.eppFunc(fn)
		result.AddChoreography(chor)
	}
	return result
}
