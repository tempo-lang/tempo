package lsp

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/tempo-lang/tempo/sym_table"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

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

func scopeToRange(scope *sym_table.Scope) protocol.Range {
	endTokenLength := scope.End().GetStop() - scope.End().GetStart()
	return protocol.Range{
		Start: protocol.Position{
			Line:      uint32(scope.Pos().GetLine() - 1),
			Character: uint32(scope.Pos().GetColumn()),
		},
		End: protocol.Position{
			Line:      uint32(scope.End().GetLine() - 1),
			Character: uint32(scope.End().GetColumn() + endTokenLength + 1),
		},
	}
}

func tokenToRange(token antlr.Token) protocol.Range {
	endTokenLength := token.GetStop() - token.GetStart()
	return protocol.Range{
		Start: protocol.Position{
			Line:      uint32(token.GetLine() - 1),
			Character: uint32(token.GetColumn()),
		},
		End: protocol.Position{
			Line:      uint32(token.GetLine() - 1),
			Character: uint32(token.GetColumn() + endTokenLength + 1),
		},
	}
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

func astNodeAtPosition(node antlr.ParserRuleContext, pos protocol.Position) (antlr.ParserRuleContext, protocol.Range) {
	for _, c := range node.GetChildren() {
		switch child := c.(type) {
		case antlr.ParserRuleContext:
			span := parserRuleToRange(child)
			if posWithinRange(pos, span) {
				if result, resultSpan := astNodeAtPosition(child, pos); result != nil {
					return result, resultSpan
				} else {
					logger.Debugf("AST Node at pos (result) %T", child)
					return child, span
				}
			}
		case antlr.TerminalNode:
			span := tokenToRange(child.GetSymbol())
			if posWithinRange(pos, span) {
				logger.Debugf("AST Node at pos (terminal) %T", child)
				return nil, protocol.Range{}
			}
		}
	}

	return node, parserRuleToRange(node)
}

func rangesOverlap(a, b protocol.Range) bool {
	posLT := func(p1, p2 protocol.Position) bool {
		if p1.Line != p2.Line {
			return p1.Line < p2.Line
		}
		return p1.Character < p2.Character
	}

	// No overlap if a ends before b starts, or b ends before a starts.
	if posLT(a.End, b.Start) || posLT(b.End, a.Start) {
		return false
	}
	return true
}
