package codegen_ts

import (
	"fmt"
	"strings"

	"github.com/tempo-lang/tempo/misc"
	"github.com/tempo-lang/tempo/projection"
)

func (gen *codegen) GenFunc(f *projection.Func) string {
	var out strings.Builder
	out.WriteString(gen.Writeln("%s {", gen.GenFuncSig(f.FuncSig)))
	gen.IncIndent()

	for _, stmt := range f.Body {
		out.WriteString(gen.GenStmt(stmt))
	}

	gen.DecIndent()
	out.WriteString(gen.Writeln("}"))
	return out.String()
}

func (gen *codegen) GenFuncParams(f *projection.FuncSig) string {
	params := []string{}
	if gen.opts.DisableTypes {
		params = append(params, "env")
	} else {
		params = append(params, "env: Env")
	}

	for _, param := range f.Params {
		if gen.opts.DisableTypes {
			params = append(params, param.Name)
		} else {
			params = append(params, fmt.Sprintf("%s: %s", param.Name, gen.GenType(param.TypeValue)))
		}
	}

	return misc.JoinStrings(params, ", ")
}

func (gen *codegen) GenFuncSig(f *projection.FuncSig) string {
	params := gen.GenFuncParams(f)
	result := fmt.Sprintf("export async function %s(%s)", f.FuncName(), params)

	if f.ReturnValue != projection.UnitType() && !gen.opts.DisableTypes {
		result += fmt.Sprintf(": Promise<%s>", gen.GenType(f.ReturnValue))
	}

	return result
}
