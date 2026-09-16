# Parser Rewrite Phase 3: Repository-Wide Parser Replacement

## Goal

Replace the generated ANTLR parser with the handwritten parser in one coordinated repository-wide change. Migrate all compiler, type-checking, projection, LSP, examples, and test code to the new AST and diagnostics, then remove generated parser code and the ANTLR runtime dependency.

Phase 3 is complete when the repository contains one Tempo parser implementation, all existing tests pass without changing valid-program behaviour, and malformed programs receive safe, useful diagnostics from the new recovery system.

## Downstream migration

- Make the Phase 2 `Parse(Source) Result` API the only Tempo parsing entry point.
- Rewrite the type checker to traverse the new AST directly. Do not add an ANTLR-context compatibility adapter.
- Preserve the existing type system, symbol-table semantics, error codes, scope construction, return analysis, and role behaviour. Parser error nodes short-circuit local semantic checks to avoid cascaded type errors while allowing unaffected declarations to be checked.
- Change compiler orchestration to:
  1. Parse source with the handwritten parser.
  2. Retain all syntax diagnostics.
  3. Type-check the recovered AST, suppressing checks rooted in invalid syntax.
  4. Return all useful syntax and independent type diagnostics.
  5. Run projection/code generation only when no error-severity diagnostics remain.
- Rewrite endpoint projection and related visitors to consume the new AST while preserving generated Go, TypeScript, JavaScript, and Java output.
- Update the LSP document model and diagnostics conversion to use source spans. Convert byte spans to UTF-16 protocol ranges through `Source`; do not perform ad hoc column arithmetic.
- Expose diagnostic code, severity, primary range, related information, and supported fixes through the LSP. Preserve the intentional preceding-token highlight for missing semicolons and apply the zero-width insertion range for quick fixes.
- Replace ANTLR character-stream inputs in public compiler helpers with the new source abstraction. If external API compatibility is required, retain a small deprecated input wrapper for one release, but it must immediately materialize a `Source` and must not expose ANTLR types.

## Big-bang cutover and removal

- Land the consumer migrations, parser switch, generated-code deletion, and dependency cleanup together so the main branch never contains two production parser paths.
- Delete generated lexer/parser/listener/visitor files, generation scripts and JARs used only for ANTLR, and old parser helpers after all consumers compile against the new AST.
- Remove the ANTLR Go runtime from `go.mod` after verifying no remaining imports with a repository-wide search.
- Retain the Phase 2 compatibility corpus and expected acceptance classifications, but replace the live ANTLR comparison with checked-in expected outcomes before removing ANTLR.
- Preserve `Tempo.g4` as the frozen language specification and historical grammar reference unless a later dedicated change chooses another canonical grammar format.
- Do not combine language syntax changes, type-system changes, diagnostic redesign, or generated-output changes with this cutover.

## Regression and integration testing

- Run every existing unit, example, type-checker, projection, code-generation, runtime, compiler fuzz-seed, and LSP test. Existing valid-program expectations must pass unchanged.
- Before deleting ANTLR, run a final side-by-side acceptance report over the complete corpus and deterministic mutation suite and require zero disagreement.
- Run all valid projection examples through the new complete pipeline and compare generated outputs byte-for-byte under the repository's existing whitespace policy.
- Port parser-position-dependent type errors to source spans while preserving error codes and logically equivalent highlighted syntax.
- Add integration fixtures containing both syntax errors and independent semantic errors. Assert safe type checking, no panic, no duplicate cascades, and diagnostics from unaffected declarations.
- Add LSP tests for ASCII, Unicode, multiline comments/strings, missing tokens, and synthetic-token fixes to verify UTF-16 ranges and edit placement.
- Run parser and full-compiler fuzz tests with the Phase 2 corpus. Assert no panic or hang through parsing, type checking, diagnostic formatting, projection gating, and LSP range conversion.
- Compare parser and compilation benchmarks before and after replacement. Investigate substantial regressions, especially malformed-input superlinear behaviour, but prioritize correctness and termination over matching ANTLR throughput exactly.

## Cutover acceptance criteria

- `go test ./...` passes after generated parser files and the ANTLR dependency are removed.
- All pre-existing valid Tempo programs retain the same acceptance, type-checking outcome, projected program, and generated code.
- The frozen compatibility corpus retains its expected accepted/rejected classification without invoking ANTLR.
- Invalid and partially written source can be parsed, traversed, type-checked where independent, formatted as diagnostics, and sent through LSP analysis without panic or hang.
- Recovery tests continue to preserve valid statements, members, and declarations following malformed syntax.
- Repository searches find no production imports, types, generated files, tools, or build steps belonging to ANTLR.
- Compiler and LSP public documentation describes the new source, AST, diagnostic, span, and fix APIs.
