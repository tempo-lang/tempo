package token

import (
	"sort"
	"unicode/utf8"
)

type Kind string
type TokenType = Kind
type Ref int

const NoRef Ref = -1

type Span struct{ Start, End int }

func (s Span) Empty() bool { return s.Start == s.End }

type SourcePos struct{ Line, Col int }
type LSPPos struct{ Line, Character int }

type Source struct {
	data  []byte
	lines []int
}

func NewSource(data []byte) *Source {
	s := &Source{data: append([]byte(nil), data...), lines: []int{0}}
	for i, b := range s.data {
		if b == '\n' {
			s.lines = append(s.lines, i+1)
		}
	}
	return s
}
func SourceFromString(text string) *Source { return NewSource([]byte(text)) }
func (s *Source) Bytes() []byte            { return append([]byte(nil), s.data...) }
func (s *Source) String() string           { return string(s.data) }
func (s *Source) Len() int                 { return len(s.data) }
func (s *Source) Text(span Span) string {
	if span.Start < 0 || span.End < span.Start || span.End > len(s.data) {
		return ""
	}
	return string(s.data[span.Start:span.End])
}
func (s *Source) Position(offset int) SourcePos {
	offset = clamp(offset, 0, len(s.data))
	line := sort.Search(len(s.lines), func(i int) bool { return s.lines[i] > offset }) - 1
	if line < 0 {
		line = 0
	}
	return SourcePos{Line: line + 1, Col: utf8.RuneCount(s.data[s.lines[line]:offset]) + 1}
}
func (s *Source) LSPPosition(offset int) LSPPos {
	offset = clamp(offset, 0, len(s.data))
	line := sort.Search(len(s.lines), func(i int) bool { return s.lines[i] > offset }) - 1
	if line < 0 {
		line = 0
	}
	units := 0
	for _, r := range string(s.data[s.lines[line]:offset]) {
		if r > 0xffff {
			units += 2
		} else {
			units++
		}
	}
	return LSPPos{Line: line, Character: units}
}
func clamp(n, lo, hi int) int {
	if n < lo {
		return lo
	}
	if n > hi {
		return hi
	}
	return n
}

type TriviaKind int

const (
	Whitespace TriviaKind = iota + 1
	LineComment
	BlockComment
)

type Trivia struct {
	Kind TriviaKind
	Text string
	Span Span
}
type TokenError struct{ Message string }

type Token struct {
	Index         Ref
	Kind          Kind
	Text          string
	Value         any
	Span          Span
	LeadingTrivia []Trivia
	Synthetic     bool
}

func NewAt(index Ref, kind Kind, text string, value any, span Span, trivia []Trivia, source *Source) Token {
	return Token{Index: index, Kind: kind, Text: text, Value: value, Span: span, LeadingTrivia: trivia}
}
func Missing(index Ref, kind Kind, offset int, source *Source) Token {
	t := NewAt(NoRef, kind, "", nil, Span{Start: offset, End: offset}, nil, source)
	t.Synthetic = true
	return t
}

const (
	BadToken   Kind = "BAD_TOKEN"
	EOF        Kind = "EOF"
	STRUCT     Kind = "STRUCT"
	INTERFACE  Kind = "INTERFACE"
	IMPLEMENTS Kind = "IMPLEMENTS"
	FUNC       Kind = "FUNC"
	RETURN     Kind = "RETURN"
	LET        Kind = "LET"
	ASYNC      Kind = "ASYNC"
	AWAIT      Kind = "AWAIT"
	IF         Kind = "IF"
	ELSE       Kind = "ELSE"
	WHILE      Kind = "WHILE"
	TRUE       Kind = "TRUE"
	FALSE      Kind = "FALSE"
	LPAREN     Kind = "LPAREN"
	RPAREN     Kind = "RPAREN"
	LSQUARE    Kind = "LSQUARE"
	RSQUARE    Kind = "RSQUARE"
	LCURLY     Kind = "LCURLY"
	RCURLY     Kind = "RCURLY"
	PLUS       Kind = "PLUS"
	MINUS      Kind = "MINUS"
	MULTIPLY   Kind = "MULTIPLY"
	DIVIDE     Kind = "DIVIDE"
	MODULO     Kind = "MODULO"
	EQUAL      Kind = "EQUAL"
	NOT_EQUAL  Kind = "NOT_EQUAL"
	LESS       Kind = "LESS"
	LESS_EQ    Kind = "LESS_EQ"
	GREATER    Kind = "GREATER"
	GREATER_EQ Kind = "GREATER_EQ"
	AND        Kind = "AND"
	OR         Kind = "OR"
	ASSIGN     Kind = "ASSIGN"
	ROLE_AT    Kind = "ROLE_AT"
	COMMA      Kind = "COMMA"
	DOT        Kind = "DOT"
	COLON      Kind = "COLON"
	SEMICOLON  Kind = "SEMICOLON"
	UNDERSCORE Kind = "UNDERSCORE"
	COM        Kind = "COM"
	STRING     Kind = "STRING"
	IDENT      Kind = "IDENT"
	FLOAT      Kind = "FLOAT"
	INT        Kind = "INT"
)
