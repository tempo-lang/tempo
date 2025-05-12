package types_test

import (
	"tempo/types"
	"testing"
)

type coerceTestCase struct {
	name      string
	fromType  *types.Type
	toType    *types.Type
	canCoerse bool
}

func TestCanCoerceTypes(t *testing.T) {

	testCases := []coerceTestCase{
		{
			"type can coerce to itself",
			types.New(types.Bool(), types.SingleRole("A")),
			types.New(types.Bool(), types.SingleRole("A")),
			true,
		},
		{
			"type can coerce to async version",
			types.New(types.Bool(), types.SingleRole("A")),
			types.New(types.NewAsync(types.Bool()), types.SingleRole("A")),
			true,
		},
		{
			"invalid coerse to anything",
			types.Invalid(),
			types.New(types.Bool(), types.NewRole([]string{"A", "B"}, true)),
			true,
		},
		{
			"shared value coerse to subset of roles",
			types.New(types.Bool(), types.NewRole([]string{"A", "B", "C"}, true)),
			types.New(types.Bool(), types.NewRole([]string{"A", "C"}, true)),
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.fromType.CanCoerceTo(tc.toType) != tc.canCoerse {
				t.Errorf("Expected %s coerse to %s to be %v", tc.fromType.ToString(), tc.toType.ToString(), tc.canCoerse)
			}
		})
	}

}
