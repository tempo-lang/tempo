// Code generated from Tempo.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Tempo
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TempoParser.
type TempoVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TempoParser#sourceFile.
	VisitSourceFile(ctx *SourceFileContext) interface{}

	// Visit a parse tree produced by TempoParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by TempoParser#valueType.
	VisitValueType(ctx *ValueTypeContext) interface{}

	// Visit a parse tree produced by TempoParser#roleTypeShared.
	VisitRoleTypeShared(ctx *RoleTypeSharedContext) interface{}

	// Visit a parse tree produced by TempoParser#roleTypeNormal.
	VisitRoleTypeNormal(ctx *RoleTypeNormalContext) interface{}

	// Visit a parse tree produced by TempoParser#closureType.
	VisitClosureType(ctx *ClosureTypeContext) interface{}

	// Visit a parse tree produced by TempoParser#closureParamList.
	VisitClosureParamList(ctx *ClosureParamListContext) interface{}

	// Visit a parse tree produced by TempoParser#struct.
	VisitStruct(ctx *StructContext) interface{}

	// Visit a parse tree produced by TempoParser#structFieldList.
	VisitStructFieldList(ctx *StructFieldListContext) interface{}

	// Visit a parse tree produced by TempoParser#structField.
	VisitStructField(ctx *StructFieldContext) interface{}

	// Visit a parse tree produced by TempoParser#interface.
	VisitInterface(ctx *InterfaceContext) interface{}

	// Visit a parse tree produced by TempoParser#interfaceMethodsList.
	VisitInterfaceMethodsList(ctx *InterfaceMethodsListContext) interface{}

	// Visit a parse tree produced by TempoParser#interfaceMethod.
	VisitInterfaceMethod(ctx *InterfaceMethodContext) interface{}

	// Visit a parse tree produced by TempoParser#func.
	VisitFunc(ctx *FuncContext) interface{}

	// Visit a parse tree produced by TempoParser#funcSig.
	VisitFuncSig(ctx *FuncSigContext) interface{}

	// Visit a parse tree produced by TempoParser#funcParamList.
	VisitFuncParamList(ctx *FuncParamListContext) interface{}

	// Visit a parse tree produced by TempoParser#funcParam.
	VisitFuncParam(ctx *FuncParamContext) interface{}

	// Visit a parse tree produced by TempoParser#funcArgList.
	VisitFuncArgList(ctx *FuncArgListContext) interface{}

	// Visit a parse tree produced by TempoParser#scope.
	VisitScope(ctx *ScopeContext) interface{}

	// Visit a parse tree produced by TempoParser#stmtVarDecl.
	VisitStmtVarDecl(ctx *StmtVarDeclContext) interface{}

	// Visit a parse tree produced by TempoParser#stmtIf.
	VisitStmtIf(ctx *StmtIfContext) interface{}

	// Visit a parse tree produced by TempoParser#stmtReturn.
	VisitStmtReturn(ctx *StmtReturnContext) interface{}

	// Visit a parse tree produced by TempoParser#stmtAssign.
	VisitStmtAssign(ctx *StmtAssignContext) interface{}

	// Visit a parse tree produced by TempoParser#stmtExpr.
	VisitStmtExpr(ctx *StmtExprContext) interface{}

	// Visit a parse tree produced by TempoParser#exprBinOp.
	VisitExprBinOp(ctx *ExprBinOpContext) interface{}

	// Visit a parse tree produced by TempoParser#exprFieldAccess.
	VisitExprFieldAccess(ctx *ExprFieldAccessContext) interface{}

	// Visit a parse tree produced by TempoParser#exprCall.
	VisitExprCall(ctx *ExprCallContext) interface{}

	// Visit a parse tree produced by TempoParser#exprCom.
	VisitExprCom(ctx *ExprComContext) interface{}

	// Visit a parse tree produced by TempoParser#exprString.
	VisitExprString(ctx *ExprStringContext) interface{}

	// Visit a parse tree produced by TempoParser#exprGroup.
	VisitExprGroup(ctx *ExprGroupContext) interface{}

	// Visit a parse tree produced by TempoParser#exprBool.
	VisitExprBool(ctx *ExprBoolContext) interface{}

	// Visit a parse tree produced by TempoParser#exprStruct.
	VisitExprStruct(ctx *ExprStructContext) interface{}

	// Visit a parse tree produced by TempoParser#exprIdent.
	VisitExprIdent(ctx *ExprIdentContext) interface{}

	// Visit a parse tree produced by TempoParser#exprNum.
	VisitExprNum(ctx *ExprNumContext) interface{}

	// Visit a parse tree produced by TempoParser#exprAwait.
	VisitExprAwait(ctx *ExprAwaitContext) interface{}

	// Visit a parse tree produced by TempoParser#exprStructField.
	VisitExprStructField(ctx *ExprStructFieldContext) interface{}

	// Visit a parse tree produced by TempoParser#identAccess.
	VisitIdentAccess(ctx *IdentAccessContext) interface{}
}
