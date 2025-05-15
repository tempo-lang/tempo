package epp

import (
	"tempo/parser"
	"tempo/projection"
)

func (epp *epp) eppSourceFile(sourceFile parser.ISourceFileContext) *projection.SourceFile {
	result := projection.NewSourceFile()
	for _, fn := range sourceFile.AllFunc_() {
		chor := epp.eppFunc(fn)
		result.AddChoreography(chor)
	}
	return result
}
