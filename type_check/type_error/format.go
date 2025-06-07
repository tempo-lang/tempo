package type_error

import (
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fatih/color"
)

func FormatError(w io.Writer, inputStream *antlr.FileStream, err Error, colorOutput bool) {
	withColor := func(colorAttr color.Attribute, format string, args ...any) string {
		if colorOutput {
			return color.New(colorAttr).Sprintf(format, args...)
		} else {
			return fmt.Sprintf(format, args...)
		}
	}

	token := err.ParserRule().GetStart()
	tokenCol := token.GetColumn()

	lineNrSpace := int(math.Ceil(math.Log10(float64(token.GetLine()))))
	lineNrStr := strings.Repeat(" ", lineNrSpace)
	errorLength := err.ParserRule().GetStop().GetStop() - err.ParserRule().GetStart().GetStart() + 1

	fmt.Fprintf(w, "%s: %s\n", withColor(color.FgRed, "error[E%d]", err.Code()), withColor(color.Bold, err.Error()))
	fmt.Fprintf(w, "%s %s %s:%d:%d\n", lineNrStr, withColor(color.FgBlue, "->"), inputStream.GetSourceName(), token.GetLine(), tokenCol+1)

	sourceLines := strings.Split(inputStream.String(), "\n")

	line := sourceLines[token.GetLine()-1]
	if colorOutput {
		line = fmt.Sprintf("%s%s%s", line[0:tokenCol], withColor(color.FgRed, line[tokenCol:tokenCol+errorLength]), line[tokenCol+errorLength:])
	}

	fmt.Fprintf(w, "%s %s\n", withColor(color.FgBlue, "%s |\n%d |", lineNrStr, token.GetLine()), line)

	space := strings.Repeat(" ", tokenCol+1)
	highlight := strings.Repeat("^", errorLength)
	fmt.Fprintf(w, "%s%s%s\n\n", withColor(color.FgBlue, "%s |", lineNrStr), space, withColor(color.FgRed, highlight))

	for _, annotation := range err.Annotations() {
		fmt.Fprintf(w, "%s: %s\n\n", withColor(color.FgBlue, string(annotation.Type)), annotation.Message)
	}
}
