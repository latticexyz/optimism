#!/bin/sh
set -exu

RETH_DATA_DIR=/reth
RETH_CHAINDATA_DIR="$RETH_DATA_DIR/db"
GENESIS_FILE_PATH="${GENESIS_FILE_PATH:-/genesis.json}"

if [ ! -d "$RETH_CHAINDATA_DIR" ]; then
	echo "$RETH_CHAINDATA_DIR missing, running init"
	echo "Initializing genesis."
	op-reth init \
		--datadir="$RETH_DATA_DIR" \
		--chain "$GENESIS_FILE_PATH"
else
	echo "$RETH_CHAINDATA_DIR exists."
fi

exec op-reth node \
  --datadir="$RETH_DATA_DIR" \
  --chain="$GENESIS_FILE_PATH" \
  --http \
  --http.addr=0.0.0.0 \
  --http.corsdomain="*" \
  --ws \
  --ws.addr=0.0.0.0 \
  --ws.origins="*" \
  --authrpc.addr=0.0.0.0 \
  "$@"
