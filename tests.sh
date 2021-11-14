#!/bin/bash

mkdir -p example/
cp -r inputs/*.go example/

for fullpath in ./example/*.go; do
    filename=`basename $fullpath`
    go run main.go -fix -- "./example/$filename" || echo "failed $filename"

    if ! cmp -s "./example/$filename" "./outputs/$filename"; then
        echo "$filename is not the same"
    fi
done

rm -rf example/
