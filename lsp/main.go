package main

import (
	"context"
	"sync"
	"tempo/parser"
	"tempo/type_check"

	"github.com/antlr4-go/antlr/v4"
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

func (s *tempoServer) textDocumentDidOpen(ctx *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	if params.TextDocument.LanguageID != "tempo" {
		return nil
	}

	logger.Infof("New file opened: %s", params.TextDocument.URI)

	analysisCtx := context.Background()
	tempoFile := newTempoFile(params.TextDocument, analysisCtx)
	s.files[params.TextDocument.URI] = tempoFile

	// analyze file in background
	go s.analyzeFile(ctx, tempoFile)

	return nil
}

func (s *tempoServer) textDocumentDidChange(ctx *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	logger.Infof("File changed: %s", params.TextDocument.URI)

	file, ok := s.files[params.TextDocument.URI]
	if !ok {
		logger.Warningf("Changed file is unknown: %s", params.TextDocument.URI)
		return nil
	}

	source := file.GetSource()

	for _, change := range params.ContentChanges {
		switch change := change.(type) {
		case protocol.TextDocumentContentChangeEventWhole:
			source = change.Text
		case protocol.TextDocumentContentChangeEvent:
			line := 0
			col := 0
			start := change.Range.Start
			startIdx := -1
			end := change.Range.End
			endIdx := -1

			for i, c := range source {
				if line == int(start.Line) && col == int(start.Character) {
					startIdx = i
				}

				if line == int(end.Line) && col == int(end.Character) {
					endIdx = i
				}

				if startIdx != -1 && endIdx != -1 {
					break
				}

				col += 1
				if c == '\n' {
					line += 1
					col = 0
				}
			}

			source = source[0:startIdx] + change.Text + source[endIdx:]
		default:
			logger.Errorf("unexpected type: %#v", change)
		}
	}

	file.ReplaceSource(source)

	go s.analyzeFile(ctx, file)

	return nil
}

func parserRuleToRange(rule antlr.ParserRuleContext) protocol.Range {
	endTokenLength := rule.GetStop().GetStop() - rule.GetStop().GetStart()
	return protocol.Range{
		Start: protocol.Position{
			Line:      uint32(rule.GetStart().GetLine() - 1),
			Character: uint32(rule.GetStart().GetColumn()),
		},
		End: protocol.Position{
			Line:      uint32(rule.GetStop().GetLine() - 1),
			Character: uint32(rule.GetStop().GetColumn() + endTokenLength + 1),
		},
	}
}

func (s *tempoServer) analyzeFile(ctx *glsp.Context, file *tempoFile) {

	logger.Infof("Analyzing file: %s", file.GetUri())

	diagnostics := []protocol.Diagnostic{}

	// parse source input
	inputStream := antlr.NewInputStream(file.GetSource())
	sourceFile, syntaxErrors := parser.Parse(inputStream)
	for _, err := range syntaxErrors {
		errorSeverity := protocol.DiagnosticSeverityError

		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      uint32(err.Line() - 1),
					Character: uint32(err.Column()),
				},
				End: protocol.Position{
					Line:      uint32(err.Line() - 1),
					Character: uint32(err.Column()),
				},
			},
			Severity: &errorSeverity,
			Message:  err.Message(),
		})
	}

	if len(syntaxErrors) == 0 {
		// type check ast
		info, typeErrors := type_check.TypeCheck(sourceFile)

		file.SetInfo(info)

		for _, err := range typeErrors {
			errorSeverity := protocol.DiagnosticSeverityError

			diagnostics = append(diagnostics, protocol.Diagnostic{
				Range:    parserRuleToRange(err.ParserRule()),
				Severity: &errorSeverity,
				Message:  err.Error(),
			})
		}
	}

	ctx.Notify(protocol.ServerTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
		URI:         file.GetUri(),
		Diagnostics: diagnostics,
	})
}

func (s *tempoServer) textDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var completionItems []protocol.CompletionItem

	return completionItems, nil
}
