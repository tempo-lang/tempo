package epp

import (
	"chorego/parser"
	"chorego/projection"
	"chorego/type_check"
	"chorego/types"
	"fmt"
)

func EppFunc(info *type_check.Info, function parser.IFuncContext) *projection.Choreography {
	sym := info.Symbols[function.Ident()]
	func_role := sym.Type().Roles()

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

	funcSym := info.Symbols[function.Ident()]
	funcType := funcSym.Type().Value().(*types.FunctionType)

	// project parameters
	for i, param := range function.FuncParamList().AllFuncParam() {
		paramType := funcType.Params()[i]
		if paramType.Roles().Contains(roleName) {
			fn.AddParam(param, paramType)
		}
	}

	// project body
	for _, stmt := range function.Scope().AllStmt() {
		switch stmt := stmt.(type) {
		case *parser.StmtAssignContext:
			sym := info.Symbols[stmt.Ident()]
			if sym.Type().Roles().Contains(roleName) {
				varName := stmt.Ident().GetText()
				expr := eppExpression(stmt.Expr())
				fn.AddStmt(projection.NewStmtAssign(varName, expr))
			}
		case *parser.StmtVarDeclContext:
			varSym := info.Symbols[stmt.Ident()]
			if varSym.Type().Roles().Contains(roleName) {
				variableName := stmt.Ident().GetText()
				varibleType, err := types.ParseValueType(stmt.ValueType())
				assertValidTree(err)

				expr := eppExpression(stmt.Expr())

				fn.AddStmt(projection.NewStmtVarDecl(variableName, varibleType, expr))
			}
		case *parser.StmtContext:
			panic("statement should never be base type")
		default:
			panic(fmt.Sprintf("unknown statement: %#v", stmt))
		}
	}
}
