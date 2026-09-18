package lsp

import (
	"testing"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/parser/token"
)

func TestSpanToRangeUTF16(t *testing.T) {
	source := token.SourceFromString("// 😀\n\"é😀\"\n")
	tests := []struct {
		name             string
		span             token.Span
		line, start, end uint32
	}{
		{"multiline-comment-boundary", token.Span{Start: 8, End: 8}, 1, 0, 0},
		{"unicode-string", token.Span{Start: 8, End: 16}, 1, 0, 5},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := spanToRange(source, test.span)
			if r.Start.Line != test.line || r.End.Line != test.line || r.Start.Character != test.start || r.End.Character != test.end {
				t.Fatalf("range = %#v", r)
			}
		})
	}
}

func TestMissingTokenDiagnosticAndFixRanges(t *testing.T) {
	source := token.SourceFromString("func@A f() { let x = 1 let y = 2; }")
	result := parser.ParseSource(source)
	if len(result.Diagnostics) == 0 || len(result.Diagnostics[0].Fixes) == 0 {
		t.Fatalf("expected missing-token diagnostic and fix: %#v", result.Diagnostics)
	}
	diagnostic := syntaxDiagnosticToDiagnostic(source, result.Diagnostics[0])
	fix := spanToRange(source, result.Diagnostics[0].Fixes[0].Span)
	if diagnostic.Code == nil || diagnostic.Source == nil || diagnostic.Severity == nil {
		t.Fatalf("incomplete LSP diagnostic: %#v", diagnostic)
	}
	if fix.Start != fix.End {
		t.Fatalf("insertion fix must be zero-width: %#v", fix)
	}
}
