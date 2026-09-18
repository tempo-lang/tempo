package lexer

import (
	"io"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/tempo-lang/tempo/parser/token"
)

type Diagnostic struct {
	Message string
	Span    token.Span
}

type Lexer struct {
	source      *token.Source
	data        []byte
	offset      int
	next        token.Ref
	diagnostics []Diagnostic
	eof         *token.Token
}

func New(r io.RuneReader) *Lexer {
	var b strings.Builder
	for {
		x, _, e := r.ReadRune()
		if e != nil {
			l := FromSource(token.SourceFromString(b.String()))
			if e != io.EOF {
				l.diagnostics = append(l.diagnostics, Diagnostic{Message: e.Error(), Span: token.Span{Start: l.source.Len(), End: l.source.Len()}})
			}
			return l
		}
		b.WriteRune(x)
	}
}

func FromString(s string) *Lexer        { return FromSource(token.SourceFromString(s)) }
func FromBytes(b []byte) *Lexer         { return FromSource(token.NewSource(b)) }
func FromSource(s *token.Source) *Lexer { return &Lexer{source: s, data: s.Bytes()} }

func (l *Lexer) Source() *token.Source     { return l.source }
func (l *Lexer) Diagnostics() []Diagnostic { return append([]Diagnostic(nil), l.diagnostics...) }

func (l *Lexer) LexAll() ([]token.Token, []Diagnostic) {
	var ts []token.Token
	for {
		t := l.ReadToken()
		ts = append(ts, t)
		if t.Kind == token.EOF {
			break
		}
	}
	return ts, l.Diagnostics()
}

func (l *Lexer) ReadToken() token.Token {
	tr := l.trivia()
	if l.offset == len(l.data) {
		if l.eof != nil {
			return *l.eof
		}
		t := token.NewAt(l.next, token.EOF, "", nil, token.Span{Start: l.offset, End: l.offset}, tr, l.source)
		l.eof = &t
		return t
	}
	s := l.offset
	c := l.data[l.offset]
	l.offset++
	if l.offset < len(l.data) {
		if k, ok := doubles[string(l.data[s:l.offset+1])]; ok {
			l.offset++
			return l.emit(k, s, nil, tr)
		}
	}
	if c == '.' && l.offset < len(l.data) && digit(l.data[l.offset]) {
		return l.number(s, tr)
	}
	if k, ok := singles[c]; ok {
		return l.emit(k, s, nil, tr)
	}
	if c == '"' {
		return l.string(s, tr)
	}
	if digit(c) {
		return l.number(s, tr)
	}
	if letter(c) || c == '_' {
		return l.ident(s, tr)
	}
	if c == 0 {
		return l.bad(s, "embedded NUL", tr)
	}
	if c >= utf8.RuneSelf {
		_, n := utf8.DecodeRune(l.data[s:])
		l.offset = s + n
	}
	return l.bad(s, "unexpected character", tr)
}

func (l *Lexer) trivia() []token.Trivia {
	var out []token.Trivia
	for l.offset < len(l.data) {
		s := l.offset
		k := token.Whitespace
		switch {
		case strings.HasPrefix(string(l.data[l.offset:]), "//"):
			k = token.LineComment
			l.offset += 2
			for l.offset < len(l.data) && l.data[l.offset] != '\r' && l.data[l.offset] != '\n' {
				l.offset++
			}
		case strings.HasPrefix(string(l.data[l.offset:]), "/*"):
			k = token.BlockComment
			l.offset += 2
			for l.offset < len(l.data) && !strings.HasPrefix(string(l.data[l.offset:]), "*/") {
				l.offset++
			}
			if l.offset < len(l.data) {
				l.offset += 2
			} else {
				l.diagnostics = append(l.diagnostics, Diagnostic{"unterminated block comment", token.Span{Start: s, End: l.offset}})
			}
		case ws(l.data[l.offset]):
			for l.offset < len(l.data) && ws(l.data[l.offset]) {
				l.offset++
			}
		default:
			return out
		}
		out = append(out, token.Trivia{Kind: k, Text: string(l.data[s:l.offset]), Span: token.Span{Start: s, End: l.offset}})
	}
	return out
}

func (l *Lexer) ident(s int, tr []token.Trivia) token.Token {
	for l.offset < len(l.data) && (letter(l.data[l.offset]) || digit(l.data[l.offset]) || l.data[l.offset] == '_') {
		l.offset++
	}
	x := string(l.data[s:l.offset])
	if x == "_" {
		return l.emit(token.UNDERSCORE, s, nil, tr)
	}
	if k, ok := keywords[x]; ok {
		var v any
		if k == token.TRUE {
			v = true
		}
		if k == token.FALSE {
			v = false
		}
		return l.emit(k, s, v, tr)
	}
	return l.emit(token.IDENT, s, x, tr)
}

