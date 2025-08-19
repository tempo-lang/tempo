package examples_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tempo-lang/tempo/examples/remote_procedure_call"
	"github.com/tempo-lang/tempo/runtime"
	"github.com/tempo-lang/tempo/simulator"
	"github.com/tempo-lang/tempo/transports"
)

func TestRemoteProcedureCallSim(t *testing.T) {
	result := simulator.Run(
		simulator.Proc("A", func(env *runtime.Env) any {
			remote_procedure_call.Start_A(env)
			return nil
		}),
		simulator.Proc("B", func(env *runtime.Env) any {
			remote_procedure_call.Start_B(env)
			return nil
		}),
	)

	expected := []simulator.Result{
		{
			Sends:    []transports.SendValue{{Value: int(10), Receivers: []string{"B"}}},
			Receives: []transports.RecvValue{{Value: int(20), Sender: "B"}},
		},
		{
			Sends:    []transports.SendValue{{Value: int(20), Receivers: []string{"A"}}},
			Receives: []transports.RecvValue{{Value: int(10), Sender: "A"}},
		},
	}

	if diff := cmp.Diff(result, expected); diff != "" {
		t.Errorf("RemoteProcedureCall result did not match expected:\n%s", diff)
	}
}
