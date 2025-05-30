package type_check

import (
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"

	"github.com/antlr4-go/antlr/v4"
)

func (tc typeChecker) checkDuplicateRoles(ctx antlr.ParserRuleContext, roleType *types.Roles) type_error.Error {
	roles := roleType.Participants()
	duplications := []string{}
	for i := range roles {
		for j := i + 1; j < len(roles); j++ {
			if roles[i] == roles[j] {
				duplications = append(duplications, roles[i])
			}
		}
	}

	// report last error if present
	if len(duplications) > 0 {
		return type_error.NewDuplicateRolesError(ctx, duplications)
	}
	return nil
}
