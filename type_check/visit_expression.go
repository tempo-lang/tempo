package type_check

import (
	"chorego/parser"
	"chorego/types"
	"strconv"
)

func (tc *typeChecker) VisitExpression(ctx *parser.ExpressionContext) any {

	if ident := ctx.Ident(); ident != nil {
		sym, err := tc.currentScope.LookupSymbol(ident)
		if err != nil {
			tc.reportError(err)
			return types.Invalid()
		}
		return sym.Type()
	}

	if num := ctx.NUMBER(); num != nil {
		_, err := strconv.Atoi(num.GetText())
		if err != nil {
			tc.reportError(types.NewInvalidNumberError(ctx))
		}

		return types.New(types.Int(), types.NewRole(nil, true))
	}

	if boolean := ctx.ExprBool(); boolean != nil {
		boolType := boolean.Accept(tc).(types.Value)
		return types.New(boolType, types.NewRole(nil, true))
	}

	panic("unexpected expression")
}

func (tc *typeChecker) VisitExprBool(ctx *parser.ExprBoolContext) any {
	return types.Bool()
}
