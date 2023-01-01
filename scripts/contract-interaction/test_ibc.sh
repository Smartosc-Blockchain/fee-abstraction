#!/bin/bash

ACCOUNT="test"
SLEEP_TIME="15"
KEYRING="test"

source scripts/vars.sh

echo ${BINARY[1]}


IBC_DENOM=$(${BINARY[1]} q ibc-transfer denom-hash transfer/channel-0/ujuno --node ${NODE[1]} -o json | jq -r .hash)