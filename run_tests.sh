#!/bin/bash
set -euo pipefail

ROOT=$(dirname "$0")

cd "${ROOT}"
golangci-lint run

for package in */; do
    go test -v -cover "${ROOT}/${package}"
done
