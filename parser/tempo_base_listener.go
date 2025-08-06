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

// EnterRoleIdent is called when production roleIdent is entered.
func (s *BaseTempoListener) EnterRoleIdent(ctx *RoleIdentContext) {}

// ExitRoleIdent is called when production roleIdent is exited.
func (s *BaseTempoListener) ExitRoleIdent(ctx *RoleIdentContext) {}

// EnterAsyncType is called when production asyncType is entered.
func (s *BaseTempoListener) EnterAsyncType(ctx *AsyncTypeContext) {}

// ExitAsyncType is called when production asyncType is exited.
func (s *BaseTempoListener) ExitAsyncType(ctx *AsyncTypeContext) {}

// EnterListType is called when production listType is entered.
func (s *BaseTempoListener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *BaseTempoListener) ExitListType(ctx *ListTypeContext) {}

// EnterClosureType is called when production closureType is entered.
func (s *BaseTempoListener) EnterClosureType(ctx *ClosureTypeContext) {}

// ExitClosureType is called when production closureType is exited.
func (s *BaseTempoListener) ExitClosureType(ctx *ClosureTypeContext) {}

// EnterNamedType is called when production namedType is entered.
func (s *BaseTempoListener) EnterNamedType(ctx *NamedTypeContext) {}

// ExitNamedType is called when production namedType is exited.
func (s *BaseTempoListener) ExitNamedType(ctx *NamedTypeContext) {}

// EnterRoleTypeShared is called when production roleTypeShared is entered.
func (s *BaseTempoListener) EnterRoleTypeShared(ctx *RoleTypeSharedContext) {}

// ExitRoleTypeShared is called when production roleTypeShared is exited.
func (s *BaseTempoListener) ExitRoleTypeShared(ctx *RoleTypeSharedContext) {}

// EnterRoleTypeNormal is called when production roleTypeNormal is entered.
func (s *BaseTempoListener) EnterRoleTypeNormal(ctx *RoleTypeNormalContext) {}

// ExitRoleTypeNormal is called when production roleTypeNormal is exited.
func (s *BaseTempoListener) ExitRoleTypeNormal(ctx *RoleTypeNormalContext) {}

// EnterClosureParamList is called when production closureParamList is entered.
func (s *BaseTempoListener) EnterClosureParamList(ctx *ClosureParamListContext) {}

// ExitClosureParamList is called when production closureParamList is exited.
func (s *BaseTempoListener) ExitClosureParamList(ctx *ClosureParamListContext) {}

// EnterClosureSig is called when production closureSig is entered.
func (s *BaseTempoListener) EnterClosureSig(ctx *ClosureSigContext) {}

// ExitClosureSig is called when production closureSig is exited.
func (s *BaseTempoListener) ExitClosureSig(ctx *ClosureSigContext) {}

// EnterStruct is called when production struct is entered.
func (s *BaseTempoListener) EnterStruct(ctx *StructContext) {}

// ExitStruct is called when production struct is exited.
func (s *BaseTempoListener) ExitStruct(ctx *StructContext) {}

// EnterStructImplements is called when production structImplements is entered.
func (s *BaseTempoListener) EnterStructImplements(ctx *StructImplementsContext) {}

// ExitStructImplements is called when production structImplements is exited.
func (s *BaseTempoListener) ExitStructImplements(ctx *StructImplementsContext) {}

// EnterStructBody is called when production structBody is entered.
func (s *BaseTempoListener) EnterStructBody(ctx *StructBodyContext) {}

// ExitStructBody is called when production structBody is exited.
func (s *BaseTempoListener) ExitStructBody(ctx *StructBodyContext) {}

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

// EnterStmtWhile is called when production stmtWhile is entered.
func (s *BaseTempoListener) EnterStmtWhile(ctx *StmtWhileContext) {}

