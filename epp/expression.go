package epp

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"
)

func (epp *epp) eppExpression(roleName string, expr parser.IExprContext) (projection.Expression, []projection.Statement) {

	exprType := epp.info.Types[expr]
	exprValue := epp.eppType(roleName, exprType)

	switch expr := expr.(type) {
	case *parser.ExprBinOpContext:
		lhs, aux := epp.eppExpression(roleName, expr.Expr(0))
		rhs, rhsAux := epp.eppExpression(roleName, expr.Expr(1))
		aux = append(aux, rhsAux...)

		if exprType.Roles().Contains(roleName) {
			operator := projection.ParseOperator(expr)
			return projection.NewExprBinaryOp(operator, lhs, rhs, exprValue), aux
		} else {
			if lhs != nil && lhs.HasSideEffects() {
				aux = append(aux, projection.NewStmtExpr(lhs))
			}
			if rhs != nil && rhs.HasSideEffects() {
				aux = append(aux, projection.NewStmtExpr(rhs))
			}
			return nil, aux
		}
	case *parser.ExprBoolContext:
		if exprType.Roles().Contains(roleName) {
			value := expr.TRUE() != nil
			return projection.NewExprBool(value), []projection.Statement{}
		} else {
			return nil, []projection.Statement{}
		}
	case *parser.ExprGroupContext:
		return epp.eppExpression(roleName, expr.Expr())
	case *parser.ExprIdentContext:
		sym := epp.info.Symbols[expr.IdentAccess().Ident()]

		if exprType.Roles().Contains(roleName) {
			name := sym.SymbolName()
			switch sym := sym.(type) {
			case *sym_table.FuncSymbol:
				funcType := exprValue.(*projection.FunctionType)

				substMap, _ := funcType.Roles().SubstituteMap(sym.Roles())

				roleSubst := substMap.Subst(roleName)
				name += "_" + roleSubst
			}

			return projection.NewExprIdent(name, exprValue), []projection.Statement{}
		} else {
			return nil, []projection.Statement{}
		}
	case *parser.ExprNumContext:
		if exprType.Roles().Contains(roleName) {
			num, err := strconv.Atoi(expr.NUMBER().GetText())
			if err != nil {
				panic(fmt.Sprintf("expected NUMBER to be convertible to int: %s", expr.GetText()))
			}
			return projection.NewExprInt(num), []projection.Statement{}
		} else {
			return nil, []projection.Statement{}
		}
	case *parser.ExprStringContext:
		if exprType.Roles().Contains(roleName) {
			str := expr.STRING().GetText()
			str = str[1 : len(str)-1] // remote quotes

			escapes := map[string]string{
				`\\`: `\`,
				`\"`: `"`,
				`\n`: "\n",
				`\r`: "\r",
				`\t`: "\t",
			}

			for from, to := range escapes {
				str = strings.ReplaceAll(str, from, to)
			}

			return projection.NewExprString(str), []projection.Statement{}
		} else {
			return nil, []projection.Statement{}
		}
	case *parser.ExprAwaitContext:
		asyncExpr, aux := epp.eppExpression(roleName, expr.Expr())
		if asyncExpr != nil {
			if innerExprAsync, innerIsFixedAsync := asyncExpr.(*projection.ExprAsync); innerIsFixedAsync {
				// await fixed async cancels out
				return innerExprAsync.Inner(), aux
			} else {
				asyncType := asyncExpr.Type().(*projection.AsyncType)
				return projection.NewExprAwait(asyncExpr, asyncType.Inner), aux
			}
		}
		return nil, aux
	case *parser.ExprComContext:
		sender := expr.GetSender().(*parser.RoleTypeNormalContext)
		senderRole := sender.Ident(0).GetText()

		receivers := parser.RoleTypeAllIdents(expr.GetReceiver())
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
			innerType := epp.eppType(senderRole, epp.info.Types[expr.Expr()])
			return projection.NewExprRecv(innerType, senderRole), aux
		}

		valueType := epp.info.Types[expr.Expr()]
		if valueType.Roles().Contains(roleName) {
			return exprValue, aux
		}

		return nil, aux
	case *parser.ExprCallContext:
		callExpr, aux := epp.eppExpression(roleName, expr.Expr())
		callType := epp.info.Types[expr.Expr()]

		callFuncValue := epp.eppType(roleName, callType)

		switch callType.Value().(type) {
		case *types.FunctionType:
			callFuncValue, _ := callFuncValue.(*projection.FunctionType)
			funcType := callType.Value().(*types.FunctionType)

			argValues := []projection.Expression{}
			for i, arg := range expr.FuncArgList().AllExpr() {
				argVal, extra := epp.eppExpression(roleName, arg)
				aux = append(aux, extra...)

				funcParam := funcType.Params()[i]
				if funcParam.Roles().Contains(roleName) {
					paramType := callFuncValue.Params[len(argValues)]
					argStored := epp.storeExpression(roleName, argVal, paramType)
					argValues = append(argValues, argStored)
				} else if argVal != nil && argVal.HasSideEffects() {
					aux = append(aux, projection.NewStmtExpr(argVal))
				}
			}

			if callType.Roles().Contains(roleName) {
				funcSym := epp.info.Symbols[callFuncValue.NameIdent()].(*sym_table.FuncSymbol)
				returnValue := callFuncValue.ReturnType
				roleSubst, _ := funcSym.Roles().SubstituteMap(callFuncValue.Roles())
				return projection.NewExprCallFunc(callExpr, roleName, argValues, returnValue, roleSubst), aux
			} else {
				return nil, aux
			}
		case *types.ClosureType:
			callFuncValue, _ := callFuncValue.(*projection.ClosureType)
			closureType := callType.Value().(*types.ClosureType)

			argValues := []projection.Expression{}
			for i, arg := range expr.FuncArgList().AllExpr() {
				argVal, extra := epp.eppExpression(roleName, arg)
				aux = append(aux, extra...)

				closureParam := closureType.Params()[i]
				if closureParam.Roles().Contains(roleName) {
					paramType := callFuncValue.Params[len(argValues)]
					argStored := epp.storeExpression(roleName, argVal, paramType)
					argValues = append(argValues, argStored)
				} else if argVal != nil && argVal.HasSideEffects() {
					aux = append(aux, projection.NewStmtExpr(argVal))
				}
			}

			if callType.Roles().Contains(roleName) {
				returnValue := callFuncValue.ReturnType
				return projection.NewExprCallClosure(callExpr, roleName, argValues, returnValue), aux
			} else {
				return nil, aux
			}
		default:
			panic("unreachable")
		}

	case *parser.ExprStructContext:
		stSym := epp.info.Symbols[expr.Ident()].(*sym_table.StructSymbol)

		defRoleSubst, _ := stSym.Type().Roles().SubstituteMap(exprType.Roles())
		exprRoleSubst, _ := exprType.Roles().SubstituteMap(stSym.Type().Roles())

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
			} else if field != nil && field.HasSideEffects() {
				aux = append(aux, projection.NewStmtExpr(field))
			}
		}

		if exprType.Roles().Contains(roleName) {
			structType := exprValue.(*projection.StructType)
			return projection.NewExprStruct(stSym.SymbolName(), exprRoleSubst.Subst(roleName), fieldNames, fields, structType), aux
		} else {
			return nil, aux
		}
	case *parser.ExprFieldAccessContext:
		baseExpr, aux := epp.eppExpression(roleName, expr.Expr())

		if exprType.Roles().Contains(roleName) {
			fieldName := expr.Ident().GetText()
			return projection.NewExprFieldAccess(baseExpr, fieldName, exprValue), aux
		} else {
			if baseExpr != nil && baseExpr.HasSideEffects() {
				aux = append(aux, projection.NewStmtExpr(baseExpr))
			}
			return nil, aux
		}
	case *parser.ExprClosureContext:
		if !exprType.Roles().Contains(roleName) {
			return nil, []projection.Statement{}
		}

		closureType := exprType.Value().(*types.ClosureType)

		params := []projection.ClosureParam{}
		for i, param := range expr.ClosureSig().FuncParamList().AllFuncParam() {
			paramType := closureType.Params()[i]
			if paramType.Roles().Contains(roleName) {
				paramValue := epp.eppType(roleName, paramType)
				params = append(params, projection.NewClosureParam(param.Ident().GetText(), paramValue))
			}
		}

		returnType := epp.eppType(roleName, closureType.ReturnType())

		body := []projection.Statement{}
		for _, stmt := range expr.Scope().AllStmt() {
			eppStmts := epp.EppStmt(roleName, stmt)
			body = append(body, eppStmts...)
		}

		return projection.NewExprClosure(params, returnType, body), []projection.Statement{}
	case *parser.ExprContext:
		panic("expr should never be base type")
	}

	panic(fmt.Sprintf("unknown expression: %#v", expr))
}
