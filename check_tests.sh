#!/bin/bash
missing=()

for file in $(find . -type f -name "*.go" ! -path "./unit_tests/*" ! -name "*_test.go"); do
  base=$(basename "$file" .go)
  found=$(find unit_tests -type f -name "*_test.go" | grep "/${base}_test.go" || true)
  if [ -z "$found" ]; then
    missing+=("$file")
  fi
done

if [ ${#missing[@]} -ne 0 ]; then
  echo "Go file without unit tests:"
  for f in "${missing[@]}"; do
    echo " - $f"
  done
  exit 1
else
  echo "All files have unit tests"
fi