// ExitStmtWhile is called when production stmtWhile is exited.
func (s *BaseTempoListener) ExitStmtWhile(ctx *StmtWhileContext) {}

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

// EnterAssignField is called when production assignField is entered.
func (s *BaseTempoListener) EnterAssignField(ctx *AssignFieldContext) {}

// ExitAssignField is called when production assignField is exited.
func (s *BaseTempoListener) ExitAssignField(ctx *AssignFieldContext) {}

// EnterAssignIndex is called when production assignIndex is entered.
func (s *BaseTempoListener) EnterAssignIndex(ctx *AssignIndexContext) {}

// ExitAssignIndex is called when production assignIndex is exited.
func (s *BaseTempoListener) ExitAssignIndex(ctx *AssignIndexContext) {}

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

// EnterExprStruct is called when production exprStruct is entered.
func (s *BaseTempoListener) EnterExprStruct(ctx *ExprStructContext) {}

// ExitExprStruct is called when production exprStruct is exited.
func (s *BaseTempoListener) ExitExprStruct(ctx *ExprStructContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseTempoListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseTempoListener) ExitExprList(ctx *ExprListContext) {}

// EnterExprIdent is called when production exprIdent is entered.
func (s *BaseTempoListener) EnterExprIdent(ctx *ExprIdentContext) {}

// ExitExprIdent is called when production exprIdent is exited.
func (s *BaseTempoListener) ExitExprIdent(ctx *ExprIdentContext) {}

// EnterExprIndex is called when production exprIndex is entered.
func (s *BaseTempoListener) EnterExprIndex(ctx *ExprIndexContext) {}

// ExitExprIndex is called when production exprIndex is exited.
func (s *BaseTempoListener) ExitExprIndex(ctx *ExprIndexContext) {}

// EnterExprAwait is called when production exprAwait is entered.
func (s *BaseTempoListener) EnterExprAwait(ctx *ExprAwaitContext) {}

// ExitExprAwait is called when production exprAwait is exited.
func (s *BaseTempoListener) ExitExprAwait(ctx *ExprAwaitContext) {}

// EnterExprClosure is called when production exprClosure is entered.
func (s *BaseTempoListener) EnterExprClosure(ctx *ExprClosureContext) {}

// ExitExprClosure is called when production exprClosure is exited.
func (s *BaseTempoListener) ExitExprClosure(ctx *ExprClosureContext) {}

// EnterExprPrimitive is called when production exprPrimitive is entered.
func (s *BaseTempoListener) EnterExprPrimitive(ctx *ExprPrimitiveContext) {}

// ExitExprPrimitive is called when production exprPrimitive is exited.
func (s *BaseTempoListener) ExitExprPrimitive(ctx *ExprPrimitiveContext) {}

// EnterExprStructField is called when production exprStructField is entered.
func (s *BaseTempoListener) EnterExprStructField(ctx *ExprStructFieldContext) {}

// ExitExprStructField is called when production exprStructField is exited.
func (s *BaseTempoListener) ExitExprStructField(ctx *ExprStructFieldContext) {}

// EnterIdentAccess is called when production identAccess is entered.
func (s *BaseTempoListener) EnterIdentAccess(ctx *IdentAccessContext) {}

// ExitIdentAccess is called when production identAccess is exited.
func (s *BaseTempoListener) ExitIdentAccess(ctx *IdentAccessContext) {}

// EnterFloat is called when production float is entered.
func (s *BaseTempoListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production float is exited.
func (s *BaseTempoListener) ExitFloat(ctx *FloatContext) {}

// EnterInt is called when production int is entered.
func (s *BaseTempoListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BaseTempoListener) ExitInt(ctx *IntContext) {}

// EnterString is called when production string is entered.
func (s *BaseTempoListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseTempoListener) ExitString(ctx *StringContext) {}

// EnterBool is called when production bool is entered.
func (s *BaseTempoListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *BaseTempoListener) ExitBool(ctx *BoolContext) {}
