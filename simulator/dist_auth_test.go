package simulator_test

import (
	"testing"

	"github.com/tempo-lang/tempo/runtime"
	"github.com/tempo-lang/tempo/simulator"
	"github.com/tempo-lang/tempo/simulator/dist_auth"
	"github.com/tempo-lang/tempo/transports"

	"github.com/google/go-cmp/cmp"
)

type clientRegistry struct {
	CheckCalls []int
	GetSalts   []int
}

// Check implements dist_auth.ClientRegistry_A.
func (c *clientRegistry) Check(env *runtime.Env, hash int) bool {
	c.CheckCalls = append(c.CheckCalls, hash)
	return true
}

// GetSalt implements dist_auth.ClientRegistry_A.
func (c *clientRegistry) GetSalt(env *runtime.Env, username int) int {
	c.GetSalts = append(c.GetSalts, username)
	return username * 2
}

type tokenGen struct {
	GenToken int
}

// GenerateToken implements dist_auth.TokenGenerator_A.
func (t *tokenGen) GenerateToken(env *runtime.Env) int {
	t.GenToken += 1
	return t.GenToken * 100
}

func TestDistAuthSim(t *testing.T) {
	credentials := dist_auth.Credentials_A{
		Username: 123,
		Password: 456,
	}

	mockRegistry := &clientRegistry{}
	mockTokenGen := &tokenGen{}

	result := simulator.Run(
		simulator.Proc("Client", func(env *runtime.Env) any {

			return dist_auth.Authenticate_Client(env, credentials)
		}),
		simulator.Proc("Service", func(env *runtime.Env) any {
			return dist_auth.Authenticate_Service(env)
		}),
		simulator.Proc("IP", func(env *runtime.Env) any {
			dist_auth.Authenticate_IP(env, mockRegistry, mockTokenGen)
			return nil
		}),
	)

	// Results
	salt := credentials.Username * 2
	hash := credentials.Password + salt
	token := 100

	expectedRegistry := &clientRegistry{
		CheckCalls: []int{hash},
		GetSalts:   []int{credentials.Username},
	}
	if diff := cmp.Diff(mockRegistry, expectedRegistry); diff != "" {
		t.Errorf("ClientRegistry did not match expected:\n%s", diff)
	}

	expectedTokenGen := &tokenGen{
		GenToken: 1,
	}
	if diff := cmp.Diff(mockTokenGen, expectedTokenGen); diff != "" {
		t.Errorf("ClientRegistry did not match expected:\n%s", diff)
	}

	expected := []simulator.Result{
		// Client
		{
			Return: dist_auth.AuthResult_C{Success: true, Token: token},
			Sends: []transports.SendValue{
				{Value: credentials.Username, Receivers: []string{"IP"}},
				{Value: hash, Receivers: []string{"IP"}},
			},
			Receives: []transports.RecvValue{
				{Value: salt, Sender: "IP"}, {Value: true, Sender: "IP"},
				{Value: token, Sender: "IP"},
			},
		},
		// Service
		{
			Return:   dist_auth.AuthResult_S{Success: true, Token: token},
			Sends:    []transports.SendValue{},
			Receives: []transports.RecvValue{{Value: true, Sender: "IP"}, {Value: token, Sender: "IP"}},
		},
		// IP
		{
			Sends: []transports.SendValue{
				{Value: salt, Receivers: []string{"Client"}},
				{Value: true, Receivers: []string{"Client", "Service"}},
				{Value: token, Receivers: []string{"Client", "Service"}},
			},
			Receives: []transports.RecvValue{{Value: 123, Sender: "Client"}, {Value: hash, Sender: "Client"}},
		},
	}

	if diff := cmp.Diff(result, expected); diff != "" {
		t.Errorf("DistAuth result did not match expected:\n%s", diff)
	}
}
