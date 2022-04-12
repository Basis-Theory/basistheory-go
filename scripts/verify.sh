#!/bin/bash
set -e

current_directory="$PWD"

cd $(dirname $0)

time {
    ./vault-checkout.sh
    ./stop-docker.sh
    ./start-docker.sh
    ./service-up.sh
    ./acceptance-test.sh
}

cd "$current_directory"
