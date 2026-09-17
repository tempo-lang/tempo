package epp

import (
	"fmt"
	"slices"

	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"
)

func (epp *epp) eppExpression(roleName string, expr ast.Expr) (projection.Expression, []projection.Statement) {

	exprType := epp.info.Types[expr]

	switch expr := expr.(type) {
	case *ast.BinaryExpr:
		return epp.eppExprBinOp(roleName, expr, exprType)
	case *ast.PrimitiveExpr:
		return epp.eppExprPrimitive(roleName, expr, exprType)
	case *ast.GroupExpr:
		return epp.eppExpression(roleName, expr.Expr)
	case *ast.IdentAccessExpr:
		return epp.eppExprIdent(roleName, expr, exprType)
	case *ast.AwaitExpr:
		return epp.eppExprAwait(roleName, expr)
	case *ast.ComExpr:
		return epp.eppExprCom(roleName, expr)
	case *ast.CallExpr:
		return epp.eppExprCall(roleName, expr, exprType)
	case *ast.StructExpr:
		return epp.eppExprStruct(roleName, expr, exprType)
	case *ast.FieldAccessExpr:
		return epp.eppExprFieldAccess(roleName, expr, exprType)
	case *ast.ClosureExpr:
		return epp.eppExprClosure(roleName, expr, exprType)
	case *ast.ListExpr:
		return epp.eppExprList(roleName, expr, exprType)
	case *ast.IndexExpr:
		return epp.eppExprIndex(roleName, expr, exprType)
	case *ast.InvalidExpr:
		panic("endpoint projection encountered an invalid expression")
	}

	panic(fmt.Sprintf("unknown expression: %#v", expr))
}

