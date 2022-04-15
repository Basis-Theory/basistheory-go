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
  echo "pulling docker"
  docker-compose pull >/dev/null 2>&1
  echo "starting docker"
  docker-compose up -d >/dev/null 2>&1
fi

echo "result: $result"

result=$?

cd "$current_directory"

exit $result
