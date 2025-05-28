package types

import (
	"fmt"
	"tempo/parser"
)

type StructType struct {
	structIdent parser.IIdentContext
	roleSubst   *RoleSubst
}

func (s *StructType) SubstituteRoles(substMap *RoleSubst) Value {
	newRoleSubst := NewRoleSubst()
	for _, from := range s.roleSubst.Roles {
		newRoleSubst.AddRole(from, substMap.Map[from])
	}

	return s
}

func (s *StructType) IsSendable() bool {
	return true
}

func (t *StructType) IsEquatable() bool {
	return true
}

func (s *StructType) IsValue() {}

func (s *StructType) ToString() string {
	return fmt.Sprintf("struct %s", s.structIdent.GetText())
}

func NewStructType(structIdent parser.IIdentContext, roles []string) Value {
	roleSubst := NewRoleSubst()
	for _, role := range roles {
		roleSubst.AddRole(role, role)
	}

	return &StructType{structIdent: structIdent, roleSubst: roleSubst}
}

func (s *StructType) Name() string {
	return s.structIdent.GetText()
}
