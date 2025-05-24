package lsp

import (
	"context"
	"sync"
	"tempo/parser"
	"tempo/type_check"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type tempoFile struct {
	lock        sync.RWMutex
	uri         protocol.DocumentUri
	source      string
	info        *type_check.Info
	ast         parser.ISourceFileContext
	analysisCtx context.Context
}

func newTempoFile(document protocol.TextDocumentItem, analysisCtx context.Context) *tempoFile {
	return &tempoFile{
		uri:         document.URI,
		source:      document.Text,
		info:        nil,
		ast:         nil,
		analysisCtx: analysisCtx,
	}
}

func (f *tempoFile) GetSource() string {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.source
}

func (f *tempoFile) GetUri() protocol.DocumentUri {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.uri
}

func (f *tempoFile) SetInfo(ast parser.ISourceFileContext, info *type_check.Info) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.ast = ast
	f.info = info
}

func (f *tempoFile) ReplaceSource(source string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.source = source
}
