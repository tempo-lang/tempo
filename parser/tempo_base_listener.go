// Code generated from Tempo.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Tempo
import "github.com/antlr4-go/antlr/v4"

// BaseTempoListener is a complete listener for a parse tree produced by TempoParser.
type BaseTempoListener struct{}

var _ TempoListener = &BaseTempoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTempoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTempoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTempoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTempoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceFile is called when production sourceFile is entered.
func (s *BaseTempoListener) EnterSourceFile(ctx *SourceFileContext) {}

// ExitSourceFile is called when production sourceFile is exited.
func (s *BaseTempoListener) ExitSourceFile(ctx *SourceFileContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseTempoListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseTempoListener) ExitIdent(ctx *IdentContext) {}

// EnterValueType is called when production valueType is entered.
func (s *BaseTempoListener) EnterValueType(ctx *ValueTypeContext) {}

// ExitValueType is called when production valueType is exited.
func (s *BaseTempoListener) ExitValueType(ctx *ValueTypeContext) {}

// EnterRoleTypeShared is called when production roleTypeShared is entered.
func (s *BaseTempoListener) EnterRoleTypeShared(ctx *RoleTypeSharedContext) {}

// ExitRoleTypeShared is called when production roleTypeShared is exited.
func (s *BaseTempoListener) ExitRoleTypeShared(ctx *RoleTypeSharedContext) {}

// EnterRoleTypeNormal is called when production roleTypeNormal is entered.
func (s *BaseTempoListener) EnterRoleTypeNormal(ctx *RoleTypeNormalContext) {}

// ExitRoleTypeNormal is called when production roleTypeNormal is exited.
func (s *BaseTempoListener) ExitRoleTypeNormal(ctx *RoleTypeNormalContext) {}

// EnterStruct is called when production struct is entered.
func (s *BaseTempoListener) EnterStruct(ctx *StructContext) {}

// ExitStruct is called when production struct is exited.
func (s *BaseTempoListener) ExitStruct(ctx *StructContext) {}

// EnterStructFieldList is called when production structFieldList is entered.
func (s *BaseTempoListener) EnterStructFieldList(ctx *StructFieldListContext) {}

// ExitStructFieldList is called when production structFieldList is exited.
func (s *BaseTempoListener) ExitStructFieldList(ctx *StructFieldListContext) {}

// EnterStructField is called when production structField is entered.
func (s *BaseTempoListener) EnterStructField(ctx *StructFieldContext) {}

// ExitStructField is called when production structField is exited.
func (s *BaseTempoListener) ExitStructField(ctx *StructFieldContext) {}

// EnterInterface is called when production interface is entered.
func (s *BaseTempoListener) EnterInterface(ctx *InterfaceContext) {}

// ExitInterface is called when production interface is exited.
func (s *BaseTempoListener) ExitInterface(ctx *InterfaceContext) {}

// EnterInterfaceMethodsList is called when production interfaceMethodsList is entered.
func (s *BaseTempoListener) EnterInterfaceMethodsList(ctx *InterfaceMethodsListContext) {}

// ExitInterfaceMethodsList is called when production interfaceMethodsList is exited.
func (s *BaseTempoListener) ExitInterfaceMethodsList(ctx *InterfaceMethodsListContext) {}

// EnterInterfaceMethod is called when production interfaceMethod is entered.
func (s *BaseTempoListener) EnterInterfaceMethod(ctx *InterfaceMethodContext) {}

// ExitInterfaceMethod is called when production interfaceMethod is exited.
func (s *BaseTempoListener) ExitInterfaceMethod(ctx *InterfaceMethodContext) {}

// EnterFunc is called when production func is entered.
func (s *BaseTempoListener) EnterFunc(ctx *FuncContext) {}

// ExitFunc is called when production func is exited.
func (s *BaseTempoListener) ExitFunc(ctx *FuncContext) {}

// EnterFuncSig is called when production funcSig is entered.
func (s *BaseTempoListener) EnterFuncSig(ctx *FuncSigContext) {}

// ExitFuncSig is called when production funcSig is exited.
func (s *BaseTempoListener) ExitFuncSig(ctx *FuncSigContext) {}

// EnterFuncParamList is called when production funcParamList is entered.
func (s *BaseTempoListener) EnterFuncParamList(ctx *FuncParamListContext) {}

// ExitFuncParamList is called when production funcParamList is exited.
func (s *BaseTempoListener) ExitFuncParamList(ctx *FuncParamListContext) {}

// EnterFuncParam is called when production funcParam is entered.
func (s *BaseTempoListener) EnterFuncParam(ctx *FuncParamContext) {}

// ExitFuncParam is called when production funcParam is exited.
func (s *BaseTempoListener) ExitFuncParam(ctx *FuncParamContext) {}

// EnterFuncArgList is called when production funcArgList is entered.
func (s *BaseTempoListener) EnterFuncArgList(ctx *FuncArgListContext) {}

// ExitFuncArgList is called when production funcArgList is exited.
func (s *BaseTempoListener) ExitFuncArgList(ctx *FuncArgListContext) {}

// EnterScope is called when production scope is entered.
func (s *BaseTempoListener) EnterScope(ctx *ScopeContext) {}

// ExitScope is called when production scope is exited.
func (s *BaseTempoListener) ExitScope(ctx *ScopeContext) {}

