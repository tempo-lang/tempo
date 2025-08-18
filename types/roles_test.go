package types_test

import (
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestRoleSubst(t *testing.T) {
	type RoleSubstSample struct {
		Name       string
		Subst      *types.RoleSubst
		From       *types.Roles
		ExpectedTo *types.Roles
	}

	samples := []RoleSubstSample{
		{
			Name:       "single role",
			Subst:      types.NewRoleSubst().AddRole("A", "B"),
			From:       types.SingleRole("A"),
			ExpectedTo: types.SingleRole("B"),
		},
		{
			Name:       "distributed role",
			Subst:      types.NewRoleSubst().AddRole("A", "X").AddRole("B", "Y"),
			From:       types.NewRole([]string{"A", "B"}, false),
			ExpectedTo: types.NewRole([]string{"X", "Y"}, false),
		},
		{
			Name:       "inverse role",
			Subst:      types.NewRoleSubst().AddRole("A", "X").AddRole("B", "Y").Inverse(),
			From:       types.NewRole([]string{"X", "Y"}, false),
			ExpectedTo: types.NewRole([]string{"A", "B"}, false),
		},
		{
			Name:       "transitive role",
			Subst:      types.NewRoleSubst().AddRole("A", "B").ApplySubst(types.NewRoleSubst().AddRole("B", "C")),
			From:       types.SingleRole("A"),
			ExpectedTo: types.SingleRole("C"),
		},
		{
			Name:       "local to shared",
			Subst:      types.NewRoleSubst().AddRole("A", "B").AddRole("A", "C"),
			From:       types.SingleRole("A"),
			ExpectedTo: types.NewRole([]string{"B", "C"}, true),
		},
		{
			Name:       "shared to local",
			Subst:      types.NewRoleSubst().AddRole("B", "A").AddRole("C", "A"),
			From:       types.NewRole([]string{"B", "C"}, true),
			ExpectedTo: types.SingleRole("A"),
		},
		{
			Name:       "shared to shared",
			Subst:      types.NewRoleSubst().AddRole("A", []string{"X", "Y", "Z"}...).AddRole("B", []string{"X", "Y", "Z"}...),
			From:       types.NewRole([]string{"A", "B"}, true),
			ExpectedTo: types.NewRole([]string{"X", "Y", "Z"}, true),
		},
	}

	for _, sample := range samples {
		t.Run(sample.Name, func(t *testing.T) {
			to := sample.From.SubstituteRoles(sample.Subst)
			if diff := cmp.Diff(to, sample.ExpectedTo, cmp.AllowUnexported(types.Roles{})); diff != "" {
				t.Errorf("Role substitution '%s' failed:\n%s", sample.Name, diff)
			}

			subst, ok := sample.From.SubstituteMap(to)
			if !ok {
				t.Fatalf("substitution '%s' failed", sample.Name)
			}

			if diff := cmp.Diff(subst, sample.Subst, nil); diff != "" {
				t.Errorf("Role substitution map wrong '%s':\n%s", sample.Name, diff)
			}
		})
	}
}
