package epp

import (
	"chorego/parser"
	"chorego/projection"
	"chorego/type_check"
	"chorego/types"
	"fmt"
)

func EppStmt(info *type_check.Info, roleName string, stmt parser.IStmtContext) projection.Statement {
	switch stmt := stmt.(type) {
	case *parser.StmtAssignContext:
		sym := info.Symbols[stmt.Ident()]
		if sym.Type().Roles().Contains(roleName) {
			varName := stmt.Ident().GetText()
			expr := eppExpression(info, roleName, stmt.Expr())
			return (projection.NewStmtAssign(varName, expr))
		}
	case *parser.StmtVarDeclContext:
		varSym := info.Symbols[stmt.Ident()]
		if varSym.Type().Roles().Contains(roleName) {
			variableName := stmt.Ident().GetText()
			varibleType, err := types.ParseValueType(stmt.ValueType())
			assertValidTree(err)

			expr := eppExpression(info, roleName, stmt.Expr())
			if _, isAsync := varibleType.Value().(*types.Async); isAsync {
				expr = projection.NewExprAsync(expr)
			}

			return (projection.NewStmtVarDecl(variableName, varibleType, expr))
		}
	case *parser.StmtContext:
		panic("statement should never be base type")
	default:
		panic(fmt.Sprintf("unknown statement: %#v", stmt))
	}

	return nil
}
