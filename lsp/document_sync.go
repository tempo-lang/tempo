package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) textDocumentDidOpen(ctx *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	if params.TextDocument.LanguageID != "tempo" {
		return nil
	}

	logger.Infof("New file opened: %s", params.TextDocument.URI)

	// analyze file in background
	go s.analyzeDocument(ctx.Notify, params.TextDocument.URI, int(params.TextDocument.Version), params.TextDocument.Text)

	return nil
}

func (s *tempoServer) textDocumentDidChange(ctx *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	logger.Infof("File changed: %s", params.TextDocument.URI)

	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		logger.Warningf("Changed file is unknown: %s", params.TextDocument.URI)
		return nil
	}

	if doc.version >= int(params.TextDocument.Version) {
		logger.Warningf("Existing document %d is newer than in change event %d", doc.version, params.TextDocument.Version)
		return nil
	}

	var newSource string

	for _, change := range params.ContentChanges {
		switch change := change.(type) {
		case protocol.TextDocumentContentChangeEventWhole:
			newSource = change.Text
		default:
			logger.Errorf("unexpected type: %#v", change)
		}
	}

	go s.analyzeDocument(ctx.Notify, params.TextDocument.URI, int(params.TextDocument.Version), newSource)

	return nil
}
