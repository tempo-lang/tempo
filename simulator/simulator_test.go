package simulator_test

import (
	"chorego/runtime"
	"chorego/simulator"
	"chorego/simulator/simple_send"
	"chorego/transports"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSimpleSendSim(t *testing.T) {
	result := simulator.Run(
		simulator.Proc("A", func(env *runtime.Env) any {
			simple_send.SimpleSend_A(env)
			return nil
		}),
		simulator.Proc("B", func(env *runtime.Env) any {
			simple_send.SimpleSend_B(env)
			return nil
		}),
	)

	expected := []simulator.Result{
		{
			Sends:    []transports.SendValue{{Value: int(10), Receivers: []string{"B"}}},
			Receives: []transports.RecvValue{},
		},
		{
			Sends:    []transports.SendValue{},
			Receives: []transports.RecvValue{{Value: int(10), Sender: "A"}},
		},
	}

	if diff := cmp.Diff(result, expected); diff != "" {
		t.Errorf("SimpleSend result did not match expected:\n%s", diff)
	}
}
