#!/bin/sh

for file in "$(dirname "$0")"/*/*.chorego; do
  echo "Processing file: $file"
  base_name=$(basename "$file" .chorego)
  go run ../main.go "$file" > "$(dirname "$file")/${base_name}.go"
done
