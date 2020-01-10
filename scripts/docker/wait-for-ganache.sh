#!/bin/bash

set -e

port="$1"
shift
cmd="$@"

echo "here"

until curl -X POST http://localhost:$port --data '{"jsonrpc":"2.0","method":"web3_clientVersion","params":[],"id":67}'; do
  echo "Ganache is unavailable - sleeping"
  sleep 3
done

echo "Ganache is up - executing command"
exec $cmd