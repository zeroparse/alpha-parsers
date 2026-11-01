#!/bin/sh
# Fetches the third-party test corpora, then vets and tests every package of the module (the refgen modules under
# */internal/refgen are not part of it). Arguments are passed to go test, for example ./test.sh -short. It tests
# against the ZeroParse versions that go.mod requires, or against a checkout of ZeroParse with a go.work that uses
# both. A failed corpus fetch stops the run, so that no test is skipped for a missing corpus.
set -e
cd "$(dirname "$0")"
go run ./hack/corpora fetch || { echo "test.sh: fetching the corpora failed" >&2; exit 1; }
go vet ./...
go test "$@" ./...
