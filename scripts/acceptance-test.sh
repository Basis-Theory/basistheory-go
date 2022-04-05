#!/bin/sh

current_directory="$PWD"

cd $(dirname $0)
cd ../

go vet
go clean -testcache
go test ./...

cd "$current_directory"
