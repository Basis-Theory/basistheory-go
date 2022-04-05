#!/bin/sh
set -e

current_directory="$PWD"

cd $(dirname $0)/..
cd vault-api

docker-compose down -v

result=$?

cd "$current_directory"

exit $result
