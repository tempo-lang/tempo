package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

// SExprGenerator controls how AST S-expressions are rendered.
//
// Indent is used for each nesting level in multiline output. An empty Indent
// defaults to two spaces. SingleLine preserves the compact representation used
// by SExpr.
type SExprGenerator struct {
	Indent     string
	SingleLine bool
}

// SExpr returns the compact, single-line representation kept for compatibility
// with existing assertions.
func SExpr(n Node, streams ...[]token.Token) string {
	return (SExprGenerator{SingleLine: true}).Generate(n, streams...)
}

// Generate returns an S-expression using the generator's formatting settings.
func (g SExprGenerator) Generate(n Node, streams ...[]token.Token) string {
	var b sexprWriter
	if len(streams) > 0 {
		b.tokens = streams[0]
	}
	writeNode(&b, n)
	compact := b.String()
	if g.SingleLine {
		return compact
	}
	indent := g.Indent
	if indent == "" {
		indent = "  "
	}
	return indentSExpr(compact, indent)
}

type sexprWriter struct {
	strings.Builder
	tokens []token.Token
}

func tok(t token.Token) string {
	if t.Synthetic {
		return "(missing " + string(t.Kind) + ")"
	}
	return strconv.Quote(t.Text)
}
func writeNode(b *sexprWriter, n Node) {
	if n == nil {
		b.WriteString("nil")
		return
	}
	switch x := n.(type) {
	case *SourceFile:
		b.WriteString("(source-file")
		for _, d := range x.Declarations {
			b.WriteByte(' ')
			writeNode(b, d)
		}
		b.WriteByte(')')
	case *Identifier:
		if x.Invalid {
			b.WriteString("(invalid-ident)")
		} else {
			b.WriteString(x.Value())
		}
	case *IdentAccessExpr:
		writeNode(b, x.Ident)
		if x.RoleType != nil {
			b.WriteByte('@')
			writeNode(b, x.RoleType)
		}
	case *RoleIdent:
		writeNode(b, x.Ident)
		if x.RoleType != nil {
			b.WriteByte('@')
			writeNode(b, x.RoleType)
		}
	case *StructExpr:
		b.WriteString("(struct-expr ")
		writeNode(b, x.RoleIdent)
		for _, f := range x.StructFields.Fields {
			b.WriteByte(' ')
			writeNode(b, f)
		}
		b.WriteByte(' ')
		b.WriteString(tok(x.StructFields.CloseToken))
		b.WriteByte(')')
	case *StructFieldExpr:
		b.WriteString("(field-init ")
		writeNode(b, x.Name)
		b.WriteByte(' ')
		writeNode(b, x.Expr)
		b.WriteByte(')')
	case *ComExpr:
		b.WriteString("(com ")
		writeNode(b, x.Sender)
		b.WriteByte(' ')
		writeNode(b, x.Receiver)
		b.WriteByte(' ')
		writeNode(b, x.Expr)
		b.WriteByte(')')
	case *ClosureExpr:
		b.WriteString("(closure ")
		writeNode(b, x.ClosureSig)
		b.WriteByte(' ')
		writeNode(b, x.Scope)
		b.WriteByte(')')
	case *ClosureSig:
		b.WriteString("(closure-sig ")
		writeNode(b, x.RoleType)
		b.WriteByte(' ')
		writeNode(b, x.Params)
		if x.ReturnType != nil {
			b.WriteByte(' ')
			writeNode(b, x.ReturnType)
		}
		b.WriteByte(')')
	case *PrimitiveExpr:
		switch v := x.Literal.(type) {
		case *IntLit:
			fmt.Fprintf(b, "(int %v)", v.Value())
		case *FloatLit:
			fmt.Fprintf(b, "(float %v)", v.Value())
		case *StringLit:
			fmt.Fprintf(b, "(string %q)", v.Value())
		case *BoolLit:
			fmt.Fprintf(b, "(bool %v)", v.Value())
		}
		if x.RoleType != nil {
			b.WriteString("@")
			writeNode(b, x.RoleType)
		}
	case *BinaryExpr:
		b.WriteString("(binary ")
		b.WriteString(x.Operator.Text)
		b.WriteString(" ")
		writeNode(b, x.Left)
		b.WriteByte(' ')
		writeNode(b, x.Right)
		b.WriteByte(')')
	case *FieldAccessExpr:
		b.WriteString("(field ")
		writeNode(b, x.Object)
		b.WriteByte(' ')
		writeNode(b, x.Field)
		b.WriteByte(')')
	case *IndexExpr:
		b.WriteString("(index ")
		writeNode(b, x.Object)
		b.WriteByte(' ')
		writeNode(b, x.Index)
		b.WriteByte(' ')
		b.WriteString(tok(x.CloseBracket))
		b.WriteByte(')')
	case *GroupExpr:
		b.WriteString("(group ")
		writeNode(b, x.Expr)
		b.WriteByte(' ')
		b.WriteString(tok(x.CloseParen))
		b.WriteByte(')')
	case *ListExpr:
		b.WriteString("(list")
		for _, e := range x.Elements {
			b.WriteByte(' ')
			writeNode(b, e)
		}
		b.WriteByte(' ')
		b.WriteString(tok(x.CloseBracket))
		b.WriteByte(')')
	case *CallExpr:
		b.WriteString("(call ")
		writeNode(b, x.Function)
		for _, a := range x.Args {
			b.WriteByte(' ')
			writeNode(b, a)
		}
		b.WriteByte(' ')
		b.WriteString(tok(x.CloseParen))
		b.WriteByte(')')
	case *AwaitExpr:
		b.WriteString("(await ")
		writeNode(b, x.Expr)
		b.WriteByte(')')
	case *InvalidExpr:
		b.WriteString("(invalid-expr")
		for _, r := range x.Skipped {
			writeSkipped(b, r)
		}
		b.WriteByte(')')
	case *LetStmt:
		b.WriteString("(let ")
		writeNode(b, x.Name)
		if x.Type != nil {
			b.WriteByte(' ')
			b.WriteString("(type ")
			writeNode(b, x.Type)
			b.WriteByte(')')
		}
		b.WriteByte(' ')
		writeNode(b, x.Expr)
		b.WriteByte(' ')
		b.WriteString(tok(x.SemiToken))
		b.WriteByte(')')
	case *ReturnStmt:
		b.WriteString("(return")
		if x.Expr != nil {
			b.WriteByte(' ')
			writeNode(b, x.Expr)
		}
		b.WriteByte(' ')
		b.WriteString(tok(x.SemiToken))
		b.WriteByte(')')
	case *AssignStmt:
		b.WriteString("(assign ")
		writeNode(b, x.LHS)
		b.WriteByte(' ')
		writeNode(b, x.RHS)
		b.WriteByte(' ')
		b.WriteString(tok(x.SemiToken))
		b.WriteByte(')')
	case *ExprStmt:
		b.WriteString("(expr ")
		writeNode(b, x.Expr)
		b.WriteByte(' ')
		b.WriteString(tok(x.SemiToken))
		b.WriteByte(')')
	case *IfStmt:
		b.WriteString("(if ")
		writeNode(b, x.Condition)
		b.WriteByte(' ')
		writeNode(b, x.ThenScope)
		if x.ElseScope != nil {
			b.WriteByte(' ')
			writeNode(b, x.ElseScope)
		}
		b.WriteByte(')')
	case *WhileStmt:
		b.WriteString("(while ")
		writeNode(b, x.Condition)
		b.WriteByte(' ')
		writeNode(b, x.Scope)
		b.WriteByte(')')
	case *Scope:
		b.WriteString("(scope ")
		b.WriteString(tok(x.OpenToken))
		for _, s := range x.Stmts {
			b.WriteByte(' ')
			writeNode(b, s)
		}
		b.WriteByte(' ')
		b.WriteString(tok(x.CloseToken))
		b.WriteByte(')')
	case *RoleType:
		if x.IsShared() {
			b.WriteString("(shared-role-type")
		} else {
			b.WriteString("(role-type")
		}
		for _, r := range x.RoleNodes {
			b.WriteByte(' ')
			writeNode(b, r)
		}
		b.WriteByte(')')
	case *Role:
		if x.Invalid {
			b.WriteString("(invalid-role")
			for _, r := range x.Skipped {
				writeSkipped(b, r)
			}
			b.WriteByte(')')
		} else {
			b.WriteString("(role ")
			b.WriteString(x.Token.Text)
			b.WriteString(")")
		}
	case *NamedType:
		b.WriteString("(named-type ")
		writeNode(b, x.RoleIdent)
		b.WriteByte(')')
	case *ClosureType:
		b.WriteString("(closure-type ")
		writeNode(b, x.RoleType)
		b.WriteByte(' ')
		writeNode(b, x.Params)
		if x.ReturnType != nil {
			b.WriteByte(' ')
			writeNode(b, x.ReturnType)
		}
		b.WriteByte(')')
	case *ClosureTypeParams:
		b.WriteString("(type-params")
		for _, p := range x.Params {
			b.WriteByte(' ')
			writeNode(b, p)
		}
		b.WriteByte(')')
	case *FuncParams:
		b.WriteString("(params")
		for _, p := range x.Params {
			b.WriteByte(' ')
			writeNode(b, p)
		}
		b.WriteByte(')')
	case *FuncParam:
		b.WriteString("(param ")
		writeNode(b, x.Name)
		b.WriteByte(' ')
		writeNode(b, x.Type)
		b.WriteByte(')')
	case *FuncSig:
		b.WriteString("(func-sig")
		if x.RoleType != nil {
			b.WriteByte(' ')
			writeNode(b, x.RoleType)
		}
		b.WriteByte(' ')
		writeNode(b, x.Name)
		b.WriteByte(' ')
		writeNode(b, x.Params)
		if x.ReturnType != nil {
			b.WriteByte(' ')
			writeNode(b, x.ReturnType)
		}
		b.WriteByte(')')
	case *Func:
		b.WriteString("(func ")
		writeNode(b, x.FuncSig)
		b.WriteByte(' ')
		writeNode(b, x.Scope)
		b.WriteByte(')')
	case *Struct:
		b.WriteString("(struct")
		if x.RoleType != nil {
			b.WriteByte(' ')
			writeNode(b, x.RoleType)
		}
		b.WriteByte(' ')
		writeNode(b, x.Name)
		for _, i := range x.Implements {
			b.WriteByte(' ')
			b.WriteString("(implements ")
			writeNode(b, i)
			b.WriteByte(')')
		}
		b.WriteByte(' ')
		writeNode(b, x.Body)
		b.WriteByte(')')
	case *StructBody:
		b.WriteString("(struct-body")
		for _, f := range x.Members {
			b.WriteByte(' ')
			writeNode(b, f)
		}
		b.WriteByte(')')
	case *StructField:
		b.WriteString("(struct-field ")
		writeNode(b, x.Name)
		b.WriteByte(' ')
		writeNode(b, x.Type)
		b.WriteByte(')')
	case *Interface:
		b.WriteString("(interface")
		if x.RoleType != nil {
			b.WriteByte(' ')
			writeNode(b, x.RoleType)
		}
		b.WriteByte(' ')
		writeNode(b, x.Name)
		b.WriteByte(' ')
		writeNode(b, x.Methods)
		b.WriteByte(')')
	case *InterfaceMethodsList:
		b.WriteString("(methods")
		for _, m := range x.Methods {
			b.WriteByte(' ')
			writeNode(b, m)
		}
		b.WriteByte(')')
	case *InterfaceMethod:
		b.WriteString("(method ")
		writeNode(b, x.FuncSig)
		b.WriteByte(')')
	case *ListType:
		b.WriteString("(list-type ")
		writeNode(b, x.Inner)
		b.WriteByte(' ')
		b.WriteString(tok(x.CloseBracket))
		b.WriteByte(')')
	case *AsyncType:
		b.WriteString("(async-type ")
		writeNode(b, x.Inner)
		b.WriteByte(')')
	case *InvalidType:
		b.WriteString("(invalid-type)")
	default:
		fmt.Fprintf(b, "(%T)", n)
	}
}
func writeSkipped(b *sexprWriter, r token.Ref) {
	i := int(r)
	if i >= 0 && i < len(b.tokens) {
		fmt.Fprintf(b, " (skipped %s %q)", b.tokens[i].Kind, b.tokens[i].Text)
		return
	}
	fmt.Fprintf(b, " (skipped %d)", r)
}

// indentSExpr keeps atoms on their parent's line and starts each nested list on
// a new, indented line. Quoted strings are copied verbatim, including escaped
// quotes and parentheses.
func indentSExpr(compact, indent string) string {
	var out strings.Builder
	depth := 0
	inString := false
	escaped := false
	for i, r := range compact {
		if inString {
			out.WriteRune(r)
			if escaped {
				escaped = false
			} else if r == '\\' {
				escaped = true
			} else if r == '"' {
				inString = false
			}
			continue
		}
		if r == '"' {
			inString = true
			out.WriteRune(r)
			continue
		}
		if r == ' ' && i+1 < len(compact) && compact[i+1] == '(' {
			continue
		}
		if r == '(' {
			if i > 0 {
				out.WriteByte('\n')
				out.WriteString(strings.Repeat(indent, depth))
			}
			depth++
			out.WriteRune(r)
			continue
		}
		if r == ')' {
			depth--
		}
		out.WriteRune(r)
	}
	return out.String()
}
