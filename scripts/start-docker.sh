#!/bin/bash
set -e

current_directory="$PWD"

cd $(dirname $0)/..
cd vault-api

export GIT_SHA=$(git rev-parse HEAD)

docker-compose pull

if [ "$IS_PR_WORKFLOW" != true ] ; then
  docker-compose up -d --build
fi

docker-compose up -d

result=$?

cd "$current_directory"

exit $result
