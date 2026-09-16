# Parser Rewrite Phase 1: Architecture and Test Foundation

## Goal

Refactor the existing implementation in `parser/new_parser` onto the final parser architecture without expanding it to the complete Tempo grammar. Preserve the syntax already supported by the new parser while establishing lossless tokenization, structured diagnostics, context-sensitive recovery, safe partial ASTs, scalable S-expression assertions, and fuzzable invariants.

Phase 1 is complete when the currently implemented lexer, role types, expressions, statements, and scopes use the new infrastructure and their tests no longer depend on constructing complete token-heavy AST values by hand.

## Source, tokens, and lexer

- Replace the `io.RuneReader`-centred state with an immutable `Source` containing the original bytes and a lazily or eagerly built line index.
- Treat byte offsets as the canonical coordinates. Derive one-based display positions and zero-based LSP UTF-16 positions from `Source`; do not store independently mutable line/column state.
- Give each token a stable index, `Kind`, exact source text, decoded value, source span, leading trivia, and `Synthetic` flag.
- Preserve whitespace and comments as trivia. Attach trivia to the following token, with EOF owning trailing trivia, so concatenating trivia and token text reconstructs the input exactly.
- Represent invalid source characters with `BadToken` and a lexical diagnostic. Do not encode diagnostics as `ILLEGAL` tokens.
- Distinguish EOF from reader/input failures, reject embedded NUL explicitly, and turn malformed strings and numeric overflow into diagnostics rather than panics.
- Implement all lexical rules already declared by `Tempo.g4`, including line and block comments, even when the corresponding parser production is deferred to Phase 2.
- Keep `Tempo.g4` identifiers and literals as the compatibility contract. Do not broaden identifier acceptance during the rewrite.

## Parser and AST architecture

- Add a buffered token cursor exposing `Current`, bounded `Peek`, `Advance`, and token-index position. The cursor must return the same EOF token on repeated reads.
- Replace `ParserErrors() []token.Token` with structured diagnostics:

  ```go
  type Diagnostic struct {
      Code        ErrorCode
      Message     string
      PrimarySpan Span
      Expected    []token.Kind
      Found       token.Kind
      Fixes       []TextEdit
  }

  type ErrorCode int

  const (
      CodeMissingToken ErrorCode = iota + 1
      CodeExpectedRole
  )
  ```

- Diagnostic codes, spans, expected/found kinds, and fixes are stable. Human-readable message wording is not a snapshot-test contract.
- Keep missing-semicolon diagnostics visually anchored to the preceding token. Supply a separate zero-width insertion edit immediately after that token.
- Keep required AST children non-nil. Use typed invalid placeholders for required expressions, statements, types, identifiers, and scopes; reserve nil for grammar-defined optional fields.
- Keep recognizable outer nodes after inner failures. Required punctuation fields contain either a real or synthetic token.
- Store the complete token stream on the source-file/test parse result so unexpected and skipped tokens remain available for lossless reconstruction.
- Refactor the implemented expression subset into a conventional Pratt loop with separate prefix, postfix, and infix handlers. Field and index postfixes must operate on any supported base expression. Validate assignment targets separately from general expressions.
- Public parsing returns a result containing the root, tokens, and diagnostics. Test-only expression, statement, type, and scope entry points must also require EOF and diagnose trailing input.

## Error recovery architecture

The recovery design deliberately uses a small set of conventions rather than a status protocol propagated through every recursive call. Parser functions return AST nodes, record diagnostics on the parser, and use caller-provided stop sets to avoid consuming enclosing syntax.

### Core contracts

```go
type TokenSet []token.Kind

func (s TokenSet) Contains(kind token.Kind) bool
func (s TokenSet) With(kind token.Kind) TokenSet
func (s TokenSet) Union(other TokenSet) TokenSet

func (p *Parser) expect(expected token.Kind, follow TokenSet) token.Ref
func (p *Parser) skipUntil(stop TokenSet) []token.Ref

func (p *Parser) parseExpr(stop TokenSet) ast.Expr
func (p *Parser) parseStmt() ast.Stmt
func (p *Parser) parseScope() *ast.Scope
```

- Parser functions return a non-nil node and do not return recovery status.
- A nested parser stops before tokens supplied by its caller and leaves those tokens unconsumed.
- The caller determines whether a boundary is valid following syntax, a missing-token location, or the beginning of its next sibling.
- Productions continue to implement their own grammar-specific recovery. `expect` and `skipUntil` provide only shared mechanics.
- `TokenSet` is intentionally simple because recovery sets are small. `Union` returns a new set so a callee cannot modify its caller's boundaries.

### Recovered syntax representation

