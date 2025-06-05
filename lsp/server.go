package lsp

import (
	"sync"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/type_check"
	"github.com/tliron/commonlog"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"

	_ "github.com/tliron/commonlog/simple"
)

const lsName = "TempoLS"

var version string = "0.0.1"

var logger commonlog.Logger

type tempoServer struct {
	lock      sync.RWMutex
	documents map[protocol.DocumentUri]*tempoDoc
}

func (s *tempoServer) GetDocument(url protocol.DocumentUri) (*tempoDoc, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	doc, ok := s.documents[url]
	return doc, ok
}

func (s *tempoServer) UpdateDocument(newDoc *tempoDoc) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if existingDoc, ok := s.documents[newDoc.uri]; ok {
		if existingDoc.version > newDoc.version {
			logger.Infof("newer version of document already exists (existing=%d, this=%d): %s",
				existingDoc.version, newDoc.version, newDoc.uri)
			return
		}
	}

	s.documents[newDoc.uri] = newDoc
}

type tempoDoc struct {
	uri     protocol.URI
	version int
	source  string
	ast     parser.ISourceFileContext
	info    *type_check.Info
}

func newTempoDoc(uri protocol.URI, version int, source string, ast parser.ISourceFileContext, info *type_check.Info) *tempoDoc {
	return &tempoDoc{
		ast:     ast,
		info:    info,
		uri:     uri,
		version: version,
		source:  source,
	}
}

func newTempoServer() *tempoServer {
	return &tempoServer{
		documents: map[protocol.DocumentUri]*tempoDoc{},
	}
}

func StartServer() {
	commonlog.Configure(2, nil)

	logger = commonlog.GetLoggerf("%s.handler", lsName)

	tempoServer := newTempoServer()

	server := server.NewServer(tempoServer.Handler(), lsName, true)

	server.RunStdio()
}

func (s *tempoServer) Handler() *protocol.Handler {
	return &protocol.Handler{
		Initialize:                    s.initialize,
		Initialized:                   s.initialized,
		Shutdown:                      s.shutdown,
		TextDocumentDidOpen:           s.textDocumentDidOpen,
		TextDocumentDidChange:         s.textDocumentDidChange,
		TextDocumentCompletion:        s.textDocumentCompletion,
		TextDocumentHover:             s.textDocumentHover,
		TextDocumentSignatureHelp:     s.signatureHelp,
		TextDocumentDocumentHighlight: s.highlight,
		TextDocumentDeclaration:       s.gotoDeclaration,
		TextDocumentDefinition:        s.gotoDefinition,
		TextDocumentReferences:        s.gotoReferences,
		TextDocumentDocumentSymbol:    s.documentSymbols,
		TextDocumentPrepareRename:     s.prepareRename,
		TextDocumentRename:            s.renameSymbol,
	}
}

func (s *tempoServer) initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	logger.Info("Initializing server...")

	capabilities := s.Handler().CreateServerCapabilities()
	capabilities.TextDocumentSync = protocol.TextDocumentSyncKindFull

	capabilities.CompletionProvider = &protocol.CompletionOptions{
		TriggerCharacters: []string{".", "@"},
	}

	capabilities.SignatureHelpProvider = &protocol.SignatureHelpOptions{
		TriggerCharacters: []string{"("},
	}

	capabilities.RenameProvider = &protocol.RenameOptions{
		PrepareProvider: toPtr(true),
	}

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

// toPtr converts a value to a pointer to the value
func toPtr[T any](val T) *T {
	return &val
}
