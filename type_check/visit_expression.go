package type_check

import (
	"chorego/parser"
	"chorego/types"
)

func (tc *typeChecker) VisitExpression(ctx *parser.ExpressionContext) any {

	if ident := ctx.Ident(); ident != nil {
		sym, err := tc.symTable.LookupSymbol(ident)
		if err != nil {
			tc.reportError(err)
			return types.Invalid()
		}
		return sym.Type()
	}

	if num := ctx.NUMBER(); num != nil {
		return types.New(types.Int(), types.NewRole(nil, true))
	}

	panic("unexpected expression")
}