- A missing required token becomes a synthetic token of the expected kind with empty text/trivia and a zero-width insertion span.
- Normally insertion occurs at the current token's start. A missing statement semicolon instead uses the preceding significant token's end for its synthetic span and insertion fix.
- A malformed required child becomes a typed invalid node; required children are never nil.
- Unexpected real tokens remain in the complete lossless token stream. `skipUntil` returns their references so an invalid node can retain the relevant skipped range when useful; attaching every skipped token to an AST node is not required.
- Synthetic tokens, invalid nodes, and skipped significant syntax represented by an invalid node appear in S-expressions. Trivia does not.

### `expect` algorithm

`expect(expected, follow)` executes in this order:

1. If current is `expected`, consume and return it.
2. If one-token lookahead is `expected`, diagnose current as unexpected with a deletion fix, consume current, and then consume and return `expected`.
3. If current is in `follow`, is EOF, or is a closing delimiter owned by the caller, diagnose a missing token and return a synthetic `expected` token without consuming current.
4. Otherwise, report one mismatch and call `skipUntil(follow.With(expected))`. Consume `expected` if reached; otherwise return a synthetic token and leave the reached follow token unconsumed.

Syntax error codes follow the existing type-error convention: define an integer `ErrorCode` block with `iota + 1`, expose `Code()`, and format user-facing codes as `error[E%d]`. Syntax codes have their own namespace, so `E1` means `CodeMissingToken` in the parser and does not reuse a type-error constant.

Each diagnostic records expected and found kinds, its primary span, and an insertion or deletion fix when unambiguous. `expect` must not emit a second diagnostic when its caller has already supplied an invalid required child for the same missing syntax.

Here, **consume** means calling `Advance` on the buffered token cursor. `Advance` moves the parser's current-token index by one and returns a reference to the token just passed; it never removes that token from the immutable token stream. Consequently, parsing and lossless storage are separate: a consumed token cannot be parsed twice, but remains available to the AST, diagnostics, and source reconstruction.

Consumption follows grammatical ownership. A child consumes tokens belonging to its construct and stops before a caller-provided boundary. The caller resumes at that same boundary and may consume it. For example, while parsing `arr[0];`, the expression for `0` stops before `]`, the index parser consumes `]`, the enclosing expression stops before `;`, and the statement parser consumes `;`. During recovery, a synthetic token satisfies a grammar slot without advancing the cursor, so the production must then move to its next grammar slot or return.

### `skipUntil` algorithm

- Stop before the first token in `stop`; the owning production decides whether to consume it.
- Always stop before EOF.
- Stop before an unmatched closing delimiter even when it is not listed, allowing an enclosing production to handle it.
- When an unexpected opener is skipped, track its matching `()`, `[]`, or `{}` so separators inside that balanced region do not become false recovery boundaries.
- Return all consumed token references and never discard a statement, member, or declaration starter included in `stop`.

Nesting awareness remains a small implementation detail of token skipping. Normal recursive parsing provides the primary delimiter ownership.

### Standard boundary sets

| Context | Tokens supplied as stops/followers |
| --- | --- |
| Role sequence | comma, matching `]` or `)`, caller stops |
| Index expression | `]`, `,`, `;`, `)`, `}`, EOF |
| Parenthesized expression | `)`, `,`, `;`, `]`, `}`, EOF |
| `if`/`while` condition | `{`, `}`, EOF, statement starters |
| Statement | `;`, `}`, EOF, statement starters |
| Scope | `}`, EOF, declaration starters |
| Struct/interface body | `}`, EOF, member and declaration starters |
| Source file | `func`, `struct`, `interface`, EOF |

Statement starters are `let`, `return`, `if`, `while`, and all expression starters. A production consumes only delimiters it opened. A sibling starter is always left unconsumed for the enclosing loop.

### Progress and cascades

- Before each parser loop iteration, record the token index. The iteration must advance input or terminate.
- Synthetic insertion satisfies the expected grammar slot immediately; the parser must not revisit and insert that slot again at the same token index.
- If a loop makes no progress, report one internal-recovery diagnostic in debug tests, call the loop's `skipUntil`, and terminate the loop if no token can safely be consumed.
- Each parser function reports only errors for syntax it owns. Callers preserve returned invalid nodes without adding generic wrapper diagnostics.
- Preserve `BadToken` lexer diagnostics. Add a parser diagnostic at the same span only when it describes a distinct missing grammar element.
- Deduplicate diagnostics by code, primary span, and expected kind.

These rules are the complete recovery protocol:

1. Never consume a boundary owned by the caller.
2. Represent missing syntax with a diagnostic and synthetic token or invalid node.
3. Skip unexpected syntax only to a context-specific boundary.
4. Every parser loop consumes input or terminates.

