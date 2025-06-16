package codegen_ts

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
	return gen.Writeln("%s = %s;", s.Name, gen.GenExpr(s.Expr))
}

func (gen *codegen) GenStmtExpr(s *projection.StmtExpr) string {
	return gen.Writeln("%s;", gen.GenExpr(s.Expr))
}

func (gen *codegen) GenStmtIf(s *projection.StmtIf) string {
	return gen.Writeln("[StmtIf]")
}

func (gen *codegen) GenStmtReturn(s *projection.StmtReturn) string {
	return gen.Writeln("return %s;", gen.GenExpr(s.Expr))
}

func (gen *codegen) GenStmtVarDecl(s *projection.StmtVarDecl) string {
	return gen.Writeln("let %s: %s = %s;", s.Name, gen.GenType(s.Expr.Type()), gen.GenExpr(s.Expr))
}

func (gen *codegen) GenStmtWhile(s *projection.StmtWhile) string {
	return gen.Writeln("[StmtWhile]")
}
