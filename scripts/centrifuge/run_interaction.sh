#!/usr/bin/env bash

# Exit on failure
set -e

pushd ./scripts/centrifuge/cent-interactions

go run . "$@"

popd