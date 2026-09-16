# Tempo grammar coverage

`Tempo.g4` is the syntax authority. Each row is exercised by the ordinary
compatibility suite and the acceptance-parity fuzz target; malformed delimiter
and required-token cases are covered by `TestGrammarAlternativeCompatibility`
and the recovery/property tests.

| Rule / labelled alternatives | Handwritten parser | Valid coverage | Recovery / parity coverage |
|---|---|---|---|
| `sourceFile` | `Parse` | repository corpus, alternative suite | trailing region + fuzz |
| `ident`, `roleIdent`, `role`, `roleTypeShared`, `roleTypeNormal` | `parseIdentifier`, `parseRoleIdent`, `parseRoleType` | role/type and corpus suites | role recovery + fuzz |
| `asyncType`, `listType`, `closureType`, `namedType` | `parseValueType` | alternative suite + corpus | type property tests + fuzz |
| closure/function parameter and argument lists | `parseClosureTypeParams`, `parseFuncParams`, postfix call parser | alternative suite + corpus | delimiter cases + fuzz |
| `exprBinOp`, `exprPrimitive`, `exprClosure`, `exprStruct` | Pratt parser / `parsePrefix` | expression and corpus suites | expression properties + fuzz |
| `exprCall`, `exprFieldAccess`, `exprIndex`, `exprList` | Pratt postfix loop / `parsePrefix` | expression and corpus suites | locality tests + fuzz |
| `exprIdent`, `exprCom`, `exprAwait`, `exprGroup` | `parsePrefix`, `parseComExpr` | alternative suite + corpus | ambiguity cases + fuzz |
| all six `stmt` alternatives, `assignExpr`, `assignSpecifier`, `scope` | `parseStmt`, `parseScope` | alternative suite + corpus | statement/scope properties + fuzz |
| `func`, `funcSig`, `closureSig` | `parseFunc`, `parseFuncSig`, `parseClosureSig` | alternative suite + corpus | required token cases + fuzz |
| struct rules | `parseStruct`, `parseStructBody` | alternative suite + corpus | member recovery + fuzz |
| interface rules | `parseInterface` | alternative suite + corpus | method recovery + fuzz |

There is no compatibility allowlist: every checked-in disagreement fails.
