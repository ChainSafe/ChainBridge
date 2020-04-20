#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


# Exit on failure
set -e

echo "Checking Binding"
git diff --exit-code bindings/ 
if [ $? -eq 0 ]
then
  echo "Bindings are good"
  exit 0
else
  echo "Bindings are out of date"
  exit 1
fi