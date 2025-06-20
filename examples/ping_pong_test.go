package examples_test

import (
	"testing"

	"github.com/tempo-lang/tempo/examples/ping_pong"
	"github.com/tempo-lang/tempo/runtime"
	"github.com/tempo-lang/tempo/simulator"
	"github.com/tempo-lang/tempo/transports"

	"github.com/google/go-cmp/cmp"
)

func TestPingPongSim(t *testing.T) {
	result := simulator.Run(
		simulator.Proc("A", func(env *runtime.Env) any {
			ping_pong.Start_A(env)
			return nil
		}),
		simulator.Proc("B", func(env *runtime.Env) any {
			ping_pong.Start_B(env)
			return nil
		}),
	)

	expected := []simulator.Result{
		{
			Sends: []transports.SendValue{
				{Value: int(4), Receivers: []string{"B"}},
				{Value: int(2), Receivers: []string{"B"}},
			},
			Receives: []transports.RecvValue{{Value: int(3), Sender: "B"}, {Value: int(1), Sender: "B"}},
		},
		{
			Sends: []transports.SendValue{
				{Value: int(3), Receivers: []string{"A"}},
				{Value: int(1), Receivers: []string{"A"}},
			},
			Receives: []transports.RecvValue{{Value: int(4), Sender: "A"}, {Value: int(2), Sender: "A"}},
		},
	}

	if diff := cmp.Diff(result, expected); diff != "" {
		t.Errorf("PingPong result did not match expected:\n%s", diff)
	}
}
