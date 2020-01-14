#!/bin/sh
# wait-for-ganache.sh

set -e

port="$1"
shift
cmd="$@"

until curl -X POST http://$host:8545 --data '{"jsonrpc":"2.0","method":"web3_clientVersion","params":[],"id":67}'; do
  >&2 echo "Ganache is unavailable - sleeping"
  sleep 1
done

>&2 echo "Ganache is up - executing command"
exec $cmd