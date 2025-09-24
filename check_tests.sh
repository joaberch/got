#!/bin/bash
missing=()

for file in $(find . -type f -name "*.go" ! -path "./unit_tests/*" ! -name "*_test.go"); do
  base=$(basename "$file" .go)
  dir=$(dirname "$file")
  test_path="unit_tests/${dir}_test/${base}_test.go"
  if [ ! -f "$test_path" ]; then
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