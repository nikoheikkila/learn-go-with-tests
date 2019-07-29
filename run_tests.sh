#!/bin/bash
set -euo pipefail

ROOT=$(dirname "$0")

for package in */; do
    go test -v "$ROOT/${package}"
done
