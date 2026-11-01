#!/bin/sh
# Fetches the third-party test corpora, then vets and tests every package and the nested refgen modules' build.
# Arguments are passed to go test, for example ./test.sh -short. Needs a go.work that uses the main module
# (go work init . ../zeroparse). A failed corpus fetch stops the run, so that no test is skipped for a missing corpus.
set -e
cd "$(dirname "$0")"
go run ./hack/corpora fetch || { echo "test.sh: fetching the corpora failed" >&2; exit 1; }
go vet ./...
go test "$@" ./...
