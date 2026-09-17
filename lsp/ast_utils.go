package lsp

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
	"github.com/tempo-lang/tempo/sym_table"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"reflect"
)

func spanToRange(source *token.Source, span token.Span) protocol.Range {
	a, b := source.LSPPosition(span.Start), source.LSPPosition(span.End)
	return protocol.Range{Start: protocol.Position{Line: uint32(a.Line), Character: uint32(a.Character)}, End: protocol.Position{Line: uint32(b.Line), Character: uint32(b.Character)}}
}
func parserRuleToRange(source *token.Source, node ast.Node) protocol.Range {
	return spanToRange(source, token.Span{Start: node.StartToken().Span.Start, End: node.EndToken().Span.End})
}
func scopeToRange(source *token.Source, scope *sym_table.Scope) protocol.Range {
	return spanToRange(source, scope.Span())
}
func posWithinRange(p protocol.Position, r protocol.Range) bool {
	return (p.Line > r.Start.Line || p.Line == r.Start.Line && p.Character >= r.Start.Character) && (p.Line < r.End.Line || p.Line == r.End.Line && p.Character <= r.End.Character)
}
func rangesOverlap(a, b protocol.Range) bool {
	return posWithinRange(a.Start, b) || posWithinRange(a.End, b) || posWithinRange(b.Start, a)
}

func astNodeAtPosition(source *token.Source, root ast.Node, pos protocol.Position) (ast.Node, protocol.Range) {
	nodes := astNodesAtPosition(source, root, pos)
	if len(nodes) == 0 {
		return nil, protocol.Range{}
	}
	best := nodes[len(nodes)-1]
	return best, parserRuleToRange(source, best)
}

func astNodesAtPosition(source *token.Source, root ast.Node, pos protocol.Position) (matches []ast.Node) {
	seen := map[uintptr]bool{}
	var walk func(reflect.Value)
	walk = func(v reflect.Value) {
		if !v.IsValid() {
			return
		}
		if v.Kind() == reflect.Interface {
			if !v.IsNil() {
				walk(v.Elem())
			}
			return
		}
		if v.Kind() == reflect.Pointer {
			if v.IsNil() || seen[v.Pointer()] {
				return
			}
			seen[v.Pointer()] = true
			if v.CanInterface() {
				if n, ok := v.Interface().(ast.Node); ok {
					if !posWithinRange(pos, parserRuleToRange(source, n)) {
						return
					}
					matches = append(matches, n)
				}
			}
			walk(v.Elem())
			return
		}
		switch v.Kind() {
		case reflect.Struct:
			for i := 0; i < v.NumField(); i++ {
				walk(v.Field(i))
			}
		case reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				walk(v.Index(i))
			}
		}
	}
	walk(reflect.ValueOf(root))
	return matches
}
