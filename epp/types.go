package epp

import "tempo/types"

func (epp *epp) eppType(roleName string, typ *types.Type) types.Value {
	if typ.Roles().Contains(roleName) {
		return typ.Value()
	} else {
		return types.Unit()
	}
}
