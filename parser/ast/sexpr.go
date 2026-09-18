package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tempo-lang/tempo/parser/token"
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
	b.writeNode(n)
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

type sexprListWriter struct {
	writer *sexprWriter
}

func (b *sexprWriter) writeList(name string, writeElements func(*sexprListWriter)) {
	b.WriteByte('(')
	b.WriteString(name)
	if writeElements != nil {
		writeElements(&sexprListWriter{writer: b})
	}
	b.WriteByte(')')
}

func (l *sexprListWriter) atom(value string) {
	l.writer.WriteByte(' ')
	l.writer.WriteString(value)
}

func (l *sexprListWriter) node(n Node) {
	l.writer.WriteByte(' ')
	l.writer.writeNode(n)
}

func (l *sexprListWriter) roleIdent(roleIdent *RoleIdent) {
	l.node(roleIdent.Ident)
	if roleIdent.RoleType != nil {
		l.node(roleIdent.RoleType)
	}
}

func (l *sexprListWriter) list(name string, writeElements func(*sexprListWriter)) {
	l.element(func(writer *sexprWriter) {
		writer.writeList(name, writeElements)
	})
}

func (l *sexprListWriter) element(write func(*sexprWriter)) {
	l.writer.WriteByte(' ')
	write(l.writer)
}

func (l *sexprListWriter) token(t token.Token) {
	if t.Synthetic {
		l.list("missing", func(list *sexprListWriter) {
			list.atom(string(t.Kind))
		})
		return
	}
	l.atom(strconv.Quote(t.Text))
}

func (l *sexprListWriter) skipped(r token.Ref) {
	i := int(r)
	l.list("skipped", func(list *sexprListWriter) {
		if i >= 0 && i < len(l.writer.tokens) {
			list.atom(string(l.writer.tokens[i].Kind))
			list.atom(strconv.Quote(l.writer.tokens[i].Text))
			return
		}
		list.atom(strconv.Itoa(i))
	})
}

