package epp

import (
	"tempo/parser"
	"tempo/projection"
	"tempo/type_check"
)

func EppSourceFile(info *type_check.Info, sourceFile parser.ISourceFileContext) *projection.SourceFile {
	result := projection.NewSourceFile()
	for _, fn := range sourceFile.AllFunc_() {
		chor := EppFunc(info, fn)
		result.AddChoreography(chor)
	}
	return result
}
