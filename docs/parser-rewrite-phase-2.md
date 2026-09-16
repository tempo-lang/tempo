# Parser Rewrite Phase 2: Complete Tempo Grammar

## Goal

Implement the complete syntax in `Tempo.g4` using the Phase 1 architecture while keeping the new parser isolated from the compiler and other existing consumers. The generated ANTLR parser remains active in production and serves as the zero-error acceptance oracle.

Phase 2 is complete when the handwritten parser implements every lexer and parser alternative, agrees with ANTLR on syntax acceptance, produces conforming ASTs for all valid forms, and provides tested local recovery for malformed forms.

## Grammar implementation

- Treat `Tempo.g4` as authoritative and freeze syntax for the duration of this phase. `Tempo.ebnf` is documentation only and must not override discrepancies in the ANTLR grammar.
- Maintain a checked grammar coverage matrix listing every rule and labelled alternative, its parser function, valid fixtures, recovery fixtures, and compatibility coverage.
- Implement the remaining syntax in vertical slices, completing AST nodes, S-expression output, valid tests, recovery tests, and ANTLR comparison before starting the next slice.
- Use this order to expose reusable building blocks early:
  1. Role identifiers and all value types.
  2. Parameter, argument, and comma-separated list helpers.
  3. Remaining prefix expressions: grouping, list literals, `await`, communication, struct construction, and closures.
  4. Remaining postfix expressions: calls combined with arbitrary field/index chains.
  5. Optional type annotations in `let` and all statement variants.
  6. Function signatures, functions, and scopes.
  7. Struct declarations, implementations, fields, and methods.
  8. Interface declarations and methods.
  9. Source-file declarations and mandatory EOF.
- Encode operator precedence and associativity explicitly in Pratt binding-power tables. Add focused fixtures documenting the intended tree for every mixed-operator pair.
- Reuse generic separated-list parsing with production-specific element parsers, delimiters, empty-list policy, trailing-comma policy, recovery sets, and S-expression output. Policies must reproduce `Tempo.g4` exactly.
- Ensure source-file parsing retains invalid top-level regions and resumes at `func`, `struct`, `interface`, or EOF.

## Compatibility harness

- Build a test-only harness that parses the same source with ANTLR and the handwritten parser.
- Define acceptance identically for both parsers: accepted means lexing and parsing complete through EOF with zero syntax diagnostics.
- Compare only acceptance, not generated parse-tree shape or diagnostic wording.
- Run the harness over:
  - Every repository `.tempo` file.
  - Tempo sources extracted from type-checker fixtures.
  - Existing compiler and projection fuzz corpora.
  - Minimal and maximal valid examples for every grammar alternative.
  - A curated invalid corpus covering every required token and delimiter.
- Every compatibility disagreement becomes a checked-in regression fixture. Because syntax is frozen, Phase 2 has no permanent allowlist: resolve each disagreement in the handwritten parser or demonstrate and fix an erroneous test extraction.
- Add differential fuzzing seeded from all valid corpus files. Include arbitrary bytes and token-level delete, duplicate, insert, and replace mutations.
- Differential fuzzing asserts only zero-error acceptance parity. Recovery quality is governed by the new parser's explicit tests rather than ANTLR's recovery behaviour.

## AST and recovery validation

- Add S-expression fixtures for every valid grammar alternative, including nested combinations rather than testing productions only in isolation.
- Verify all optional and repeated elements in empty, singleton, and multiple-element forms.
- Add recovery fixtures for missing, extra, and replaced openers, closers, commas, colons, assignment operators, semicolons, names, types, expressions, bodies, and EOF.
- For each recovery fixture, assert structured diagnostics and a recovered S-expression containing valid constructs after the error.
- Add delimiter-stress cases involving nested closures, lists, calls, indexes, struct expressions, scopes, struct bodies, and interface bodies.
- Test ambiguous-looking prefixes explicitly, including role types versus list expressions, identifiers versus role-annotated identifiers and struct expressions, and function signatures versus closures.
- For every advertised fix, apply the edit and assert that it changes the intended range. Curated single-error fixtures must parse cleanly after applying the fix.
- Run the Phase 1 structural properties across the full grammar: lossless token reconstruction, non-nil required children, valid token references and spans, deterministic parsing, bounded diagnostics, safe traversal, and loop progress.

## Test organization and gates

- Keep small production tests table-driven and S-expression based.
- Keep larger examples as files under parser testdata with adjacent expected S-expression and diagnostic data.
- Partition tests into fast unit tests, corpus compatibility tests, mutation tests, and fuzz targets so normal `go test` remains useful while extended jobs run in CI.
- Run all valid corpus compatibility tests in the normal test suite.
- Run a deterministic bounded mutation sample in the normal suite and the full mutation matrix in extended CI.
- Run fuzz smoke tests in CI with a fixed time budget; retain every discovered failure in the ordinary regression corpus.
- Track benchmarks for lexing, valid parsing, and heavily malformed parsing. Performance is informational unless a change causes an unbounded or order-of-magnitude regression.

## Acceptance criteria

- Every `Tempo.g4` rule and alternative is implemented and marked complete in the coverage matrix.
- The complete valid corpus parses with zero handwritten-parser diagnostics.
- ANTLR and the handwritten parser have zero unexplained zero-error acceptance disagreements across checked-in fixtures and the deterministic mutation suite.
- All AST, diagnostic, recovery-locality, lossless, progress, and fuzz invariants pass for the complete grammar.
- Source-file parsing always terminates at EOF and preserves declarations following malformed input whenever a grammatical boundary can be identified.
- No production code outside `parser/new_parser` imports or invokes the handwritten parser in this phase.
- The existing compiler continues using ANTLR and all repository tests remain unchanged and passing.

