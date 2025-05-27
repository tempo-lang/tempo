package simulator_test

import (
	"math"
	"tempo/runtime"
	"tempo/simulator"
	"tempo/simulator/diffie_hellman"
	"tempo/transports"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type mathImpl struct{}

// Exp implements diffie_hellman.Math_A.
func (m *mathImpl) Exp(env *runtime.Env, base int, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}

func TestDiffieHellmanSim(t *testing.T) {

	result := simulator.Run(
		simulator.Proc("A", func(env *runtime.Env) any {
			return diffie_hellman.DiffieHellman_A(env, &mathImpl{})
		}),
		simulator.Proc("B", func(env *runtime.Env) any {
			return diffie_hellman.DiffieHellman_B(env, &mathImpl{})
		}),
	)

	expected := []simulator.Result{
		{
			Return:   diffie_hellman.Secret_A{A: 18},
			Sends:    []transports.SendValue{{Value: int(4), Receivers: []string{"B"}}},
			Receives: []transports.RecvValue{{Value: int(10), Sender: "B"}},
		},
		{
			Return:   diffie_hellman.Secret_B{B: 18},
			Sends:    []transports.SendValue{{Value: int(10), Receivers: []string{"A"}}},
			Receives: []transports.RecvValue{{Value: int(4), Sender: "A"}},
		},
	}

	if diff := cmp.Diff(result, expected); diff != "" {
		t.Errorf("DiffieHellman result did not match expected:\n%s", diff)
	}
}
