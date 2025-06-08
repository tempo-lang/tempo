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

	// EnterAsyncType is called when entering the asyncType production.
	EnterAsyncType(c *AsyncTypeContext)

	// EnterListType is called when entering the listType production.
	EnterListType(c *ListTypeContext)

	// EnterClosureType is called when entering the closureType production.
	EnterClosureType(c *ClosureTypeContext)

	// EnterNamedType is called when entering the namedType production.
	EnterNamedType(c *NamedTypeContext)

	// EnterRoleTypeShared is called when entering the roleTypeShared production.
	EnterRoleTypeShared(c *RoleTypeSharedContext)

	// EnterRoleTypeNormal is called when entering the roleTypeNormal production.
	EnterRoleTypeNormal(c *RoleTypeNormalContext)

	// EnterClosureParamList is called when entering the closureParamList production.
	EnterClosureParamList(c *ClosureParamListContext)

	// EnterClosureSig is called when entering the closureSig production.
	EnterClosureSig(c *ClosureSigContext)

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

	// EnterStmtWhile is called when entering the stmtWhile production.
	EnterStmtWhile(c *StmtWhileContext)

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

	// EnterExprGroup is called when entering the exprGroup production.
	EnterExprGroup(c *ExprGroupContext)

	// EnterExprStruct is called when entering the exprStruct production.
	EnterExprStruct(c *ExprStructContext)

	// EnterExprIdent is called when entering the exprIdent production.
	EnterExprIdent(c *ExprIdentContext)

	// EnterExprAwait is called when entering the exprAwait production.
	EnterExprAwait(c *ExprAwaitContext)

	// EnterExprClosure is called when entering the exprClosure production.
	EnterExprClosure(c *ExprClosureContext)

	// EnterExprPrimitive is called when entering the exprPrimitive production.
	EnterExprPrimitive(c *ExprPrimitiveContext)

	// EnterExprStructField is called when entering the exprStructField production.
	EnterExprStructField(c *ExprStructFieldContext)

	// EnterIdentAccess is called when entering the identAccess production.
	EnterIdentAccess(c *IdentAccessContext)

	// EnterFloat is called when entering the float production.
	EnterFloat(c *FloatContext)

	// EnterInt is called when entering the int production.
	EnterInt(c *IntContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterBool is called when entering the bool production.
	EnterBool(c *BoolContext)

	// ExitSourceFile is called when exiting the sourceFile production.
	ExitSourceFile(c *SourceFileContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitAsyncType is called when exiting the asyncType production.
	ExitAsyncType(c *AsyncTypeContext)

	// ExitListType is called when exiting the listType production.
	ExitListType(c *ListTypeContext)

	// ExitClosureType is called when exiting the closureType production.
	ExitClosureType(c *ClosureTypeContext)

	// ExitNamedType is called when exiting the namedType production.
	ExitNamedType(c *NamedTypeContext)

	// ExitRoleTypeShared is called when exiting the roleTypeShared production.
	ExitRoleTypeShared(c *RoleTypeSharedContext)

	// ExitRoleTypeNormal is called when exiting the roleTypeNormal production.
	ExitRoleTypeNormal(c *RoleTypeNormalContext)

	// ExitClosureParamList is called when exiting the closureParamList production.
	ExitClosureParamList(c *ClosureParamListContext)

	// ExitClosureSig is called when exiting the closureSig production.
	ExitClosureSig(c *ClosureSigContext)

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

	// ExitStmtWhile is called when exiting the stmtWhile production.
	ExitStmtWhile(c *StmtWhileContext)

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

	// ExitExprGroup is called when exiting the exprGroup production.
	ExitExprGroup(c *ExprGroupContext)

	// ExitExprStruct is called when exiting the exprStruct production.
	ExitExprStruct(c *ExprStructContext)

	// ExitExprIdent is called when exiting the exprIdent production.
	ExitExprIdent(c *ExprIdentContext)

	// ExitExprAwait is called when exiting the exprAwait production.
	ExitExprAwait(c *ExprAwaitContext)

	// ExitExprClosure is called when exiting the exprClosure production.
	ExitExprClosure(c *ExprClosureContext)

	// ExitExprPrimitive is called when exiting the exprPrimitive production.
	ExitExprPrimitive(c *ExprPrimitiveContext)

	// ExitExprStructField is called when exiting the exprStructField production.
	ExitExprStructField(c *ExprStructFieldContext)

	// ExitIdentAccess is called when exiting the identAccess production.
	ExitIdentAccess(c *IdentAccessContext)

	// ExitFloat is called when exiting the float production.
	ExitFloat(c *FloatContext)

	// ExitInt is called when exiting the int production.
	ExitInt(c *IntContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitBool is called when exiting the bool production.
	ExitBool(c *BoolContext)
}
