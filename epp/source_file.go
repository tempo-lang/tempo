package epp

import (
	"chorego/parser"
	"chorego/projection"
)

func EppSourceFile(sourceFile parser.ISourceFileContext) *projection.SourceFile {
	result := projection.NewSourceFile()
	for _, fn := range sourceFile.AllFunc_() {
		chor := EppFunc(fn)
		result.AddChoreography(chor)
	}
	return result
}
