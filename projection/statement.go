package projection

import (
	"chorego/types"

	"github.com/dave/jennifer/jen"
)

type Statement interface {
	Codegen() jen.Statement
	IsStatement()
}

type StmtVarDecl struct {
	Name string
	Type *types.Type
	Expr Expression
}

func NewStmtVarDecl(name string, typeName *types.Type, expr Expression) Statement {
	return &StmtVarDecl{
		Name: name,
		Type: typeName,
		Expr: expr,
	}
}

func (decl *StmtVarDecl) Codegen() jen.Statement {

	genDecl := jen.Var().Id(decl.Name).Id(CodegenType(decl.Type)).Op("=").Add(decl.Expr.Codegen())
	fixUnused := jen.Id("_").Op("=").Id(decl.Name).Comment("Suppress unused variable error")

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
