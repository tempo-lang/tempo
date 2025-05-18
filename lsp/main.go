package main

import (
	"context"
	"sync"
	"tempo/type_check"

	"github.com/tliron/commonlog"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"

	_ "github.com/tliron/commonlog/simple"
)

const lsName = "TempoLS"

var version string = "0.0.1"

var logger commonlog.Logger

type tempoFile struct {
	lock        sync.RWMutex
	uri         protocol.DocumentUri
	source      string
	info        *type_check.Info
	analysisCtx context.Context
}

func newTempoFile(document protocol.TextDocumentItem, analysisCtx context.Context) *tempoFile {
	return &tempoFile{
		uri:         document.URI,
		source:      document.Text,
		info:        nil,
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

func (f *tempoFile) SetInfo(info *type_check.Info) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.info = info
}

func (f *tempoFile) ReplaceSource(source string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.source = source
}

type tempoServer struct {
	files map[protocol.DocumentUri]*tempoFile
}

func newTempoServer() *tempoServer {
	return &tempoServer{
		files: map[protocol.DocumentUri]*tempoFile{},
	}
}

func main() {
	commonlog.Configure(2, nil)

	logger = commonlog.GetLoggerf("%s.handler", lsName)

	tempoServer := newTempoServer()

	server := server.NewServer(tempoServer.Handler(), lsName, true)

	server.RunStdio()
}

func (s *tempoServer) Handler() *protocol.Handler {
	return &protocol.Handler{
		Initialize:             s.initialize,
		Initialized:            s.initialized,
		Shutdown:               s.shutdown,
		TextDocumentDidOpen:    s.textDocumentDidOpen,
		TextDocumentDidChange:  s.textDocumentDidChange,
		TextDocumentCompletion: s.textDocumentCompletion,
	}
}

func (s *tempoServer) initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	logger.Info("Initializing server...")

	capabilities := s.Handler().CreateServerCapabilities()

	logger.Infof("Capabilities: %#v", capabilities)

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func (s *tempoServer) initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}

func (s *tempoServer) shutdown(context *glsp.Context) error {
	return nil
}

func (s *tempoServer) textDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	var completionItems []protocol.CompletionItem

	return completionItems, nil
}
