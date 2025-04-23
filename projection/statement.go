package projection

import (
	"chorego/type_check"
	"slices"

	"github.com/dave/jennifer/jen"
)

type Statement interface {
	Codegen() jen.Statement
	IsStatement()
}

type StmtVarDecl struct {
	Name string
	Type string
	Expr Expression
}

func NewStmtVarDecl(name string, typeName string, expr Expression) *StmtVarDecl {
	return &StmtVarDecl{
		Name: name,
		Type: typeName,
		Expr: expr,
	}
}

func (decl *StmtVarDecl) Codegen() jen.Statement {

	typeName := decl.Type
	if slices.Contains(type_check.BuiltinTypes(), decl.Type) {
		typeName = BuiltinTypeGo(decl.Type)
	}

	genDecl := jen.Var().Id(decl.Name).Id(typeName).Op("=").Add(decl.Expr.Codegen())
	fixUnused := jen.Id(decl.Name).Op("=").Id(decl.Name).Comment("Suppress unused variable error")

	return []jen.Code{genDecl, fixUnused}
}

func (decl *StmtVarDecl) IsStatement() {}
