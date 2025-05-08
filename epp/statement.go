package epp

import (
	"chorego/parser"
	"chorego/projection"
	"chorego/type_check"
	"chorego/types"
	"fmt"
)

func EppStmt(info *type_check.Info, roleName string, stmt parser.IStmtContext) (result []projection.Statement) {
	result = []projection.Statement{}

	switch stmt := stmt.(type) {
	case *parser.StmtAssignContext:
		sym := info.Symbols[stmt.Ident()]

		expr, aux := eppExpression(info, roleName, stmt.Expr())
		for _, e := range aux {
			result = append(result, projection.NewStmtExpr(e))
		}

		if sym.Type().Roles().Contains(roleName) {
			varName := stmt.Ident().GetText()
			result = append(result, projection.NewStmtAssign(varName, expr))
			return
		}
	case *parser.StmtVarDeclContext:
		varSym := info.Symbols[stmt.Ident()]

		expr, aux := eppExpression(info, roleName, stmt.Expr())
		for _, e := range aux {
			result = append(result, projection.NewStmtExpr(e))
		}

		if varSym.Type().Roles().Contains(roleName) {
			variableName := stmt.Ident().GetText()
			varibleType := varSym.Type()

			_, exprIsAsync := expr.Type().(*types.Async)
			if _, isAsync := varibleType.Value().(*types.Async); isAsync && !exprIsAsync {
				expr = projection.NewExprAsync(expr)
			}

			result = append(result, projection.NewStmtVarDecl(variableName, varibleType, expr))
			return
		}
	case *parser.StmtContext:
		panic("statement should never be base type")
	default:
		panic(fmt.Sprintf("unknown statement: %#v", stmt))
	}

	return
}
