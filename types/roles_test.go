package types_test

import (
	"chorego/types"
	"slices"
	"testing"
)

func TestRoleIntersect(t *testing.T) {

	roles := []*types.Roles{
		types.NewRole([]string{"A", "B", "C"}, true),
		types.NewRole([]string{"B", "C", "D"}, true),
	}

	intersect, err := types.RoleIntersect(nil, roles...)
	if err != nil {
		t.Fatalf("expected roles to be intersectable, got error: %v", err)
	}

	expected := types.NewRole([]string{"B", "C"}, true)

	if slices.Compare(intersect.Participants(), expected.Participants()) != 0 {
		t.Errorf("Expected participants: %v, got %v", expected.Participants(), intersect.Participants())
	}
}
