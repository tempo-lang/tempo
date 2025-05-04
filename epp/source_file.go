package epp

import (
	"chorego/parser"
	"chorego/projection"
	"chorego/type_check"
)

func EppSourceFile(info *type_check.Info, sourceFile parser.ISourceFileContext) *projection.SourceFile {
	result := projection.NewSourceFile()
	for _, fn := range sourceFile.AllFunc_() {
		chor := EppFunc(info, fn)
		result.AddChoreography(chor)
	}
	return result
}
