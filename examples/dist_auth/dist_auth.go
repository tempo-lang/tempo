// Code generated by tempo, DO NOT EDIT.
package dist_auth

import runtime "github.com/tempo-lang/tempo/runtime"

// Projection of interface ClientRegistry
type ClientRegistry_A interface {
	GetSalt(env *runtime.Env, username string) string
	Check(env *runtime.Env, hash string) bool
}

// Projection of interface TokenGenerator
type TokenGenerator_A interface {
	GenerateToken(env *runtime.Env) string
}

// Projection of interface Hasher
type Hasher_A interface {
	CalcHash(env *runtime.Env, salt string, password string) string
}

// Projection of struct Credentials
type Credentials_A struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// Projection of struct AuthResult
type AuthResult_C struct {
	Success bool   `json:"Success"`
	Token   string `json:"Token"`
}
type AuthResult_S struct {
	Success bool   `json:"Success"`
	Token   string `json:"Token"`
}

// Projection of choreography Authenticate
func Authenticate_Client(env *runtime.Env, credentials Credentials_A, hasher Hasher_A) AuthResult_C {
	runtime.Send(env, credentials.Username, "IP")
	var salt *runtime.Async[string] = runtime.Recv[string](env, "IP")
	_ = salt
	var tmp0 *runtime.Async[string] = runtime.FixedAsync(hasher.CalcHash(env.Subst("Client", "A"), runtime.GetAsync(salt), credentials.Password))
	runtime.Send(env, runtime.GetAsync(tmp0), "IP")
	var valid *runtime.Async[bool] = runtime.Recv[bool](env, "IP")
	_ = valid
	if runtime.GetAsync(valid) {
		var token *runtime.Async[string] = runtime.Recv[string](env, "IP")
		_ = token
		return AuthResult_C{
			Success: true,
			Token:   runtime.GetAsync(token),
		}
	} else {
		return AuthResult_C{
			Success: false,
			Token:   "",
		}
	}
}
func Authenticate_Service(env *runtime.Env) AuthResult_S {
	var valid *runtime.Async[bool] = runtime.Recv[bool](env, "IP")
	_ = valid
	if runtime.GetAsync(valid) {
		var token *runtime.Async[string] = runtime.Recv[string](env, "IP")
		_ = token
		return AuthResult_S{
			Success: true,
			Token:   runtime.GetAsync(token),
		}
	} else {
		return AuthResult_S{
			Success: false,
			Token:   "",
		}
	}
}
func Authenticate_IP(env *runtime.Env, registry ClientRegistry_A, tokenGen TokenGenerator_A) {
	var username *runtime.Async[string] = runtime.Recv[string](env, "Client")
	_ = username
	var tmp1 *runtime.Async[string] = runtime.FixedAsync(registry.GetSalt(env.Subst("IP", "A"), runtime.GetAsync(username)))
	runtime.Send(env, runtime.GetAsync(tmp1), "Client")
	var hash *runtime.Async[string] = runtime.Recv[string](env, "Client")
	_ = hash
	var tmp2 *runtime.Async[bool] = runtime.FixedAsync(registry.Check(env.Subst("IP", "A"), runtime.GetAsync(hash)))
	runtime.Send(env, runtime.GetAsync(tmp2), "Client", "Service")
	var valid *runtime.Async[bool] = tmp2
	_ = valid
	if runtime.GetAsync(valid) {
		var tmp3 *runtime.Async[string] = runtime.FixedAsync(tokenGen.GenerateToken(env.Subst("IP", "A")))
		runtime.Send(env, runtime.GetAsync(tmp3), "Client", "Service")
	}
}
