#!/bin/bash

current_directory="$PWD"

cd $(dirname $0)
cd ../

go vet
go clean -testcache
go test ./tests

result=$?

cd "$current_directory"

exit $result
