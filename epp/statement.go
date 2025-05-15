package epp

import (
	"fmt"
	"tempo/parser"
	"tempo/projection"
	"tempo/types"
)

func (epp *epp) EppStmt(roleName string, stmt parser.IStmtContext) (result []projection.Statement) {
	result = []projection.Statement{}

	switch stmt := stmt.(type) {
	case *parser.StmtAssignContext:
		sym := epp.info.Symbols[stmt.Ident()]

		expr, aux := epp.eppExpression(roleName, stmt.Expr())
		result = aux

		if sym.Type().Roles().Contains(roleName) {
			varName := stmt.Ident().GetText()
			result = append(result, projection.NewStmtAssign(varName, expr))
			return
		} else if expr != nil && expr.HasSideEffects() {
			result = append(result, projection.NewStmtExpr(expr))
			return
		}
	case *parser.StmtVarDeclContext:
		varSym := epp.info.Symbols[stmt.Ident()]

		expr, aux := epp.eppExpression(roleName, stmt.Expr())
		result = aux

		if varSym.Type().Roles().Contains(roleName) {
			variableName := stmt.Ident().GetText()
			varibleType := varSym.Type()

			_, exprIsAsync := expr.Type().(*types.Async)
			if _, isAsync := varibleType.Value().(*types.Async); isAsync && !exprIsAsync {
				expr = projection.NewExprAsync(expr)
			}

			result = append(result, projection.NewStmtVarDecl(variableName, expr))
			return
		} else if expr != nil && expr.HasSideEffects() {
			result = append(result, projection.NewStmtExpr(expr))
			return
		}
	case *parser.StmtIfContext:
		guard, aux := epp.eppExpression(roleName, stmt.Expr())
		result = aux

		guardType := epp.info.Types[stmt.Expr()]

		if guardType.Roles().Contains(roleName) {
			thenBranch := []projection.Statement{}
			for _, s := range stmt.Scope(0).AllStmt() {
				thenBranch = append(thenBranch, epp.EppStmt(roleName, s)...)
			}

			elseBranch := []projection.Statement{}
			if elseScope := stmt.Scope(1); elseScope != nil {
				for _, s := range elseScope.AllStmt() {
					elseBranch = append(elseBranch, epp.EppStmt(roleName, s)...)
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
