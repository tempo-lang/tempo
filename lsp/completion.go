package lsp

import (
	"tempo/parser"
	"tempo/sym_table"
	"tempo/types"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) textDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	file, ok := s.files[params.TextDocument.URI]
	if !ok {
		return nil, nil
	}
	file.lock.RLock()
	defer file.lock.RUnlock()

	leaf, _ := astNodeAtPosition(file.ast, params.Position)
	if leaf == nil {
		return nil, nil
	}

	scope := file.info.GlobalScope.Innermost(leaf.GetStart())

	var completionItems []protocol.CompletionItem
	foundItems := false
	showGlobalSymbols := false

	var node antlr.Tree = leaf
	for node != nil {
		switch node := node.(type) {
		case *parser.ExprFieldAccessContext:
			logger.Debugf("Visit Expression (FieldAccess)")
			items, ok := completionItemsForFieldAccess(file, node)
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
			items, ok := completionItemsForRoleType(file, node)
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
					items, ok := completionItemsForFieldAccess(file, fieldAccess)
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
					items, ok := completionItemsForRoleType(file, roleType)
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
		k := protocol.CompletionItemKindProperty
		kind = &k
	case *sym_table.FuncSymbol:
		k := protocol.CompletionItemKindFunction
		kind = &k
	case *sym_table.StructSymbol:
		k := protocol.CompletionItemKindStruct
		kind = &k
	case *sym_table.StructFieldSymbol:
		k := protocol.CompletionItemKindField
		kind = &k
	case *sym_table.VariableSymbol:
		k := protocol.CompletionItemKindVariable
		kind = &k
	case *sym_table.InterfaceMethodSymbol:
		k := protocol.CompletionItemKindMethod
		kind = &k
	case *sym_table.InterfaceSymbol:
		k := protocol.CompletionItemKindInterface
		kind = &k
	}

	symTypeName := sym.Type().ToString()

	return protocol.CompletionItem{
		Label:  sym.SymbolName(),
		Kind:   kind,
		Detail: &symTypeName,
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

func completionItemsForFieldAccess(file *tempoFile, fieldAccess *parser.ExprFieldAccessContext) ([]protocol.CompletionItem, bool) {
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

func completionItemsForRoleType(file *tempoFile, roleType parser.IRoleTypeContext) ([]protocol.CompletionItem, bool) {
	logger.Infof("Finding completions for RoleType")

	scope := file.info.GlobalScope.Innermost(roleType.GetStart())
	if scope == nil || scope.GetFunc() == nil {
		return nil, false
	}

	funcSym := scope.GetFunc().Roles()
	if funcSym == nil {
		return nil, false
	}

	completionItems := []protocol.CompletionItem{}
	for _, role := range funcSym.Participants() {
		detail := "role"

		completionItems = append(completionItems, protocol.CompletionItem{
			Label:  role,
			Detail: &detail,
		})
	}

	return completionItems, true
}
