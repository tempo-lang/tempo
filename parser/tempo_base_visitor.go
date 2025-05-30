// Code generated from Tempo.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Tempo
import "github.com/antlr4-go/antlr/v4"

type BaseTempoVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTempoVisitor) VisitSourceFile(ctx *SourceFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitValueType(ctx *ValueTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitRoleTypeShared(ctx *RoleTypeSharedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitRoleTypeNormal(ctx *RoleTypeNormalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitStruct(ctx *StructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitStructFieldList(ctx *StructFieldListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitStructField(ctx *StructFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitInterface(ctx *InterfaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitInterfaceMethodsList(ctx *InterfaceMethodsListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitInterfaceMethod(ctx *InterfaceMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitFunc(ctx *FuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitFuncSig(ctx *FuncSigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitFuncParamList(ctx *FuncParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitFuncParam(ctx *FuncParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitFuncArgList(ctx *FuncArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitScope(ctx *ScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitStmtVarDecl(ctx *StmtVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitStmtIf(ctx *StmtIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitStmtReturn(ctx *StmtReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitStmtAssign(ctx *StmtAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitStmtExpr(ctx *StmtExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitExprBinOp(ctx *ExprBinOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitExprFieldAccess(ctx *ExprFieldAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitExprCall(ctx *ExprCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitExprCom(ctx *ExprComContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitExprGroup(ctx *ExprGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitExprBool(ctx *ExprBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitExprStruct(ctx *ExprStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitExprIdent(ctx *ExprIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitExprNum(ctx *ExprNumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitExprAwait(ctx *ExprAwaitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitExprStructField(ctx *ExprStructFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTempoVisitor) VisitIdentAccess(ctx *IdentAccessContext) interface{} {
	return v.VisitChildren(ctx)
}
