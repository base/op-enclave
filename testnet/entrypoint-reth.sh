#!/bin/bash

RETH_DATA_DIR=${RETH_DATA_DIR:-/data}
echo "$L2_ENGINE_JWT" > /tmp/engine.jwt

exec ./reth \
  --datadir="$RETH_DATA_DIR" \
  --chain="$GENESIS_FILE_PATH" \
  --http \
  --http.corsdomain="*" \
  --http.vhosts="*" \
  --http.addr=0.0.0.0 \
  --http.port=8545 \
  --http.api=web3,debug,eth,net,engine \
  --authrpc.addr=0.0.0.0 \
  --authrpc.port=8551 \
  --authrpc.vhosts="*" \
  --authrpc.jwtsecret=/tmp/engine.jwt \
  --port=30303 \
  --db.read-transaction-timeout=0
