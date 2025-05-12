package epp

import (
	"fmt"
	"tempo/parser"
	"tempo/projection"
	"tempo/type_check"
	"tempo/types"
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
	case *parser.StmtIfContext:
		guard, aux := eppExpression(info, roleName, stmt.Expr())
		for _, e := range aux {
			result = append(result, projection.NewStmtExpr(e))
		}

		guardType := info.Types[stmt.Expr()]

		if guardType.Roles().Contains(roleName) {
			thenBranch := []projection.Statement{}
			for _, s := range stmt.Scope(0).AllStmt() {
				thenBranch = append(thenBranch, EppStmt(info, roleName, s)...)
			}

			elseBranch := []projection.Statement{}
			if elseScope := stmt.Scope(1); elseScope != nil {
				for _, s := range elseScope.AllStmt() {
					elseBranch = append(elseBranch, EppStmt(info, roleName, s)...)
				}
			}

			result = append(result, projection.NewStmtIf(guard, thenBranch, elseBranch))
		}

	case *parser.StmtContext:
		panic("statement should never be base type")
	default:
		panic(fmt.Sprintf("unknown statement: %#v", stmt))
	}

	return
}
