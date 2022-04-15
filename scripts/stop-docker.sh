#!/bin/bash
set -e

current_directory="$PWD"

cd $(dirname $0)/..
cd vault-api

if [ "$IS_PR_WORKFLOW" != true ]
then
  docker-compose down -v
else
  echo "stopping docker"
  docker-compose down -v >/dev/null 2>&1
fi

result=$?

cd "$current_directory"

exit $result
