package epp

import (
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"
)

func (epp *epp) eppType(roleName string, typ *types.Type) types.Value {
	if typ.Roles().Contains(roleName) {

		switch typeValue := typ.Value().(type) {
		case *types.StructType:
			structSym := epp.info.GlobalScope.LookupParent(typeValue.Name()).(*sym_table.StructSymbol)
			structSym.Type().Roles().Participants()

			roleSubstMap, ok := typ.Roles().SubstituteMap(structSym.Type().Roles())
			if !ok {
				panic("role substitution should succeed")
			}

			return projection.NewStructType(typeValue, roleSubstMap.Subst(roleName))
		case *types.InterfaceType:
			interfaceSym := epp.info.GlobalScope.LookupParent(typeValue.Name()).(*sym_table.InterfaceSymbol)
			interfaceSym.Type().Roles().Participants()

			roleSubstMap, ok := typ.Roles().SubstituteMap(interfaceSym.Type().Roles())
			if !ok {
				panic("role substitution should succeed")
			}

			return projection.NewInterfaceType(typeValue, roleSubstMap.Subst(roleName))
		case *types.FunctionType:

			params := []types.Value{}
			for _, param := range typeValue.Params() {
				paramType := epp.eppType(roleName, param)

				if paramType != types.Unit() {
					params = append(params, paramType)
				}
			}

			returnType := epp.eppType(roleName, typeValue.ReturnType())

			return projection.NewFunctionType(typeValue, params, returnType)
		}

		return typ.Value()
	} else {
		return types.Unit()
	}
}
