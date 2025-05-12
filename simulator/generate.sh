#!/bin/sh

for file in "$(dirname "$0")"/*/*.tempo; do
  echo "Processing file: $file"
  base_name=$(basename "$file" .tempo)
  go run ../main.go "$file" > "$(dirname "$file")/${base_name}.go"
done
