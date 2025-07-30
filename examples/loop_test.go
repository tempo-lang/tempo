package examples_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tempo-lang/tempo/examples/loop"
	"github.com/tempo-lang/tempo/runtime"
	"github.com/tempo-lang/tempo/simulator"
)

func TestLoopSim(t *testing.T) {
	t.Skip()

	result := simulator.Run(
		simulator.Proc("A", func(env *runtime.Env) any {
			loop.Main_A(env)
			return nil
		}),
		simulator.Proc("B", func(env *runtime.Env) any {
			loop.Main_B(env)
			return nil
		}),
	)

	expected := []simulator.Result{}

	if diff := cmp.Diff(result, expected); diff != "" {
		t.Errorf("Loop result did not match expected:\n%s", diff)
	}
}
