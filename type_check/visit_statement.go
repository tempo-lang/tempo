package type_check

import (
	"chorego/parser"
	"chorego/type_check/sym_table"
	"chorego/type_check/type_error"
	"chorego/type_check/types"
)

func (tc *typeChecker) VisitStatement(ctx *parser.StatementContext) any {
	if varDecl := ctx.StmtVarDecl(); varDecl != nil {
		return varDecl.Accept(tc)
	}

	panic("unexpected statement")
}

func (tc *typeChecker) VisitStmtVarDecl(ctx *parser.StmtVarDeclContext) any {

	exprType := ctx.Expression().Accept(tc).(types.Type)
	declType, err := ParseValueType(ctx.ValueType())
	if err != nil {
		tc.reportError(err)
	} else if exprType != declType {
		tc.reportError(type_error.NewTypeMismatchError(ctx.ValueType(), declType, ctx.Expression(), exprType))
	}

	if err := tc.symTable.InsertSymbol(sym_table.NewVariableSymbol(ctx)); err != nil {
		tc.reportError(err)
	}

	return tc.VisitChildren(ctx)
}
