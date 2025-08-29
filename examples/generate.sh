#!/bin/sh

go build ..

for file in "$(dirname "$0")"/*/*.tempo; do
  echo "Processing file: $file"
  base_name=$(basename "$file" .tempo)
  ./tempo build --lang=go --package="$base_name" "$file" > "$(dirname "$file")/${base_name}.go"
  ./tempo build --lang=ts --runtime="../../typescript/runtime.ts" "$file" > "$(dirname "$file")/${base_name}.ts"
  ./tempo build --lang=java --package="$base_name" "$file" > "$(dirname "$file")/Choreography.java"
done

rm ./tempo
