package type_check

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

func (tc typeChecker) checkDuplicateRoles(ctx ast.Node, roleType *types.Roles) type_error.Error {
	roles := roleType.Participants()
	duplications := []string{}
	for i := range roles {
		for j := i + 1; j < len(roles); j++ {
			if roles[i] == roles[j] && roles[i] != "_" && roles[j] != "_" {
				duplications = append(duplications, roles[i])
			}
		}
	}

	// report last error if present
	if len(duplications) > 0 {
		return type_error.NewDuplicateRoles(ctx, duplications)
	}
	return nil
}
