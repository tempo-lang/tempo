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
	Type types.Type
	Expr Expression
}

func NewStmtVarDecl(name string, typeName types.Type, expr Expression) *StmtVarDecl {
	return &StmtVarDecl{
		Name: name,
		Type: typeName,
		Expr: expr,
	}
}

func (decl *StmtVarDecl) Codegen() jen.Statement {

	genDecl := jen.Var().Id(decl.Name).Id(CodegenType(decl.Type)).Op("=").Add(decl.Expr.Codegen())
	fixUnused := jen.Id(decl.Name).Op("=").Id(decl.Name).Comment("Suppress unused variable error")

	return []jen.Code{genDecl, fixUnused}
}

func (decl *StmtVarDecl) IsStatement() {}
