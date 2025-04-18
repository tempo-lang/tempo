package process

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

type Func struct {
	Process *Process
	Name    string
}

func (f *Func) Codegen(file *jen.File) {
	// file.Commentf("Projection of function %s at role %s", f.Name, f.Process.Name)

	file.Func().Id(fmt.Sprintf("%s_%s", f.Name, f.Process.Name)).Params().Block()
}
