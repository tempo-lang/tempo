package ast

import (
	"fmt"
	"reflect"

	"github.com/tempo-lang/tempo/parser/new_parser/token"
)

var (
	tokenValueType = reflect.TypeOf(token.Token{})
	tokenRefType   = reflect.TypeOf(token.Ref(0))
)

// Validate checks required AST children, token references, and all spans
// reachable from an AST. Nil is accepted only for grammar-defined optional
// fields; slices may be nil, but their elements may not be.
func Validate(root Node, tokens []token.Token, sourceLen int) error {
	if root == nil {
		return fmt.Errorf("root is nil")
	}
	for i, tok := range tokens {
		if int(tok.Index) != i {
			return fmt.Errorf("token %d has index %d", i, tok.Index)
		}
		if !validSpan(tok.Span, sourceLen) {
			return fmt.Errorf("token %d has invalid span", i)
		}
	}
	return validateValue(reflect.ValueOf(root), tokens, sourceLen, map[uintptr]bool{}, "root", false)
}

func validateValue(v reflect.Value, tokens []token.Token, sourceLen int, seen map[uintptr]bool, path string, optional bool) error {
	if !v.IsValid() {
		return nil
	}
	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			if optional {
				return nil
			}
			return fmt.Errorf("%s is a nil required child", path)
		}
		return validateValue(v.Elem(), tokens, sourceLen, seen, path, optional)
	}
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			if optional {
				return nil
			}
			return fmt.Errorf("%s is a nil required child", path)
		}
		ptr := v.Pointer()
		if seen[ptr] {
			return nil
		}
		seen[ptr] = true

		// Check children before calling boundary methods, because those methods
		// legitimately assume their required children are present.
		if err := validateValue(v.Elem(), tokens, sourceLen, seen, path, false); err != nil {
			return err
		}
		if node, ok := v.Interface().(Node); ok {
			start, end := node.StartToken(), node.EndToken()
			if start.Kind == "" || end.Kind == "" {
				return fmt.Errorf("%s (%T) has an empty boundary token", path, node)
			}
			if !validSpan(start.Span, sourceLen) || !validSpan(end.Span, sourceLen) || start.Span.Start > end.Span.End {
				return fmt.Errorf("%s (%T) has unordered boundaries", path, node)
			}
		}
		return nil
	}
	if v.Type() == tokenValueType {
		return validateToken(v.Interface().(token.Token), tokens, sourceLen, path, optional)
	}
	if v.Type() == tokenRefType {
		ref := v.Interface().(token.Ref)
		if ref < 0 || int(ref) >= len(tokens) {
			return fmt.Errorf("%s has invalid token reference %d", path, ref)
		}
		return nil
	}

	switch v.Kind() {
	case reflect.Struct:
		typ := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := typ.Field(i)
			fieldPath := path + "." + field.Name
			if err := validateValue(v.Field(i), tokens, sourceLen, seen, fieldPath, optionalField(typ.Name(), field.Name, v)); err != nil {
				return err
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if err := validateValue(v.Index(i), tokens, sourceLen, seen, fmt.Sprintf("%s[%d]", path, i), false); err != nil {
				return err
			}
		}
	}
	return nil
}

func validateToken(tok token.Token, tokens []token.Token, sourceLen int, path string, optional bool) error {
	if tok.Kind == "" {
		if optional {
			return nil
		}
		return fmt.Errorf("%s is an empty required token", path)
	}
	if !validSpan(tok.Span, sourceLen) {
		return fmt.Errorf("%s has invalid span", path)
	}
	if tok.Synthetic {
		if tok.Index != token.NoRef {
			return fmt.Errorf("%s synthetic token has reference %d", path, tok.Index)
		}
		return nil
	}
	if tok.Index < 0 || int(tok.Index) >= len(tokens) {
		return fmt.Errorf("%s has invalid token reference %d", path, tok.Index)
	}
	streamToken := tokens[int(tok.Index)]
	if streamToken.Kind != tok.Kind || streamToken.Text != tok.Text || streamToken.Span != tok.Span {
		return fmt.Errorf("%s does not match token stream entry %d", path, tok.Index)
	}
	return nil
}

func validSpan(span token.Span, sourceLen int) bool {
	return span.Start >= 0 && span.End >= span.Start && span.End <= sourceLen
}

func optionalField(parent, name string, parentValue reflect.Value) bool {
	switch parent + "." + name {
	case "InvalidExpr.PartialExpr", "InvalidStmt.PartialStmt",
		"StringLit.RoleType", "BoolLit.RoleType", "PrimitiveExpr.RoleType",
		"IdentAccessExpr.RoleType",
		"ClosureSig.ReturnType", "FuncSig.RoleType", "FuncSig.ReturnType",
		"IfStmt.ElseScope", "ReturnStmt.Expr", "Struct.RoleType",
		"ClosureType.ReturnType", "RoleIdent.RoleType",
		"PrimitiveExpr.RoleAt", "IdentAccessExpr.RoleAt", "RoleIdent.RoleAt",
		"LetStmt.Colon", "IfStmt.ElseToken":
		return true
	case "LetStmt.Type":
		// A type annotation is optional only when its colon is absent.
		return parentValue.FieldByName("Colon").Interface().(token.Token).Kind == ""
	}
	return false
}
