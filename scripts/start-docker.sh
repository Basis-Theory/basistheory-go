#!/bin/bash
set -e

current_directory="$PWD"

cd $(dirname $0)/..
cd vault-api

export GIT_SHA=$(git rev-parse HEAD)

if [ "$IS_PR_WORKFLOW" != true ]
then
  docker-compose pull
  docker-compose up -d
else
  docker-compose pull >/dev/null 2>&1
  docker-compose up -d >/dev/null 2>&1
fi

result=$?

cd "$current_directory"

exit $result
