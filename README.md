# Ready-Made Parsers

Parsers for common languages, written in ZeroParse and generated into Go with `zp gen`. They are packages of the one Go
module `zeroparse.dev/alpha-parsers` and depend only on the standard library, the ZeroParse library
(`zeroparse.dev/zp/codegen/go/lib`) and `zeroparse.dev/zp/core`: import one and parse, with nothing to generate or compile at run time.

| Language | Module | Specification | Checked against |
|:--|:--|:--|:--|
| CEL | [`zeroparse.dev/alpha-parsers/cel`](cel/) | The [CEL language definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md), as cel-go parses it | All 2,527 conformance tests of cel-spec; cel-go v0.32.0 (trees, offsets and literal values on 32,794 expressions; fuzz tests) |
| CSV | [`zeroparse.dev/alpha-parsers/csv`](csv/) | [RFC 4180](https://www.rfc-editor.org/rfc/rfc4180), with the leniencies of `encoding/csv` | 73 edge cases, `encoding/csv` (differential and fuzz tests) |
| CUE | [`zeroparse.dev/alpha-parsers/cue`](cue/) | [The CUE Language Specification](https://cuelang.org/docs/reference/spec/), as `cue/parser` v0.17.1 parses it | `cue/parser`: 4,061 sources of the CUE repository, 32,488 mutants, 15,000 generated inputs and 4,126 literals (trees, offsets and comments; no difference); fuzz tests |
| DuckDB SQL | [`zeroparse.dev/alpha-parsers/duckdb`](duckdb/) | DuckDB 1.5.6's SQL dialect | DuckDB 1.5.6 on 158,605 texts (acceptance, statement splitting and the AST of SELECT statements; 13 known deviations) |
| Go | [`zeroparse.dev/alpha-parsers/golang`](golang/) | [The Go Programming Language Specification](https://go.dev/ref/spec) (go1.27), as `go/parser` parses it; `ParseFile` returns `go/ast` | `go/parser`: every file of GOROOT/src (whole trees, positions included), the programs embedded in its tests, mutations and fuzz tests (no difference) |
| JSON | [`zeroparse.dev/alpha-parsers/json`](json/) | [RFC 8259](https://www.rfc-editor.org/rfc/rfc8259) | [JSONTestSuite](https://github.com/nst/JSONTestSuite) (every `y_` and `n_` case), `encoding/json` (differential and fuzz tests) |
| Python | [`zeroparse.dev/alpha-parsers/python`](python/) | Python 3.14, as CPython 3.14.0 parses it | CPython's `ast.dump`, positions included, on the standard library, `Lib/test` and its 61,540 code snippets (no difference); mutation and fuzz tests |
| TypeScript | [`zeroparse.dev/alpha-parsers/typescript`](typescript/) | TypeScript 5.9 and TSX, as the TypeScript 5.9.3 compiler parses them | The compiler's trees on its own test cases (14,799 of 14,804 files agree), its `lib` declaration files and mutated sources |
| XML | [`zeroparse.dev/alpha-parsers/xml`](xml/) | [XML 1.0 (Fifth Edition)](https://www.w3.org/TR/xml/), non-validating, with [Namespaces in XML 1.0](https://www.w3.org/TR/xml-names/) | [W3C XML Conformance Test Suite](https://www.w3.org/XML/Test/) (no failure among the XML 1.0 tests that need no external entity), `encoding/xml` (differential and fuzz tests) |
| YAML | [`zeroparse.dev/alpha-parsers/yaml`](yaml/) | [YAML 1.2.2](https://yaml.org/spec/1.2.2/), with the core schema for Go values | [yaml-test-suite](https://github.com/yaml/yaml-test-suite) (all 402 tests: events and JSON values) |

## Install

The module `zeroparse.dev/alpha-parsers` is not published, so `go get zeroparse.dev/alpha-parsers/json` does not work
from another module. It requires a published version of `zeroparse.dev/zp`, so a checkout
of this repository is all a parser needs. Use it in one of two ways.

With a `go.work` file next to your module's `go.mod`:

```bash
go work init .
go work use ~/src/alpha-parsers
```

or with a `replace` directive in your own `go.mod`:

```text
require zeroparse.dev/alpha-parsers v0.0.0

replace zeroparse.dev/alpha-parsers => /path/to/alpha-parsers
```

(`/path/to/alpha-parsers` is the full path of your checkout; `go.mod` does not expand `~`.) Run `go mod tidy` afterwards
to record the requirements of the module.

```go
import "zeroparse.dev/alpha-parsers/json"

v, err := json.Decode(`{"name": "zp", "tags": ["peg", "operators-rule"]}`) // map[string]any{...}
ast, err := json.ParseValue(src)                                      // typed values with positions
```

Every parser has the API of a parser generated with `zp gen -types -recognize` (see
[code generation](https://github.com/zeroparse/zp/blob/main/docs/guide/generated-go-api.md#the-generated-api)), plus functions written for its language:

| | |
|:--|:--|
| `ParseValue(input, opts...)` | The document as the Go types the grammar declares, each with its `Span` in the input |
| `Parse(input, opts...)` | The document as a tree of `*Node`, as the engine returns it |
| `Recognize(input, opts...)` | Only checks the input, without building anything |
| `*SyntaxError` | The position (`Line`, `Col`, `Pos`) and the expected tokens of a syntax error |

Positions are in code points by default; pass `Bytes` for byte offsets.

## Layout

Each directory holds:

| File | Contents |
|:--|:--|
| `<language>.zp` | The grammar. It can be used with the engine too (`zp.CompileSource`), for streaming, incremental parsing or the `zp` tools. |
| `parser.go`, `api.go`, `types/types.go` | The parser generated from the grammar, the names it re-exports from the library, and the Go types of the grammar (the package `types`). The parser imports the library package `zeroparse.dev/zp/codegen/go/lib` (generated with `-lib import`). Do not edit them. |
| other `.go` files | Code written for the language: documentation, conversions to Go values, helpers. Those in `types/` hold the methods of the types of the grammar (a method can only be declared in the package of its type); what needs the parser or the decoder is a function of the module instead. |
| `*_test.go` | Tests against the language's specification, conformance suites and reference implementations; examples; benchmarks |
| `testdata/*.txt`, `*.golden` | Inputs and the expected trees and errors, checked on the generated parser here and on every backend of the engine by [parsers_test.go](parsers_test.go) |

## Development

From the root of the repository:

```bash
go generate .       # regenerate every parser (after changing a grammar, the library or the generator)
go test .           # generated files up to date; golden files on every backend of the engine
go test . -update   # rewrite the golden files from the engine
./test.sh           # fetch the test corpora, then vet and test every package
go test ./json/...  # the tests of one parser
```

The generation commands are the `go:generate` lines of [generate.go](generate.go); the up-to-date test reads them, so a
new parser needs only a new line there. The third-party corpora of the tests are fetched, not stored here (see
[development](https://github.com/zeroparse/zp/blob/main/docs/development.md#third-party-test-corpora)); a test whose corpus is not fetched is skipped with the command to run. Run `./test.sh` against a
checkout of ZeroParse (with a `go.work` that uses both) after changing its engine or generator.

## Status

alpha-parsers is a preview of the ready-made parsers, kept apart from the ZeroParse repository. It requires
development versions of ZeroParse, and it is discarded when ZeroParse publishes its first release.
The grammars under `testdata/grammars/` of ZeroParse are the copies that its tests read; the ones here are the sources
the parsers are generated from.
