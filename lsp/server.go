package lsp

import (
	"context"
	"fmt"
	"sync"
	"tempo/parser"
	"tempo/type_check"
	"tempo/types"

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
		Initialize:             s.initialize,
		Initialized:            s.initialized,
		Shutdown:               s.shutdown,
		TextDocumentDidOpen:    s.textDocumentDidOpen,
		TextDocumentDidChange:  s.textDocumentDidChange,
		TextDocumentCompletion: s.textDocumentCompletion,
		TextDocumentHover:      s.textDocumentHover,
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

func (s *tempoServer) textDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {

	file, ok := s.files[params.TextDocument.URI]
	if !ok {
		return nil, nil
	}

	file.lock.RLock()
	defer file.lock.RUnlock()

	leaf, nodeRange := astNodeAtPosition(file.ast, params.Position)

	node := leaf
	for node != nil {
		switch node := node.(type) {
		case *parser.StmtVarDeclContext:
			if sym, ok := file.info.Symbols[node.Ident()]; ok {
				exprRange := parserRuleToRange(node)
				return &protocol.Hover{
					Contents: fmt.Sprintf("let %s: %s", sym.SymbolName(), sym.Type().ToString()),
					Range:    &exprRange,
				}, nil
			}
		case parser.IExprContext:
			if exprType, ok := file.info.Types[node]; ok {

				if len(exprType.Roles().Participants()) == 0 {
					scope := file.info.GlobalScope.Innermost(node.GetStart())

					exprType = types.New(
						exprType.Value(),
						types.NewRole(scope.Roles().Participants(), true),
					)
				}

				exprRange := parserRuleToRange(node)
				return &protocol.Hover{
					Contents: exprType.ToString(),
					Range:    &exprRange,
				}, nil
			}
		}

		node = node.GetParent()
	}

	debugNode := "```"
	node = leaf
	for node != nil {
		debugNode = fmt.Sprintf("%s\n\n%#v", debugNode, node)
		node = node.GetParent()
	}
	debugNode = fmt.Sprintf("%s\n```", debugNode)

	return &protocol.Hover{
		Contents: &protocol.MarkupContent{
			Kind:  protocol.MarkupKindMarkdown,
			Value: debugNode,
		},
		Range: &nodeRange,
	}, nil
}

func posWithinRange(pos protocol.Position, span protocol.Range) bool {
	if span.Start.Line <= pos.Line && span.End.Line >= pos.Line {
		if span.Start.Line == pos.Line && pos.Character < span.Start.Character {
			return false
		}

		if span.End.Line == pos.Line && span.End.Character < pos.Character {
			return false
		}

		return true
	}

	return false
}

func astNodeAtPosition(node antlr.ParserRuleContext, pos protocol.Position) (antlr.Tree, protocol.Range) {
	for _, c := range node.GetChildren() {
		switch child := c.(type) {
		case antlr.ParserRuleContext:
			span := parserRuleToRange(child)
			if posWithinRange(pos, span) {
				return astNodeAtPosition(child, pos)
			}
		case antlr.TerminalNode:
			span := tokenToRange(child.GetSymbol())
			if posWithinRange(pos, span) {
				return child, span
			}
		}
	}
	return node, parserRuleToRange(node)
}
