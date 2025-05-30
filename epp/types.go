package epp

import (
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"
)

func (epp *epp) eppType(roleName string, typ *types.Type) types.Value {
	if typ.Roles().Contains(roleName) {

		if structType, ok := typ.Value().(*types.StructType); ok {
			structSym := epp.info.GlobalScope.LookupParent(structType.Name()).(*sym_table.StructSymbol)
			structSym.Type().Roles().Participants()

			roleSubstMap, ok := typ.Roles().SubstituteMap(structSym.Type().Roles())
			if !ok {
				panic("role substitution should succeed")
			}

			return projection.NewStructType(structType, roleSubstMap.Subst(roleName))
		}

		if interfaceType, ok := typ.Value().(*types.InterfaceType); ok {
			interfaceSym := epp.info.GlobalScope.LookupParent(interfaceType.Name()).(*sym_table.InterfaceSymbol)
			interfaceSym.Type().Roles().Participants()

			roleSubstMap, ok := typ.Roles().SubstituteMap(interfaceSym.Type().Roles())
			if !ok {
				panic("role substitution should succeed")
			}

			return projection.NewInterfaceType(interfaceType, roleSubstMap.Subst(roleName))
		}

		return typ.Value()
	} else {
		return types.Unit()
	}
}
