package lsp

import (
	"github.com/antlr4-go/antlr/v4"
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

func astNodeAtPosition(node antlr.ParserRuleContext, pos protocol.Position) (antlr.TerminalNode, protocol.Range) {
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

	return nil, parserRuleToRange(node)
}
