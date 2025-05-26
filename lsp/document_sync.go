package lsp

import (
	"context"
	"slices"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

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
	newSource := file.GetSource()

	// sort edits to be in reverse order
	slices.SortFunc(params.ContentChanges, func(a, b any) int {
		aa, ok := a.(protocol.TextDocumentContentChangeEvent)
		if !ok {
			return 0
		}

		bb, ok := b.(protocol.TextDocumentContentChangeEvent)
		if !ok {
			return 0
		}

		if aa.Range.Start.Line != bb.Range.Start.Line {
			return int(bb.Range.Start.Line) - int(aa.Range.Start.Line)
		}

		return int(bb.Range.Start.Character) - int(aa.Range.Start.Character)
	})

	for _, change := range params.ContentChanges {
		switch change := change.(type) {
		case protocol.TextDocumentContentChangeEventWhole:
			newSource = change.Text
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

			newSource = source[0:startIdx] + change.Text + newSource[endIdx:]
		default:
			logger.Errorf("unexpected type: %#v", change)
		}
	}

	file.ReplaceSource(newSource)

	go s.analyzeFile(ctx, file)

	return nil
}
