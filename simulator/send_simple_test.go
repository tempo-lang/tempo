package simulator_test

import (
	"chorego/runtime"
	"chorego/simulator"
	"chorego/simulator/send_simple"
	"chorego/transports"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSendSimpleSim(t *testing.T) {
	result := simulator.Run(
		simulator.Proc("A", func(env *runtime.Env) any {
			send_simple.SimpleSend_A(env)
			return nil
		}),
		simulator.Proc("B", func(env *runtime.Env) any {
			send_simple.SimpleSend_B(env)
			return nil
		}),
	)

	expected := []simulator.Result{
		{
			Sends:    []transports.SendValue{{Value: 10, Receivers: []string{"B"}}},
			Receives: []transports.RecvValue{},
		},
		{
			Sends:    []transports.SendValue{},
			Receives: []transports.RecvValue{{Value: 10, Sender: "A"}},
		},
	}

	if diff := cmp.Diff(result, expected); diff != "" {
		t.Errorf("SendSimple result did not match expected:\n%s", diff)
	}
}
