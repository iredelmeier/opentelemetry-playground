#!/bin/bash

set -eu

ROOT_DIR="$(realpath "$(dirname "${BASH_SOURCE[0]}")/..")"

while read -r module; do
  {
    cd "$module"

    # Avoid leaving behind build artifacts
    if [[ "$(find . -maxdepth 1 -name "main.go" | wc -l)" -gt 0 ]]; then
      output_dir="$(mktemp -d)"
      output_file="${output_dir}/tmp"

      go build -o "${output_file}" ./...
    else
      go build ./...
    fi
  }
done < <(find "$ROOT_DIR" -name "go.mod" -printf "%h\\n")
