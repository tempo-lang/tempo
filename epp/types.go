package epp

import (
	"fmt"

	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"
)

func (epp *epp) eppType(roleName string, typ types.Type) projection.Type {
	if typ.Roles().Contains(roleName) {

		switch typeValue := typ.(type) {
		case *types.StructType:
			structSym := epp.info.GlobalScope.LookupParent(typeValue.Name()).(*sym_table.StructSymbol)
			structSym.Type().Roles().Participants()

			roleSubstMap, ok := typ.Roles().SubstituteMap(structSym.Type().Roles())
			if !ok {
				panic("role substitution should succeed")
			}

			return projection.NewStructType(typeValue, roleSubstMap.Subst(roleName)[0])
		case *types.InterfaceType:
			interfaceSym := epp.info.GlobalScope.LookupParent(typeValue.Name()).(*sym_table.InterfaceSymbol)
			interfaceSym.Type().Roles().Participants()

			roleSubstMap, ok := typ.Roles().SubstituteMap(interfaceSym.Type().Roles())
			if !ok {
				panic("role substitution should succeed")
			}

			return projection.NewInterfaceType(typeValue, roleSubstMap.Subst(roleName)[0])
		case *types.FunctionType:
			params := []projection.Type{}
			for _, param := range typeValue.Params() {
				paramType := epp.eppType(roleName, param)

				if paramType != projection.UnitType() {
					params = append(params, paramType)
				}
			}

			returnType := epp.eppType(roleName, typeValue.ReturnType())

			return projection.NewFunctionType(typeValue, params, returnType)
		case *types.ClosureType:
			params := []projection.Type{}
			for _, param := range typeValue.Params() {
				paramType := epp.eppType(roleName, param)

				if paramType != projection.UnitType() {
					params = append(params, paramType)
				}
			}

			returnType := epp.eppType(roleName, typeValue.ReturnType())

			return projection.NewClosureType(params, returnType)
		case *types.AsyncType:
			innerType := epp.eppType(roleName, typeValue.Inner())
			return projection.NewAsyncType(innerType)
		case *types.ListType:
			innerType := epp.eppType(roleName, typeValue.Inner())
			return projection.NewListType(innerType)
		case *types.UnitType:
			return projection.UnitType()
		}

		if builtin, ok := typ.(types.Builtin); ok {
			switch builtin.(type) {
			case *types.BoolType:
				return projection.BoolType
			case *types.FloatType:
				return projection.FloatType
			case *types.IntType:
				return projection.IntType
			case *types.StringType:
				return projection.StringType
			default:
				panic(fmt.Sprintf("unexpected types.Builtin: %#v", builtin))
			}
		}

		panic(fmt.Sprintf("failed to epp type: %#v\n", typ))
	} else {
		return projection.UnitType()
	}
}
