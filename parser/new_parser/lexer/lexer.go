package lexer

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"

	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

type Lexer struct {
	sourceReader io.RuneReader
	pos          token.SourcePos
	rune         rune
}

func New(sourceReader io.RuneReader) *Lexer {
	l := &Lexer{sourceReader: sourceReader, pos: token.SourcePos{Line: 1, Col: 0}}
	l.advance()
	return l
}

func FromString(source string) *Lexer {
	return New(strings.NewReader(source))
}

func (l *Lexer) advance() {
	ch, _, err := l.sourceReader.ReadRune()
	if err != nil {
		l.rune = 0 // io.EOF or error
		return
	}
	l.rune = ch
	if ch == '\n' {
		l.pos.Line++
		l.pos.Col = 0
	} else {
		l.pos.Col++
	}
}

func (l *Lexer) ReadToken() token.Token {
	// Skip whitespace
	for l.rune != 0 && isWhitespace(l.rune) {
		l.advance()
	}

	if l.rune == 0 {
		return token.New(token.EOF, "", l.pos, nil)
	}

	startPos := l.pos
	ch := l.rune

	switch ch {
	// Arithmetic
	case '+':
		l.advance()
		return token.New(token.PLUS, "+", startPos, nil)
	case '-':
		l.advance()
		if l.rune == '>' {
			l.advance()
			return token.New(token.COM, "->", startPos, nil)
		}
		return token.New(token.MINUS, "-", startPos, nil)
	case '*':
		l.advance()
		return token.New(token.MULTIPLY, "*", startPos, nil)
	case '/':
		l.advance()
		return token.New(token.DIVIDE, "/", startPos, nil)
	case '%':
		l.advance()
		return token.New(token.MODULO, "%", startPos, nil)

	// Parenthesis and brackets
	case '(':
		l.advance()
		return token.New(token.LPAREN, "(", startPos, nil)
	case ')':
		l.advance()
		return token.New(token.RPAREN, ")", startPos, nil)
	case '[':
		l.advance()
		return token.New(token.LSQUARE, "[", startPos, nil)
	case ']':
		l.advance()
		return token.New(token.RSQUARE, "]", startPos, nil)
	case '{':
		l.advance()
		return token.New(token.LCURLY, "{", startPos, nil)
	case '}':
		l.advance()
		return token.New(token.RCURLY, "}", startPos, nil)

	// Comparison
	case '=':
		l.advance()
		if l.rune == '=' {
			l.advance()
			return token.New(token.EQUAL, "==", startPos, nil)
		}
		return token.New(token.ASSIGN, "=", startPos, nil)
	case '!':
		l.advance()
		if l.rune == '=' {
			l.advance()
			return token.New(token.NOT_EQUAL, "!=", startPos, nil)
		}
		return token.Error(string(ch), startPos, "expected !=")
	case '<':
		l.advance()
		if l.rune == '=' {
			l.advance()
			return token.New(token.LESS_EQ, "<=", startPos, nil)
		}
		return token.New(token.LESS, "<", startPos, nil)
	case '>':
		l.advance()
		if l.rune == '=' {
			l.advance()
			return token.New(token.GREATER_EQ, ">=", startPos, nil)
		}
		return token.New(token.GREATER, ">", startPos, nil)

	// Logical
	case '&':
		l.advance()
		if l.rune == '&' {
			l.advance()
			return token.New(token.AND, "&&", startPos, nil)
		}
		return token.Error(string(ch), startPos, "unknown symbol `&`")
	case '|':
		l.advance()
		if l.rune == '|' {
			l.advance()
			return token.New(token.OR, "||", startPos, nil)
		}
		return token.Error(string(ch), startPos, "unknown symbol `|`")

	// Assignment and symbols
	case ',':
		l.advance()
		return token.New(token.COMMA, ",", startPos, nil)
	case '.':
		l.advance()
		if isDigit(l.rune) {
			return l.readNumber(true)
		}
		return token.New(token.DOT, ".", startPos, nil)
	case ':':
		l.advance()
		return token.New(token.COLON, ":", startPos, nil)
	case ';':
		l.advance()
		return token.New(token.SEMICOLON, ";", startPos, nil)
	case '@':
		l.advance()
		return token.New(token.ROLE_AT, "@", startPos, nil)
	case '_':
		return l.readIdentifier()
	case '"':
		return l.readStringLiteral()

	default:
		if unicode.IsLetter(ch) {
			return l.readIdentifier()
		}

		if isDigit(ch) {
			return l.readNumber(false)
		}

		l.advance()
		return token.Error(string(ch), startPos, "unexpected character")
	}
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func (l *Lexer) readStringLiteral() token.Token {
	var literalBuilder strings.Builder
	var valueBuilder strings.Builder

	literalBuilder.WriteRune(l.rune)
	startPos := l.pos
	l.advance() // consume opening quote

	for l.rune != 0 {
		literalBuilder.WriteRune(l.rune)
		if l.rune == '"' {
			l.advance()
			return token.New(token.STRING, literalBuilder.String(), startPos, valueBuilder.String())
		}
		if l.rune == '\\' {
			l.advance()
			switch l.rune {
			case '"':
				valueBuilder.WriteString(`\"`)
			case '\\':
				valueBuilder.WriteString(`\\`)
			case 'n':
				valueBuilder.WriteString("\n")
			case 't':
				valueBuilder.WriteString("\t")
			case '\n':
				// ignore escaped newline
			default:
				valueBuilder.WriteRune(l.rune)
			}

			literalBuilder.WriteRune(l.rune)
			l.advance()
			continue
		}
		valueBuilder.WriteRune(l.rune)
		l.advance()
	}

	return token.Error(literalBuilder.String(), startPos, "unexpected end of file, expected \"")
}

func (l *Lexer) readIdentifier() token.Token {
	startPos := l.pos

	var sb strings.Builder
	for l.rune != 0 && (unicode.IsLetter(l.rune) || unicode.IsNumber(l.rune) || l.rune == '_') {
		sb.WriteRune(l.rune)
		l.advance()
	}

	return token.New(token.IDENT, sb.String(), startPos, sb.String())
}

func (l *Lexer) readNumber(prefixDot bool) token.Token {
	startPos := l.pos
	containsDot := prefixDot
	var sb strings.Builder
	if prefixDot {
		startPos.Col -= 1
		sb.WriteRune('.')
	}

	for isDigit(l.rune) || (l.rune == '.' && !containsDot) {
		if l.rune == '.' {
			containsDot = true
		}
		sb.WriteRune(l.rune)
		l.advance()
	}

	literal := sb.String()
	if containsDot {
		value, err := strconv.ParseFloat(literal, 64)
		if err != nil {
			panic(fmt.Sprintf("lexer: literal should be a valid float: %v", err))
		}
		return token.New(token.FLOAT, literal, startPos, value)
	} else {
		value, err := strconv.ParseInt(literal, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("lexer: literal should be a valid int: %v", err))
		}
		return token.New(token.INT, literal, startPos, int(value))
	}
}
