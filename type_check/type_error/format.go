package type_error

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
	"io"
	"strings"
)

// FormatError renders a semantic diagnostic from the AST's byte span.
func FormatError(w io.Writer, source *token.Source, sourceName string, err Error, colorOutput bool) {
	n := err.ParserRule()
	span := token.Span{Start: n.StartToken().Span.Start, End: n.EndToken().Span.End}
	start, end := source.Position(span.Start), source.Position(span.End)
	lines := strings.Split(source.String(), "\n")
	line := ""
	if start.Line > 0 && start.Line <= len(lines) {
		line = lines[start.Line-1]
	}
	length := end.Col - start.Col
	if end.Line != start.Line || length < 1 {
		length = 1
	}
	red := func(s string, a ...any) string {
		if colorOutput {
			return color.RedString(s, a...)
		}
		return fmt.Sprintf(s, a...)
	}
	fmt.Fprintf(w, "%s: %s\n -> %s:%d:%d\n%d | %s\n  | %s%s\n\n", red("error[E%d]", err.Code()), err.Error(), sourceName, start.Line, start.Col, start.Line, line, strings.Repeat(" ", max(start.Col-1, 0)), red("%s", strings.Repeat("^", length)))
	for _, a := range err.Annotations() {
		fmt.Fprintf(w, "%s: %s\n\n", a.Type, a.Message)
	}
}
