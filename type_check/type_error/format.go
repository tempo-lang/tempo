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
	redColor := func(format string, args ...any) string {
		if colorOutput {
			return color.RedString(format, args...)
		} else {
			return fmt.Sprintf(format, args...)
		}
	}

	blueColor := func(format string, args ...any) string {
		if colorOutput {
			return color.BlueString(format, args...)
		} else {
			return fmt.Sprintf(format, args...)
		}
	}

	token := err.ParserRule().GetStart()
	tokenCol := token.GetColumn()

	lineNrSpace := int(math.Ceil(math.Log10(float64(token.GetLine()))))
	lineNrStr := strings.Repeat(" ", lineNrSpace)
	errorLength := err.ParserRule().GetStop().GetStop() - err.ParserRule().GetStart().GetStart() + 1

	fmt.Fprintf(w, "%s: %s\n", redColor("error"), err.Error())
	fmt.Fprintf(w, "%s %s %s:%d:%d\n", lineNrStr, blueColor("->"), inputStream.GetSourceName(), token.GetLine(), tokenCol+1)

	sourceLines := strings.Split(inputStream.String(), "\n")

	line := sourceLines[token.GetLine()-1]
	if colorOutput {
		line = fmt.Sprintf("%s%s%s", line[0:tokenCol], redColor(line[tokenCol:tokenCol+errorLength]), line[tokenCol+errorLength:])
	}

	fmt.Fprintf(w, "%s %s\n", blueColor("%s |\n%d |", lineNrStr, token.GetLine()), line)

	space := strings.Repeat(" ", tokenCol+1)
	highlight := strings.Repeat("^", errorLength)
	fmt.Fprintf(w, "%s%s%s\n\n", blueColor("%s |", lineNrStr), space, redColor(highlight))

	for _, annotation := range err.Annotations() {
		fmt.Fprintf(w, "%s: %s\n\n", blueColor(string(annotation.Type)), annotation.Message)
	}
}
