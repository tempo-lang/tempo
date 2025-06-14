package codegen_ts

import (
	"fmt"
	"io"
	"strings"

	"github.com/tempo-lang/tempo/projection"
)

type codegen struct {
	indent int
	opts   Options
	w      io.Writer
}

func (gen *codegen) Writeln(format string, args ...any) {
	repeat := strings.Repeat("  ", gen.indent)
	fmt.Fprintf(gen.w, repeat+format+"\n", args...)
}

func (gen *codegen) IncIndent() {
	gen.indent += 1
}

func (gen *codegen) DecIndent() {
	gen.indent -= 1
}

type Options struct {
	DisableTypes bool
}

func Codegen(w io.Writer, sourceFile *projection.SourceFile, opts *Options) {
	gen := codegen{
		indent: 0,
		w:      w,
	}
	if opts != nil {
		gen.opts = *opts
	}

	gen.GenSourceFile(sourceFile)
}
