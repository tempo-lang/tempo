package lsp

import (
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
	files map[protocol.DocumentUri]*tempoFile
}

func newTempoServer() *tempoServer {
	return &tempoServer{
		files: map[protocol.DocumentUri]*tempoFile{},
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
		Initialize:                s.initialize,
		Initialized:               s.initialized,
		Shutdown:                  s.shutdown,
		TextDocumentDidOpen:       s.textDocumentDidOpen,
		TextDocumentDidChange:     s.textDocumentDidChange,
		TextDocumentCompletion:    s.textDocumentCompletion,
		TextDocumentHover:         s.textDocumentHover,
		TextDocumentSignatureHelp: s.signatureHelp,
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
