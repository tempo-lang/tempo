package type_error

import (
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fatih/color"
)

func FormatError(w io.Writer, inputStream *antlr.InputStream, sourceName string, err Error, colorOutput bool) {
	withColor := func(colorAttr color.Attribute, format string, args ...any) string {
		if colorOutput {
			return color.New(colorAttr).Sprintf(format, args...)
		} else {
			return fmt.Sprintf(format, args...)
		}
	}

	tokenStart := err.ParserRule().GetStart()
	tokenEnd := err.ParserRule().GetStop()

	// Swap start and end tokens if their indices are reversed
	if tokenStart.GetTokenIndex() > tokenEnd.GetTokenIndex() {
		tokenStart, tokenEnd = tokenEnd, tokenStart
	}

	sourceLines := strings.Split(inputStream.String(), "\n")
	line := sourceLines[tokenStart.GetLine()-1]

	lineNrSpace := int(math.Ceil(math.Log10(float64(tokenStart.GetLine() + 1))))
	lineNrStr := strings.Repeat(" ", lineNrSpace)

	tokenCol := tokenStart.GetColumn()

	errorLength := tokenEnd.GetStop() - tokenStart.GetStart() + 1
	if tokenStart.GetLine() != tokenEnd.GetLine() {
		errorLength = len(line) - tokenCol
	}

	fmt.Fprintf(w, "%s: %s\n", withColor(color.FgRed, "error[E%d]", err.Code()), withColor(color.Bold, "%s", err.Error()))
	fmt.Fprintf(w, "%s %s %s:%d:%d\n", lineNrStr, withColor(color.FgBlue, "->"), sourceName, tokenStart.GetLine(), tokenCol+1)

	if colorOutput {
		line = fmt.Sprintf("%s%s%s", line[0:tokenCol], withColor(color.FgRed, "%s", line[tokenCol:tokenCol+errorLength]), line[tokenCol+errorLength:])
	}

	fmt.Fprintf(w, "%s %s\n", withColor(color.FgBlue, "%s |\n%d |", lineNrStr, tokenStart.GetLine()), line)

	space := strings.Repeat(" ", tokenCol+1)
	highlight := strings.Repeat("^", errorLength)
	fmt.Fprintf(w, "%s%s%s\n\n", withColor(color.FgBlue, "%s |", lineNrStr), space, withColor(color.FgRed, "%s", highlight))

	for _, annotation := range err.Annotations() {
		fmt.Fprintf(w, "%s: %s\n\n", withColor(color.FgBlue, "%s", string(annotation.Type)), annotation.Message)
	}
}
