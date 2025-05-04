package epp

import (
	"chorego/parser"
	"chorego/projection"
	"chorego/type_check"
	"chorego/types"
	"fmt"
)

func EppFunc(info *type_check.Info, function parser.IFuncContext) *projection.Choreography {
	obj := info.Symbols[function.Ident()]

	func_role := obj.Type().Roles()

	choreography := projection.NewChoreography(function.Ident().GetText())

	for _, role := range func_role.Participants() {
		eppFuncRole(info, choreography, function, role)
	}

	return choreography
}

func assertValidTree(err types.Error) {
	if err != nil {
		panic(fmt.Sprintf("expected valid syntax tree: %v", err))
	}
}

func eppFuncRole(info *type_check.Info, choreography *projection.Choreography, function parser.IFuncContext, roleName string) {
	fn := choreography.AddFunc(roleName, function)

	obj := info.Symbols[function.Ident()]
	funcType := obj.Type().Value().(*types.FunctionType)

	// project parameters
	for i, param := range function.FuncParamList().AllFuncParam() {
		paramType := funcType.Params()[i]
		if paramType.Roles().Contains(types.SingleRole(roleName)) {
			fn.AddParam(param, paramType)
		}
	}

	// project body
	for _, stmt := range function.Scope().AllStatement() {
		if varDecl := stmt.StmtVarDecl(); varDecl != nil {
			if ValueExistsAtRole(varDecl.ValueType(), roleName) {
				variableName := varDecl.Ident().GetText()
				varibleType, err := types.ParseValueType(varDecl.ValueType())
				assertValidTree(err)

				expr := eppExpression(varDecl.Expression())

				fn.AddStmt(projection.NewStmtVarDecl(variableName, varibleType, expr))
			}
		}
	}
}
