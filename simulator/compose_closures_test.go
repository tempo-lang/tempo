package simulator_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tempo-lang/tempo/runtime"
	"github.com/tempo-lang/tempo/simulator"
	"github.com/tempo-lang/tempo/simulator/compose_closures"
	"github.com/tempo-lang/tempo/transports"
)

func TestComposeClosuresSim(t *testing.T) {
	result := simulator.Run(
		simulator.Proc("A", func(env *runtime.Env) any {
			compose_closures.Start_A(env, 1)
			return nil
		}),
		simulator.Proc("B", func(env *runtime.Env) any {
			compose_closures.Start_B(env)
			return nil
		}),
		simulator.Proc("C", func(env *runtime.Env) any {
			return compose_closures.Start_C(env)
		}),
	)

	expected := []simulator.Result{
		// A
		{
			Sends:    []transports.SendValue{{Value: int(2), Receivers: []string{"B"}}},
			Receives: []transports.RecvValue{},
		},
		// B
		{
			Sends:    []transports.SendValue{{Value: int(3), Receivers: []string{"C"}}},
			Receives: []transports.RecvValue{{Value: int(2), Sender: "A"}},
		},
		// C
		{
			Return:   int(3),
			Sends:    []transports.SendValue{},
			Receives: []transports.RecvValue{{Value: int(3), Sender: "B"}},
		},
	}

	if diff := cmp.Diff(result, expected); diff != "" {
		t.Errorf("ComposeClosures result did not match expected:\n%s", diff)
	}
}
