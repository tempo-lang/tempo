package epp

import (
	"tempo/projection"
	"tempo/sym_table"
	"tempo/types"
)

func (epp *epp) eppType(roleName string, typ *types.Type) types.Value {
	if typ.Roles().Contains(roleName) {

		if structType, ok := typ.Value().(*types.StructType); ok {
			structSym := epp.info.GlobalScope.LookupParent(structType.Name()).(*sym_table.StructSymbol)
			structSym.Type().Roles().Participants()

			return projection.NewStructType(structType, "ROLE")
		}

		return typ.Value()
	} else {
		return types.Unit()
	}
}