func (l *Lexer) number(s int, tr []token.Trivia) token.Token {
	dot := l.data[s] == '.'
	for l.offset < len(l.data) && digit(l.data[l.offset]) {
		l.offset++
	}
	if !dot && l.offset < len(l.data) && l.data[l.offset] == '.' {
		dot = true
		l.offset++
		for l.offset < len(l.data) && digit(l.data[l.offset]) {
			l.offset++
		}
	}
	x := string(l.data[s:l.offset])
	if dot {
		v, e := strconv.ParseFloat(x, 64)
		if e != nil {
			l.diagnostics = append(l.diagnostics, Diagnostic{"numeric overflow", token.Span{Start: s, End: l.offset}})
			return l.emit(token.FLOAT, s, nil, tr)
		}
		return l.emit(token.FLOAT, s, v, tr)
	}
	v, e := strconv.ParseInt(x, 10, 64)
	if e != nil {
		l.diagnostics = append(l.diagnostics, Diagnostic{"numeric overflow", token.Span{Start: s, End: l.offset}})
		return l.emit(token.INT, s, nil, tr)
	}
	return l.emit(token.INT, s, int(v), tr)
}

func (l *Lexer) string(s int, tr []token.Trivia) token.Token {
	var v strings.Builder
	for l.offset < len(l.data) {
		c := l.data[l.offset]
		l.offset++
		if c == '"' {
			return l.emit(token.STRING, s, v.String(), tr)
		}
		if c == '\\' {
			if l.offset == len(l.data) {
				break
			}
			e := l.data[l.offset]
			l.offset++
			switch e {
			case 'n':
				v.WriteByte('\n')
			case 'r':
				v.WriteByte('\r')
			case 't':
				v.WriteByte('\t')
			case '"':
				v.WriteByte('"')
			case '\\':
				v.WriteByte('\\')
			case '\n':
			default:
				return l.bad(s, "invalid string escape", tr)
			}
			continue
		}
		if c < 0x20 {
			return l.bad(s, "invalid control character in string", tr)
		}
		v.WriteByte(c)
	}
	return l.bad(s, "unterminated string", tr)
}

func (l *Lexer) emit(k token.Kind, s int, v any, tr []token.Trivia) token.Token {
	t := token.NewAt(l.next, k, string(l.data[s:l.offset]), v, token.Span{Start: s, End: l.offset}, tr, l.source)
	l.next++
	return t
}

func (l *Lexer) bad(s int, m string, tr []token.Trivia) token.Token {
	l.diagnostics = append(l.diagnostics, Diagnostic{m, token.Span{Start: s, End: l.offset}})
	return l.emit(token.BadToken, s, token.TokenError{Message: m}, tr)
}

func ws(c byte) bool     { return c == ' ' || c == '\t' || c == '\r' || c == '\n' }
func digit(c byte) bool  { return c >= '0' && c <= '9' }
func letter(c byte) bool { return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' }

var doubles = map[string]token.Kind{
	"->": token.COM,
	"==": token.EQUAL,
	"!=": token.NOT_EQUAL,
	"<=": token.LESS_EQ,
	">=": token.GREATER_EQ,
	"&&": token.AND,
	"||": token.OR,
}

var singles = map[byte]token.Kind{'+': token.PLUS,
	'-': token.MINUS,
	'*': token.MULTIPLY,
	'/': token.DIVIDE,
	'%': token.MODULO,
	'(': token.LPAREN,
	')': token.RPAREN,
	'[': token.LSQUARE,
	']': token.RSQUARE,
	'{': token.LCURLY,
	'}': token.RCURLY,
	'=': token.ASSIGN,
	'<': token.LESS,
	'>': token.GREATER,
	',': token.COMMA,
	'.': token.DOT,
	':': token.COLON,
	';': token.SEMICOLON,
	'@': token.ROLE_AT,
}

var keywords = map[string]token.Kind{
	"struct":     token.STRUCT,
	"interface":  token.INTERFACE,
	"implements": token.IMPLEMENTS,
	"func":       token.FUNC,
	"return":     token.RETURN,
	"let":        token.LET,
	"async":      token.ASYNC,
	"await":      token.AWAIT,
	"if":         token.IF,
	"else":       token.ELSE,
	"while":      token.WHILE,
	"true":       token.TRUE,
	"false":      token.FALSE,
}
