package projection

import (
	"github.com/dave/jennifer/jen"
)

type Statement interface {
	Codegen() jen.Statement
	IsStatement()
}

type StmtVarDecl struct {
	Name string
	Expr Expression
}

func NewStmtVarDecl(name string, expr Expression) Statement {
	return &StmtVarDecl{
		Name: name,
		Expr: expr,
	}
}

func (decl *StmtVarDecl) Codegen() jen.Statement {

	genDecl := jen.Var().Id(decl.Name).Add(CodegenType(decl.Expr.Type())).Op("=").Add(decl.Expr.Codegen())
	fixUnused := jen.Id("_").Op("=").Id(decl.Name)

	return []jen.Code{genDecl, fixUnused}
}

func (decl *StmtVarDecl) IsStatement() {}

type StmtAssign struct {
	Name string
	Expr Expression
}

func (s *StmtAssign) Codegen() jen.Statement {
	return []jen.Code{jen.Id(s.Name).Op("=").Add(s.Expr.Codegen())}
}

func (s *StmtAssign) IsStatement() {}

func NewStmtAssign(name string, expr Expression) Statement {
	return &StmtAssign{
		Name: name,
		Expr: expr,
	}
}

type StmtExpr struct {
	Expr Expression
}

func (s *StmtExpr) Codegen() jen.Statement {
	if s.Expr.ReturnsValue() {
		return []jen.Code{jen.Id("_").Op("=").Add(s.Expr.Codegen())}
	} else {
		return []jen.Code{s.Expr.Codegen()}
	}
}

func (s *StmtExpr) IsStatement() {}

func NewStmtExpr(expr Expression) Statement {
	return &StmtExpr{
		Expr: expr,
	}
}

type StmtIf struct {
	Guard      Expression
	ThenBranch []Statement
	ElseBranch []Statement
}

func NewStmtIf(guard Expression, thenBranch, elseBranch []Statement) Statement {
	return &StmtIf{
		Guard:      guard,
		ThenBranch: thenBranch,
		ElseBranch: elseBranch,
	}
}

func (s *StmtIf) Codegen() jen.Statement {
	thenStmts := jen.Statement{}
	for _, stmt := range s.ThenBranch {
		thenStmts = append(thenStmts, stmt.Codegen()...)
	}

	ifStmt := jen.If(s.Guard.Codegen()).Block(thenStmts...)

	if elseStmts := s.ElseBranch; len(elseStmts) > 0 {
		elseStmts := jen.Statement{}
		for _, stmt := range s.ElseBranch {
			elseStmts = append(elseStmts, stmt.Codegen()...)
		}

		ifStmt = ifStmt.Else().Block(elseStmts...)
	}

	return []jen.Code{ifStmt}
}

func (s *StmtIf) IsStatement() {}

type StmtReturn struct {
	Expr Expression
}

func NewStmtReturn(expr Expression) Statement {
	return &StmtReturn{
		Expr: expr,
	}
}

func (s *StmtReturn) Codegen() jen.Statement {
	if s.Expr != nil {
		return []jen.Code{jen.Return(s.Expr.Codegen())}
	} else {
		return []jen.Code{jen.Return()}
	}
}

func (s *StmtReturn) IsStatement() {}
