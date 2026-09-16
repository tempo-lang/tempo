package new_parser

import (
	"strings"
	"testing"

	"github.com/tempo-lang/tempo/parser/new_parser/ast"
	"github.com/tempo-lang/tempo/parser/new_parser/lexer"
	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

func reconstruct(ts []token.Token) string {
	var b strings.Builder
	for _, t := range ts {
		for _, v := range t.LeadingTrivia {
			b.WriteString(v.Text)
		}
		b.WriteString(t.Text)
	}
	return b.String()
}
func TestLexerLossless(t *testing.T) {
	cases := []string{"", " a\r\n// hi\n/* x */+b", "@#\x00", "\"x\\n\"", "/*unterminated"}
	for _, s := range cases {
		l := lexer.FromString(s)
		ts, _ := l.LexAll()
		if got := reconstruct(ts); got != s {
			t.Fatalf("%q reconstructed as %q", s, got)
		}
	}
}
func TestLexerTokenTable(t *testing.T) {
	src := "struct interface implements func return let async await if else while true false ( ) [ ] { } + - * / % == != < <= > >= && || = @ , . : ; _ -> name 1 1. .5"
	want := []token.Kind{token.STRUCT, token.INTERFACE, token.IMPLEMENTS, token.FUNC, token.RETURN, token.LET, token.ASYNC, token.AWAIT, token.IF, token.ELSE, token.WHILE, token.TRUE, token.FALSE, token.LPAREN, token.RPAREN, token.LSQUARE, token.RSQUARE, token.LCURLY, token.RCURLY, token.PLUS, token.MINUS, token.MULTIPLY, token.DIVIDE, token.MODULO, token.EQUAL, token.NOT_EQUAL, token.LESS, token.LESS_EQ, token.GREATER, token.GREATER_EQ, token.AND, token.OR, token.ASSIGN, token.ROLE_AT, token.COMMA, token.DOT, token.COLON, token.SEMICOLON, token.UNDERSCORE, token.COM, token.IDENT, token.INT, token.FLOAT, token.FLOAT, token.EOF}
	ts, ds := lexer.FromString(src).LexAll()
	if len(ds) != 0 {
		t.Fatal(ds)
	}
	if len(ts) != len(want) {
		t.Fatalf("%d tokens, want %d", len(ts), len(want))
	}
	for i, k := range want {
		if ts[i].Kind != k || int(ts[i].Index) != i {
			t.Fatalf("%d: %#v, want %s", i, ts[i], k)
		}
	}
}
func TestLexerDiagnostics(t *testing.T) {
	for _, src := range []string{"#", "\x00", "\"bad\\q\"", "\"unterminated", "/*unterminated", "999999999999999999999999999999999999"} {
		l := lexer.FromString(src)
		ts, ds := l.LexAll()
		if len(ds) == 0 {
			t.Errorf("%q: missing diagnostic", src)
		}
		if reconstruct(ts) != src {
			t.Errorf("%q: not lossless", src)
		}
	}
}
func TestKeywordAndIdentifierContract(t *testing.T) {
	ts, _ := lexer.FromString("letter _name letx å").LexAll()
	want := []token.Kind{token.IDENT, token.IDENT, token.IDENT, token.BadToken, token.EOF}
	for i, k := range want {
		if ts[i].Kind != k {
			t.Fatalf("%d: %s", i, ts[i].Kind)
		}
	}
}
func TestSourceCoordinates(t *testing.T) {
	s := token.SourceFromString("a😀\r\nb")
	if got := s.Position(5); got.Line != 1 || got.Col != 3 {
		t.Fatal(got)
	}
	if got := s.LSPPosition(5); got.Line != 0 || got.Character != 3 {
		t.Fatal(got)
	}
}
func TestCursorStableEOF(t *testing.T) {
	ts, _ := lexer.FromString("x").LexAll()
	c := NewCursor(ts)
	c.Advance()
	a := c.Current()
	c.Advance()
	b := c.Current()
	if a.Index != b.Index || a.Kind != token.EOF {
		t.Fatal(a, b)
	}
}
func TestExpressionSuites(t *testing.T) {
	cases := map[string]string{"1 + 2 * 3": "(binary + (int 1) (binary * (int 2) (int 3)))", "(1).x[0]": "(index (field (group (int 1) \")\") x) (int 0) \"]\")", "foo(1, 2)": "(call foo (int 1) (int 2) \")\")"}
	for src, want := range cases {
		r := ParseExpression(src)
		if got := ast.SExpr(r.Node); got != want {
			t.Errorf("%s: %s", src, got)
		}
		if len(r.Diagnostics) > 0 {
			t.Errorf("%s: %v", src, r.Diagnostics)
		}
	}
}

func TestSExprGeneratorFormatting(t *testing.T) {
	r := ParseExpression("1 + 2 * 3")

	compact := (ast.SExprGenerator{SingleLine: true}).Generate(r.Node)
	if compact != ast.SExpr(r.Node) {
		t.Fatalf("single-line output changed: %q", compact)
	}

	want := "(binary +\n--(int 1)\n--(binary *\n----(int 2)\n----(int 3)))"
	if got := (ast.SExprGenerator{Indent: "--"}).Generate(r.Node); got != want {
		t.Fatalf("multiline output did not match.\nWant:\n%s\nGot:\n%s", want, got)
	}
}

func TestSExprRoleAnnotationsRemainNested(t *testing.T) {
	r := ParseExpression("foo@A(1@A)")
	want := `(call (at foo (role-type (role A))) (at (int 1) (role-type (role A))) ")")`
	if got := ast.SExpr(r.Node); got != want {
		t.Fatalf("role annotations escaped their expression nodes.\nWant: %s\nGot:  %s", want, got)
	}
}

func TestRecoveryLocality(t *testing.T) {
	r := ParseScopeText("{let x = arr[0; let y = 2;}")
	got := ast.SExpr(r.Node)
	if !strings.Contains(got, "(missing RSQUARE)") || !strings.Contains(got, "(let y (int 2)") {
		t.Fatal(got)
	}
}
func TestRecoveryPreservesExpressionSentinel(t *testing.T) {
	r := ParseScopeText("{#; 42;}")
	s := r.Node.(*ast.Scope)
	if len(s.Stmts) != 1 || ast.SExpr(s.Stmts[0]) != "(expr (int 42) \";\")" {
		t.Fatal(ast.SExpr(s))
	}
}
func TestMissingSemicolonFix(t *testing.T) {
	r := ParseScopeText("{let x = 1 let y = 2;}")
	if len(r.Diagnostics) == 0 {
		t.Fatal("missing diagnostic")
	}
	d := r.Diagnostics[0]
	if d.Code != CodeMissingToken || d.Expected[0] != token.SEMICOLON || d.Found != token.LET || len(d.Fixes) != 1 || d.Fixes[0].NewText != ";" {
		t.Fatal(d)
	}
}
func TestASTInvariantsOnPartialResults(t *testing.T) {
	results := []ProductionResult{
		ParseExpression("foo[1 + ]"),
		ParseStatement("let = ;"),
		ParseType("async []"),
		ParseScopeText("{if { let x = ;"),
	}
	for _, r := range results {
		if err := ast.Validate(r.Node, r.Tokens, r.Source.Len()); err != nil {
			t.Errorf("%s: %v", ast.SExpr(r.Node, r.Tokens), err)
		}
	}
}
func TestASTInvariantsRejectNilRequiredChild(t *testing.T) {
	r := ParseExpression("1 + 2")
	binary := r.Node.(*ast.BinaryExpr)
	binary.Left = nil
	if err := ast.Validate(binary, r.Tokens, r.Source.Len()); err == nil || !strings.Contains(err.Error(), "root.Left is a nil required child") {
		t.Fatalf("got %v", err)
	}
}
func TestASTInvariantsRejectInvalidTokenReference(t *testing.T) {
	r := ParseExpression("1")
	primitive := r.Node.(*ast.PrimitiveExpr)
	literal := primitive.Literal.(*ast.IntLit)
	literal.IntToken.Index = token.Ref(len(r.Tokens) + 1)
	if err := ast.Validate(primitive, r.Tokens, r.Source.Len()); err == nil || !strings.Contains(err.Error(), "invalid token reference") {
		t.Fatalf("got %v", err)
	}
}
func TestASTInvariantsAllowOptionalChild(t *testing.T) {
	r := ParseStatement("return;")
	if err := ast.Validate(r.Node, r.Tokens, r.Source.Len()); err != nil {
		t.Fatal(err)
	}
}
func TestRoleRecovery(t *testing.T) {
	p := FromString("[A,123,B]")
	r := p.parseRoleType(TokenSet{token.EOF})
	p.expectEOF()
	if len(r.RoleNodes) != 3 || !r.RoleNodes[1].Invalid || len(p.Diagnostics()) != 1 {
		t.Fatal(ast.SExpr(r, p.Tokens()), p.Diagnostics())
	}
	if got := ast.SExpr(r, p.Tokens()); !strings.Contains(got, `(skipped INT "123")`) {
		t.Fatal(got)
	}
}
func FuzzLexer(f *testing.F) {
	f.Add("let x = 1;")
	f.Fuzz(func(t *testing.T, s string) {
		l := lexer.FromString(s)
		ts, _ := l.LexAll()
		if reconstruct(ts) != s {
			t.Fatal("not lossless")
		}
	})
}
func FuzzExpression(f *testing.F) {
	f.Add("a[0]+1")
	f.Fuzz(func(t *testing.T, s string) {
		r := ParseExpression(s)
		_ = ast.SExpr(r.Node)
		if reconstruct(r.Tokens) != s {
			t.Fatal("not lossless")
		}
	})
}
func FuzzStatement(f *testing.F) {
	f.Add("let x = 1;")
	f.Fuzz(func(t *testing.T, s string) {
		r := ParseStatement(s)
		_ = ast.SExpr(r.Node)
		if err := ast.Validate(r.Node, r.Tokens, len(s)); err != nil {
			t.Fatal(err)
		}
	})
}
func FuzzType(f *testing.F) {
	f.Add("async [T]")
	f.Fuzz(func(t *testing.T, s string) {
		r := ParseType(s)
		_ = ast.SExpr(r.Node)
		if err := ast.Validate(r.Node, r.Tokens, len(s)); err != nil {
			t.Fatal(err)
		}
	})
}
func FuzzScope(f *testing.F) {
	f.Add("{let x = 1;}")
	f.Fuzz(func(t *testing.T, s string) {
		r := ParseScopeText(s)
		_ = ast.SExpr(r.Node)
		if err := ast.Validate(r.Node, r.Tokens, len(s)); err != nil {
			t.Fatal(err)
		}
	})
}
func BenchmarkExpression(b *testing.B) {
	for b.Loop() {
		ParseExpression("a.b[1 + 2 * 3]")
	}
}
func BenchmarkScope(b *testing.B) {
	for b.Loop() {
		ParseScopeText("{let x = a.b[1 + 2 * 3]; return x;}")
	}
}
