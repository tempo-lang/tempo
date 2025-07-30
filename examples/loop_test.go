package examples_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tempo-lang/tempo/examples/loop"
	"github.com/tempo-lang/tempo/runtime"
	"github.com/tempo-lang/tempo/simulator"
	"github.com/tempo-lang/tempo/transports"
)

func TestLoopSim(t *testing.T) {
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

	expected := []simulator.Result{
		{
			Sends: []transports.SendValue{
				{Value: bool(true), Receivers: []string{"B"}},
				{Value: bool(true), Receivers: []string{"B"}},
				{Value: bool(true), Receivers: []string{"B"}},
				{Value: bool(true), Receivers: []string{"B"}},
				{Value: bool(true), Receivers: []string{"B"}},
				{Value: bool(false), Receivers: []string{"B"}},
			},
			Receives: []transports.RecvValue{
				{Value: string("ping"), Sender: "B"}, {Value: string("ping"), Sender: "B"},
				{Value: string("ping"), Sender: "B"}, {Value: string("ping"), Sender: "B"},
				{Value: string("ping"), Sender: "B"},
			},
		},
		{
			Sends: []transports.SendValue{
				{Value: string("ping"), Receivers: []string{"A"}},
				{Value: string("ping"), Receivers: []string{"A"}},
				{Value: string("ping"), Receivers: []string{"A"}},
				{Value: string("ping"), Receivers: []string{"A"}},
				{Value: string("ping"), Receivers: []string{"A"}},
			},
			Receives: []transports.RecvValue{
				{Value: bool(true), Sender: "A"}, {Value: bool(true), Sender: "A"},
				{Value: bool(true), Sender: "A"}, {Value: bool(true), Sender: "A"},
				{Value: bool(true), Sender: "A"}, {Value: bool(false), Sender: "A"},
			},
		},
	}

	if diff := cmp.Diff(result, expected); diff != "" {
		t.Errorf("Loop result did not match expected:\n%s", diff)
	}
}
