package projection

type Statement interface {
	IsStatement()
}

type StmtVarDecl struct {
	Name     string
	Expr     Expression
	IsUnused bool
}

func NewStmtVarDecl(name string, expr Expression, isUnused bool) Statement {
	return &StmtVarDecl{
		Name:     name,
		Expr:     expr,
		IsUnused: isUnused,
	}
}

func (decl *StmtVarDecl) IsStatement() {}

type StmtAssign struct {
	Name string
	Expr Expression
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

func (s *StmtIf) IsStatement() {}

type StmtWhile struct {
	Cond  Expression
	Stmts []Statement
}

func NewStmtWhile(cond Expression, stmts []Statement) Statement {
	return &StmtWhile{
		Cond:  cond,
		Stmts: stmts,
	}
}

func (s *StmtWhile) IsStatement() {}

type StmtReturn struct {
	Expr Expression
}

func NewStmtReturn(expr Expression) Statement {
	return &StmtReturn{
		Expr: expr,
	}
}

func (s *StmtReturn) IsStatement() {}
