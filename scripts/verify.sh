#!/bin/sh
set -e

current_directory="$PWD"

cd $(dirname $0)

pwd

./stop-docker.sh
./start-docker.sh
./service-up.sh
./acceptance-test.sh

cd "$current_directory"
