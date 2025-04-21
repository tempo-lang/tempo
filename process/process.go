package process

import (
	"chorego/parser"

	"github.com/dave/jennifer/jen"
)

type Process struct {
	Name  string
	Funcs []*Func
}

func (p *Process) AddFunc(funcCtx parser.IFuncContext) *Func {
	fn := &Func{
		Process: p,
		FuncCtx: funcCtx,
		Name:    funcCtx.Ident().GetText(),
	}

	p.Funcs = append(p.Funcs, fn)
	return fn
}

func (p *Process) Codegen(file *jen.File) {
	file.Commentf("Projection of process %s", p.Name)

	for _, fn := range p.Funcs {
		fn.Codegen(file)
	}
}
