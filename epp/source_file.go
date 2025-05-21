package epp

import (
	"tempo/parser"
	"tempo/projection"
)

func (epp *epp) eppSourceFile(sourceFile parser.ISourceFileContext) *projection.SourceFile {
	result := projection.NewSourceFile()

	for _, st := range sourceFile.AllStruct_() {
		eppSt := epp.eppStruct(st)
		result.AddStruct(eppSt)
	}

	for _, fn := range sourceFile.AllFunc_() {
		chor := epp.eppFunc(fn)
		result.AddChoreography(chor)
	}
	return result
}
