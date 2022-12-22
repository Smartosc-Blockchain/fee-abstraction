#!/bin/bash

source scripts/vars.sh

CONTRACT_ADDRESS=$1
ACCOUNT="test"
SLEEP_TIME="15"
KEYRING="test"

QUERY_MESSAGE=$(jq --null-input '{"list_channels":{}}')

# query wasm
CHANNEL_ID=$(${BINARY[0]} query wasm contract-state smart "$CONTRACT_ADDRESS" "$QUERY_MESSAGE" --node "${NODE[0]}" --chain-id "${CHAINID[0]}" --output json | jq -r ".data.channels[0].id")
echo $CHANNEL_ID