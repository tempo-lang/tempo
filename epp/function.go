package epp

import (
	"chorego/parser"
	"chorego/projection"
	"chorego/types"
	"fmt"
)

func EppFunc(function parser.IFuncContext) *projection.Choreography {
	func_role := function.RoleTypeNormal()

	choreography := projection.NewChoreography(function.Ident().GetText())

	for _, role := range func_role.AllIdent() {
		eppFuncRole(choreography, function, role)
	}

	return choreography
}

func assertValidTree(err types.Error) {
	if err != nil {
		panic(fmt.Sprintf("expected valid syntax tree: %v", err))
	}
}

func eppFuncRole(choreography *projection.Choreography, function parser.IFuncContext, role parser.IIdentContext) {
	roleName := role.ID().GetText()
	fn := choreography.AddFunc(roleName, function)

	// project parameters
	for _, param := range function.FuncParamList().AllFuncParam() {
		if ValueExistsAtRole(param.ValueType(), roleName) {
			valType, err := types.ParseValueType(param.ValueType())
			assertValidTree(err)
			fn.AddParam(param, valType)
		}
	}

	// project body
	for _, stmt := range function.Scope().AllStatement() {
		if varDecl := stmt.StmtVarDecl(); varDecl != nil {
			if ValueExistsAtRole(varDecl.ValueType(), roleName) {
				variableName := varDecl.Ident().GetText()
				varibleType, err := types.ParseValueType(varDecl.ValueType())
				assertValidTree(err)

				expr := eppExpression(role, varDecl.Expression())

				fn.AddStmt(projection.NewStmtVarDecl(variableName, varibleType, expr))
			}
		}
	}
}
