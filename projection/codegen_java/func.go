package codegen_java

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenFunc(f *projection.Func) string {
	out := gen.Writeln("%s throws Exception {", gen.GenFuncSig(f.FuncSig))
	gen.IncIndent()

	for _, stmt := range f.Body {
		out += gen.GenStmt(stmt)
	}

	gen.DecIndent()
	out += gen.Writeln("}")
	return out
}

func (gen *codegen) GenFuncParams(f *projection.FuncSig) string {
	gen.AddImport(javaPkgTempoEnv)
	params := []string{"Env env"}

	for _, param := range f.Params {
		params = append(params, fmt.Sprintf("%s %s", gen.GenType(param.TypeValue), param.Name))
	}

	return misc.JoinStrings(params, ", ")
}

func (gen *codegen) GenFuncSig(f *projection.FuncSig) string {
	params := gen.GenFuncParams(f)
	result := fmt.Sprintf("public static %s %s(%s)", gen.GenType(f.ReturnValue), f.FuncName(), params)

	return result
}
