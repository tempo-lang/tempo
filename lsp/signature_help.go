package lsp

import (
	"fmt"
	"strings"

	"github.com/tempo-lang/tempo/parser/ast"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *tempoServer) signatureHelp(_ *glsp.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	doc, ok := s.GetDocument(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}
	var call *ast.CallExpr
	for _, n := range astNodesAtPosition(doc.source, doc.ast, params.Position) {
		if c, ok := n.(*ast.CallExpr); ok {
			call = c
		}
	}
	if call == nil {
		return nil, nil
	}
	t := doc.info.Types[call.Function]
	labels := []string{}
	parameters := []protocol.ParameterInformation{}
	var title string
	switch f := t.(type) {
	case *types.FunctionType:
		sym := doc.info.Symbols[f.NameIdent()].(*sym_table.FuncSymbol)
		subst, _ := sym.Roles().SubstituteMap(f.Roles())
		for _, p := range sym.Params() {
			labels = append(labels, fmt.Sprintf("%s: %s", p.SymbolName(), p.Type().SubstituteRoles(subst).ToString()))
		}
		title = fmt.Sprintf("func@%s %s(%s)", sym.Roles().SubstituteRoles(subst).ToString(), sym.SymbolName(), strings.Join(labels, ", "))
		if f.ReturnType() != types.Unit() {
			title += " " + f.ReturnType().ToString()
		}
	case *types.ClosureType:
		for _, p := range f.Params() {
			labels = append(labels, p.ToString())
		}
		title = f.ToString()
	default:
		return nil, nil
	}
	for _, label := range labels {
		parameters = append(parameters, protocol.ParameterInformation{Label: label})
	}
	var active *protocol.UInteger
	for i, arg := range call.Args {
		if posWithinRange(params.Position, parserRuleToRange(doc.source, arg)) {
			v := uint32(i)
			active = &v
		}
	}
	return &protocol.SignatureHelp{Signatures: []protocol.SignatureInformation{{Label: title, Parameters: parameters, ActiveParameter: active}}}, nil
}
