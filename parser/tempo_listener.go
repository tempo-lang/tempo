// Code generated from Tempo.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Tempo
import "github.com/antlr4-go/antlr/v4"

// TempoListener is a complete listener for a parse tree produced by TempoParser.
type TempoListener interface {
	antlr.ParseTreeListener

	// EnterSourceFile is called when entering the sourceFile production.
	EnterSourceFile(c *SourceFileContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterValueType is called when entering the valueType production.
	EnterValueType(c *ValueTypeContext)

	// EnterRoleTypeShared is called when entering the roleTypeShared production.
	EnterRoleTypeShared(c *RoleTypeSharedContext)

	// EnterRoleTypeNormal is called when entering the roleTypeNormal production.
	EnterRoleTypeNormal(c *RoleTypeNormalContext)

	// EnterClosureType is called when entering the closureType production.
	EnterClosureType(c *ClosureTypeContext)

	// EnterClosureParamList is called when entering the closureParamList production.
	EnterClosureParamList(c *ClosureParamListContext)

	// EnterStruct is called when entering the struct production.
	EnterStruct(c *StructContext)

	// EnterStructFieldList is called when entering the structFieldList production.
	EnterStructFieldList(c *StructFieldListContext)

	// EnterStructField is called when entering the structField production.
	EnterStructField(c *StructFieldContext)

	// EnterInterface is called when entering the interface production.
	EnterInterface(c *InterfaceContext)

	// EnterInterfaceMethodsList is called when entering the interfaceMethodsList production.
	EnterInterfaceMethodsList(c *InterfaceMethodsListContext)

	// EnterInterfaceMethod is called when entering the interfaceMethod production.
	EnterInterfaceMethod(c *InterfaceMethodContext)

	// EnterFunc is called when entering the func production.
	EnterFunc(c *FuncContext)

	// EnterFuncSig is called when entering the funcSig production.
	EnterFuncSig(c *FuncSigContext)

	// EnterFuncParamList is called when entering the funcParamList production.
	EnterFuncParamList(c *FuncParamListContext)

	// EnterFuncParam is called when entering the funcParam production.
	EnterFuncParam(c *FuncParamContext)

	// EnterFuncArgList is called when entering the funcArgList production.
	EnterFuncArgList(c *FuncArgListContext)

	// EnterScope is called when entering the scope production.
	EnterScope(c *ScopeContext)

	// EnterStmtVarDecl is called when entering the stmtVarDecl production.
	EnterStmtVarDecl(c *StmtVarDeclContext)

	// EnterStmtIf is called when entering the stmtIf production.
	EnterStmtIf(c *StmtIfContext)

	// EnterStmtReturn is called when entering the stmtReturn production.
	EnterStmtReturn(c *StmtReturnContext)

	// EnterStmtAssign is called when entering the stmtAssign production.
	EnterStmtAssign(c *StmtAssignContext)

	// EnterStmtExpr is called when entering the stmtExpr production.
	EnterStmtExpr(c *StmtExprContext)

	// EnterExprBinOp is called when entering the exprBinOp production.
	EnterExprBinOp(c *ExprBinOpContext)

	// EnterExprFieldAccess is called when entering the exprFieldAccess production.
	EnterExprFieldAccess(c *ExprFieldAccessContext)

	// EnterExprCall is called when entering the exprCall production.
	EnterExprCall(c *ExprCallContext)

	// EnterExprCom is called when entering the exprCom production.
	EnterExprCom(c *ExprComContext)

	// EnterExprString is called when entering the exprString production.
	EnterExprString(c *ExprStringContext)

	// EnterExprGroup is called when entering the exprGroup production.
	EnterExprGroup(c *ExprGroupContext)

	// EnterExprBool is called when entering the exprBool production.
	EnterExprBool(c *ExprBoolContext)

	// EnterExprStruct is called when entering the exprStruct production.
	EnterExprStruct(c *ExprStructContext)

	// EnterExprIdent is called when entering the exprIdent production.
	EnterExprIdent(c *ExprIdentContext)

	// EnterExprNum is called when entering the exprNum production.
	EnterExprNum(c *ExprNumContext)

	// EnterExprAwait is called when entering the exprAwait production.
	EnterExprAwait(c *ExprAwaitContext)

	// EnterExprStructField is called when entering the exprStructField production.
	EnterExprStructField(c *ExprStructFieldContext)

	// EnterIdentAccess is called when entering the identAccess production.
	EnterIdentAccess(c *IdentAccessContext)

	// ExitSourceFile is called when exiting the sourceFile production.
	ExitSourceFile(c *SourceFileContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitValueType is called when exiting the valueType production.
	ExitValueType(c *ValueTypeContext)

	// ExitRoleTypeShared is called when exiting the roleTypeShared production.
	ExitRoleTypeShared(c *RoleTypeSharedContext)

	// ExitRoleTypeNormal is called when exiting the roleTypeNormal production.
	ExitRoleTypeNormal(c *RoleTypeNormalContext)

	// ExitClosureType is called when exiting the closureType production.
	ExitClosureType(c *ClosureTypeContext)

	// ExitClosureParamList is called when exiting the closureParamList production.
	ExitClosureParamList(c *ClosureParamListContext)

	// ExitStruct is called when exiting the struct production.
	ExitStruct(c *StructContext)

	// ExitStructFieldList is called when exiting the structFieldList production.
	ExitStructFieldList(c *StructFieldListContext)

	// ExitStructField is called when exiting the structField production.
	ExitStructField(c *StructFieldContext)

	// ExitInterface is called when exiting the interface production.
	ExitInterface(c *InterfaceContext)

	// ExitInterfaceMethodsList is called when exiting the interfaceMethodsList production.
	ExitInterfaceMethodsList(c *InterfaceMethodsListContext)

	// ExitInterfaceMethod is called when exiting the interfaceMethod production.
	ExitInterfaceMethod(c *InterfaceMethodContext)

	// ExitFunc is called when exiting the func production.
	ExitFunc(c *FuncContext)

	// ExitFuncSig is called when exiting the funcSig production.
	ExitFuncSig(c *FuncSigContext)

	// ExitFuncParamList is called when exiting the funcParamList production.
	ExitFuncParamList(c *FuncParamListContext)

	// ExitFuncParam is called when exiting the funcParam production.
	ExitFuncParam(c *FuncParamContext)

	// ExitFuncArgList is called when exiting the funcArgList production.
	ExitFuncArgList(c *FuncArgListContext)

	// ExitScope is called when exiting the scope production.
	ExitScope(c *ScopeContext)

	// ExitStmtVarDecl is called when exiting the stmtVarDecl production.
	ExitStmtVarDecl(c *StmtVarDeclContext)

	// ExitStmtIf is called when exiting the stmtIf production.
	ExitStmtIf(c *StmtIfContext)

	// ExitStmtReturn is called when exiting the stmtReturn production.
	ExitStmtReturn(c *StmtReturnContext)

	// ExitStmtAssign is called when exiting the stmtAssign production.
	ExitStmtAssign(c *StmtAssignContext)

	// ExitStmtExpr is called when exiting the stmtExpr production.
	ExitStmtExpr(c *StmtExprContext)

	// ExitExprBinOp is called when exiting the exprBinOp production.
	ExitExprBinOp(c *ExprBinOpContext)

	// ExitExprFieldAccess is called when exiting the exprFieldAccess production.
	ExitExprFieldAccess(c *ExprFieldAccessContext)

	// ExitExprCall is called when exiting the exprCall production.
	ExitExprCall(c *ExprCallContext)

	// ExitExprCom is called when exiting the exprCom production.
	ExitExprCom(c *ExprComContext)

	// ExitExprString is called when exiting the exprString production.
	ExitExprString(c *ExprStringContext)

	// ExitExprGroup is called when exiting the exprGroup production.
	ExitExprGroup(c *ExprGroupContext)

	// ExitExprBool is called when exiting the exprBool production.
	ExitExprBool(c *ExprBoolContext)

	// ExitExprStruct is called when exiting the exprStruct production.
	ExitExprStruct(c *ExprStructContext)

	// ExitExprIdent is called when exiting the exprIdent production.
	ExitExprIdent(c *ExprIdentContext)

	// ExitExprNum is called when exiting the exprNum production.
	ExitExprNum(c *ExprNumContext)

	// ExitExprAwait is called when exiting the exprAwait production.
	ExitExprAwait(c *ExprAwaitContext)

	// ExitExprStructField is called when exiting the exprStructField production.
	ExitExprStructField(c *ExprStructFieldContext)

	// ExitIdentAccess is called when exiting the identAccess production.
	ExitIdentAccess(c *IdentAccessContext)
}
