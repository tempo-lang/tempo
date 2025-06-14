package codegen_ts

import (
	"fmt"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenFunc(f *projection.Func) {
	gen.Writeln("%s {", gen.GenFuncSig(f.FuncSig))
	gen.IncIndent()

	for _, stmt := range f.Body {
		gen.GenStmt(stmt)
	}

	gen.DecIndent()
	gen.Writeln("}")
}

func (gen *codegen) GenFuncSig(f *projection.FuncSig) string {

	params := []string{"env: Env"}
	for _, param := range f.Params {
		if gen.opts.DisableTypes {
			params = append(params, param.Name)
		} else {
			params = append(params, fmt.Sprintf("%s: %s", param.Name, gen.GenType(param.TypeValue)))
		}
	}

	result := fmt.Sprintf("export async function %s_%s(%s)", f.Name, f.Role, misc.JoinStrings(params, ", "))

	if f.ReturnValue != projection.UnitType() && !gen.opts.DisableTypes {
		result += fmt.Sprintf(": Promise<%s>", gen.GenType(f.ReturnValue))
	}

	return result
}
