// Code generated by tempo, DO NOT EDIT.
package choreography

import runtime "github.com/tempo-lang/tempo/runtime"

// Projection of choreography call
func call_X(env *runtime.Env, val int) {}

// Projection of choreography foo
func foo_A(env *runtime.Env) {
	var x func(int) = func(val int) {
		call_X(env.Subst("A", "X"), val)
	}
	_ = x
	var val int = 10
	_ = val
	x(val)
}
func foo_B(env *runtime.Env) {
	var y func(int) = func(val int) {
		call_X(env.Subst("B", "X"), val)
	}
	_ = y
	var val int = 10
	_ = val
	y(val)
}
