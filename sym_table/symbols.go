package sym_table

import (
	"iter"

	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/types"
)

type Symbol interface {
	SymbolName() string
	Ident() *ast.Identifier
	Type() types.Type
	Parent() *Scope
	IsAssignable() bool

	AccessReads() []*ast.Identifier
	AccessWrites() []*ast.Identifier
	AddRead(ident *ast.Identifier)
	AddWrite(ident *ast.Identifier)
}

type baseSymbol struct {
	ident        *ast.Identifier
	symType      types.Type
	parent       *Scope
	accessReads  []*ast.Identifier
	accessWrites []*ast.Identifier
}

func newBaseSymbol(ident *ast.Identifier, symType types.Type, parent *Scope) baseSymbol {
	return baseSymbol{
		ident:        ident,
		symType:      symType,
		parent:       parent,
		accessReads:  []*ast.Identifier{},
		accessWrites: []*ast.Identifier{},
	}
}

func (s *baseSymbol) SymbolName() string {
	return s.ident.Value()
}
func (s *baseSymbol) Ident() *ast.Identifier {
	return s.ident
}
func (s *baseSymbol) Type() types.Type {
	return s.symType
}
func (s *baseSymbol) Parent() *Scope {
	return s.parent
}
func (s *baseSymbol) AccessReads() []*ast.Identifier {
	return s.accessReads
}
func (s *baseSymbol) AccessWrites() []*ast.Identifier {
	return s.accessWrites
}
func (s *baseSymbol) AddRead(ident *ast.Identifier) {
	s.accessReads = append(s.accessReads, ident)
}
func (s *baseSymbol) AddWrite(ident *ast.Identifier) {
	s.accessWrites = append(s.accessWrites, ident)
}

type FuncSymbol struct {
	baseSymbol
	funcCtx *ast.FuncSig
	scope   *Scope
	params  []*FuncParamSymbol
}

func (f *FuncSymbol) FuncSig() *ast.FuncSig {
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

func (f *FuncSymbol) ReturnCtx() ast.ValueType {
	return f.funcCtx.ReturnType
}

func (f *FuncSymbol) Roles() *types.Roles {
	return f.Type().Roles()
}

func NewFuncSymbol(fn *ast.FuncSig, scope *Scope, funcType types.Type) Symbol {
	return &FuncSymbol{
		baseSymbol: newBaseSymbol(fn.Name, funcType, scope.Parent()),
		funcCtx:    fn,
		scope:      scope,
		params:     []*FuncParamSymbol{},
	}
}

type FuncParamSymbol struct {
	baseSymbol
	param *ast.FuncParam
}

func NewFuncParamSymbol(param *ast.FuncParam, parent *Scope, paramType types.Type) Symbol {
	return &FuncParamSymbol{
		baseSymbol: newBaseSymbol(param.Name, paramType, parent),
		param:      param,
	}
}

func (param *FuncParamSymbol) IsAssignable() bool {
	return true
}

func (param *FuncParamSymbol) Param() *ast.FuncParam {
	return param.param
}

type VariableSymbol struct {
	baseSymbol
	decl *ast.LetStmt
}

func NewVariableSymbol(decl *ast.LetStmt, parent *Scope, varType types.Type) Symbol {
	return &VariableSymbol{
		baseSymbol: newBaseSymbol(decl.Name, varType, parent),
		decl:       decl,
	}
}

func (v *VariableSymbol) IsAssignable() bool {
	return true
}

func (v *VariableSymbol) VarDecl() *ast.LetStmt {
	return v.decl
}

type StructSymbol struct {
	baseSymbol
	structCtx *ast.Struct
	scope     *Scope
	fields    []*StructFieldSymbol
	methods   []*FuncSymbol
}

type StructFieldSymbol struct {
	baseSymbol
	field        *ast.StructField
	parentStruct *StructSymbol
	fieldType    types.Type
}

func NewStructSymbol(structCtx *ast.Struct, scope *Scope, structType types.Type) Symbol {
	return &StructSymbol{
		baseSymbol: newBaseSymbol(structCtx.Name, structType, scope.Parent()),
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

func (s *StructSymbol) StructCtx() *ast.Struct {
	return s.structCtx
}

func NewStructFieldSymbol(field *ast.StructField, parentStruct *StructSymbol, fieldType types.Type) Symbol {
	return &StructFieldSymbol{
		baseSymbol:   newBaseSymbol(field.Name, fieldType, parentStruct.Scope()),
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

func (f *StructFieldSymbol) Field() *ast.StructField {
	return f.field
}

type InterfaceSymbol struct {
	baseSymbol
	interfaceCtx *ast.Interface
	scope        *Scope
	methods      map[string]*FuncSymbol
}

func NewInterfaceSymbol(interfaceCtx *ast.Interface, scope *Scope, interfaceType types.Type) Symbol {
	return &InterfaceSymbol{
		baseSymbol:   newBaseSymbol(interfaceCtx.Name, interfaceType, scope.Parent()),
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

type TypeSymbol struct {
	baseSymbol
	name string
}

func (t *TypeSymbol) IsAssignable() bool {
	return false
}

func (s *TypeSymbol) SymbolName() string {
	return s.name
}

func NewTypeSymbol(castType types.Type, globalScope *Scope) Symbol {
	return &TypeSymbol{
		baseSymbol: newBaseSymbol(nil, castType, globalScope),
		name:       castType.ToString(),
	}
}
