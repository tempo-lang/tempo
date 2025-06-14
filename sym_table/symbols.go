package sym_table

import (
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

func (f *FuncSymbol) FuncValue() *types.FunctionType {
	return f.Type().(*types.FunctionType)
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
	return f.FuncValue().ReturnType()
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

func (s *StructSymbol) AddField(field *StructFieldSymbol) {
	s.fields = append(s.fields, field)

	stType := s.baseSymbol.symType.(*types.StructType)
	stType.AddField(field.SymbolName(), field.fieldType)
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
	// methods      []*InterfaceMethodSymbol
}

// type InterfaceMethodSymbol struct {
// 	baseSymbol
// 	method       parser.IInterfaceMethodContext
// 	parentStruct *InterfaceSymbol
// 	methodType   types.Type
// }

func NewInterfaceSymbol(interfaceCtx parser.IInterfaceContext, scope *Scope, interfaceType types.Type) Symbol {
	return &InterfaceSymbol{
		baseSymbol:   newBaseSymbol(interfaceCtx.Ident(), interfaceType, scope.Parent()),
		interfaceCtx: interfaceCtx,
		scope:        scope,
		// methods:      []*InterfaceMethodSymbol{},
	}
}

func (i *InterfaceSymbol) IsAssignable() bool {
	return false
}

func (i *InterfaceSymbol) Scope() *Scope {
	return i.scope
}

// func (i *InterfaceSymbol) Methods() []*InterfaceMethodSymbol {
// 	return i.methods
// }

// func (i *InterfaceSymbol) AddMethod(method *FuncSymbol) {
// 	i.methods = append(i.methods, method)

// 	infType := i.baseSymbol.symType.(*types.InterfaceType)
// 	infType.AddField(method.SymbolName(), method.methodType)
// }

// func NewInterfaceMethodSymbol(method parser.IInterfaceMethodContext, parentInterface *InterfaceSymbol, methodType types.Type) Symbol {
// 	return &InterfaceMethodSymbol{
// 		baseSymbol:   newBaseSymbol(method.FuncSig().GetName(), methodType, parentInterface.Scope()),
// 		method:       method,
// 		parentStruct: parentInterface,
// 		methodType:   methodType,
// 	}
// }

// func (m *InterfaceMethodSymbol) IsAssignable() bool {
// 	return false
// }

// func (m *InterfaceMethodSymbol) Interface() *InterfaceSymbol {
// 	return m.parentStruct
// }

// func (m *InterfaceMethodSymbol) Method() parser.IInterfaceMethodContext {
// 	return m.method
// }
