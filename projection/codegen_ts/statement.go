package codegen_ts

import (
	"fmt"

	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenStmt(s projection.Statement) {
	switch s := s.(type) {
	case *projection.StmtAssign:
		gen.GenStmtAssign(s)
		return
	case *projection.StmtExpr:
		gen.GenStmtExpr(s)
		return
	case *projection.StmtIf:
		gen.GenStmtIf(s)
		return
	case *projection.StmtReturn:
		gen.GenStmtReturn(s)
		return
	case *projection.StmtVarDecl:
		gen.GenStmtVarDecl(s)
		return
	case *projection.StmtWhile:
		gen.GenStmtWhile(s)
		return
	}

	panic(fmt.Sprintf("unexpected projection.Statement: %#v", s))
}

func (gen *codegen) GenStmtAssign(s *projection.StmtAssign) {
	gen.Writeln("%s = %s;", s.Name, gen.GenExpr(s.Expr))
}

func (gen *codegen) GenStmtExpr(s *projection.StmtExpr) {
	gen.Writeln("%s;", gen.GenExpr(s.Expr))
}

func (gen *codegen) GenStmtIf(s *projection.StmtIf) {
	gen.Writeln("[StmtIf]")
}

func (gen *codegen) GenStmtReturn(s *projection.StmtReturn) {
	gen.Writeln("return %s;", gen.GenExpr(s.Expr))
}

func (gen *codegen) GenStmtVarDecl(s *projection.StmtVarDecl) {
	gen.Writeln("let %s: %s = %s;", s.Name, gen.GenType(s.Expr.Type()), gen.GenExpr(s.Expr))
}

func (gen *codegen) GenStmtWhile(s *projection.StmtWhile) {
	gen.Writeln("[StmtWhile]")
}
