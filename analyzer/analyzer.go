package analyzer

import (
	"chorego/parser"
)

type AnalyzerListener struct {
	parser.BaseChoregoListener
	ErrorListener ErrorListener
}

func New() *AnalyzerListener {
	return &AnalyzerListener{
		ErrorListener: &DefaultErrorListener{},
	}
}

func (a *AnalyzerListener) EnterFunc(ctx *parser.FuncContext) {
	a.checkFuncDuplicateRoles(ctx)
}
