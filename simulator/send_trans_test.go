package simulator_test

import (
	"chorego/runtime"
	"chorego/simulator"
	"chorego/simulator/send_trans"
	"chorego/transports"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSendTransSim(t *testing.T) {
	result := simulator.Run(
		simulator.Proc("A", func(env *runtime.Env) any {
			send_trans.Trans_A(env)
			return nil
		}),
		simulator.Proc("B", func(env *runtime.Env) any {
			send_trans.Trans_B(env)
			return nil
		}),
		simulator.Proc("C", func(env *runtime.Env) any {
			send_trans.Trans_C(env)
			return nil
		}),
	)

	expected := []simulator.Result{
		{
			Sends:    []transports.SendValue{{Value: 10, Receivers: []string{"B"}}},
			Receives: []transports.RecvValue{},
		},
		{
			Sends:    []transports.SendValue{{Value: 10, Receivers: []string{"C"}}},
			Receives: []transports.RecvValue{{Value: 10, Sender: "A"}},
		},
		{
			Sends:    []transports.SendValue{},
			Receives: []transports.RecvValue{{Value: 10, Sender: "B"}},
		},
	}

	if diff := cmp.Diff(result, expected); diff != "" {
		t.Errorf("SendTrans result did not match expected:\n%s", diff)
	}
}
