package new_parser

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/lexer"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

type Span = token.Span
type Source = token.Source

func NewSource(data []byte) *Source { return token.NewSource(data) }

type ErrorCode int

const (
	CodeMissingToken ErrorCode = iota + 1
	CodeExpectedRole
	CodeUnexpectedToken
	CodeExpectedExpression
	CodeExpectedIdentifier
	CodeInvalidAssignmentTarget
	CodeLexical
	CodeTrailingInput
	CodeInternalRecovery
)

func (c ErrorCode) Code() int      { return int(c) }
func (c ErrorCode) String() string { return fmt.Sprintf("error[E%d]", c) }

type TextEdit struct {
	Span    Span
	NewText string
}
type Diagnostic struct {
	Code        ErrorCode
	Message     string
	PrimarySpan Span
	Expected    []token.Kind
	Found       token.Kind
	Fixes       []TextEdit
}

func (d Diagnostic) String() string {
	return fmt.Sprintf("%s %d:%d found=%s expected=%v fixes=%v", d.Code, d.PrimarySpan.Start, d.PrimarySpan.End, d.Found, d.Expected, d.Fixes)
}

type TokenSet []token.Kind

func (s TokenSet) Contains(k token.Kind) bool {
	return slices.Contains(s, k)
}
func (s TokenSet) With(k token.Kind) TokenSet {
	if s.Contains(k) {
		return append(TokenSet(nil), s...)
	}
	r := append(TokenSet(nil), s...)
	return append(r, k)
}
func (s TokenSet) Union(o TokenSet) TokenSet {
	r := append(TokenSet(nil), s...)
	for _, k := range o {
		if !r.Contains(k) {
			r = append(r, k)
		}
	}
	return r
}

type Cursor struct {
	tokens []token.Token
	pos    int
}

func NewCursor(tokens []token.Token) *Cursor { return &Cursor{tokens: tokens} }
func (c *Cursor) Position() int              { return c.pos }
func (c *Cursor) Current() token.Token       { return c.Peek(0) }
func (c *Cursor) Peek(n int) token.Token {
	if len(c.tokens) == 0 {
		return token.Token{Kind: token.EOF}
	}
	i := max(c.pos+n, 0)
	if i >= len(c.tokens) {
		i = len(c.tokens) - 1
	}
	return c.tokens[i]
}

func (c *Cursor) Advance() token.Ref {
	t := c.Current()
	if t.Kind != token.EOF {
		c.pos++
	}
	return t.Index
}

type Result struct {
	Root        *ast.SourceFile
	Source      *token.Source
	Tokens      []token.Token
	Diagnostics []Diagnostic
}

type ProductionResult struct {
	Node        ast.Node
	Source      *token.Source
	Tokens      []token.Token
	Diagnostics []Diagnostic
	Remaining   token.Kind
}

type Parser struct {
	source      *token.Source
	tokens      []token.Token
	cursor      *Cursor
	diagnostics []Diagnostic
	previous    token.Token
}

func New(r io.RuneReader) *Parser {
	l := lexer.New(r)
	return fromLexer(l)
}

func FromString(s string) *Parser { return FromSource(token.SourceFromString(s)) }

func FromSource(s *token.Source) *Parser {
	return fromLexer(lexer.FromSource(s))
}

func fromLexer(l *lexer.Lexer) *Parser {
	ts, lds := l.LexAll()
	p := &Parser{source: l.Source(), tokens: ts, cursor: NewCursor(ts)}
	for _, d := range lds {
		p.add(Diagnostic{Code: CodeLexical, Message: d.Message, PrimarySpan: d.Span, Found: token.BadToken})
	}
	return p
}

func (p *Parser) Tokens() []token.Token     { return append([]token.Token(nil), p.tokens...) }
func (p *Parser) Diagnostics() []Diagnostic { return append([]Diagnostic(nil), p.diagnostics...) }

func (p *Parser) current() token.Token   { return p.cursor.Current() }
func (p *Parser) peek(n int) token.Token { return p.cursor.Peek(n) }
func (p *Parser) advance() token.Token {
	t := p.current()
	p.cursor.Advance()
	p.previous = t
	return t
}

func (p *Parser) curTokenIs(k ...token.Kind) bool {
	for _, x := range k {
		if p.current().Kind == x {
			return true
		}
	}
	return false
}

func (p *Parser) add(d Diagnostic) {
	for _, x := range p.diagnostics {
		if x.Code == d.Code && x.PrimarySpan == d.PrimarySpan && first(x.Expected) == first(d.Expected) {
			return
		}
	}
	p.diagnostics = append(p.diagnostics, d)
}

func (p *Parser) trailingComma(close token.Kind) {
	p.add(Diagnostic{Code: CodeUnexpectedToken, Message: "trailing comma", PrimarySpan: p.current().Span, Expected: []token.Kind{token.IDENT}, Found: close})
}

func first(s []token.Kind) token.Kind {
	if len(s) > 0 {
		return s[0]
	}
	return ""
}

func spelling(k token.Kind) string {
	return map[token.Kind]string{token.SEMICOLON: ";", token.RSQUARE: "]", token.RPAREN: ")", token.RCURLY: "}", token.LCURLY: "{", token.ASSIGN: "=", token.IDENT: "identifier"}[k]
}

