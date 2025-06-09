package types_test

import (
	"testing"

	"github.com/tempo-lang/tempo/types"
)

type coerceTestCase struct {
	name      string
	fromType  types.Type
	toType    types.Type
	canCoerse bool
}

func TestCanCoerceTypes(t *testing.T) {

	testCases := []coerceTestCase{
		{
			"type can coerce to itself",
			types.Bool([]string{"A"}),
			types.Bool([]string{"A"}),
			true,
		},
		{
			"type can coerce to async version",
			types.Bool([]string{"A"}),
			types.Async(types.Bool([]string{"A"})),
			true,
		},
		{
			"invalid coerse to anything",
			types.Invalid(),
			types.Bool([]string{"A"}),
			true,
		},
		{
			"shared value coerse to subset of roles",
			types.Bool([]string{"A"}),
			types.Bool([]string{"A"}),
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if _, canCoerce := tc.fromType.CoerceTo(tc.toType); canCoerce != tc.canCoerse {
				t.Errorf("Expected %s coerse to %s to be %v", tc.fromType.ToString(), tc.toType.ToString(), tc.canCoerse)
			}
		})
	}

}