### Normative recovery examples

Missing semicolon before a sibling statement:

```tempo
{let x = 1 let y = 2;}
```

```text
(scope "{"
  (let x (int 1) (missing SEMICOLON))
  (let y (int 2) ";")
  "}")
```

Emit `E1 (CodeMissingToken)`, expected `SEMICOLON`, found `LET`; anchor it to `1`, insert `;` after `1`, and leave `let` unconsumed so the scope parses the second statement.

Missing index closer:

```tempo
{let x = arr[0; let y = 2;}
```

```text
(scope "{"
  (let x (index arr (int 0) (missing RSQUARE)) ";")
  (let y (int 2) ";")
  "}")
```

Emit `E1 (CodeMissingToken)`, expected `RSQUARE`, found `SEMICOLON`; synthesize `]` without consuming `;`. Statement parsing then consumes the real semicolon.

Invalid role-list member:

```tempo
[A,123,B]
```

```text
(shared-role-type "[" (role A)
  (invalid-role (skipped INT "123"))
  (role B) "]")
```

Emit `E2 (CodeExpectedRole)`, expected `IDENT` or `UNDERSCORE`, found `INT`; skip only `123`, stop before the comma, and continue with `B`.

Malformed nested expression:

```tempo
{let x = foo[1 + ]; let y = 2;}
```

```text
(scope "{"
  (let x (index foo (binary + (int 1)
    (invalid-expr before RSQUARE)) "]") ";")
  (let y (int 2) ";")
  "}")
```

Emit one expression diagnostic at `]`. The invalid operand leaves `]` untouched, the index parser consumes it, and no secondary missing-bracket or invalid-statement diagnostic is emitted.

Unexpected EOF in a scope:

```tempo
{let x = 1;
```

```text
(scope "{" (let x (int 1) ";") (missing RCURLY))
```

Emit `E1 (CodeMissingToken)`, expected `RCURLY`, found `EOF`; insert a zero-width `}` at EOF, preserve the completed statement, and do not repeat the EOF diagnostic.

## S-expression test representation


- Implement a deterministic S-expression printer for AST assertions. It includes node kinds, significant token text, decoded literal values, synthetic-token markers, and invalid/skipped syntax, but excludes ordinary whitespace and incidental coordinate fields.
- Keep diagnostics out of the AST S-expression; print them through a separate stable diagnostic formatter containing code, span, found/expected kinds, and fixes.
- Use compact forms such as:

  ```text
  (let x (= 1))
  (binary + 1 (binary * 2 3))
  (index arr (invalid-expr ";"))
  (missing SEMICOLON)
  ```

- Provide helpers that parse a fixture and compare:
  1. The S-expression.
  2. Structured diagnostics.
  3. Remaining/trailing input status.
- Rewrite the existing parser tests as table-driven production suites using these helpers. Keep separate focused tests for exact token spans, trivia, decoded values, and source reconstruction.
- Organize cases by grammar production rather than one large expected-AST table. Every production suite contains valid forms, precedence/nesting cases, malformed required elements, local recovery, and trailing-input cases.

## Phase 1 tests

- Lexer tables cover every token, keyword boundaries, comments, whitespace, CRLF, invalid characters, malformed escapes, unterminated strings/comments, embedded NUL, and numeric overflow.
- AST invariant tests walk every result and verify non-nil required children, valid token references, ordered in-bounds spans, and safe `StartToken`/`EndToken` access.
- Lossless tests reconstruct the exact source from tokens and trivia for both valid and invalid inputs.
- Recovery-locality tests place an error before a valid sentinel statement and assert that the sentinel remains in the AST and its first token is not consumed.
- Diagnostic tests verify codes, primary spans, expected/found tokens, fixes, and cascade suppression.
- Add fuzz targets for the lexer and each implemented parser entry point. Assert no panic, deterministic output, parser progress, safe AST traversal, valid spans, and exact source reconstruction.
- Keep a small benchmark for representative expressions and scopes. Record results for regression visibility, but do not make performance a Phase 1 pass/fail gate.

## Acceptance criteria

- `go test ./parser/new_parser/...` passes with rewritten S-expression-based suites.
- Existing supported syntax produces equivalent logical AST structure to the pre-refactor implementation.
- Invalid input cannot panic, hang, create invalid spans, or leave required AST children nil.
- A syntax error in one statement does not prevent a following valid statement from being represented.
- Lexer/parser diagnostics are structured and no longer stored as error tokens.
- Exact input can be reconstructed from the token stream for every test and fuzz input.
- Phase 1 does not modify compiler, type-checker, projection, code-generation, or LSP consumers.