// EnterStmtVarDecl is called when production stmtVarDecl is entered.
func (s *BaseTempoListener) EnterStmtVarDecl(ctx *StmtVarDeclContext) {}

// ExitStmtVarDecl is called when production stmtVarDecl is exited.
func (s *BaseTempoListener) ExitStmtVarDecl(ctx *StmtVarDeclContext) {}

// EnterStmtIf is called when production stmtIf is entered.
func (s *BaseTempoListener) EnterStmtIf(ctx *StmtIfContext) {}

// ExitStmtIf is called when production stmtIf is exited.
func (s *BaseTempoListener) ExitStmtIf(ctx *StmtIfContext) {}

// EnterStmtReturn is called when production stmtReturn is entered.
func (s *BaseTempoListener) EnterStmtReturn(ctx *StmtReturnContext) {}

// ExitStmtReturn is called when production stmtReturn is exited.
func (s *BaseTempoListener) ExitStmtReturn(ctx *StmtReturnContext) {}

// EnterStmtAssign is called when production stmtAssign is entered.
func (s *BaseTempoListener) EnterStmtAssign(ctx *StmtAssignContext) {}

// ExitStmtAssign is called when production stmtAssign is exited.
func (s *BaseTempoListener) ExitStmtAssign(ctx *StmtAssignContext) {}

// EnterStmtExpr is called when production stmtExpr is entered.
func (s *BaseTempoListener) EnterStmtExpr(ctx *StmtExprContext) {}

// ExitStmtExpr is called when production stmtExpr is exited.
func (s *BaseTempoListener) ExitStmtExpr(ctx *StmtExprContext) {}

// EnterExprBinOp is called when production exprBinOp is entered.
func (s *BaseTempoListener) EnterExprBinOp(ctx *ExprBinOpContext) {}

// ExitExprBinOp is called when production exprBinOp is exited.
func (s *BaseTempoListener) ExitExprBinOp(ctx *ExprBinOpContext) {}

// EnterExprFieldAccess is called when production exprFieldAccess is entered.
func (s *BaseTempoListener) EnterExprFieldAccess(ctx *ExprFieldAccessContext) {}

// ExitExprFieldAccess is called when production exprFieldAccess is exited.
func (s *BaseTempoListener) ExitExprFieldAccess(ctx *ExprFieldAccessContext) {}

// EnterExprCall is called when production exprCall is entered.
func (s *BaseTempoListener) EnterExprCall(ctx *ExprCallContext) {}

// ExitExprCall is called when production exprCall is exited.
func (s *BaseTempoListener) ExitExprCall(ctx *ExprCallContext) {}

// EnterExprCom is called when production exprCom is entered.
func (s *BaseTempoListener) EnterExprCom(ctx *ExprComContext) {}

// ExitExprCom is called when production exprCom is exited.
func (s *BaseTempoListener) ExitExprCom(ctx *ExprComContext) {}

// EnterExprGroup is called when production exprGroup is entered.
func (s *BaseTempoListener) EnterExprGroup(ctx *ExprGroupContext) {}

// ExitExprGroup is called when production exprGroup is exited.
func (s *BaseTempoListener) ExitExprGroup(ctx *ExprGroupContext) {}

// EnterExprBool is called when production exprBool is entered.
func (s *BaseTempoListener) EnterExprBool(ctx *ExprBoolContext) {}

// ExitExprBool is called when production exprBool is exited.
func (s *BaseTempoListener) ExitExprBool(ctx *ExprBoolContext) {}

// EnterExprStruct is called when production exprStruct is entered.
func (s *BaseTempoListener) EnterExprStruct(ctx *ExprStructContext) {}

// ExitExprStruct is called when production exprStruct is exited.
func (s *BaseTempoListener) ExitExprStruct(ctx *ExprStructContext) {}

// EnterExprIdent is called when production exprIdent is entered.
func (s *BaseTempoListener) EnterExprIdent(ctx *ExprIdentContext) {}

// ExitExprIdent is called when production exprIdent is exited.
func (s *BaseTempoListener) ExitExprIdent(ctx *ExprIdentContext) {}

// EnterExprNum is called when production exprNum is entered.
func (s *BaseTempoListener) EnterExprNum(ctx *ExprNumContext) {}

// ExitExprNum is called when production exprNum is exited.
func (s *BaseTempoListener) ExitExprNum(ctx *ExprNumContext) {}

// EnterExprAwait is called when production exprAwait is entered.
func (s *BaseTempoListener) EnterExprAwait(ctx *ExprAwaitContext) {}

// ExitExprAwait is called when production exprAwait is exited.
func (s *BaseTempoListener) ExitExprAwait(ctx *ExprAwaitContext) {}

// EnterExprStructField is called when production exprStructField is entered.
func (s *BaseTempoListener) EnterExprStructField(ctx *ExprStructFieldContext) {}

// ExitExprStructField is called when production exprStructField is exited.
func (s *BaseTempoListener) ExitExprStructField(ctx *ExprStructFieldContext) {}

// EnterIdentAccess is called when production identAccess is entered.
func (s *BaseTempoListener) EnterIdentAccess(ctx *IdentAccessContext) {}

// ExitIdentAccess is called when production identAccess is exited.
func (s *BaseTempoListener) ExitIdentAccess(ctx *IdentAccessContext) {}
