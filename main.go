// The main entry for the Tempo compiler binary.
// This package only consists of a `main` function which immediately calls [cmd.Execute].
package main

import (
	"github.com/tempo-lang/tempo/cmd"
)

func main() {
	cmd.Execute()
}
