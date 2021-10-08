#!/bin/sh

if [ "$SKIP_LINT" = true ] || [ "$SKIP_LINT" = 1 ]
then
  echo SKIP_LINT is set, skipping golangci-lint run...
else
  golangci-lint run --fix
fi

go test ./...
