#!/bin/bash
set -euo pipefail

NAMESPACE="github.com/nikoheikkila/"

for package in */; do
    go test -v "${NAMESPACE}${package}"
done
