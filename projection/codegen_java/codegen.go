package codegen_java

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/type_check"
)

type javaPkg string

const (
	javaPkgTempoEnv          javaPkg = "tempo.runtime.Env"
	javaPkgArrayList         javaPkg = "java.util.ArrayList"
	javaPkgList              javaPkg = "java.util.List"
	javaPkgFuture            javaPkg = "java.util.concurrent.Future"
	javaPkgComplatableFuture javaPkg = "java.util.concurrent.CompletableFuture"
	javaPkgStream            javaPkg = "java.util.stream.Stream"
	javaPkgCollectors        javaPkg = "java.util.stream.Collectors"
)

type codegen struct {
	indent  int
	opts    Options
	info    *type_check.Info
	imports map[javaPkg]struct{}
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

func (gen *codegen) AddImport(pkg ...javaPkg) {
	for _, p := range pkg {
		gen.imports[p] = struct{}{}
	}
}

type Options struct {
	Package   *string
	ClassName string
}

func DefaultOptions() Options {
	return Options{
		Package:   nil,
		ClassName: "Choreography",
	}
}

func Codegen(info *type_check.Info, sourceFile *projection.SourceFile, opts *Options) string {
	gen := codegen{
		indent:  0,
		info:    info,
		imports: map[javaPkg]struct{}{},
	}
	if opts != nil {
		gen.opts = *opts
	} else {
		gen.opts = DefaultOptions()
	}

	code := gen.GenSourceFile(sourceFile)

	out := ""
	for _, pkg := range slices.Sorted(maps.Keys(gen.imports)) {
		out += gen.Writeln("import %s;", pkg)
	}
	out += gen.Writeln("")

	out += code

	return out
}
