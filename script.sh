#!/bin/bash

run_process() {
    echo "Running $1"
    cd $(pwd)/services/$1 &&
    go mod tidy -compat=1.24.0 &&
    cd -
}

run_process bff
run_process media
run_process series
run_process tmdb
run_process movie
run_process season
run_process episode
run_process ../shared

