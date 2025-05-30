package types_test

import (
	"slices"
	"testing"

	"github.com/tempo-lang/tempo/types"
)

func TestRoleIntersect(t *testing.T) {

	roles := []*types.Roles{
		types.NewRole([]string{"A", "B", "C"}, true),
		types.NewRole([]string{"B", "C", "D"}, true),
	}

	intersect, ok := types.RoleIntersect(roles...)
	if !ok {
		t.Fatalf("expected roles to be intersectable")
	}

	expected := types.NewRole([]string{"B", "C"}, true)

	if slices.Compare(intersect.Participants(), expected.Participants()) != 0 {
		t.Errorf("Expected participants: %v, got %v", expected.Participants(), intersect.Participants())
	}
}
