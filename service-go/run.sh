#!/usr/bin/env bash
set -x
serviceimage=$1
logLevel=""
if [ -n "$2" ]; then
 logLevel=$2
fi
logGrpc=""
if [ -n "$3" ]; then
 logGrpc=$3
fi
docker container run --network="host" ${serviceimage} $logLevel $logGrpc &
set +x
