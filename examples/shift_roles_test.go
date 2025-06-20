package examples_test

import (
	"testing"

	"github.com/tempo-lang/tempo/examples/shift_roles"
	"github.com/tempo-lang/tempo/runtime"
	"github.com/tempo-lang/tempo/simulator"
	"github.com/tempo-lang/tempo/transports"

	"github.com/google/go-cmp/cmp"
)

func TestShiftRolesSim(t *testing.T) {
	result := simulator.Run(
		simulator.Proc("A", func(env *runtime.Env) any {
			shift_roles.ShiftRoles_A(env, 5)
			return nil
		}),
		simulator.Proc("B", func(env *runtime.Env) any {
			shift_roles.ShiftRoles_B(env, 5)
			return nil
		}),
		simulator.Proc("C", func(env *runtime.Env) any {
			shift_roles.ShiftRoles_C(env, 5)
			return nil
		}),
		simulator.Proc("D", func(env *runtime.Env) any {
			shift_roles.ShiftRoles_D(env, 5)
			return nil
		}),
	)

	expected := []simulator.Result{
		// A
		{
			Sends: []transports.SendValue{
				{Value: 5, Receivers: []string{"B"}},
				{Value: 1, Receivers: []string{"B"}},
			},
			Receives: []transports.RecvValue{{Value: 2, Sender: "D"}},
		},
		// B
		{
			Sends: []transports.SendValue{
				{Value: 4, Receivers: []string{"C"}},
			},
			Receives: []transports.RecvValue{{Value: 5, Sender: "A"}, {Value: 1, Sender: "A"}},
		},
		// C
		{
			Sends:    []transports.SendValue{{Value: 3, Receivers: []string{"D"}}},
			Receives: []transports.RecvValue{{Value: 4, Sender: "B"}},
		},
		// D
		{
			Sends:    []transports.SendValue{{Value: 2, Receivers: []string{"A"}}},
			Receives: []transports.RecvValue{{Value: 3, Sender: "C"}},
		},
	}

	if diff := cmp.Diff(result, expected); diff != "" {
		t.Errorf("ShiftRoles result did not match expected:\n%s", diff)
	}
}
