// Code generated by tempo, DO NOT EDIT.
package choreography

import runtime "github.com/tempo-lang/tempo/runtime"

// Projection of choreography call
func call_X(env *runtime.Env, val int) {}

// Projection of choreography call2
func call2_Y(env *runtime.Env, param int) {}

// Projection of choreography foo
func foo_A(env *runtime.Env) {
	var x func(int) = func(val int) {
		call_X(env.Subst("A", "X"), val)
	}
	_ = x
	x = func(param int) {
		call2_Y(env.Subst("A", "Y"), param)
	}
	x(10)
}
func foo_B(env *runtime.Env) {}
