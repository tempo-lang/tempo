package sym_table

import (
	"iter"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/types"
)

type Symbol interface {
	SymbolName() string
	Ident() parser.IIdentContext
	Type() types.Type
	Parent() *Scope
	IsAssignable() bool

	AccessReads() []parser.IIdentContext
	AccessWrites() []parser.IIdentContext
	AddRead(ident parser.IIdentContext)
	AddWrite(ident parser.IIdentContext)
}

type baseSymbol struct {
	ident        parser.IIdentContext
	symType      types.Type
	parent       *Scope
	accessReads  []parser.IIdentContext
	accessWrites []parser.IIdentContext
}

func newBaseSymbol(ident parser.IIdentContext, symType types.Type, parent *Scope) baseSymbol {
	return baseSymbol{
		ident:        ident,
		symType:      symType,
		parent:       parent,
		accessReads:  []parser.IIdentContext{},
		accessWrites: []parser.IIdentContext{},
	}
}

func (s *baseSymbol) SymbolName() string {
	return s.ident.GetText()
}
func (s *baseSymbol) Ident() parser.IIdentContext {
	return s.ident
}
func (s *baseSymbol) Type() types.Type {
	return s.symType
}
func (s *baseSymbol) Parent() *Scope {
	return s.parent
}
func (s *baseSymbol) AccessReads() []parser.IIdentContext {
	return s.accessReads
}
func (s *baseSymbol) AccessWrites() []parser.IIdentContext {
	return s.accessWrites
}
func (s *baseSymbol) AddRead(ident parser.IIdentContext) {
	s.accessReads = append(s.accessReads, ident)
}
func (s *baseSymbol) AddWrite(ident parser.IIdentContext) {
	s.accessWrites = append(s.accessWrites, ident)
}

type FuncSymbol struct {
	baseSymbol
	funcCtx parser.IFuncSigContext
	scope   *Scope
	params  []*FuncParamSymbol
}

func (f *FuncSymbol) FuncSig() parser.IFuncSigContext {
	return f.funcCtx
}

func (f *FuncSymbol) FuncType() *types.FunctionType {
	return f.Type().(*types.FunctionType)
}

func (f *FuncSymbol) CallableType() types.CallableType {
	return f.FuncType()
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

func (f *FuncSymbol) ReturnType() types.Type {
	return f.FuncType().ReturnType()
}

func (f *FuncSymbol) ReturnCtx() parser.IValueTypeContext {
	return f.funcCtx.GetReturnType()
}

func (f *FuncSymbol) Roles() *types.Roles {
	return f.Type().Roles()
}

func NewFuncSymbol(fn parser.IFuncSigContext, scope *Scope, funcType types.Type) Symbol {
	return &FuncSymbol{
		baseSymbol: newBaseSymbol(fn.GetName(), funcType, scope.Parent()),
		funcCtx:    fn,
		scope:      scope,
		params:     []*FuncParamSymbol{},
	}
}

type FuncParamSymbol struct {
	baseSymbol
	param parser.IFuncParamContext
}

func NewFuncParamSymbol(param parser.IFuncParamContext, parent *Scope, paramType types.Type) Symbol {
	return &FuncParamSymbol{
		baseSymbol: newBaseSymbol(param.Ident(), paramType, parent),
		param:      param,
	}
}

func (param *FuncParamSymbol) IsAssignable() bool {
	return true
}

func (param *FuncParamSymbol) Param() parser.IFuncParamContext {
	return param.param
}

type VariableSymbol struct {
	baseSymbol
	decl *parser.StmtVarDeclContext
}

func NewVariableSymbol(decl *parser.StmtVarDeclContext, parent *Scope, varType types.Type) Symbol {
	return &VariableSymbol{
		baseSymbol: newBaseSymbol(decl.Ident(), varType, parent),
		decl:       decl,
	}
}

func (v *VariableSymbol) IsAssignable() bool {
	return true
}

func (v *VariableSymbol) VarDecl() *parser.StmtVarDeclContext {
	return v.decl
}

type StructSymbol struct {
	baseSymbol
	structCtx parser.IStructContext
	scope     *Scope
	fields    []*StructFieldSymbol
	methods   []*FuncSymbol
}

type StructFieldSymbol struct {
	baseSymbol
	field        parser.IStructFieldContext
	parentStruct *StructSymbol
	fieldType    types.Type
}

func NewStructSymbol(structCtx parser.IStructContext, scope *Scope, structType types.Type) Symbol {
	return &StructSymbol{
		baseSymbol: newBaseSymbol(structCtx.Ident(), structType, scope.Parent()),
		structCtx:  structCtx,
		scope:      scope,
		fields:     []*StructFieldSymbol{},
		methods:    []*FuncSymbol{},
	}
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

func (s *StructSymbol) Field(name string) (*StructFieldSymbol, bool) {
	for _, field := range s.fields {
		if field.SymbolName() == name {
			return field, true
		}
	}
	return nil, false
}

func (s *StructSymbol) Methods() []*FuncSymbol {
	return s.methods
}

func (s *StructSymbol) Method(name string) (*FuncSymbol, bool) {
	for _, method := range s.methods {
		if method.SymbolName() == name {
			return method, true
		}
	}
	return nil, false
}

func (s *StructSymbol) AddField(field *StructFieldSymbol) {
	s.fields = append(s.fields, field)
}

func (s *StructSymbol) AddMethod(method *FuncSymbol) {
	s.methods = append(s.methods, method)
}

func (s *StructSymbol) StructCtx() parser.IStructContext {
	return s.structCtx
}

func NewStructFieldSymbol(field parser.IStructFieldContext, parentStruct *StructSymbol, fieldType types.Type) Symbol {
	return &StructFieldSymbol{
		baseSymbol:   newBaseSymbol(field.Ident(), fieldType, parentStruct.Scope()),
		field:        field,
		parentStruct: parentStruct,
		fieldType:    fieldType,
	}
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

type InterfaceSymbol struct {
	baseSymbol
	interfaceCtx parser.IInterfaceContext
	scope        *Scope
	methods      map[string]*FuncSymbol
}

func NewInterfaceSymbol(interfaceCtx parser.IInterfaceContext, scope *Scope, interfaceType types.Type) Symbol {
	return &InterfaceSymbol{
		baseSymbol:   newBaseSymbol(interfaceCtx.Ident(), interfaceType, scope.Parent()),
		interfaceCtx: interfaceCtx,
		scope:        scope,
		methods:      map[string]*FuncSymbol{},
	}
}

func (i *InterfaceSymbol) IsAssignable() bool {
	return false
}

func (i *InterfaceSymbol) Scope() *Scope {
	return i.scope
}

func (i *InterfaceSymbol) Methods() iter.Seq2[string, *FuncSymbol] {
	return func(yield func(string, *FuncSymbol) bool) {
		for name, method := range i.methods {
			if !yield(name, method) {
				return
			}
		}
	}
}

func (i *InterfaceSymbol) Method(name string) (*FuncSymbol, bool) {
	fn, found := i.methods[name]
	return fn, found
}

func (i *InterfaceSymbol) AddMethod(fnSym *FuncSymbol) {
	i.methods[fnSym.SymbolName()] = fnSym
}
