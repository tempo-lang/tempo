package codegen_ts

import (
	"fmt"
	"strings"

	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/type_check"
)

type codegen struct {
	indent int
	opts   Options
	info   *type_check.Info
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

func Codegen(info *type_check.Info, sourceFile *projection.SourceFile, opts *Options) string {
	gen := codegen{
		indent: 0,
		info:   info,
	}
	if opts != nil {
		gen.opts = *opts
	} else {
		gen.opts = DefaultOptions()
	}

	return gen.GenSourceFile(sourceFile)
}