func (b *sexprWriter) writeNode(n Node) {
	if n == nil {
		b.WriteString("nil")
		return
	}

	switch x := n.(type) {
	case *SourceFile:
		b.writeList("source-file", func(list *sexprListWriter) {
			for _, declaration := range x.Declarations {
				list.node(declaration)
			}
		})
	case *Identifier:
		if x.Invalid {
			b.writeList("invalid-ident", nil)
		} else {
			b.WriteString(x.Value())
		}
	case *IdentAccessExpr:
		if x.RoleType == nil {
			b.writeNode(x.Ident)
			break
		}
		b.writeList("at", func(list *sexprListWriter) {
			list.node(x.Ident)
			list.node(x.RoleType)
		})
	case *RoleIdent:
		b.writeList("role-ident", func(list *sexprListWriter) {
			list.roleIdent(x)
		})
	case *StructExpr:
		b.writeList("struct-expr", func(list *sexprListWriter) {
			list.roleIdent(x.RoleIdent)
			for _, field := range x.StructFields.Fields {
				list.node(field)
			}
			list.token(x.StructFields.CloseToken)
		})
	case *StructFieldExpr:
		b.writeList("field-init", func(list *sexprListWriter) {
			list.node(x.Name)
			list.node(x.Expr)
		})
	case *ComExpr:
		b.writeList("com", func(list *sexprListWriter) {
			list.node(x.Sender)
			list.node(x.Receiver)
			list.node(x.Expr)
		})
	case *ClosureExpr:
		b.writeList("closure", func(list *sexprListWriter) {
			list.node(x.ClosureSig)
			list.node(x.Scope)
		})
	case *ClosureSig:
		b.writeList("closure-sig", func(list *sexprListWriter) {
			list.node(x.RoleType)
			list.node(x.Params)
			if x.ReturnType != nil {
				list.node(x.ReturnType)
			}
		})
	case *PrimitiveExpr:
		writeLiteral := func(writer *sexprWriter) {
			switch literal := x.Literal.(type) {
			case *IntLit:
				writer.writeList("int", func(value *sexprListWriter) { value.atom(fmt.Sprint(literal.Value())) })
			case *FloatLit:
				writer.writeList("float", func(value *sexprListWriter) { value.atom(fmt.Sprint(literal.Value())) })
			case *StringLit:
				writer.writeList("string", func(value *sexprListWriter) { value.atom(strconv.Quote(literal.Value())) })
			case *BoolLit:
				writer.writeList("bool", func(value *sexprListWriter) { value.atom(fmt.Sprint(literal.Value())) })
			}
		}
		if x.RoleType != nil {
			b.writeList("at", func(list *sexprListWriter) {
				list.element(writeLiteral)
				list.node(x.RoleType)
			})
		} else {
			writeLiteral(b)
		}
	case *BinaryExpr:
		b.writeList("binary", func(list *sexprListWriter) {
			list.atom(x.Operator.Text)
			list.node(x.Left)
			list.node(x.Right)
		})
	case *FieldAccessExpr:
		b.writeList("field", func(list *sexprListWriter) {
			list.node(x.Object)
			list.node(x.Field)
		})
	case *IndexExpr:
		b.writeList("index", func(list *sexprListWriter) {
			list.node(x.Object)
			list.node(x.Index)
			list.token(x.CloseBracket)
		})
	case *GroupExpr:
		b.writeList("group", func(list *sexprListWriter) {
			list.node(x.Expr)
			list.token(x.CloseParen)
		})
	case *ListExpr:
		b.writeList("list", func(list *sexprListWriter) {
			for _, element := range x.Elements {
				list.node(element)
			}
			list.token(x.CloseBracket)
		})
	case *CallExpr:
		b.writeList("call", func(list *sexprListWriter) {
			list.node(x.Function)
			for _, argument := range x.Args {
				list.node(argument)
			}
			list.token(x.CloseParen)
		})
	case *AwaitExpr:
		b.writeList("await", func(list *sexprListWriter) { list.node(x.Expr) })
	case *InvalidExpr:
		b.writeList("invalid-expr", func(list *sexprListWriter) {
			for _, skipped := range x.Skipped {
				list.skipped(skipped)
			}
		})
	case *LetStmt:
		b.writeList("let", func(list *sexprListWriter) {
			list.node(x.Name)
			if x.Type != nil {
				list.list("type", func(typeList *sexprListWriter) { typeList.node(x.Type) })
			}
			list.node(x.Expr)
			list.token(x.SemiToken)
		})
	case *ReturnStmt:
		b.writeList("return", func(list *sexprListWriter) {
			if x.Expr != nil {
				list.node(x.Expr)
			}
			list.token(x.SemiToken)
		})
	case *AssignStmt:
		b.writeList("assign", func(list *sexprListWriter) {
			list.node(x.LHS)
			list.node(x.RHS)
			list.token(x.SemiToken)
		})
	case *ExprStmt:
		b.writeList("expr", func(list *sexprListWriter) {
			list.node(x.Expr)
			list.token(x.SemiToken)
		})
	case *IfStmt:
		b.writeList("if", func(list *sexprListWriter) {
			list.node(x.Condition)
			list.node(x.ThenScope)
			if x.ElseScope != nil {
				list.node(x.ElseScope)
			}
		})
	case *WhileStmt:
		b.writeList("while", func(list *sexprListWriter) {
			list.node(x.Condition)
			list.node(x.Scope)
		})
	case *Scope:
		b.writeList("scope", func(list *sexprListWriter) {
			list.token(x.OpenToken)
			for _, statement := range x.Stmts {
				list.node(statement)
			}
			list.token(x.CloseToken)
		})
	case *RoleType:
		name := "role-type"
		if x.IsShared() {
			name = "shared-role-type"
		}
		b.writeList(name, func(list *sexprListWriter) {
			for _, role := range x.RoleNodes {
				list.node(role)
			}
		})
	case *Role:
		if x.Invalid {
			b.writeList("invalid-role", func(list *sexprListWriter) {
				for _, skipped := range x.Skipped {
					list.skipped(skipped)
				}
			})
		} else {
			b.writeList("role", func(list *sexprListWriter) { list.atom(x.Token.Text) })
		}
	case *NamedType:
		b.writeList("named-type", func(list *sexprListWriter) { list.roleIdent(x.RoleIdent) })
	case *ClosureType:
		b.writeList("closure-type", func(list *sexprListWriter) {
			list.node(x.RoleType)
			list.node(x.Params)
			if x.ReturnType != nil {
				list.node(x.ReturnType)
			}
		})
	case *ClosureTypeParams:
		b.writeList("type-params", func(list *sexprListWriter) {
			for _, parameter := range x.Params {
				list.node(parameter)
			}
		})
	case *FuncParams:
		b.writeList("params", func(list *sexprListWriter) {
			for _, parameter := range x.Params {
				list.node(parameter)
			}
		})
	case *FuncParam:
		b.writeList("param", func(list *sexprListWriter) {
			list.node(x.Name)
			list.node(x.Type)
		})
	case *FuncSig:
		b.writeList("func-sig", func(list *sexprListWriter) {
			if x.RoleType != nil {
				list.node(x.RoleType)
			}
			list.node(x.Name)
			list.node(x.Params)
			if x.ReturnType != nil {
				list.node(x.ReturnType)
			}
		})
	case *Func:
		b.writeList("func", func(list *sexprListWriter) {
			list.node(x.FuncSig)
			list.node(x.Scope)
		})
	case *Struct:
		b.writeList("struct", func(list *sexprListWriter) {
			if x.RoleType != nil {
				list.node(x.RoleType)
			}
			list.node(x.Name)
			for _, implemented := range x.Implements {
				list.list("implements", func(implements *sexprListWriter) { implements.roleIdent(implemented) })
			}
			list.node(x.Body)
		})
	case *StructBody:
		b.writeList("struct-body", func(list *sexprListWriter) {
			for _, member := range x.Members {
				list.node(member)
			}
		})
	case *StructField:
		b.writeList("struct-field", func(list *sexprListWriter) {
			list.node(x.Name)
			list.node(x.Type)
		})
	case *Interface:
		b.writeList("interface", func(list *sexprListWriter) {
			if x.RoleType != nil {
				list.node(x.RoleType)
			}
			list.node(x.Name)
			list.node(x.Methods)
		})
	case *InterfaceMethodsList:
		b.writeList("methods", func(list *sexprListWriter) {
			for _, method := range x.Methods {
				list.node(method)
			}
		})
	case *InterfaceMethod:
		b.writeList("method", func(list *sexprListWriter) { list.node(x.FuncSig) })
	case *ListType:
		b.writeList("list-type", func(list *sexprListWriter) {
			list.node(x.Inner)
			list.token(x.CloseBracket)
		})
	case *AsyncType:
		b.writeList("async-type", func(list *sexprListWriter) { list.node(x.Inner) })
	case *InvalidType:
		b.writeList("invalid-type", nil)
	default:
		fmt.Fprintf(b, "(%T)", n)
	}
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