func (p *Parser) missing(k token.Kind, at int, primary Span, found token.Kind) token.Token {
	p.add(Diagnostic{Code: CodeMissingToken, Message: "missing " + string(k), PrimarySpan: primary, Expected: []token.Kind{k}, Found: found, Fixes: []TextEdit{{Span: Span{Start: at, End: at}, NewText: spelling(k)}}})
	return token.Missing(token.Ref(len(p.tokens)), k, at, p.source)
}

func (p *Parser) expect(k token.Kind, follow TokenSet) token.Token {
	if p.current().Kind == k {
		return p.advance()
	}
	if p.peek(1).Kind == k {
		bad := p.advance()
		p.add(Diagnostic{Code: CodeUnexpectedToken, Message: "unexpected token", PrimarySpan: bad.Span, Expected: []token.Kind{k}, Found: bad.Kind, Fixes: []TextEdit{{Span: bad.Span, NewText: ""}}})
		return p.advance()
	}
	cur := p.current()
	if follow.Contains(cur.Kind) || cur.Kind == token.EOF || closing(cur.Kind) {
		return p.missing(k, cur.Span.Start, cur.Span, cur.Kind)
	}
	p.add(Diagnostic{Code: CodeUnexpectedToken, Message: "unexpected token", PrimarySpan: cur.Span, Expected: []token.Kind{k}, Found: cur.Kind})
	p.skipUntil(follow.With(k))
	if p.current().Kind == k {
		return p.advance()
	}
	return token.Missing(token.Ref(len(p.tokens)), k, p.current().Span.Start, p.source)
}

func closing(k token.Kind) bool { return k == token.RPAREN || k == token.RSQUARE || k == token.RCURLY }

func (p *Parser) skipUntil(stop TokenSet) []token.Ref {
	var out []token.Ref
	var stack []token.Kind
	for p.current().Kind != token.EOF {
		kind := p.current().Kind
		if len(stack) == 0 && (stop.Contains(kind) || closing(kind)) {
			break
		}
		switch kind {
		case token.LPAREN:
			stack = append(stack, token.RPAREN)
		case token.LSQUARE:
			stack = append(stack, token.RSQUARE)
		case token.LCURLY:
			stack = append(stack, token.RCURLY)
		default:
			if len(stack) > 0 && kind == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
		}
		out = append(out, p.current().Index)
		p.advance()
	}
	return out
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	if p.current().Kind == token.IDENT {
		return &ast.Identifier{Token: p.advance()}
	}
	cur := p.current()
	p.add(Diagnostic{Code: CodeExpectedIdentifier, Message: "expected identifier", PrimarySpan: cur.Span, Expected: []token.Kind{token.IDENT}, Found: cur.Kind})
	return &ast.Identifier{Token: token.Missing(token.Ref(len(p.tokens)), token.IDENT, cur.Span.Start, p.source), Invalid: true}
}

func (p *Parser) expectEOF() {
	if p.current().Kind != token.EOF {
		cur := p.current()
		p.add(Diagnostic{Code: CodeTrailingInput, Message: "trailing input", PrimarySpan: cur.Span, Expected: []token.Kind{token.EOF}, Found: cur.Kind})
		p.skipUntil(TokenSet{token.EOF})
	}
}

func (p *Parser) production(node ast.Node) ProductionResult {
	p.expectEOF()
	return ProductionResult{Node: node, Source: p.source, Tokens: p.Tokens(), Diagnostics: p.Diagnostics(), Remaining: p.current().Kind}
}

func Parse(text string) Result {
	p := FromString(text)
	root := &ast.SourceFile{Tokens: p.Tokens()}
	if len(p.tokens) > 0 {
		root.FirstToken = p.tokens[0]
		root.LastToken = p.tokens[len(p.tokens)-1]
	}
	for p.current().Kind != token.EOF {
		pos := p.cursor.Position()
		switch p.current().Kind {
		case token.FUNC:
			d := p.parseFunc()
			root.Functions = append(root.Functions, d)
			root.Declarations = append(root.Declarations, d)
		case token.STRUCT:
			d := p.parseStruct()
			root.Structs = append(root.Structs, d)
			root.Declarations = append(root.Declarations, d)
		case token.INTERFACE:
			d := p.parseInterface()
			root.Interfaces = append(root.Interfaces, d)
			root.Declarations = append(root.Declarations, d)
		default:
			cur := p.current()
			p.add(Diagnostic{Code: CodeUnexpectedToken, Message: "expected declaration", PrimarySpan: cur.Span, Expected: []token.Kind{token.FUNC, token.STRUCT, token.INTERFACE}, Found: cur.Kind})
			p.skipUntil(declarationStart)
		}
		if p.cursor.Position() == pos {
			p.advance()
		}
	}
	return Result{Root: root, Source: p.source, Tokens: p.Tokens(), Diagnostics: p.Diagnostics()}
}

func FormatDiagnostics(ds []Diagnostic) string {
	var b strings.Builder
	for i, d := range ds {
		if i > 0 {
			b.WriteByte('\n')
		}
		b.WriteString(d.String())
	}
	return b.String()
}

func ParseExpression(s string) ProductionResult {
	p := FromString(s)
	return p.production(p.parseExpr(TokenSet{token.EOF}))
}

func ParseStatement(s string) ProductionResult {
	p := FromString(s)
	return p.production(p.parseStmt())
}

func ParseScopeText(s string) ProductionResult {
	p := FromString(s)
	return p.production(p.parseScope())
}
