package lsp

import (
	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) textDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	leaf, _ := astNodeAtPosition(doc.ast, params.Position)
	if leaf == nil {
		return nil, nil
	}

	scope := doc.info.GlobalScope.Innermost(leaf.GetStart())

	var completionItems []protocol.CompletionItem
	foundItems := false
	showGlobalSymbols := false

	var node antlr.Tree = leaf
	for node != nil {
		switch node := node.(type) {
		case *parser.ExprFieldAccessContext:
			logger.Debugf("Visit Expression (FieldAccess)")
			items, ok := completionItemsForFieldAccess(doc, node)
			if ok {
				completionItems = items
				foundItems = true
			}
		case *parser.ExprIdentContext:
			logger.Debugf("Visit Expression (Identifier)")
			showGlobalSymbols = true
		case parser.IExprContext:
			logger.Debugf("Visit Expression: %T", node)
		case parser.IStmtContext:
			logger.Debugf("Visit Statement: %T", node)
		case *parser.ScopeContext:
			logger.Debugf("Visit Scope")
		case *parser.FuncContext:
			logger.Debugf("Visit Function")
		case parser.IRoleTypeContext:
			logger.Debugf("Visit Role Type: %T", node)
			items, ok := completionItemsForRoleType(doc, node)
			if ok {
				completionItems = items
				foundItems = true
			}
		case *antlr.TerminalNodeImpl:
			logger.Debugf("Visit Terminal: %s", node.GetText())
		case *antlr.BaseParserRuleContext:
			// Attempt to recover ExprFieldAccess
			if leaf.GetText() == "." {
				logger.Debug("Attempting to recover ExprFieldAccess")
				if fieldAccess, ok := node.GetChild(0).GetParent().(*parser.ExprFieldAccessContext); ok {
					items, ok := completionItemsForFieldAccess(doc, fieldAccess)
					if ok {
						completionItems = items
						foundItems = true
					}
				}
			}

			// Attempt to recover RoleType
			if leaf.GetText() == "@" {
				logger.Debug("Attempting to recover RoleType")
				if roleType, ok := node.GetChild(node.GetChildCount() - 1).(parser.IRoleTypeContext); ok {
					items, ok := completionItemsForRoleType(doc, roleType)
					if ok {
						completionItems = items
						foundItems = true
					}
				}
			}

			logger.Debugf("Visit Base: %#v", node)
		default:
			logger.Debugf("Visit Other: %T", node)
		}

		node = node.GetParent()
		if foundItems {
			break
		}
	}

	if !foundItems && showGlobalSymbols {
		completionItems = completionItemsAllInScope(scope)
	}

	return protocol.CompletionList{
		IsIncomplete: false,
		Items:        completionItems,
	}, nil
}

func symbolToCompletionItem(sym sym_table.Symbol) protocol.CompletionItem {

	var kind *protocol.CompletionItemKind = nil
	switch sym.(type) {
	case *sym_table.FuncParamSymbol:
		kind = toPtr(protocol.CompletionItemKindProperty)
	case *sym_table.FuncSymbol:
		kind = toPtr(protocol.CompletionItemKindFunction)
	case *sym_table.StructSymbol:
		kind = toPtr(protocol.CompletionItemKindStruct)
	case *sym_table.StructFieldSymbol:
		kind = toPtr(protocol.CompletionItemKindField)
	case *sym_table.VariableSymbol:
		kind = toPtr(protocol.CompletionItemKindVariable)
	case *sym_table.InterfaceMethodSymbol:
		kind = toPtr(protocol.CompletionItemKindMethod)
	case *sym_table.InterfaceSymbol:
		kind = toPtr(protocol.CompletionItemKindInterface)
	}

	return protocol.CompletionItem{
		Label:  sym.SymbolName(),
		Kind:   kind,
		Detail: toPtr(sym.Type().ToString()),
	}
}

func completionItemsAllInScope(scope *sym_table.Scope) []protocol.CompletionItem {
	completionItems := []protocol.CompletionItem{}
	cursorScope := scope
	for cursorScope != nil {
		for sym := range cursorScope.Symbols() {
			completionItems = append(completionItems, symbolToCompletionItem(sym))
		}
		cursorScope = cursorScope.Parent()
	}
	return completionItems
}

func completionItemsForFieldAccess(file *tempoDoc, fieldAccess *parser.ExprFieldAccessContext) ([]protocol.CompletionItem, bool) {
	logger.Infof("Finding completions for ExprFieldAccess")

	fieldType, ok := file.info.Types[fieldAccess.Expr()]
	if !ok {
		logger.Debugf("")
		return nil, false
	}

	switch fieldType := fieldType.Value().(type) {
	case *types.StructType:
		scope := file.info.GlobalScope.Innermost(fieldAccess.GetStart())

		structSym := scope.LookupParent(fieldType.Name()).(*sym_table.StructSymbol)

		completionItems := []protocol.CompletionItem{}
		for _, field := range structSym.Fields() {
			completionItems = append(completionItems, symbolToCompletionItem(field))
		}
		return completionItems, true
	case *types.InterfaceType:
		infSym := file.info.Symbols[fieldType.Ident()].(*sym_table.InterfaceSymbol)

		completionItems := []protocol.CompletionItem{}
		for _, method := range infSym.Methods() {
			completionItems = append(completionItems, symbolToCompletionItem(method))
		}
		return completionItems, true
	}

	return nil, false
}

func completionItemsForRoleType(file *tempoDoc, roleType parser.IRoleTypeContext) ([]protocol.CompletionItem, bool) {
	logger.Infof("Finding completions for RoleType")

	scope := file.info.GlobalScope.Innermost(roleType.GetStart())
	if scope == nil {
		return nil, false
	}

	completionItems := []protocol.CompletionItem{}
	for _, role := range scope.Roles().Participants() {
		detail := "role"

		completionItems = append(completionItems, protocol.CompletionItem{
			Label:  role,
			Detail: &detail,
		})
	}

	return completionItems, true
}
