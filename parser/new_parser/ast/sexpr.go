package ast

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

var sexprTokens []token.Token
var sexprMu sync.Mutex

func SExpr(n Node, streams ...[]token.Token) string {
	sexprMu.Lock()
	defer sexprMu.Unlock()
	var b strings.Builder
	old := sexprTokens
	if len(streams) > 0 {
		sexprTokens = streams[0]
	} else {
		sexprTokens = nil
	}
	writeNode(&b, n)
	sexprTokens = old
	return b.String()
}
func tok(t token.Token) string {
	if t.Synthetic {
		return "(missing " + string(t.Kind) + ")"
	}
	return strconv.Quote(t.Text)
}
func writeNode(b *strings.Builder, n Node) {
	if n == nil {
		b.WriteString("nil")
		return
	}
	switch x := n.(type) {
	case *Identifier:
		if x.Invalid {
			b.WriteString("(invalid-ident)")
		} else {
			b.WriteString(x.Value())
		}
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
		writeNode(b, x.RoleIdent.Ident)
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
func writeSkipped(b *strings.Builder, r token.Ref) {
	i := int(r)
	if i >= 0 && i < len(sexprTokens) {
		fmt.Fprintf(b, " (skipped %s %q)", sexprTokens[i].Kind, sexprTokens[i].Text)
		return
	}
	fmt.Fprintf(b, " (skipped %d)", r)
}
