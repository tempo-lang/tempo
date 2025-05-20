package type_error

import (
	"fmt"
	"tempo/parser"
	"tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

type SymbolAlreadyExists struct {
	ExistingSymbol parser.IIdentContext
	NewSymbol      parser.IIdentContext
}

func NewSymbolAlreadyExistsError(existing parser.IIdentContext, newSym parser.IIdentContext) Error {
	return &SymbolAlreadyExists{
		ExistingSymbol: existing,
		NewSymbol:      newSym,
	}
}

func (s *SymbolAlreadyExists) IsTypeError() {}

func (s *SymbolAlreadyExists) Error() string {
	return fmt.Sprintf("symbol '%s' already declared", s.NewSymbol.GetText())
}

func (e *SymbolAlreadyExists) ParserRule() antlr.ParserRuleContext {
	return e.NewSymbol
}

type UnknownSymbolError struct {
	SymName parser.IIdentContext
}

func NewUnknownSymbolError(symName parser.IIdentContext) Error {
	return &UnknownSymbolError{
		SymName: symName,
	}
}

func (e *UnknownSymbolError) Error() string {
	return fmt.Sprintf("unknown symbol '%s'", e.SymName.GetText())
}

func (e *UnknownSymbolError) IsTypeError() {}

func (e *UnknownSymbolError) ParserRule() antlr.ParserRuleContext {
	return e.SymName
}

type UnassignableSymbolError struct {
	Assign *parser.StmtAssignContext
	Type   *types.Type
}

func (u *UnassignableSymbolError) Error() string {
	return fmt.Sprintf("can not assign value to '%s' of type '%s'", u.Assign.Ident().GetText(), u.Type.ToString())
}

func (u *UnassignableSymbolError) IsTypeError() {}

func (u *UnassignableSymbolError) ParserRule() antlr.ParserRuleContext {
	return u.Assign
}

func NewUnassignableSymbolError(assign *parser.StmtAssignContext, symType *types.Type) Error {
	return &UnassignableSymbolError{
		Assign: assign,
		Type:   symType,
	}
}
