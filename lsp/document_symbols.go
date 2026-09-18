package lsp

import (
	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/parser/token"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) documentSymbols(context *glsp.Context, params *protocol.DocumentSymbolParams) (any, error) {
	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	symbols := documentSymbolsInScope(doc.source, doc.info.GlobalScope)

	return symbols, nil
}

func documentSymbolsInScope(source *token.Source, scope *sym_table.Scope) []protocol.DocumentSymbol {
	result := []protocol.DocumentSymbol{}

	for sym := range scope.Symbols() {
		if sym.SymbolName() == "" {
			continue
		}

		symNameRange := parserRuleToRange(source, sym.Ident())
		newSym := protocol.DocumentSymbol{
			Name:           sym.SymbolName(),
			Detail:         misc.ToPtr(sym.Type().ToString()),
			Range:          symNameRange,
			SelectionRange: symNameRange,
			Children:       []protocol.DocumentSymbol{},
		}

		switch sym := sym.(type) {
		case *sym_table.FuncSymbol:
			newSym.Kind = protocol.SymbolKindFunction
			newSym.Range = scopeToRange(source, sym.Scope())
			newSym.Children = documentSymbolsInScope(source, sym.Scope())
		// case *sym_table.InterfaceMethodSymbol:
		// 	newSym.Kind = protocol.SymbolKindMethod
		case *sym_table.InterfaceSymbol:
			newSym.Kind = protocol.SymbolKindInterface
			newSym.Range = scopeToRange(source, sym.Scope())
			newSym.Children = documentSymbolsInScope(source, sym.Scope())
		case *sym_table.StructFieldSymbol:
			newSym.Kind = protocol.SymbolKindField
		case *sym_table.StructSymbol:
			newSym.Kind = protocol.SymbolKindStruct
			newSym.Range = scopeToRange(source, sym.Scope())
			newSym.Children = documentSymbolsInScope(source, sym.Scope())
		case *sym_table.FuncParamSymbol:
			// skip to not clutter
			continue
		case *sym_table.VariableSymbol:
			// skip to not clutter
			continue
		default:
			logger.Errorf("unexpected symbol in documentSymbolsInScope: %#v", sym)
		}

		result = append(result, newSym)
	}

	return result
}
