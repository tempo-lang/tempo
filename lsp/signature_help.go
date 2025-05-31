package lsp

// func (s *tempoServer) signatureHelp(context *glsp.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
// 	file, ok := s.files[params.TextDocument.URI]
// 	if !ok {
// 		return nil, nil
// 	}
// 	file.lock.RLock()
// 	defer file.lock.RUnlock()

// 	leaf, _ := astNodeAtPosition(file.ast, params.Position)
// 	if leaf == nil {
// 		return nil, nil
// 	}

// 	var exprCall *parser.ExprCallContext = nil

// 	var node antlr.Tree = leaf
// 	for node != nil {
// 		switch node := node.(type) {
// 		case *parser.ExprCallContext:
// 			exprCall = node
// 		default:
// 		}

// 		node = node.GetParent()
// 		if exprCall != nil {
// 			break
// 		}
// 	}

// 	if exprCall == nil {
// 		return nil, nil
// 	}

// 	funcType := file.info.Types[exprCall.Expr()]
// 	funcValue, ok := funcType.Value().(*types.FunctionType)
// 	if !ok {
// 		logger.Debugf("can not give type hint for non-function: %T", funcType.Value())
// 		return nil, nil
// 	}

// 	funcSym := file.info.Symbols[funcValue.FuncIdent()].(*sym_table.FuncSymbol)

// 	parameters := []protocol.ParameterInformation{}

// 	roleSubst := funcValue.RoleSubstitution().Inverse()

// 	paramLabels := []string{}
// 	for _, param := range funcSym.Params() {
// 		paramLabel := fmt.Sprintf("%s: %s", param.SymbolName(), param.Type().SubstituteRoles(roleSubst).ToString())
// 		paramLabels = append(paramLabels, paramLabel)
// 		parameters = append(parameters, protocol.ParameterInformation{
// 			Label:         paramLabel,
// 			Documentation: nil,
// 		})
// 	}

// 	var activeParameter *protocol.UInteger = nil
// 	for i, arg := range exprCall.FuncArgList().AllExpr() {
// 		argRange := parserRuleToRange(arg)
// 		if posWithinRange(params.Position, argRange) {
// 			p := uint32(i)
// 			activeParameter = &p
// 		}
// 	}

// 	funcLabel := fmt.Sprintf("func@%s %s(%s)", funcSym.Roles().SubstituteRoles(roleSubst).ToString(), funcSym.SymbolName(), misc.JoinStrings(paramLabels, ", "))
// 	if funcSym.FuncValue().ReturnType().Value() != types.Unit() {
// 		funcLabel = funcLabel + " " + funcSym.FuncValue().ReturnType().SubstituteRoles(roleSubst).ToString()
// 	}

// 	return &protocol.SignatureHelp{
// 		Signatures: []protocol.SignatureInformation{
// 			{
// 				Label:           funcLabel,
// 				Documentation:   nil,
// 				Parameters:      parameters,
// 				ActiveParameter: activeParameter,
// 			},
// 		},
// 	}, nil
// }
