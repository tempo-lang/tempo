package epp

import (
	"fmt"
	"slices"
	"strconv"
	"tempo/parser"
	"tempo/projection"
	"tempo/sym_table"
	"tempo/type_check"
	"tempo/types"
)

func (epp *epp) eppExpression(roleName string, expr parser.IExprContext) (projection.Expression, []projection.Statement) {

	exprType := epp.info.Types[expr]

	switch expr := expr.(type) {
	case *parser.ExprBinOpContext:
		lhs, aux := epp.eppExpression(roleName, expr.Expr(0))
		rhs, rhsAux := epp.eppExpression(roleName, expr.Expr(1))
		aux = append(aux, rhsAux...)

		if exprType.Roles().Contains(roleName) {
			operator := projection.ParseOperator(expr)
			return projection.NewExprBinaryOp(operator, lhs, rhs, exprType.Value()), aux
		} else {
			return nil, aux
		}
	case *parser.ExprBoolContext:
		value := expr.TRUE() != nil
		return projection.NewExprBool(value), []projection.Statement{}
	case *parser.ExprGroupContext:
		return epp.eppExpression(roleName, expr.Expr())
	case *parser.ExprIdentContext:
		sym := epp.info.Symbols[expr.Ident()]

		if sym.Type().Roles().Contains(roleName) {
			appendRole := false
			switch sym.Type().Value().(type) {
			case *types.FunctionType:
				appendRole = true
			}

			name := sym.SymbolName()
			if appendRole {
				name += "_" + roleName
			}

			return projection.NewExprIdent(name, exprType.Value()), []projection.Statement{}
		} else {
			return nil, []projection.Statement{}
		}
	case *parser.ExprNumContext:
		num, err := strconv.Atoi(expr.GetText())
		if err != nil {
			panic(fmt.Sprintf("expected NUMBER to be convertible to int: %s", expr.GetText()))
		}
		return projection.NewExprInt(num), []projection.Statement{}
	case *parser.ExprAwaitContext:
		inner, aux := epp.eppExpression(roleName, expr.Expr())
		if inner != nil {
			if innerExprAsync, innerIsFixedAsync := inner.(*projection.ExprAsync); innerIsFixedAsync {
				// await fixed async cancels out
				return innerExprAsync.Inner(), aux
			} else {
				asyncType := epp.info.Types[expr.Expr()].Value().(*types.Async)
				return projection.NewExprAwait(inner, asyncType.Inner()), aux
			}
		} else {
			return nil, aux
		}
	case *parser.ExprComContext:
		sender := expr.RoleType(0).(*parser.RoleTypeNormalContext)
		senderRole := sender.Ident(0).GetText()

		receivers := parser.RoleTypeAllIdents(expr.RoleType(1))
		inner, aux := epp.eppExpression(roleName, expr.Expr())

		receiverRoles := []string{}
		for _, receiver := range receivers {
			receiverRoles = append(receiverRoles, receiver.GetText())
		}

		exprValue := projection.NewExprAsync(inner)
		if roleName == senderRole {
			if inner.HasSideEffects() {
				tmpName := epp.nextTmpName()
				aux = append(aux, projection.NewStmtVarDecl(tmpName, exprValue))
				exprValue = projection.NewExprIdent(tmpName, exprValue.Type())
				aux = append(aux, projection.NewStmtExpr(projection.NewExprSend(projection.NewExprAwait(exprValue, inner.Type()), receiverRoles)))
			} else {
				aux = append(aux, projection.NewStmtExpr(projection.NewExprSend(inner, receiverRoles)))
			}
		}

		// receiver
		if slices.Contains(receiverRoles, roleName) {
			innerType := epp.info.Types[expr.Expr()]
			return projection.NewExprRecv(innerType.Value(), senderRole), aux
		}

		valueType := epp.info.Types[expr.Expr()]
		if valueType.Roles().Contains(roleName) {
			return exprValue, aux
		}

		return nil, aux
	case *parser.ExprCallContext:
		funcSym := epp.info.Symbols[expr.Ident()].(*sym_table.FuncSymbol)

		funcRoles := parser.RoleTypeAllIdents(funcSym.FuncSig().RoleType())
		callRoles, _ := type_check.ParseRoleType(expr.RoleType())

		argRoleSubst := map[string]string{}
		paramRoleSubst := map[string]string{}
		for i, callRole := range callRoles.Participants() {
			funcRole := funcRoles[i].GetText()
			argRoleSubst[callRole] = funcRole
			paramRoleSubst[funcRole] = callRole
		}

		aux := []projection.Statement{}
		argValues := []projection.Expression{}

		for i, arg := range expr.FuncArgList().AllExpr() {
			argVal, extra := epp.eppExpression(roleName, arg)
			aux = append(aux, extra...)

			if funcSym.Params()[i].Type().Roles().Contains(argRoleSubst[roleName]) {
				argValues = append(argValues, argVal)
			}
		}

		if callRoles.Contains(roleName) {
			returnType := funcSym.FuncValue().ReturnType().SubstituteRoles(paramRoleSubst)
			returnValue := epp.eppType(roleName, returnType)

			roleSubs := []projection.ExprCallRoleSubs{}
			for _, callRole := range callRoles.Participants() {
				roleSubs = append(roleSubs, projection.ExprCallRoleSubs{
					From: callRole,
					To:   argRoleSubst[callRole],
				})
			}

			return projection.NewExprCall(funcSym.SymbolName(), argRoleSubst[roleName], argValues, returnValue, roleSubs), aux
		}

		return nil, aux
	case *parser.ExprStructContext:
		stSym := epp.info.Symbols[expr.Ident()].(*sym_table.StructSymbol)

		defRoleSubst := map[string]string{}
		exprRoleSubst := map[string]string{}
		for i, defRole := range stSym.Type().Roles().Participants() {
			exprRole := exprType.Roles().Participants()[i]
			defRoleSubst[defRole] = exprRole
			exprRoleSubst[exprRole] = defRole
		}

		aux := []projection.Statement{}

		fields := map[string]projection.Expression{}
		fieldNames := []string{}

		fieldNamesIdents := expr.ExprStructField().AllIdent()
		for i, fieldExpr := range expr.ExprStructField().AllExpr() {
			fieldName := fieldNamesIdents[i].GetText()

			field, a := epp.eppExpression(roleName, fieldExpr)
			aux = append(aux, a...)

			containsRole := stSym.Fields()[i].Type().Roles().
				SubstituteRoles(defRoleSubst).Contains(roleName)

			if containsRole {
				fields[fieldName] = field
				fieldNames = append(fieldNames, fieldName)
			}
		}

		structType := projection.NewStructType(exprType.Value().(*types.StructType), exprRoleSubst[roleName])

		return projection.NewExprStruct(stSym.SymbolName(), exprRoleSubst[roleName], fieldNames, fields, structType), aux
	case *parser.ExprFieldAccessContext:
		baseExpr, aux := epp.eppExpression(roleName, expr.Expr())
		fieldName := expr.Ident().GetText()

		return projection.NewExprFieldAccess(baseExpr, fieldName, exprType.Value()), aux
	case *parser.ExprContext:
		panic("expr should never be base type")
	}

	panic(fmt.Sprintf("unknown expression: %#v", expr))
}
