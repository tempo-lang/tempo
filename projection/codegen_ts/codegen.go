package codegen_ts

import (
	"fmt"
	"strings"

	"github.com/tempo-lang/tempo/projection"
)

type codegen struct {
	indent int
	opts   Options
}

func (gen *codegen) WriteIndent() string {
	return strings.Repeat("  ", gen.indent)
}

func (gen *codegen) Writeln(format string, args ...any) string {
	return fmt.Sprintf(gen.WriteIndent()+format+"\n", args...)
}

func (gen *codegen) IncIndent() {
	gen.indent += 1
}

func (gen *codegen) DecIndent() {
	gen.indent -= 1
}

type Options struct {
	DisableTypes bool
	RuntimePath  string
}

func DefaultOptions() Options {
	return Options{
		DisableTypes: false,
		RuntimePath:  "@tempo-lang/tempo/runtime",
	}
}

func Codegen(sourceFile *projection.SourceFile, opts *Options) string {
	gen := codegen{
		indent: 0,
	}
	if opts != nil {
		gen.opts = *opts
	} else {
		gen.opts = DefaultOptions()
	}

	return gen.GenSourceFile(sourceFile)
}
