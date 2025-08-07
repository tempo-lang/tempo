package epp

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"
)

func (epp *epp) EppStmt(roleName string, stmt parser.IStmtContext) (result []projection.Statement) {
	result = []projection.Statement{}

	switch stmt := stmt.(type) {
	case *parser.StmtAssignContext:
		sym := epp.info.Symbols[stmt.Ident()]
		assignType := sym.Type()
		specifiers := []projection.AssignSpecifier{}
		for _, specifier := range stmt.AllAssignSpecifier() {
			switch specifier := specifier.(type) {
			case *parser.AssignFieldContext:
				specifiers = append(specifiers, projection.AssignSpecifier{
					Kind:      projection.AssignField,
					FieldName: specifier.Ident().GetText(),
				})
				assignType, _ = epp.info.Field(assignType, specifier.Ident().GetText())
			case *parser.AssignIndexContext:
				indexExpr, aux := epp.eppExpression(roleName, specifier.Expr())
				result = append(result, aux...)
				specifiers = append(specifiers, projection.AssignSpecifier{
					Kind:      projection.AssignIndex,
					IndexExpr: indexExpr,
				})
				assignType = assignType.(*types.ListType).Inner()
			default:
				panic(fmt.Sprintf("unexpected parser.IAssignSpecifierContext: %#v", specifier))
			}
		}

		expr, aux := epp.eppExpression(roleName, stmt.Expr())
		result = append(result, aux...)

		if assignType.Roles().Contains(roleName) {

			varibleType := epp.eppType(roleName, assignType)
			expr = epp.storeExpression(roleName, expr, varibleType)

			varName := stmt.Ident().GetText()

			result = append(result, projection.NewStmtAssign(varName, specifiers, expr))
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
			varibleType := epp.eppType(roleName, varSym.Type())

			expr = epp.storeExpression(roleName, expr, varibleType)

			result = append(result, projection.NewStmtVarDecl(variableName, expr, varibleType, true))
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
	case *parser.StmtWhileContext:
		cond, aux := epp.eppExpression(roleName, stmt.Expr())
		result = aux

		condType := epp.info.Types[stmt.Expr()]

		if condType.Roles().Contains(roleName) {
			stmts := []projection.Statement{}
			for _, s := range stmt.Scope().AllStmt() {
				stmts = append(stmts, epp.EppStmt(roleName, s)...)
			}

			result = append(result, projection.NewStmtWhile(cond, stmts))
		}
	case *parser.StmtExprContext:
		expr, aux := epp.eppExpression(roleName, stmt.Expr())
		result = aux
		if expr != nil && expr.HasSideEffects() {
			result = append(result, projection.NewStmtExpr(expr))
		}
	case *parser.StmtReturnContext:
		var expr projection.Expression = nil
		if stmt.Expr() != nil {
			expr, result = epp.eppExpression(roleName, stmt.Expr())
		}

		scope := epp.info.GlobalScope.Innermost(stmt.GetStart())
		funcReturnRoles := scope.GetCallableEnv().ReturnType().Roles()
		if funcReturnRoles.Contains(roleName) {
			result = append(result, projection.NewStmtReturn(expr))
		} else if expr != nil && expr.HasSideEffects() {
			result = append(result, projection.NewStmtExpr(expr))
		}
	case *parser.StmtContext:
		panic("statement should never be base type")
	default:
		panic(fmt.Sprintf("unknown statement: %#v", stmt))
	}

	return
}

func (epp *epp) storeExpression(roleName string, expr projection.Expression, storeType projection.Type) projection.Expression {
	if _, ok := expr.Type().(*projection.FunctionType); ok {
		expr = epp.convertFuncToClosure(roleName, expr)
	}

	_, exprIsAsync := expr.Type().(*projection.AsyncType)
	if _, isAsync := storeType.(*projection.AsyncType); isAsync && !exprIsAsync {
		expr = projection.NewExprAsync(expr)
	}

	return projection.NewExprPassValue(expr)
}

func (epp *epp) convertFuncToClosure(roleName string, funcExpr projection.Expression) projection.Expression {
	funcType, ok := funcExpr.Type().(*projection.FunctionType)
	if !ok {
		panic(fmt.Sprintf("can not convert non-function to closure: %#v", funcExpr.Type()))
	}

	funcSym := epp.info.Symbols[funcType.NameIdent()].(*sym_table.FuncSymbol)
	closureParams := []projection.ClosureParam{}

	paramRoleSubst, _ := funcType.Roles().SubstituteMap(funcSym.Roles())
	argRoleSubst := paramRoleSubst.Inverse()

	paramRole := paramRoleSubst.Subst(roleName)

	for _, param := range funcSym.Params() {
		if param.Type().Roles().Contains(paramRole) {
			paramType := funcType.Params[len(closureParams)]
			closureParams = append(closureParams, projection.NewClosureParam(param.SymbolName(), paramType))
		}
	}

	argValues := []projection.Expression{}
	for _, param := range closureParams {
		argValues = append(argValues, projection.NewExprIdent(param.Name, param.Type))
	}

	callExpr := projection.NewExprCallFunc(funcExpr, roleName, argValues, funcType.ReturnType, argRoleSubst)
	var callStmt projection.Statement
	if funcType.ReturnType != projection.UnitType() {
		callStmt = projection.NewStmtReturn(callExpr)
	} else {
		callStmt = projection.NewStmtExpr(callExpr)
	}

	body := []projection.Statement{
		callStmt,
	}

	return projection.NewExprClosure(closureParams, funcType.ReturnType, body)
}
