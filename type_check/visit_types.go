package type_check

import (
	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/type_check/type_error"
	"github.com/tempo-lang/tempo/types"
)

func (tc *typeChecker) visitValueType(node ast.ValueType) types.Type {
	if node == nil {
		return types.Invalid()
	}
	t, err := tc.parseValueType(node)
	if err != nil {
		tc.reportError(err)
		return types.Invalid()
	}
	if t.Roles().IsHidden() {
		tc.reportError(type_error.NewHiddenTypeSignature(node, t))
	}
	if c, ok := node.(*ast.ClosureType); ok && !t.Roles().IsComplete() {
		tc.reportError(type_error.NewUnexpectedHiddenRoles(c.RoleType))
	}
	return t
}
