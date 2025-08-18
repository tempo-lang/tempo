package projection

type Statement interface {
	IsStatement()
}

// Declares a new variable and assigns the result of evaluating `Expr` to it.
// If the `Expr` returns a reference, then the value is copied first in order to maintain value semantics.
type StmtVarDecl struct {
	// The name of the new variable.
	Name string
	// The expression that the variable is initially assigned with.
	Expr Expression
	// The type of the variable. Can be different from the Expr type in case of subtyping.
	Type Type
	// IsUnused if true if the variable is never read.
	// False does not necessarily imply that it is used.
	IsUnused bool
}

func NewStmtVarDecl(name string, expr Expression, varType Type, isUnused bool) Statement {
	return &StmtVarDecl{
		Name:     name,
		Expr:     expr,
		Type:     varType,
		IsUnused: isUnused,
	}
}

func (decl *StmtVarDecl) IsStatement() {}

// Evaluates the expression and assigns the result to the variable identified by `Name`.
type StmtAssign struct {
	Name              string
	Specifiers        []AssignSpecifier
	IsMethodAttribute bool
	Expr              Expression
}

type AssignSpecifierKind int

const (
	AssignField AssignSpecifierKind = iota
	AssignIndex
	AssignSelf
)

type AssignSpecifier struct {
	Kind      AssignSpecifierKind
	FieldName string
	IndexExpr Expression
}

func (s *StmtAssign) IsStatement() {}

func NewStmtAssign(name string, specifiers []AssignSpecifier, IsMethodAttribute bool, expr Expression) Statement {
	return &StmtAssign{
		Name:              name,
		Specifiers:        specifiers,
		IsMethodAttribute: IsMethodAttribute,
		Expr:              expr,
	}
}

// Evaluates the expression and discards the result.
type StmtExpr struct {
	Expr Expression
}

func (s *StmtExpr) IsStatement() {}

func NewStmtExpr(expr Expression) Statement {
	return &StmtExpr{
		Expr: expr,
	}
}

// Evaluates the if-statement.
// The `Guard` is of type [BoolType].
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

// Evaluate the condition and run the `Stmts` repeatedly as long as the condition is true.
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

// Return from the current function or closure with the value obtained by evaluating `Expr`.
// `Expr` may be nil if the function does not return a value.
type StmtReturn struct {
	Expr Expression
}

func NewStmtReturn(expr Expression) Statement {
	return &StmtReturn{
		Expr: expr,
	}
}

func (s *StmtReturn) IsStatement() {}
