// Package parsers generates and checks the parsers in its subdirectories, packages of this module
// (zeroparse.dev/alpha-parsers/<language>) that depend only on the standard library, the ZeroParse library (zeroparse.dev/zp/codegen/go/lib) and core (zeroparse.dev/zp/core).
//
// Each subdirectory holds the grammar (<language>.zp), the parser generated from it (parser.go, api.go, and types/types.go; they import the library), code
// written by hand around the generated API, and tests against the language's reference implementation
// or conformance suite. The tests of this package check that every generated file is up to date and that
// every backend of the engine gives the results recorded in the subdirectories' golden files, which the
// generated parsers are checked against in their own packages. From the root of the repository,
// regenerate the parsers with go generate ., test one parser with go test ./<language>/..., and
// all of them with ./test.sh.
package parsers

//go:generate go run zeroparse.dev/zp/cmd/zp gen -g cel/cel.zp -module zeroparse.dev/alpha-parsers/cel -types -typed-only -nodoc -lib import -o cel
//go:generate go run zeroparse.dev/zp/cmd/zp gen -g csv/csv.zp -module zeroparse.dev/alpha-parsers/csv -types -recognize -limits -nodoc -lib import -o csv
//go:generate go run zeroparse.dev/zp/cmd/zp gen -g cue/cue.zp -module zeroparse.dev/alpha-parsers/cue -types -typed-only -nodoc -lib import -o cue
//go:generate go run zeroparse.dev/zp/cmd/zp gen -g duckdb/duckdb.zp -module zeroparse.dev/alpha-parsers/duckdb -types -typed-only -nodoc -lib import -o duckdb
//go:generate go run zeroparse.dev/zp/cmd/zp gen -g golang/golang.zp -expectations -module zeroparse.dev/alpha-parsers/golang -types -typed-only -nodoc -lib import -o golang
//go:generate go run zeroparse.dev/zp/cmd/zp gen -g json/json.zp -module zeroparse.dev/alpha-parsers/json -types -recognize -limits -nodoc -lib import -o json
//go:generate go run zeroparse.dev/zp/cmd/zp gen -g python/python.zp -expectations -module zeroparse.dev/alpha-parsers/python -types -typed-only -nodoc -lib import -o python
//go:generate go run zeroparse.dev/zp/cmd/zp gen -g typescript/typescript.zp -expectations -module zeroparse.dev/alpha-parsers/typescript -types -typed-only -nodoc -lib import -o typescript
//go:generate go run zeroparse.dev/zp/cmd/zp gen -g xml/xml.zp -expectations -module zeroparse.dev/alpha-parsers/xml -types -recognize -limits -nodoc -lib import -o xml
//go:generate go run zeroparse.dev/zp/cmd/zp gen -g yaml/yaml.zp -module zeroparse.dev/alpha-parsers/yaml -types -recognize -limits -nodoc -lib import -o yaml
