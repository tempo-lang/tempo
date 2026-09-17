package epp

import (
	"fmt"
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"
)

func (epp *epp) EppStmt(role string, stmt ast.Stmt) (out []projection.Statement) {
	switch s := stmt.(type) {
	case *ast.InvalidStmt:
		panic("endpoint projection encountered an invalid statement")
	case *ast.LetStmt:
		sym := epp.info.Symbols[s.Name]
		expr, aux := epp.eppExpression(role, s.Expr)
		out = aux
		if sym.Type().Roles().Contains(role) {
			t := epp.eppType(role, sym.Type())
			out = append(out, projection.NewStmtVarDecl(s.Name.Value(), epp.storeExpression(role, expr, t), t, true))
		} else if expr != nil && expr.HasSideEffects() {
			out = append(out, projection.NewStmtExpr(expr))
		}
	case *ast.ExprStmt:
		expr, aux := epp.eppExpression(role, s.Expr)
		out = aux
		if expr != nil && expr.HasSideEffects() {
			out = append(out, projection.NewStmtExpr(expr))
		}
	case *ast.IfStmt:
		guard, aux := epp.eppExpression(role, s.Condition)
		out = aux
		if epp.info.Types[s.Condition].Roles().Contains(role) {
			out = append(out, projection.NewStmtIf(guard, epp.eppScope(role, s.ThenScope), epp.eppScope(role, s.ElseScope)))
		}
	case *ast.WhileStmt:
		cond, aux := epp.eppExpression(role, s.Condition)
		out = aux
		if epp.info.Types[s.Condition].Roles().Contains(role) {
			out = append(out, projection.NewStmtWhile(cond, epp.eppScope(role, s.Scope)))
		}
	case *ast.ReturnStmt:
		var expr projection.Expression
		if s.Expr != nil {
			expr, out = epp.eppExpression(role, s.Expr)
		}
		scope := epp.info.GlobalScope.Innermost(s.StartToken().Span.Start)
		if scope.GetCallableEnv().ReturnType().Roles().Contains(role) {
			out = append(out, projection.NewStmtReturn(expr))
		} else if expr != nil && expr.HasSideEffects() {
			out = append(out, projection.NewStmtExpr(expr))
		}
	case *ast.AssignStmt:
		return epp.eppAssign(role, s)
	default:
		panic(fmt.Sprintf("unknown statement: %#v", stmt))
	}
	return out
}

func (epp *epp) eppScope(role string, scope *ast.Scope) (out []projection.Statement) {
	if scope == nil {
		return nil
	}
	for _, s := range scope.Stmts {
		out = append(out, epp.EppStmt(role, s)...)
	}
	return
}

func assignmentParts(expr ast.Expr) (*ast.Identifier, []ast.Expr) {
	switch e := expr.(type) {
	case *ast.IdentAccessExpr:
		return e.Ident, nil
	case *ast.FieldAccessExpr:
		id, p := assignmentParts(e.Object)
		return id, append(p, e)
	case *ast.IndexExpr:
		id, p := assignmentParts(e.Object)
		return id, append(p, e)
	}
	return nil, nil
}

func (epp *epp) eppAssign(role string, s *ast.AssignStmt) (out []projection.Statement) {
	id, parts := assignmentParts(s.LHS)
	if id == nil {
		panic("endpoint projection encountered an invalid assignment target")
	}
	sym := epp.info.Symbols[id]
	target := sym.Type()
	specs := []projection.AssignSpecifier{}
	for _, part := range parts {
		switch p := part.(type) {
		case *ast.FieldAccessExpr:
			specs = append(specs, projection.AssignSpecifier{Kind: projection.AssignField, FieldName: p.Field.Value()})
			target, _ = epp.info.Field(target, p.Field.Value())
		case *ast.IndexExpr:
			index, aux := epp.eppExpression(role, p.Index)
			out = append(out, aux...)
			specs = append(specs, projection.AssignSpecifier{Kind: projection.AssignIndex, IndexExpr: index})
			target = target.(*types.ListType).Inner()
		}
	}
	expr, aux := epp.eppExpression(role, s.RHS)
	out = append(out, aux...)
	if target.Roles().Contains(role) {
		t := epp.eppType(role, target)
		expr = epp.storeExpression(role, expr, t)
		structScope := epp.info.GlobalScope.Innermost(s.StartToken().Span.Start).GetStruct()
		attr := structScope != nil && sym.Parent() == structScope.Scope()
		out = append(out, projection.NewStmtAssign(id.Value(), specs, attr, expr))
	} else if expr != nil && expr.HasSideEffects() {
		out = append(out, projection.NewStmtExpr(expr))
	}
	return
}

func (epp *epp) storeExpression(role string, expr projection.Expression, storeType projection.Type) projection.Expression {
	if _, ok := expr.Type().(*projection.FunctionType); ok {
		expr = epp.convertFuncToClosure(role, expr)
	}
	_, async := expr.Type().(*projection.AsyncType)
	if _, want := storeType.(*projection.AsyncType); want && !async {
		expr = projection.NewExprAsync(expr)
	}
	return projection.NewExprPassValue(expr)
}

func (epp *epp) convertFuncToClosure(role string, expr projection.Expression) projection.Expression {
	ft, ok := expr.Type().(*projection.FunctionType)
	if !ok {
		panic(fmt.Sprintf("cannot close %T", expr.Type()))
	}
	sym := epp.info.Symbols[ft.NameIdent()].(*sym_table.FuncSymbol)
	params := []projection.ClosureParam{}
	forward, _ := ft.Roles().SubstituteMap(sym.Roles())
	inverse := forward.Inverse()
	for _, p := range sym.Params() {
		if p.Type().Roles().SubstituteRoles(inverse).Contains(role) {
			params = append(params, projection.NewClosureParam(p.SymbolName(), ft.Params()[len(params)]))
		}
	}
	args := []projection.Expression{}
	for _, p := range params {
		args = append(args, projection.NewExprIdent(p.Name, p.Type))
	}
	call := projection.NewExprCallFunc(expr, role, args, ft.ReturnType(), sym.Roles(), ft.Roles())
	var body projection.Statement = projection.NewStmtExpr(call)
	if ft.ReturnType() != projection.UnitType() {
		body = projection.NewStmtReturn(call)
	}
	return projection.NewExprClosure(params, ft.ReturnType(), []projection.Statement{body})
}
