// Code generated by tempo, DO NOT EDIT.
package choreography

import runtime "github.com/tempo-lang/tempo/runtime"

// Projection of choreography send
func send_X(env *runtime.Env, value int) {
	runtime.Send(env, value, "Y")
}
func send_Y(env *runtime.Env) int {
	return runtime.GetAsync(runtime.Recv[int](env, "X"))
}

// Projection of choreography foo
func foo_A(env *runtime.Env) {
	var x int = 10
	_ = x
	send_X(env.Subst("A", "X", "B", "Y"), x)
}
func foo_B(env *runtime.Env) {
	var Y int = send_Y(env.Subst("A", "X", "B", "Y"))
	_ = Y
}
