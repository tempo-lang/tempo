package simulator_test

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"testing"

	"github.com/tempo-lang/tempo/runtime"
	"github.com/tempo-lang/tempo/simulator"
	"github.com/tempo-lang/tempo/simulator/dist_auth"
	"github.com/tempo-lang/tempo/transports"

	"github.com/google/go-cmp/cmp"
)

type clientRegistry struct {
	CheckCalls []string
	GetSalts   []string
}

// Check implements dist_auth.ClientRegistry_A.
func (c *clientRegistry) Check(env *runtime.Env, hash string) bool {
	c.CheckCalls = append(c.CheckCalls, hash)
	return true
}

// GetSalt implements dist_auth.ClientRegistry_A.
func (c *clientRegistry) GetSalt(env *runtime.Env, username string) string {
	hash := sha256.Sum256([]byte(username))
	salt := base64.StdEncoding.EncodeToString(hash[0:8])
	c.GetSalts = append(c.GetSalts, salt)
	return salt
}

type hasher struct {
	hashes []string
}

// CalcHash implements dist_auth.Hasher_A.
func (h *hasher) CalcHash(env *runtime.Env, salt string, password string) string {
	hashBytes := sha256.Sum256([]byte(salt + password))
	hash := base64.StdEncoding.EncodeToString(hashBytes[:])
	h.hashes = append(h.hashes, hash)
	return hash
}

type tokenGen struct {
	Tokens []string
}

// GenerateToken implements dist_auth.TokenGenerator_A.
func (t *tokenGen) GenerateToken(env *runtime.Env) string {
	rng := rand.New(rand.NewSource(int64(len(t.Tokens))))

	token := fmt.Sprintf("%d", rng.Int63())
	t.Tokens = append(t.Tokens, token)

	return token
}

func TestDistAuthSim(t *testing.T) {
	credentials := dist_auth.Credentials_A{
		Username: "username",
		Password: "password",
	}

	mockRegistry := &clientRegistry{}
	mockTokenGen := &tokenGen{}
	mockHasher := &hasher{}

	result := simulator.Run(
		simulator.Proc("Client", func(env *runtime.Env) any {

			return dist_auth.Authenticate_Client(env, credentials, mockHasher)
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
	salt := "FveKfWMX8QI="
	hash := "b26lVgiKAEqqFJfP5KU8AVFIM7eTQO6GsJ8+fido3vw="
	token := "8717895732742165505"

	expectedRegistry := &clientRegistry{
		CheckCalls: []string{hash},
		GetSalts:   []string{salt},
	}
	if diff := cmp.Diff(mockRegistry, expectedRegistry); diff != "" {
		t.Errorf("ClientRegistry did not match expected:\n%s", diff)
	}

	expectedTokenGen := &tokenGen{
		Tokens: []string{token},
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
			Receives: []transports.RecvValue{{Value: credentials.Username, Sender: "Client"}, {Value: hash, Sender: "Client"}},
		},
	}

	if diff := cmp.Diff(result, expected); diff != "" {
		t.Errorf("DistAuth result did not match expected:\n%s", diff)
	}
}
