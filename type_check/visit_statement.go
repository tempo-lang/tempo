package type_check

import (
	"github.com/tempo-lang/tempo/parser/ast"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

func (tc *typeChecker) visitStmt(stmt ast.Stmt) bool {
	switch s := stmt.(type) {
	case *ast.InvalidStmt:
		return false
	case *ast.LetStmt:
		tc.visitLet(s)
	case *ast.AssignStmt:
		tc.visitAssign(s)
	case *ast.ExprStmt:
		tc.visitExpr(s.Expr)
	case *ast.IfStmt:
		return tc.visitIf(s)
	case *ast.WhileStmt:
		tc.visitWhile(s)
	case *ast.ReturnStmt:
		return tc.visitReturn(s)
	}
	return false
}

func (tc *typeChecker) visitLet(s *ast.LetStmt) {
	var declared types.Type
	failed := false
	invalidDeclaration := false
	previous := tc.currentTypeHint
	if s.Type != nil {
		declared = tc.visitValueType(s.Type)
		failed = declared.IsInvalid()
		if rt, ok := findRoleType(s.Type); ok && !tc.checkRolesInScope(rt) {
			failed = true
		}
		invalidDeclaration = failed
		if !failed {
			tc.currentTypeHint = declared
		}
	}
	exprType := tc.visitExpr(s.Expr)
	tc.currentTypeHint = previous
	if f, ok := exprType.(*types.FunctionType); ok {
		exprType = f.ToClosure()
	}
	stmtType := tc.coerceExprToScope(s.Expr, exprType)
	if s.Type != nil {
		if !failed {
			if t, ok := stmtType.CoerceTo(declared); ok {
				stmtType = t
			} else {
				tc.reportError(type_error.NewInvalidAssignType(s.Expr, declared, stmtType))
				failed = true
			}
		}
		if failed && !invalidDeclaration {
			stmtType = types.Invalid()
		}
	} else if _, unit := stmtType.(*types.UnitType); unit {
		tc.reportError(type_error.NewAssignUnitValue(s.Expr))
		stmtType = types.Invalid()
	}
	if len(stmtType.Roles().Participants()) == 0 {
		stmtType = stmtType.ReplaceSharedRoles(tc.currentScope.Roles().Participants())
	}
	tc.insertSymbol(sym_table.NewVariableSymbol(s, tc.currentScope, stmtType))
}

func (tc *typeChecker) visitIf(s *ast.IfStmt) bool {
	guard := tc.visitExpr(s.Condition)
	if _, ok := guard.CoerceTo(types.Bool(guard.Roles().Participants())); !ok {
		tc.reportError(type_error.NewInvalidValue(s.Condition, guard, types.Bool(guard.Roles().Participants())))
	}
	guard = tc.coerceExprToScope(s.Condition, guard)
	tc.currentScope = tc.currentScope.MakeChild(nodeSpan(s.ThenScope), guard.Roles().Participants())
	thenReturn := tc.visitScope(s.ThenScope)
	tc.currentScope = tc.currentScope.Parent()
	if s.ElseScope == nil {
		return false
	}
	tc.currentScope = tc.currentScope.MakeChild(nodeSpan(s.ElseScope), guard.Roles().Participants())
	elseReturn := tc.visitScope(s.ElseScope)
	tc.currentScope = tc.currentScope.Parent()
	return thenReturn && elseReturn
}

func (tc *typeChecker) visitWhile(s *ast.WhileStmt) {
	cond := tc.visitExpr(s.Condition)
	if types.BuiltinKind(cond) != types.BuiltinBool {
		tc.reportError(type_error.NewInvalidValue(s.Condition, cond, types.Bool(nil)))
	}
	cond = tc.coerceExprToScope(s.Condition, cond)
	tc.currentScope = tc.currentScope.MakeChild(nodeSpan(s.Scope), cond.Roles().Participants())
	tc.visitScope(s.Scope)
	tc.currentScope = tc.currentScope.Parent()
}

func (tc *typeChecker) visitReturn(s *ast.ReturnStmt) bool {
	env := tc.currentScope.GetCallableEnv()
	expected := env.ReturnType()
	if missing := env.Scope().Roles().SubtractParticipants(tc.currentScope.Roles().Participants()); len(missing) > 0 {
		tc.reportError(type_error.NewReturnNotAllRoles(s, missing))
	}
	if s.Expr == nil {
		if expected != types.Unit() {
			tc.reportError(type_error.NewReturnValueMissing(env, s))
		}
		return true
	}
	actual := tc.visitExpr(s.Expr)
	if _, ok := actual.CoerceTo(expected); !ok {
		tc.reportError(type_error.NewIncompatibleTypes(s.Expr, actual, expected))
	}
	return true
}

func (tc *typeChecker) visitAssign(s *ast.AssignStmt) {
	root, ok := assignmentRoot(s.LHS)
	if !ok {
		return
	}
	sym, err := tc.lookupSymbol(root)
	if err != nil {
		tc.reportError(err)
		return
	}
	tc.info.Symbols[root] = sym
	sym.AddWrite(root)
	if !sym.IsAssignable() {
		tc.reportError(type_error.NewUnassignableSymbol(s, sym.Type()))
		return
	}
	target := tc.assignmentType(s.LHS, sym.Type())
	for _, indexExpr := range assignmentIndexes(s.LHS) {
		index := tc.visitExpr(indexExpr)
		expected := types.Int(target.Roles().Participants())
		if _, ok := index.CoerceTo(expected); !ok {
			tc.reportError(type_error.NewInvalidValue(indexExpr, index, expected))
		}
	}
	tc.checkExprInScope(s.LHS, target.Roles())
	value := tc.visitExpr(s.RHS)
	if _, ok := value.CoerceTo(target); !ok {
		tc.reportError(type_error.NewInvalidAssignType(s.RHS, target, value))
	}
}

func assignmentIndexes(expr ast.Expr) (result []ast.Expr) {
	switch e := expr.(type) {
	case *ast.FieldAccessExpr:
		return assignmentIndexes(e.Object)
	case *ast.IndexExpr:
		result = assignmentIndexes(e.Object)
		return append(result, e.Index)
	}
	return nil
}

func assignmentRoot(expr ast.Expr) (*ast.Identifier, bool) {
	switch e := expr.(type) {
	case *ast.IdentAccessExpr:
		return e.Ident, true
	case *ast.FieldAccessExpr:
		return assignmentRoot(e.Object)
	case *ast.IndexExpr:
		return assignmentRoot(e.Object)
	default:
		return nil, false
	}
}

func (tc *typeChecker) assignmentType(expr ast.Expr, current types.Type) types.Type {
	switch e := expr.(type) {
	case *ast.IdentAccessExpr:
		return current
	case *ast.FieldAccessExpr:
		base := tc.assignmentType(e.Object, current)
		st, ok := base.(*types.StructType)
		if !ok {
			tc.reportError(type_error.NewFieldAccessUnknownField(e.Field, base))
			return types.Invalid()
		}
		field, ok := tc.info.Field(st, e.Field.Value())
		if !ok {
			tc.reportError(type_error.NewFieldAccessUnknownField(e.Field, base))
			return types.Invalid()
		}
		return field
	case *ast.IndexExpr:
		base := tc.assignmentType(e.Object, current)
		list, ok := base.(*types.ListType)
		if !ok {
			tc.reportError(type_error.NewIndexWrongBaseType(e.Object, base))
			return types.Invalid()
		}
		return list.Inner()
	default:
		return types.Invalid()
	}
}
