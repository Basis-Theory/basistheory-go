#!/bin/bash

current_directory="$PWD"

cd $(dirname $0)
cd ../

echo "running clean operation"
go clean -testcache
echo "running go tests"
go test ./tests

result=$?

echo "result: $result"

cd "$current_directory"

exit $result
