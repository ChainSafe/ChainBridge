#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


# Exit on failure
set -e

echo "" > coverage.txt

for d in $(go list ./...); do
    go test -race -coverprofile=profile.out -covermode=atomic "$d"
    if [[ -f profile.out ]]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
