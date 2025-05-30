package projection

import "github.com/dave/jennifer/jen"

const RUNTIME_PATH = "github.com/tempo-lang/tempo/runtime"

func RuntimeFunc(name string) *jen.Statement {
	return jen.Qual(RUNTIME_PATH, name)
}
