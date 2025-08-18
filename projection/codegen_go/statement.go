package codegen_go

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/tempo-lang/tempo/projection"
)

func GenStatement(stmt projection.Statement) jen.Statement {
	switch s := stmt.(type) {
	case *projection.StmtAssign:
		return GenStmtAssign(s)
	case *projection.StmtExpr:
		return GenStmtExpr(s)
	case *projection.StmtIf:
		return GenStmtIf(s)
	case *projection.StmtReturn:
		return GenStmtReturn(s)
	case *projection.StmtVarDecl:
		return GenStmtVarDecl(s)
	case *projection.StmtWhile:
		return GenStmtWhile(s)
	}
	panic(fmt.Sprintf("unexpected projection.Statement: %#v", stmt))
}

func GenStmtVarDecl(decl *projection.StmtVarDecl) jen.Statement {
	result := []jen.Code{
		jen.Var().Id(decl.Name).Add(GenType(decl.Type)).Op("=").Add(GenExpression(decl.Expr)),
	}

	if decl.IsUnused {
		result = append(result, jen.Id("_").Op("=").Id(decl.Name))
	}

	return result
}

func GenStmtAssign(s *projection.StmtAssign) jen.Statement {
	base := jen.Id(s.Name)
	if s.IsMethodAttribute {
		base = jen.Id("self").Dot(s.Name)
	}

	for _, specifier := range s.Specifiers {
		switch specifier.Kind {
		case projection.AssignField:
			base = base.Dot(specifier.FieldName)
		case projection.AssignIndex:
			indexExpr := GenExpression(specifier.IndexExpr)
			base = base.Index(indexExpr)
		default:
			panic(fmt.Sprintf("unexpected projection.AssignSpecifierKind: %#v", specifier.Kind))
		}
	}

	return []jen.Code{base.Op("=").Add(GenExpression(s.Expr))}
}

func GenStmtExpr(s *projection.StmtExpr) jen.Statement {
	if s.Expr.ReturnsValue() {
		return []jen.Code{jen.Id("_").Op("=").Add(GenExpression(s.Expr))}
	} else {
		return []jen.Code{GenExpression(s.Expr)}
	}
}

func GenStmtIf(s *projection.StmtIf) jen.Statement {
	thenStmts := jen.Statement{}
	for _, stmt := range s.ThenBranch {
		thenStmts = append(thenStmts, GenStatement(stmt)...)
	}

	ifStmt := jen.If(GenExpression(s.Guard)).Block(thenStmts...)

	if elseStmts := s.ElseBranch; len(elseStmts) > 0 {
		elseStmts := jen.Statement{}
		for _, stmt := range s.ElseBranch {
			elseStmts = append(elseStmts, GenStatement(stmt)...)
		}

		ifStmt = ifStmt.Else().Block(elseStmts...)
	}

	return []jen.Code{ifStmt}
}

func GenStmtWhile(s *projection.StmtWhile) jen.Statement {
	body := jen.Statement{}
	for _, stmt := range s.Stmts {
		body = append(body, GenStatement(stmt)...)
	}
	return []jen.Code{
		jen.For(GenExpression(s.Cond)).Block(body...),
	}
}

func GenStmtReturn(s *projection.StmtReturn) jen.Statement {
	if s.Expr != nil {
		return []jen.Code{jen.Return(GenExpression(s.Expr))}
	} else {
		return []jen.Code{jen.Return()}
	}
}
