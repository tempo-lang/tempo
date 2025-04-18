package process

import "github.com/dave/jennifer/jen"

type Process struct {
	Name  string
	Funcs []*Func
}

func (p *Process) AddFunc(name string) *Func {
	fn := &Func{
		Process: p,
		Name:    name,
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