func (epp *epp) eppExprBinOp(roleName string, expr *ast.BinaryExpr, exprType types.Type) (projection.Expression, []projection.Statement) {
	lhs, aux := epp.eppExpression(roleName, expr.Left)
	rhs, rhsAux := epp.eppExpression(roleName, expr.Right)
	aux = append(aux, rhsAux...)

	if exprType.Roles().Contains(roleName) {
		exprValue := epp.eppType(roleName, exprType)
		operator := projection.Operator(expr.Operator.Text)
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
}

func (epp *epp) eppExprPrimitive(roleName string, expr *ast.PrimitiveExpr, exprType types.Type) (projection.Expression, []projection.Statement) {
	if exprType.Roles().Contains(roleName) {
		switch lit := expr.Literal.(type) {
		case *ast.BoolLit:
			return projection.NewExprBool(lit.Value()), nil
		case *ast.FloatLit:
			return projection.NewExprFloat(lit.Value()), nil
		case *ast.IntLit:
			return projection.NewExprInt(lit.Value()), nil
		case *ast.StringLit:
			return projection.NewExprString(lit.Value()), nil
		}

		panic(fmt.Sprintf("unknown literal: %#v", expr.Literal))
	} else {
		return nil, []projection.Statement{}
	}
}

func (epp *epp) eppExprIdent(roleName string, expr *ast.IdentAccessExpr, exprType types.Type) (projection.Expression, []projection.Statement) {
	sym := epp.info.Symbols[expr.Ident]

	if exprType.Roles().Contains(roleName) {
		name := sym.SymbolName()

		exprValue := epp.eppType(roleName, exprType)

		// check if identifier is struct attribute
		structScope := epp.info.GlobalScope.Innermost(expr.StartToken().Span.Start).GetStruct()
		if structScope != nil && sym.Parent() == structScope.Scope() {
			selfExpr := projection.NewExprSelf(epp.eppType(roleName, structScope.Type()))
			return projection.NewExprFieldAccess(selfExpr, name, exprValue), []projection.Statement{}
		}

		switch sym := sym.(type) {
		case *sym_table.FuncSymbol:
			funcType := exprValue.(*projection.FunctionType)

			if !sym.Roles().IsUnnamedRole() {
				substMap, ok := funcType.Roles().SubstituteMap(sym.Roles())
				if !ok {
					panic("type check ensures substitution is valid")
				}

				roleSubst := substMap.Subst(roleName)[0]
				name += "_" + roleSubst
			}
		}

		return projection.NewExprIdent(name, exprValue), []projection.Statement{}
	} else {
		return nil, []projection.Statement{}
	}
}

func (epp *epp) eppExprAwait(roleName string, expr *ast.AwaitExpr) (projection.Expression, []projection.Statement) {
	asyncExpr, aux := epp.eppExpression(roleName, expr.Expr)
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
}

func (epp *epp) eppExprCom(roleName string, expr *ast.ComExpr) (projection.Expression, []projection.Statement) {
	senderRole := expr.Sender.Roles()[0].Token.Text
	inner, aux := epp.eppExpression(roleName, expr.Expr)

	receiverRoles := []string{}
	for _, receiver := range expr.Receiver.Roles() {
		receiverRoles = append(receiverRoles, receiver.Token.Text)
	}

	isReceiver := slices.Contains(receiverRoles, roleName)

	exprValue := projection.NewExprAsync(inner)
	if roleName == senderRole {
		if !isReceiver {
			exprValue = projection.NewExprSend(inner, receiverRoles)
		} else {
			aux = append(aux, projection.NewStmtExpr(projection.NewExprSend(inner, receiverRoles)))
		}
	}

	if isReceiver {
		innerType := epp.eppType(senderRole, epp.info.Types[expr.Expr])
		return projection.NewExprRecv(innerType, senderRole), aux
	}

	valueType := epp.info.Types[expr.Expr]
	if valueType.Roles().Contains(roleName) {
		return exprValue, aux
	}

	// roles that are not a part of the com
	return nil, aux
}

func (epp *epp) eppExprCall(roleName string, expr *ast.CallExpr, exprType types.Type) (projection.Expression, []projection.Statement) {
	// Special case for type casting
	if exprIdent, isIdent := expr.Function.(*ast.IdentAccessExpr); isIdent {
		if sym, ok := epp.info.Symbols[exprIdent.Ident]; ok {
			if _, ok := sym.(*sym_table.TypeSymbol); ok {
				arg := expr.Args[0]
				inner, aux := epp.eppExpression(roleName, arg)

				if exprType.Roles().Contains(roleName) {
					castedType := epp.eppType(roleName, exprType)
					return projection.NewExprTypeCast(inner, castedType), aux
				} else {
					return nil, aux
				}
			}
		}
	}

	callExpr, aux := epp.eppExpression(roleName, expr.Function)
	callType := epp.info.Types[expr.Function]

	callFuncValue := epp.eppType(roleName, callType)

	switch callType.(type) {
	case *types.FunctionType:
		callFuncValue, _ := callFuncValue.(*projection.FunctionType)
		funcType := callType.(*types.FunctionType)

		argValues := []projection.Expression{}
		for i, arg := range expr.Args {
			argVal, extra := epp.eppExpression(roleName, arg)
			aux = append(aux, extra...)

			funcParam := funcType.Params()[i]
			if funcParam.Roles().Contains(roleName) {
				paramType := callFuncValue.Params()[len(argValues)]
				argStored := epp.storeExpression(roleName, argVal, paramType)
				argValues = append(argValues, argStored)
			} else if argVal != nil && argVal.HasSideEffects() {
				aux = append(aux, projection.NewStmtExpr(argVal))
			}
		}

		if callType.Roles().Contains(roleName) {
			funcSym := epp.info.Symbols[callFuncValue.NameIdent()].(*sym_table.FuncSymbol)
			returnValue := callFuncValue.ReturnType()

			callExpr := projection.NewExprCallFunc(callExpr, roleName, argValues, returnValue, funcSym.Roles(), callType.Roles())
			if returnValue == projection.UnitType() {
				aux = append(aux, projection.NewStmtExpr(callExpr))
				return nil, aux
			} else {
				return callExpr, aux
			}
		} else {
			return nil, aux
		}
	case *types.ClosureType:
		callFuncValue, _ := callFuncValue.(*projection.ClosureType)
		closureType := callType.(*types.ClosureType)

		argValues := []projection.Expression{}
		for i, arg := range expr.Args {
			argVal, extra := epp.eppExpression(roleName, arg)
			aux = append(aux, extra...)

			closureParam := closureType.Params()[i]
			if closureParam.Roles().Contains(roleName) {
				paramType := callFuncValue.Params()[len(argValues)]
				argStored := epp.storeExpression(roleName, argVal, paramType)
				argValues = append(argValues, argStored)
			} else if argVal != nil && argVal.HasSideEffects() {
				aux = append(aux, projection.NewStmtExpr(argVal))
			}
		}

		if callType.Roles().Contains(roleName) {
			returnValue := callFuncValue.ReturnType()
			callExpr := projection.NewExprCallClosure(callExpr, roleName, argValues, returnValue)
			if returnValue == projection.UnitType() {
				aux = append(aux, projection.NewStmtExpr(callExpr))
				return nil, aux
			} else {
				return callExpr, aux
			}
		} else {
			return nil, aux
		}
	default:
		panic("unreachable")
	}
}

func (epp *epp) eppExprStruct(roleName string, expr *ast.StructExpr, exprType types.Type) (projection.Expression, []projection.Statement) {
	stSym := epp.info.Symbols[expr.RoleIdent.Ident].(*sym_table.StructSymbol)
	defRoleSubst, _ := stSym.Type().Roles().SubstituteMap(exprType.Roles())

	aux := []projection.Statement{}
	fields := map[string]projection.Expression{}
	fieldNames := []string{}

	for _, init := range expr.StructFields.Fields {
		fieldName, fieldExpr := init.Name.Value(), init.Expr
		symField, ok := stSym.Field(fieldName)
		if !ok {
			panic("assuming field exists when expr is well-typed")
		}

		field, a := epp.eppExpression(roleName, fieldExpr)
		aux = append(aux, a...)

		containsRole := symField.Type().Roles().SubstituteRoles(defRoleSubst).Contains(roleName)
		if containsRole {
			fields[fieldName] = field
			fieldNames = append(fieldNames, fieldName)
		} else if field != nil && field.HasSideEffects() {
			aux = append(aux, projection.NewStmtExpr(field))
		}
	}

	if exprType.Roles().Contains(roleName) {
		structType := epp.eppType(roleName, exprType).(*projection.StructType)
		return projection.NewExprStruct(structType, fieldNames, fields), aux
	} else {
		return nil, aux
	}
}

func (epp *epp) eppExprFieldAccess(roleName string, expr *ast.FieldAccessExpr, exprType types.Type) (projection.Expression, []projection.Statement) {
	baseExpr, aux := epp.eppExpression(roleName, expr.Object)

	if exprType.Roles().Contains(roleName) {
		exprValue := epp.eppType(roleName, exprType)
		fieldName := expr.Field.Value()
		fieldExpr := epp.eppField(baseExpr, fieldName, exprValue)
		return fieldExpr, aux
	} else {
		if baseExpr != nil && baseExpr.HasSideEffects() {
			aux = append(aux, projection.NewStmtExpr(baseExpr))
		}
		return nil, aux
	}
}

func (epp *epp) eppExprClosure(roleName string, expr *ast.ClosureExpr, exprType types.Type) (projection.Expression, []projection.Statement) {
	if !exprType.Roles().Contains(roleName) {
		return nil, []projection.Statement{}
	}

	closureType := exprType.(*types.ClosureType)

	params := []projection.ClosureParam{}
	for i, param := range expr.ClosureSig.Params.Params {
		paramType := closureType.Params()[i]
		if paramType.Roles().Contains(roleName) {
			paramValue := epp.eppType(roleName, paramType)
			params = append(params, projection.NewClosureParam(param.Name.Value(), paramValue))
		}
	}

	returnType := epp.eppType(roleName, closureType.ReturnType())

	body := []projection.Statement{}
	for _, stmt := range expr.Scope.Stmts {
		eppStmts := epp.EppStmt(roleName, stmt)
		body = append(body, eppStmts...)
	}

	return projection.NewExprClosure(params, returnType, body), []projection.Statement{}
}

func (epp *epp) eppExprList(roleName string, expr *ast.ListExpr, exprType types.Type) (projection.Expression, []projection.Statement) {
	items := []projection.Expression{}
	aux := []projection.Statement{}
	for _, item := range expr.Elements {
		itemExpr, ax := epp.eppExpression(roleName, item)
		aux = append(aux, ax...)
		items = append(items, itemExpr)
	}

	if exprType.Roles().Contains(roleName) {
		exprValue := epp.eppType(roleName, exprType)
		return projection.NewExprList(items, exprValue), aux
	} else {
		for _, item := range items {
			if item != nil && item.HasSideEffects() {
				aux = append(aux, projection.NewStmtExpr(item))
			}
		}
		return nil, aux
	}
}

func (epp *epp) eppExprIndex(roleName string, expr *ast.IndexExpr, exprType types.Type) (projection.Expression, []projection.Statement) {
	baseExpr, aux := epp.eppExpression(roleName, expr.Object)

	indexExpr, a := epp.eppExpression(roleName, expr.Index)
	aux = append(aux, a...)

	if exprType.Roles().Contains(roleName) {
		return projection.NewExprIndex(baseExpr, indexExpr), aux
	} else {
		if baseExpr != nil && baseExpr.HasSideEffects() {
			aux = append(aux, projection.NewStmtExpr(baseExpr))
		}
		if indexExpr != nil && indexExpr.HasSideEffects() {
			aux = append(aux, projection.NewStmtExpr(indexExpr))
		}
		return nil, aux
	}
}
