// Code generated by tempo, DO NOT EDIT.
package choreography

import runtime "github.com/tempo-lang/tempo/runtime"

// Projection of choreography binOps
func binOps_A(env *runtime.Env) {
	var a int = 1 + 2 - 3*4/5%6
	_ = a
	var b bool = true == true
	_ = b
	var c bool = false != false
	_ = c
	var d bool = 10 < 20
	_ = d
	var e bool = 10 <= 20
	_ = e
	var f bool = 10 > 20
	_ = f
	var g bool = 10 >= 20
	_ = g
}
