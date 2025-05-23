package sym_table

import (
	"tempo/parser"
	"tempo/types"
)

type Symbol interface {
	SymbolName() string
	Ident() parser.IIdentContext
	Type() *types.Type
	Parent() *Scope
	IsAssignable() bool
}

type FuncSymbol struct {
	funcCtx  parser.IFuncContext
	scope    *Scope
	funcType *types.Type
	params   []*FuncParamSymbol
}

func (f *FuncSymbol) Parent() *Scope {
	return f.scope.parent
}

func (f *FuncSymbol) Ident() parser.IIdentContext {
	return f.funcCtx.Ident()
}

func (f *FuncSymbol) SymbolName() string {
	return f.funcCtx.Ident().GetText()
}

func (f *FuncSymbol) Type() *types.Type {
	return f.funcType
}

func (f *FuncSymbol) Func() parser.IFuncContext {
	return f.funcCtx
}

func (f *FuncSymbol) FuncValue() *types.FunctionType {
	return f.funcType.Value().(*types.FunctionType)
}

func (f *FuncSymbol) Scope() *Scope {
	return f.scope
}

func (f *FuncSymbol) IsAssignable() bool {
	return false
}

func (f *FuncSymbol) Params() []*FuncParamSymbol {
	return f.params
}

func (f *FuncSymbol) AddParam(param *FuncParamSymbol) {
	f.params = append(f.params, param)
}

func (f *FuncSymbol) Roles() *types.Roles {
	return f.funcType.Roles()
}

func NewFuncSymbol(fn parser.IFuncContext, scope *Scope, funcType *types.Type) Symbol {
	return &FuncSymbol{
		funcCtx:  fn,
		scope:    scope,
		funcType: funcType,
		params:   []*FuncParamSymbol{},
	}
}

type FuncParamSymbol struct {
	param     parser.IFuncParamContext
	paramType *types.Type
	parent    *Scope
}

func NewFuncParamSymbol(param parser.IFuncParamContext, parent *Scope, paramType *types.Type) Symbol {
	return &FuncParamSymbol{param: param, parent: parent, paramType: paramType}
}

func (param *FuncParamSymbol) SymbolName() string {
	return param.param.Ident().GetText()
}

func (param *FuncParamSymbol) Ident() parser.IIdentContext {
	return param.param.Ident()
}

func (f *FuncParamSymbol) Type() *types.Type {
	return f.paramType
}

func (param *FuncParamSymbol) Parent() *Scope {
	return param.parent
}

func (param *FuncParamSymbol) IsAssignable() bool {
	return true
}

func (param *FuncParamSymbol) Param() parser.IFuncParamContext {
	return param.param
}

type VariableSymbol struct {
	decl    *parser.StmtVarDeclContext
	parent  *Scope
	varType *types.Type
}

func NewVariableSymbol(decl *parser.StmtVarDeclContext, parent *Scope, varType *types.Type) Symbol {
	return &VariableSymbol{decl: decl, parent: parent, varType: varType}
}

func (v *VariableSymbol) SymbolName() string {
	return v.decl.Ident().GetText()
}

func (v *VariableSymbol) Ident() parser.IIdentContext {
	return v.decl.Ident()
}

func (v *VariableSymbol) Type() *types.Type {
	return v.varType
}

func (v *VariableSymbol) Parent() *Scope {
	return v.parent
}

func (v *VariableSymbol) IsAssignable() bool {
	return true
}

func (v *VariableSymbol) VarDecl() *parser.StmtVarDeclContext {
	return v.decl
}

type StructSymbol struct {
	structCtx  parser.IStructContext
	scope      *Scope
	structType *types.Type
	fields     []*StructFieldSymbol
}

type StructFieldSymbol struct {
	field        parser.IStructFieldContext
	parentStruct *StructSymbol
	fieldType    *types.Type
}

func NewStructSymbol(structCtx parser.IStructContext, scope *Scope, structType *types.Type) Symbol {
	return &StructSymbol{
		structCtx:  structCtx,
		scope:      scope,
		structType: structType,
		fields:     []*StructFieldSymbol{},
	}
}

func (s *StructSymbol) SymbolName() string {
	return s.structCtx.Ident().GetText()
}

func (s *StructSymbol) Ident() parser.IIdentContext {
	return s.structCtx.Ident()
}

func (s *StructSymbol) Type() *types.Type {
	return s.structType
}

func (s *StructSymbol) Parent() *Scope {
	return s.scope.parent
}

func (s *StructSymbol) IsAssignable() bool {
	return false
}

func (s *StructSymbol) Scope() *Scope {
	return s.scope
}

func (s *StructSymbol) Fields() []*StructFieldSymbol {
	return s.fields
}

func (s *StructSymbol) AddField(field *StructFieldSymbol) {
	s.fields = append(s.fields, field)
}

func NewStructFieldSymbol(field parser.IStructFieldContext, parentStruct *StructSymbol, fieldType *types.Type) Symbol {
	return &StructFieldSymbol{
		field:        field,
		parentStruct: parentStruct,
		fieldType:    fieldType,
	}
}

func (f *StructFieldSymbol) SymbolName() string {
	return f.field.Ident().GetText()
}

func (f *StructFieldSymbol) Ident() parser.IIdentContext {
	return f.field.Ident()
}

func (f *StructFieldSymbol) Type() *types.Type {
	return f.fieldType
}

func (f *StructFieldSymbol) Parent() *Scope {
	return f.parentStruct.scope.parent
}

func (f *StructFieldSymbol) IsAssignable() bool {
	return true
}

func (f *StructFieldSymbol) Struct() *StructSymbol {
	return f.parentStruct
}

func (f *StructFieldSymbol) Field() parser.IStructFieldContext {
	return f.field
}
