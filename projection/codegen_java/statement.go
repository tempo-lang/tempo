package codegen_java

import (
	"fmt"

	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenStmt(s projection.Statement) string {
	switch s := s.(type) {
	case *projection.StmtAssign:
		return gen.GenStmtAssign(s)
	case *projection.StmtExpr:
		return gen.GenStmtExpr(s)
	case *projection.StmtIf:
		return gen.GenStmtIf(s)
	case *projection.StmtReturn:
		return gen.GenStmtReturn(s)
	case *projection.StmtVarDecl:
		return gen.GenStmtVarDecl(s)
	case *projection.StmtWhile:
		return gen.GenStmtWhile(s)
	}

	panic(fmt.Sprintf("unexpected projection.Statement: %#v", s))
}

func (gen *codegen) GenStmtAssign(s *projection.StmtAssign) string {
	base := s.Name
	if s.IsMethodAttribute {
		base = fmt.Sprintf("this.%s", s.Name)
	}

	for _, specifier := range s.Specifiers {
		switch specifier.Kind {
		case projection.AssignField:
			base = fmt.Sprintf("%s.%s", base, specifier.FieldName)
		case projection.AssignIndex:
			indexExpr := gen.GenExpr(specifier.IndexExpr)
			base = fmt.Sprintf("%s[%s]", base, indexExpr)
		default:
			panic(fmt.Sprintf("unexpected projection.AssignSpecifierKind: %#v", specifier.Kind))
		}
	}

	return gen.Writeln("%s = %s;", base, gen.GenExpr(s.Expr))
}

func (gen *codegen) GenStmtExpr(s *projection.StmtExpr) string {
	return gen.Writeln("%s;", gen.GenExpr(s.Expr))
}

func (gen *codegen) GenStmtIf(s *projection.StmtIf) string {
	out := gen.Writeln("if (%s) {", gen.GenExpr(s.Guard))

	gen.IncIndent()
	for _, stmt := range s.ThenBranch {
		out += gen.GenStmt(stmt)
	}
	gen.DecIndent()

	if len(s.ElseBranch) > 0 {
		out += gen.Writeln("} else {")

		gen.IncIndent()
		for _, stmt := range s.ElseBranch {
			out += gen.GenStmt(stmt)
		}
		gen.DecIndent()
	}

	out += gen.Writeln("}")

	return out
}

func (gen *codegen) GenStmtReturn(s *projection.StmtReturn) string {
	if s.Expr != nil {
		return gen.Writeln("return %s;", gen.GenExpr(s.Expr))
	} else {
		return gen.Writeln("return;")
	}
}

func (gen *codegen) GenStmtVarDecl(s *projection.StmtVarDecl) string {
	return gen.Writeln("%s %s = %s;", gen.GenType(s.Type), s.Name, gen.GenExpr(s.Expr))
}

func (gen *codegen) GenStmtWhile(s *projection.StmtWhile) string {
	out := gen.Writeln("while (%s) {", gen.GenExpr(s.Cond))

	gen.IncIndent()
	for _, stmt := range s.Stmts {
		out += gen.GenStmt(stmt)
	}
	gen.DecIndent()

	out += gen.Writeln("}")

	return out
}
