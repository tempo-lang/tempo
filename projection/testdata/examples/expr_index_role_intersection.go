// Code generated by tempo, DO NOT EDIT.
package choreography

import runtime "github.com/tempo-lang/tempo/runtime"

// Projection of choreography foo
func foo_A(env *runtime.Env) {
	var x []int = []int{1, 2, 3}
	_ = x
}
func foo_B(env *runtime.Env) {
	var x []int = []int{1, 2, 3}
	_ = x
	var i int = 1
	_ = i
	var y int = x[i]
	_ = y
}
func foo_C(env *runtime.Env) {
	var i int = 1
	_ = i
}
