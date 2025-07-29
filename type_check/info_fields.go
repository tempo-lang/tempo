package type_check

import (
	"fmt"
	"iter"

	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"
)

// Fields returns an iterator over all fields that can be accessed on this type.
func (info *Info) Fields(typ types.Type) iter.Seq2[string, types.Type] {
	fieldKeys := func(keys []string) iter.Seq2[string, types.Type] {
		return func(yield func(string, types.Type) bool) {
			for _, key := range keys {
				fieldType, _ := info.Field(typ, key)
				if !yield(key, fieldType) {
					return
				}
			}
		}
	}
	switch typ := typ.(type) {
	case *types.AsyncType:
		return fieldKeys([]string{})
	case *types.BoolType:
		return fieldKeys([]string{})
	case *types.ClosureType:
		return fieldKeys([]string{})
	case *types.FloatType:
		return fieldKeys([]string{})
	case *types.FunctionType:
		return fieldKeys([]string{})
	case *types.IntType:
		return fieldKeys([]string{})
	case *types.InterfaceType:
		infSym := info.Symbols[typ.Ident()].(*sym_table.InterfaceSymbol)
		return func(yield func(string, types.Type) bool) {
			for name, sym := range infSym.Methods() {
				if !yield(name, sym.FuncType().SubstituteRoles(typ.SubstMap())) {
					return
				}
			}
		}
	case *types.InvalidType:
		return fieldKeys([]string{})
	case *types.ListType:
		return fieldKeys([]string{"length"})
	case *types.StringType:
		return fieldKeys([]string{})
	case *types.StructType:
		stSym := info.Symbols[typ.Ident()].(*sym_table.StructSymbol)
		return func(yield func(string, types.Type) bool) {
			for _, sym := range stSym.Fields() {
				if !yield(sym.SymbolName(), sym.Type().SubstituteRoles(typ.SubstMap())) {
					return
				}
			}
			for _, method := range stSym.Methods() {
				if !yield(method.SymbolName(), method.Type().SubstituteRoles(typ.SubstMap())) {
					return
				}
			}
		}
	case *types.UnitType:
		return fieldKeys([]string{})
	default:
		panic(fmt.Sprintf("unexpected types.Type: %#v", typ))
	}
}

// Field returns the field with the given name, the returned boolean indicates whether the field was found.
func (info *Info) Field(typ types.Type, name string) (types.Type, bool) {
	switch typ := typ.(type) {
	case *types.AsyncType:
		return types.Invalid(), false
	case *types.BoolType:
		return types.Invalid(), false
	case *types.ClosureType:
		return types.Invalid(), false
	case *types.FloatType:
		return types.Invalid(), false
	case *types.FunctionType:
		return types.Invalid(), false
	case *types.IntType:
		return types.Invalid(), false
	case *types.InterfaceType:
		infSym := info.Symbols[typ.Ident()].(*sym_table.InterfaceSymbol)
		if method := infSym.Method(name); method != nil {
			return method.FuncType().SubstituteRoles(typ.SubstMap()), true
		}
		return types.Invalid(), false
	case *types.InvalidType:
		return types.Invalid(), false
	case *types.ListType:
		if name == "length" {
			participants := typ.Inner().Roles().Participants()
			return types.Int(participants), true
		}
		return types.Invalid(), false
	case *types.StringType:
		return types.Invalid(), false
	case *types.StructType:
		stSym := info.Symbols[typ.Ident()].(*sym_table.StructSymbol)
		if field, found := stSym.Field(name); found {
			return field.Type().SubstituteRoles(typ.SubstMap()), true
		}
		if method, found := stSym.Method(name); found {
			return method.Type().SubstituteRoles(typ.SubstMap()), true
		}
		return types.Invalid(), false
	case *types.UnitType:
		return types.Invalid(), false
	default:
		panic(fmt.Sprintf("unexpected types.Type: %#v", typ))
	}
}
